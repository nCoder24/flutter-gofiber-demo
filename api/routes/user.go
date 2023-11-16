package routes

import (
	"demo/services"

	"github.com/gofiber/fiber/v2"
)

func getUsers(c *fiber.Ctx) error {
	users, err := services.GetUsers()
	if err != nil {
		return err
	}

	return c.JSON(users)
}

func addUser(c *fiber.Ctx) error {
	return c.SendString("User Added")
}

func SetUserRoutes(router fiber.Router) {
	router.Get("/", getUsers)
	router.Post("/", addUser)
}
