package errors

import "github.com/gofiber/fiber/v2"

var (
	InvalidCredentials = fiber.NewError(400, "Invalid Credentials")
	UserNotFound       = fiber.NewError(404, "User Not Found")

	CouldNotAddUser = fiber.NewError(500, "Could Not Add User")
	CouldNotGetUser = fiber.NewError(500, "Could Not Get User")
)
