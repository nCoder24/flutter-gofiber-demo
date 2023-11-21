package config

import (
	"github.com/joho/godotenv"
)

func GetConfig() {
	err := godotenv.Load(".env")

	if err != nil {
		panic(err)
	}
}
