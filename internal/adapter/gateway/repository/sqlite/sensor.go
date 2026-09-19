package sqliterepo

import (
	"context"
	"fmt"

	"dmc/internal/domain/sensor"
	"dmc/internal/infrastructure/database/sqlc"
)

type sensorRepository struct {
	queries *sqlc.Queries
}

func NewSensorRepository(queries *sqlc.Queries) *sensorRepository {
	return &sensorRepository{queries}
}

func (r *sensorRepository) Create(ctx context.Context, sens sensor.Sensor) error {
	if err := r.queries.CreateSensorData(ctx, sqlc.CreateSensorDataParams{
		Temperature: sens.Temperature,
		Humidity:    sens.Humidity,
		CreatedAt:   sens.CreatedAt,
	}); err != nil {
		return fmt.Errorf("sqliterepo create: %w", err)
	}
	return nil
}
