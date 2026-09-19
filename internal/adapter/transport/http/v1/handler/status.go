package httphandlerv1

import (
	"context"

	"dmc/internal/application/port/inbound"
	"dmc/pkg/constants"
	"github.com/gofiber/fiber/v3"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	responsev1 "dmc/pkg/response/v1"
)

type StatusHandler struct {
	statusUC inbound.StatusUseCase
}

func NewStatusHandler(statusUC inbound.StatusUseCase) *StatusHandler {
	return &StatusHandler{statusUC}
}

func (h *StatusHandler) Check(c fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.Context(), constants.RequestTimeout)
	defer cancel()
	l := *log.Ctx(ctx)
	l.WithLevel(zerolog.GlobalLevel()).Msg("check")
	backend, database, mqtt := h.statusUC.Check(ctx)
	return c.Status(fiber.StatusOK).JSON(responsev1.Response{
		Data: map[string]any{
			"backend":  backend,
			"database": database,
			"mqtt":     mqtt,
		},
	})
}
