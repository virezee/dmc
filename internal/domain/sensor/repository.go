package sensor

import "context"

type Repository interface {
	Create(ctx context.Context, sensor Sensor) error
}
