package parser

import (
	"demo/api/constant"

	"github.com/gofiber/fiber/v2"
)

func ParseBody[R any](ctx *fiber.Ctx) error {
	resource := new(R)

	if err := ctx.BodyParser(resource); err != nil {
		return err
	}

	ctx.Locals(constant.Resource, *resource)
	return ctx.Next()
}
