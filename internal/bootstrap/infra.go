package bootstrap

import (
	"context"
	"database/sql"

	"github.com/rs/zerolog/log"

	"dmc/internal/infrastructure/database"
	"dmc/internal/infrastructure/database/sqlc"
	"dmc/internal/infrastructure/mqtt"
)

type Infra struct {
	DB   *sql.DB
	SQLC *sqlc.Queries
	MQTT *mqtt.MQTT
}

func NewInfra(ctx context.Context) *Infra {
	db, err := database.NewSQLite(ctx, "./data/dmc.db")
	if err != nil {
		log.Fatal().Err(err).Msg("failed to initialize sqlite")
	}
	mqtt, err := mqtt.NewMQTT("tcp://mqtt:1883", "dmc-api")
	if err != nil {
		log.Fatal().Err(err).Msg("failed to initialize mqtt")
	}
	return &Infra{
		DB:   db,
		SQLC: sqlc.New(db),
		MQTT: mqtt,
	}
}

func (i *Infra) Shutdown() {
	i.MQTT.Close()
	if err := i.DB.Close(); err != nil {
		log.Error().Err(err).Msg("failed to close sqlite")
	}
}
