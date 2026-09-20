package router

import (
	"dmc/internal/adapter/transport/http/v1/handler"
	"dmc/pkg/response"
	"github.com/oaswrap/spec/adapter/fiberv3openapi"
	"github.com/oaswrap/spec/option"

	requestv1 "dmc/internal/adapter/transport/http/v1/request"
	responsev1 "dmc/pkg/response/v1"
)

func RegisterDeviceRoutes(r fiberv3openapi.Router, hdlr *httphandlerv1.DeviceHandler) {
	r.Post("/device-control", hdlr.SendCommand).With(deviceDoc()...)
}

func deviceDoc() []option.OperationOption {
	return []option.OperationOption{
		option.Tags("Device"),
		option.Summary("Control greenhouse device"),
		option.Description(
			"Publishes an ON or OFF command for a greenhouse device through MQTT.",
		),
		option.OperationID("deviceControl"),
		option.Request(
			new(requestv1.DeviceControl),
			option.ContentDescription(
				"Provide the device ID and the control command.",
			),
			option.ContentRequired(),
		),
		option.Response(
			202,
			nil,
			option.ContentDescription(
				"The command was successfully published to the MQTT broker. This does not confirm that the device received or executed the command.",
			),
		),
		option.Response(
			400,
			new(struct {
				Code   string               `json:"code"`
				Errors []responsev1.Message `json:"errors,omitempty"`
			}),
			option.ContentDescription(
				"The request payload could not be read or one or more fields failed validation.",
			),
			option.ContentNamedExample(
				"invalidPayload",
				responsev1.Message{
					Code: response.CodeInvalidPayload,
				},
				option.ExampleSummary("The request payload is invalid"),
				option.ExampleDescription(
					"The JSON is malformed, a field has an invalid type, or an unknown field was provided.",
				),
			),
			option.ContentNamedExample(
				"validationFailed",
				map[string]any{
					"code": response.CodeValidationFailed,
					"errors": []responsev1.Message{
						{
							Code: response.CodeValidationRequired,
							Params: map[string]any{
								"field": "device_id",
							},
						},
						{
							Code: response.CodeValidationInvalid,
							Params: map[string]any{
								"field": "command",
							},
						},
					},
				},
				option.ExampleSummary("A device control field failed validation"),
				option.ExampleDescription(
					"`device_id` is required and `command` must be either `ON` or `OFF`.",
				),
			),
		),
		option.Response(
			500,
			new(struct {
				Code   string         `json:"code"`
				Params map[string]any `json:"params,omitempty"`
			}),
			option.ContentDescription(
				"The command could not be published because an internal MQTT error occurred.",
			),
			option.ContentExample(
				responsev1.Message{
					Code: response.CodeInternalError,
					Params: map[string]any{
						"ref": "a1B2c3D4e5",
					},
				},
			),
		),
	}
}
