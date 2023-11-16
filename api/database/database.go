package database

import "demo/models"

func GetUsers() ([]models.User, error) {
	return []models.User{
		{Name: "Sauma"},
		{Name: "Utsab"},
	}, nil
}
