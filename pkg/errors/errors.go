package errors

import "github.com/gofiber/fiber/v2"

var (
	FailedToFetchAccount  = fiber.NewError(500, "Failed To Fetch Account")
	FailedToCreateAccount = fiber.NewError(500, "Failed To Create Account")

	AccountNotFound = fiber.NewError(404, "Account Not Found")
)