package main

import (
	"demo/api/server"
	"demo/config"
	"demo/core/operator"
	"demo/database"
	"demo/database/connection"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	log.SetFlags(log.Llongfile)
	config := config.GetConfig()
	app := fiber.New()

	db := database.GetDB(connection.Connect(config.MongodbURI))
	operators := operator.InitOperators(db)

	server.Init(app, operators)
	log.Fatal(app.Listen(config.Port))
}
