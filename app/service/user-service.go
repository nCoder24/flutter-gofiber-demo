package service

import (
	"demo/core/models"
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
