package main

import (
	"demo/api/server"
	"demo/config"
	"demo/core/service"
	"demo/database"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config.GetConfig()
	app := fiber.New()

	dbURI := os.Getenv("MONGODB_URI")
	port := ":" + os.Getenv("PORT")

	db := database.GetDB(database.Connect(dbURI))
	services := service.InitServices(db)

	server.Init(app, services)
	app.Listen(port)
}
