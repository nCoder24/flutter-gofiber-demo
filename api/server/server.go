package server

import (
	"demo/api/handler"
	"demo/api/router"
	"demo/core/service"
	"demo/database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func initialize(app *fiber.App, services service.Services) {
	app.Use(logger.New())

	handler := handler.Handler{
		User: handler.NewUserHandler(services.UserService),
	}

	router.RegisterRoutes(app, handler)
}

func New(db database.DB) *fiber.App {
	app := fiber.New()
	initialize(app, service.InitServices(db))
	return app
}
