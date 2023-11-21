package router

import (
	"demo/api/handler"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App, handler handler.Handler) {
	router := app.Group("/user")
	router.Post("/", handler.User.CreateAccount)
	router.Get("/:username", handler.User.GetAccount)
	router.Post("/login", handler.User.Login)
}
