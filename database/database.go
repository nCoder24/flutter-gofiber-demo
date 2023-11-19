package database

import "demo/domain/models"

func GetUsers() ([]models.User, error) {
	return []models.User{
		{Name: "Sauma"},
		{Name: "Utsab"},
	}, nil
}
