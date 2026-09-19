package router

import (
	"dmc/internal/adapter/transport/http/v1/handler"
	"github.com/gofiber/fiber/v3"
)

func RegisterSensorRoutes(r fiber.Router, hdlr *httphandlerv1.SensorHandler) {
	r.Post("/sensor-data", hdlr.Receive)
}
