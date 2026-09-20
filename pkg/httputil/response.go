package httputil

import (
	"context"

	"github.com/gofiber/fiber/v3"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"dmc/pkg/cryptox"
	"dmc/pkg/response"

	responsev1 "dmc/pkg/response/v1"
)

func InternalServerErrorWithErr(ctx context.Context, c fiber.Ctx, err error) error {
	ref := cryptox.GenerateRef()
	log.Ctx(ctx).Error().Stack().Err(err).Str("ref", ref).Send()
	return internalError(c, ref)
}

func InternalServerErrorWithLoggerErr(l zerolog.Logger, c fiber.Ctx, err error) error {
	ref := cryptox.GenerateRef()
	l.Error().Stack().Err(err).Str("ref", ref).Send()
	return internalError(c, ref)
}

func internalError(c fiber.Ctx, ref string) error {
	return c.Status(fiber.StatusInternalServerError).JSON(responsev1.Message{
		Code:   response.CodeInternalError,
		Params: map[string]any{"ref": ref},
	})
}
