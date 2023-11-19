package services

import (
	"demo/database"
	"demo/domain/models"
)

func GetUsers() ([]models.User, error) {
	return database.GetUsers()
}
