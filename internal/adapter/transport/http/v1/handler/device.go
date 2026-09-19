package httphandlerv1

import (
	"context"

	"dmc/internal/adapter/transport/http/v1/request"
	"dmc/internal/application/port/inbound"
	"dmc/internal/domain/device"
	"dmc/pkg/constants"
	"dmc/pkg/httputil"
	"github.com/gofiber/fiber/v3"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type DeviceHandler struct {
	devUC inbound.DeviceUseCase
}

func NewDeviceHandler(devUC inbound.DeviceUseCase) *DeviceHandler {
	return &DeviceHandler{devUC}
}

func (h *DeviceHandler) SendCommand(c fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.Context(), constants.RequestTimeout)
	defer cancel()
	req := httputil.Bind[requestv1.DeviceControl](c)
	if req == nil {
		return nil
	}
	id := req.DeviceID
	com := req.Command
	l := log.Ctx(ctx).With().Str("device_id", id).Str("command", com).Logger()
	l.WithLevel(zerolog.GlobalLevel()).Msg("send command")
	if err := h.devUC.SendCommand(ctx, device.Device{
		ID:      id,
		Command: com,
	}); err != nil {
		return httputil.InternalServerErrorWithLoggerErr(l, c, err)
	}
	return c.SendStatus(fiber.StatusNoContent)
}
