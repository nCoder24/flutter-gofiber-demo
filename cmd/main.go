package main

import (
	"demo/api/server"
	"demo/config"
	"demo/database"
	"os"
)

func main() {
	config.Config()

	dbURI := os.Getenv("MONGODB_URI")
	port := ":" + os.Getenv("PORT")

	db := database.GetDB(database.Connect(dbURI))
	server.New(db).Listen(port)
}
