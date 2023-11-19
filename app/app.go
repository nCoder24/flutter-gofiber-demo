package app

import (
	"demo/app/api/router"
	"demo/app/service"
	"demo/db"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func CreateApp(db db.DB) *fiber.App {
	app := fiber.New()
	app.Use(logger.New())

	services := service.InitServices(service.InitUserService(db))
	router.RegisterRoutes(app, services)

	return app
}
