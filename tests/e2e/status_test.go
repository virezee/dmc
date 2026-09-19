package e2e

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"dmc/internal/bootstrap"
)

func TestStatus(t *testing.T) {
	ctx := context.Background()
	if err := os.MkdirAll("data", 0o755); err != nil {
		t.Fatal(err)
	}
	infra := bootstrap.NewInfra(ctx)
	t.Cleanup(func() {
		infra.Shutdown()
		_ = os.RemoveAll("data")
	})
	app := bootstrap.NewApp(ctx, infra)
	req := httptest.NewRequest(
		http.MethodGet,
		"/status",
		nil,
	)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = resp.Body.Close() }()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected %d, got %d", http.StatusOK, resp.StatusCode)
	}
	var result struct {
		Data struct {
			Backend  bool `json:"backend"`
			Database bool `json:"database"`
			MQTT     bool `json:"mqtt"`
		} `json:"data"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		t.Fatal(err)
	}
	if !result.Data.Backend {
		t.Fatal("expected backend status true")
	}
	if !result.Data.Database {
		t.Fatal("expected database status true")
	}
	if !result.Data.MQTT {
		t.Fatal("expected mqtt status true")
	}
}
