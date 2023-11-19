package router

import (
	"demo/app/service"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App, services service.Services) {
	registerUserRoutes(app.Group("user"), services)
}
