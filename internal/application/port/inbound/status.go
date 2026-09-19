package inbound

import "context"

type StatusUseCase interface {
	Check(ctx context.Context) (backend, database, mqtt bool)
}
