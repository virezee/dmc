package main

import (
	"context"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/rs/zerolog/log"

	"dmc/internal/bootstrap"
	"dmc/pkg/logger"
)

func main() {
	time.Local = time.UTC
	logger.App()
	ctx, stop := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
	)
	defer stop()
	infra := bootstrap.NewInfra(ctx)
	defer infra.Shutdown()
	app := bootstrap.NewApp(ctx, infra)
	if err := app.Listen(":3000", fiber.ListenConfig{
		GracefulContext: ctx,
	}); err != nil {
		log.Fatal().Stack().Err(err).Send()
	}
}
