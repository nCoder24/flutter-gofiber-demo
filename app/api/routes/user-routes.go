package routes

import (
	"demo/app/api/handlers"
	"demo/app/services"

	"github.com/gofiber/fiber/v2"
)

func registerUserRoutes(router fiber.Router, services services.Services) {
	handler := handlers.NewUserHandler(services.UserService)

	router.Get("/", handler.GetUsers)
	router.Post("/", handler.AddUser)
}
