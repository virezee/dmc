package e2e

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

	"dmc/internal/bootstrap"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func TestDeviceControl(t *testing.T) {
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
	t.Run("success", func(t *testing.T) {
		topic := "greenhouse/control/fan-01"
		messageCh := make(chan string, 1)

		client := mqtt.NewClient(
			mqtt.NewClientOptions().
				AddBroker("tcp://localhost:1883").
				SetClientID("dmc-e2e-device-subscriber"),
		)
		token := client.Connect()
		if !token.Wait() || token.Error() != nil {
			t.Fatal(token.Error())
		}
		defer client.Disconnect(250)
		token = client.Subscribe(topic, 0, func(_ mqtt.Client, msg mqtt.Message) {
			messageCh <- string(msg.Payload())
		})
		if !token.Wait() || token.Error() != nil {
			t.Fatal(token.Error())
		}
		payload := `{"device_id":"fan-01","command":"ON"}`
		req := httptest.NewRequest(
			http.MethodPost,
			"/device-control",
			strings.NewReader(payload),
		)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)
		if err != nil {
			t.Fatal(err)
		}
		defer func() { _ = resp.Body.Close() }()
		if resp.StatusCode != http.StatusNoContent {
			t.Fatalf("expected %d, got %d", http.StatusNoContent, resp.StatusCode)
		}
		select {
		case message := <-messageCh:
			if message != "ON" {
				t.Fatalf("expected ON, got %s", message)
			}
		case <-time.After(2 * time.Second):
			t.Fatal("timeout waiting for mqtt message")
		}
	})
	t.Run("invalid command", func(t *testing.T) {
		payload := `{"device_id":"fan-01","command":"INVALID"}`
		req := httptest.NewRequest(
			http.MethodPost,
			"/device-control",
			strings.NewReader(payload),
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
	t.Run("missing device id", func(t *testing.T) {
		payload := `{"command":"ON"}`
		req := httptest.NewRequest(
			http.MethodPost,
			"/device-control",
			strings.NewReader(payload),
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
