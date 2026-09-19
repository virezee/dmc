package inbound

import (
	"context"

	"dmc/internal/domain/sensor"
)

type SensorUseCase interface {
	Receive(ctx context.Context, sensor sensor.Sensor) error
}
