package server

import (
	"demo/api/handler"
	"demo/api/router"
	"demo/core/operator"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Init(app *fiber.App, operators operator.Services) {
	app.Use(logger.New())

	handler := handler.Handler{
		User: handler.NewUserHandler(operators.UserService),
	}

	router.RegisterRoutes(app, handler)
}
