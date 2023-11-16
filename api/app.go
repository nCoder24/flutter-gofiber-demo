package main

import (
	"demo/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())

	routes.SetUserRoutes(app.Group("user"))

	app.Listen(":8080")
}
