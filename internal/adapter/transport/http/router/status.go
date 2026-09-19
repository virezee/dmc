package router

import (
	"github.com/oaswrap/spec/adapter/fiberv3openapi"
	"github.com/oaswrap/spec/option"

	httphandlerv1 "dmc/internal/adapter/transport/http/v1/handler"
)

func RegisterStatusRoutes(r fiberv3openapi.Router, hdlr *httphandlerv1.StatusHandler) {
	r.Get("/status", hdlr.Check).With(statusDoc()...)
}

func statusDoc() []option.OperationOption {
	return []option.OperationOption{
		option.Tags("System"),
		option.Summary("Check system status"),
		option.Description(
			"Returns the current status of the backend service, database connection, and MQTT connection.",
		),
		option.OperationID("status"),
		option.Response(
			200,
			new(struct {
				Data struct {
					Backend  bool `json:"backend"`
					Database bool `json:"database"`
					MQTT     bool `json:"mqtt"`
				} `json:"data"`
			}),
			option.ContentDescription(
				"Returns the health status of the backend, database, and MQTT connection.",
			),
			option.ContentExample(
				map[string]any{
					"data": map[string]any{
						"backend":  true,
						"database": true,
						"mqtt":     true,
					},
				},
			),
		),
	}
}
