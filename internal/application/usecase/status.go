package usecase

import (
	"context"
	"database/sql"

	"dmc/internal/application/port/outbound"
)

type statusUseCase struct {
	db   *sql.DB
	mqtt outbound.MQTT
}

func NewStatusUseCase(db *sql.DB, mqtt outbound.MQTT) *statusUseCase {
	return &statusUseCase{db, mqtt}
}

func (u *statusUseCase) Check(ctx context.Context) (backend, database, mqtt bool) {
	backend = true
	database = u.db.PingContext(ctx) == nil
	mqtt = u.mqtt.IsConnected()
	return backend, database, mqtt
}
