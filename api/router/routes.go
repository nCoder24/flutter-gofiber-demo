package router

import (
	"demo/api/handler"
	"demo/api/middleware/parser"
	"demo/api/middleware/validator"
	"demo/core/models"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App, handler handler.Handler) {
	router := app.Group("/user")

	router.Post("/",
		parser.ParseBody[models.UserDetails],
		validator.Validate,
		handler.User.CreateUser,
	)

	router.Get("/:username", handler.User.GetUser)
	router.Post("/login", handler.User.Login)
}
