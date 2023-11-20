package router

import (
	"demo/api/handler"
	"demo/core/service"

	"github.com/gofiber/fiber/v2"
)

func registerUserRoutes(router fiber.Router, services service.Services) {
	handler := handler.NewUserHandler(services.UserService)

	router.Get("/:username", handler.GetAccount)
	router.Post("/", handler.CreateAccount)
	router.Post("/login", handler.Login)
}
