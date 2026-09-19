package usecase

import (
	"context"
	"time"

	"dmc/internal/domain/sensor"
)

type sensorUseCase struct {
	sensRepo sensor.Repository
}

func NewSensorUseCase(sensRepo sensor.Repository) *sensorUseCase {
	return &sensorUseCase{sensRepo}
}

func (u *sensorUseCase) Receive(ctx context.Context, sens sensor.Sensor) error {
	sens.CreatedAt = time.Now().Format(time.RFC3339)
	return u.sensRepo.Create(ctx, sens)
}
