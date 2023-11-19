package main

import (
	"demo/app"
)

func main() {
	app.CreateApp().Listen(":8080")
}
