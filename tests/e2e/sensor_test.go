package e2e

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

	"dmc/internal/bootstrap"
)

func TestSensor(t *testing.T) {
	ctx := context.Background()
	if err := os.MkdirAll("data", 0o755); err != nil {
		t.Fatal(err)
	}
	infra := bootstrap.NewInfra(ctx)
	t.Cleanup(func() {
		infra.Shutdown()
		_ = os.RemoveAll("./data")
	})
	migration, err := os.ReadFile("../../migrations/20260919100208_init.sql")
	if err != nil {
		t.Fatal(err)
	}
	if _, err := infra.DB.ExecContext(ctx, string(migration)); err != nil {
		t.Fatal(err)
	}
	app := bootstrap.NewApp(ctx, infra)
	t.Run("success", func(t *testing.T) {
		payload := map[string]float64{
			"temperature": 25.5,
			"humidity":    68.2,
		}
		body, err := json.Marshal(payload)
		if err != nil {
			t.Fatal(err)
		}
		req := httptest.NewRequest(
			http.MethodPost,
			"/sensor-data",
			strings.NewReader(string(body)),
		)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)
		if err != nil {
			t.Fatal(err)
		}
		defer func() { _ = resp.Body.Close() }()
		if resp.StatusCode != http.StatusCreated {
			t.Fatalf("expected %d, got %d", http.StatusCreated, resp.StatusCode)
		}
		var temperature, humidity float64
		var createdAt string
		err = infra.DB.QueryRowContext(ctx, `
			SELECT temperature, humidity, created_at
			FROM sensor_data
			ORDER BY id DESC
			LIMIT 1
		`).Scan(&temperature, &humidity, &createdAt)
		if err != nil {
			t.Fatal(err)
		}
		if temperature != 25.5 {
			t.Fatalf("expected temperature 25.5, got %v", temperature)
		}
		if humidity != 68.2 {
			t.Fatalf("expected humidity 68.2, got %v", humidity)
		}
		if _, err := time.Parse(time.RFC3339, createdAt); err != nil {
			t.Fatalf("invalid created_at: %q: %v", createdAt, err)
		}
	})
	t.Run("invalid payload", func(t *testing.T) {
		req := httptest.NewRequest(
			http.MethodPost,
			"/sensor-data",
			strings.NewReader(`{"temperature":"invalid","humidity":68.2}`),
		)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)
		if err != nil {
			t.Fatal(err)
		}
		defer func() { _ = resp.Body.Close() }()
		if resp.StatusCode != http.StatusBadRequest {
			t.Fatalf("expected %d, got %d", http.StatusBadRequest, resp.StatusCode)
		}
	})
	t.Run("missing required field", func(t *testing.T) {
		req := httptest.NewRequest(
			http.MethodPost,
			"/sensor-data",
			strings.NewReader(`{"temperature":25.5}`),
		)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)
		if err != nil {
			t.Fatal(err)
		}
		defer func() { _ = resp.Body.Close() }()
		if resp.StatusCode != http.StatusBadRequest {
			t.Fatalf("expected %d, got %d", http.StatusBadRequest, resp.StatusCode)
		}
	})
}
