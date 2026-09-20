package bootstrap

import (
	"context"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/recover"

	"dmc/internal/adapter/transport/http/router"
	"dmc/internal/application/usecase"
	"dmc/pkg/httputil"

	mqttclient "dmc/internal/adapter/gateway/mqtt"
	sqliterepo "dmc/internal/adapter/gateway/repository/sqlite"
	httphandlerv1 "dmc/internal/adapter/transport/http/v1/handler"
)

func NewApp(ctx context.Context, infra *Infra) *fiber.App {
	sensRepo := sqliterepo.NewSensorRepository(infra.SQLC)
	devMQTT := mqttclient.NewMQTTClient(infra.MQTT)
	sensUC := usecase.NewSensorUseCase(sensRepo)
	statsUC := usecase.NewStatusUseCase(infra.DB, devMQTT)
	sensHdlr := httphandlerv1.NewSensorHandler(sensUC)
	statsHdlr := httphandlerv1.NewStatusHandler(statsUC)
	devUC := usecase.NewDeviceUseCase(devMQTT)
	devHdlr := httphandlerv1.NewDeviceHandler(devUC)
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c fiber.Ctx, err error) error {
			return httputil.InternalServerErrorWithErr(c.Context(), c, err)
		},
	})
	app.Use(recover.New())
	r := newSpecRouter(app)
	router.RegisterSensorRoutes(r, sensHdlr)
	router.RegisterStatusRoutes(r, statsHdlr)
	router.RegisterDeviceRoutes(r, devHdlr)
	return app
}
