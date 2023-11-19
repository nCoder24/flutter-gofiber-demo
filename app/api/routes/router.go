package routes

import (
	"demo/app/services"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App, services services.Services) {
	registerUserRoutes(app.Group("user"), services)
}
