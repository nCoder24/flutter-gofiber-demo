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
		parser.ParseBody[models.UserData],
		validator.Validate,
		handler.User.AddNewUser,
	)

	router.Get("/:username", handler.User.GetUser)
}
