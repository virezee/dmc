package router

import (
	"github.com/oaswrap/spec/adapter/fiberv3openapi"
	"github.com/oaswrap/spec/option"

	"dmc/pkg/response"

	httphandlerv1 "dmc/internal/adapter/transport/http/v1/handler"
	requestv1 "dmc/internal/adapter/transport/http/v1/request"
)

func RegisterSensorRoutes(r fiberv3openapi.Router, hdlr *httphandlerv1.SensorHandler) {
	r.Post("/sensor-data", hdlr.Receive).With(sensorDoc()...)
}

func sensorDoc() []option.OperationOption {
	return []option.OperationOption{
		option.Tags("Sensor"),
		option.Summary("Ingest sensor data"),
		option.Description(
			"Receives temperature and humidity data from the greenhouse sensor system and stores it in the database.",
		),
		option.OperationID("sensorData"),
		option.Request(
			new(requestv1.SensorData),
			option.ContentDescription(
				"Provide the current temperature and humidity readings.",
			),
			option.ContentRequired(),
		),
		option.Response(
			201,
			nil,
			option.ContentDescription(
				"The sensor data was successfully stored.",
			),
		),
		option.Response(
			400,
			new(struct {
				Code   string `json:"code"`
				Errors []struct {
					Code   string `json:"code"`
					Params struct {
						Field string `json:"field"`
					} `json:"params"`
				} `json:"errors"`
			}),
			option.ContentDescription(
				"The request payload is invalid or validation failed.",
			),
			option.ContentNamedExample(
				"invalidPayload",
				map[string]any{
					"code": response.CodeInvalidPayload,
				},
				option.ExampleSummary("The request payload is invalid"),
				option.ExampleDescription(
					"The JSON body is malformed, contains an unsupported field, or contains a value with an invalid type.",
				),
			),
			option.ContentNamedExample(
				"validationFailed",
				map[string]any{
					"code": response.CodeValidationFailed,
					"errors": []map[string]any{
						{
							"code": response.CodeValidationRequired,
							"params": map[string]any{
								"field": "temperature",
							},
						},
						{
							"code": response.CodeValidationRequired,
							"params": map[string]any{
								"field": "humidity",
							},
						},
					},
				},
				option.ExampleSummary("Required sensor fields are missing"),
				option.ExampleDescription(
					"Each invalid field is returned in the errors list.",
				),
			),
		),
		option.Response(
			500,
			new(struct {
				Code   string `json:"code"`
				Params struct {
					Ref string `json:"ref"`
				} `json:"params"`
			}),
			option.ContentDescription(
				"The sensor data could not be stored because an internal database error occurred.",
			),
			option.ContentExample(
				map[string]any{
					"code": response.CodeInternalError,
					"params": map[string]any{
						"ref": "a1B2c3D4e5",
					},
				},
			),
		),
	}
}
