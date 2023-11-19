package main

import (
	"demo/app"
	"demo/config"
	"demo/db"
	"os"
)

func main() {
	config.Config()

	dbURI := os.Getenv("MONGODB_URI")
	port := ":" + os.Getenv("PORT")

	db := db.GetDB(db.Connect(dbURI))
	app.CreateApp(db).Listen(port)
}
