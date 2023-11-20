package server

import (
	"demo/api/router"
	"demo/core/service"
	"demo/database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func CreateServer(db database.DB) *fiber.App {
	app := fiber.New()
	app.Use(logger.New())

	router.RegisterRoutes(app, service.InitServices(db))

	return app
}
