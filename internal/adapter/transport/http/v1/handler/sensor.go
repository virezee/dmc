package httphandlerv1

import (
	"context"

	"dmc/internal/application/port/inbound"
	"dmc/internal/domain/sensor"
	"dmc/pkg/constants"
	"dmc/pkg/httputil"
	"github.com/gofiber/fiber/v3"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	requestv1 "dmc/internal/adapter/transport/http/v1/request"
)

type SensorHandler struct {
	sensUC inbound.SensorUseCase
}

func NewSensorHandler(sensUC inbound.SensorUseCase) *SensorHandler {
	return &SensorHandler{sensUC}
}

func (h *SensorHandler) Receive(c fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.Context(), constants.RequestTimeout)
	defer cancel()
	req := httputil.Bind[requestv1.SensorData](c)
	if req == nil {
		return nil
	}
	temp := req.Temperature
	hum := req.Humidity
	l := log.Ctx(ctx).With().Float64("temperature", temp).Float64("humidity", hum).Logger()
	l.WithLevel(zerolog.GlobalLevel()).Msg("receive")
	if err := h.sensUC.Receive(ctx, sensor.Sensor{
		Temperature: temp,
		Humidity:    hum,
	}); err != nil {
		return httputil.InternalServerErrorWithLoggerErr(l, c, err)
	}
	return c.SendStatus(fiber.StatusCreated)
}
