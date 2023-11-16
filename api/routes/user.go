package routes

import "github.com/gofiber/fiber/v2"

func getUsers(c *fiber.Ctx) error {
	return c.SendString("Users")
}

func addUser(c *fiber.Ctx) error {
	return c.SendString("User Added")
}

func SetUserRoutes(router fiber.Router) {
	router.Get("/", getUsers)
	router.Post("/", addUser)
}
