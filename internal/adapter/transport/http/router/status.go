package router

import (
	"github.com/gofiber/fiber/v3"

	httphandlerv1 "dmc/internal/adapter/transport/http/v1/handler"
)

func RegisterStatusRoutes(r fiber.Router, hdlr *httphandlerv1.StatusHandler) {
	r.Get("/status", hdlr.Check)
}
