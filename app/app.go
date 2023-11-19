package app

import (
	"demo/app/api/routes"
	"demo/app/services"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func CreateApp() *fiber.App {
	app := fiber.New()
	app.Use(logger.New())

	services := services.InitServices(services.InitUserService())
	routes.RegisterRoutes(app, services)

	return app
}
