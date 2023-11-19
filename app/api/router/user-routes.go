package router

import (
	"demo/app/api/handler"
	"demo/app/service"

	"github.com/gofiber/fiber/v2"
)

func registerUserRoutes(router fiber.Router, services service.Services) {
	handler := handler.NewUserHandler(services.UserService)

	router.Get("/:username", handler.GetAccount)
	router.Post("/", handler.AddUser)
}
