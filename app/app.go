package app

import (
	"demo/app/api/router"
	"demo/app/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func CreateApp() *fiber.App {
	app := fiber.New()
	app.Use(logger.New())

	services := service.InitServices(service.InitUserService())
	router.RegisterRoutes(app, services)

	return app
}
