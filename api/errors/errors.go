package errors

import "github.com/gofiber/fiber/v2"

var (
	InvalidCredentials = fiber.NewError(400, "Invalid Credentials")
)
