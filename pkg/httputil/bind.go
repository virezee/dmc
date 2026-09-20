package httputil

import (
	"bytes"
	"encoding/json"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"

	"dmc/pkg/response"

	responsev1 "dmc/pkg/response/v1"
)

var validate = func() *validator.Validate {
	v := validator.New(validator.WithRequiredStructEnabled())
	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		if name, _, _ := strings.Cut(fld.Tag.Get("json"), ","); name != "" && name != "-" {
			return name
		}
		return fld.Name
	})
	return v
}()

func Bind[T any](c fiber.Ctx) *T {
	var req T
	if body := c.Body(); len(body) > 0 {
		decoder := json.NewDecoder(bytes.NewReader(body))
		decoder.DisallowUnknownFields()
		if err := decoder.Decode(&req); err != nil {
			_ = c.Status(fiber.StatusBadRequest).JSON(responsev1.Message{
				Code: response.CodeInvalidPayload,
			})
			return nil
		}
	}
	if !validateStruct(c, &req) {
		return nil
	}
	return &req
}

func validateStruct(c fiber.Ctx, req any) bool {
	err := validate.Struct(req)
	if err == nil {
		return true
	}
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		_ = c.Status(fiber.StatusBadRequest).JSON(responsev1.Message{
			Code: response.CodeInvalidPayload,
		})
		return false
	}
	fields := make([]responsev1.Message, len(errs))
	for i, e := range errs {
		r := responsev1.Message{Params: map[string]any{"field": e.Field()}}
		switch e.Tag() {
		case "required":
			r.Code = response.CodeValidationRequired
		default:
			r.Code = response.CodeValidationInvalid
		}
		fields[i] = r
	}
	_ = c.Status(fiber.StatusBadRequest).JSON(responsev1.Response{
		Message: responsev1.Message{Code: response.CodeValidationFailed},
		Errors:  fields,
	})
	return false
}
