package validator

import (
	"demo/api/constant"
	"demo/api/errors"

	"github.com/gofiber/fiber/v2"
)

type ValidatableResource interface {
	IsValid() bool
}

func Validate(ctx *fiber.Ctx) error {
	resource := ctx.Locals(constant.Resource).(ValidatableResource)

	if !resource.IsValid() {
		return errors.InvalidCredentials
	}

	return ctx.Next()
}
