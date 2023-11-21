package server

import (
	"demo/api/handler"
	"demo/api/router"
	"demo/core/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Init(app *fiber.App, services service.Services) {
	app.Use(logger.New())

	handler := handler.Handler{
		User: handler.NewUserHandler(services.UserService),
	}

	router.RegisterRoutes(app, handler)
}
