package router

import (
	"dmc/internal/adapter/transport/http/v1/handler"
	"github.com/gofiber/fiber/v3"
)

func RegisterDeviceRoutes(r fiber.Router, hdlr *httphandlerv1.DeviceHandler) {
	r.Post("/device-control", hdlr.SendCommand)
}
