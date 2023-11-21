package main

import (
	"demo/api/server"
	"demo/config"
	"demo/core/service"
	"demo/database"
	"demo/database/connection"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config := config.GetConfig()
	app := fiber.New()

	db := database.GetDB(connection.Connect(config.MongodbURI))
	services := service.InitServices(db)

	server.Init(app, services)
	app.Listen(config.Port)
}
