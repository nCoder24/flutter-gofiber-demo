package services

import (
	"demo/database"
	"demo/models"
)

func GetUsers() ([]models.User, error) {
	return database.GetUsers()
}
