package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	MongodbURI string
	Port       string
}

func GetConfig() Config {
	err := godotenv.Load(".env")

	if err != nil {
		panic(err)
	}

	return Config{
		MongodbURI: os.Getenv("MONGODB_URI"),
		Port:       os.Getenv("MONGODB_URI"),
	}
}
