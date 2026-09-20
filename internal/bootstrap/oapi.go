package bootstrap

import (
	"github.com/gofiber/fiber/v3"
	"github.com/oaswrap/spec-ui/config"
	"github.com/oaswrap/spec/adapter/fiberv3openapi"
	"github.com/oaswrap/spec/openapi"
	"github.com/oaswrap/spec/option"
)

func newSpecRouter(app *fiber.App) fiberv3openapi.Generator {
	return fiberv3openapi.NewRouter(app, specOptions()...)
}

func specOptions() []option.OpenAPIOption {
	return []option.OpenAPIOption{
		option.WithDocsPath("/docs/v1"),
		option.WithSpecPath("/docs/v1/openapi.json"),
		option.WithScalar(config.Scalar{DarkMode: true, Theme: "deepSpace"}),
		option.WithOpenAPIVersion(openapi.Version320),
		option.WithTitle("DMC Greenhouse API"),
		option.WithInfoSummary("Greenhouse sensor data, device control, and system health API."),
		option.WithDescription(
			"API for greenhouse sensor data ingestion, device control through MQTT, and system health monitoring.",
		),
		option.WithContact(openapi.Contact{Name: "Zee"}),
		option.WithLicense(openapi.License{Name: "Proprietary"}),
		option.WithVersion("1.0.0"),
		option.WithJSONSchemaDialect("https://spec.openapis.org/oas/3.2/dialect/2025-09-17"),
		option.WithServer("http://localhost:3000",
			option.ServerName("Local"),
			option.ServerDescription("Local environment (developer machine).")),
		option.WithTags(
			openapi.Tag{
				Name:        "Sensor",
				Summary:     "Greenhouse sensor data.",
				Description: "Receives and stores sensor data from the greenhouse system.",
				Kind:        "nav",
			},
			openapi.Tag{
				Name:        "Device",
				Summary:     "Greenhouse device control.",
				Description: "Sends ON or OFF commands to greenhouse devices through MQTT.",
				Kind:        "nav",
			},
			openapi.Tag{
				Name:        "System",
				Summary:     "System health status.",
				Description: "Reports backend, database, and MQTT connection status.",
				Kind:        "nav",
			},
		),
	}
}
