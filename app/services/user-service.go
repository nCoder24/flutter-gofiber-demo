package services

import (
	"demo/domain/models"
)

type UserService struct {
}

func InitUserService() UserService {
	return UserService{}
}

func (user *UserService) GetAllAccounts() ([]models.User, error) {
	return []models.User{
		{Name: "Sauma"},
		{Name: "Utsab"},
	}, nil
}
