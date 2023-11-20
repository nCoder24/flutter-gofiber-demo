package service

import (
	"demo/core/models"
	"demo/db"
)

type UserService struct {
	db db.UserDB
}

func initUserService(db db.DB) UserService {
	return UserService{db.UserDB}
}

func (service *UserService) GetAccount(username string) (models.Account, error) {
	data, err := service.db.GetAccountByUsername(username)

	if err != nil {
		return models.Account{}, err
	}

	return models.AccountWith(data.Username, data.Password), nil
}

func (service *UserService) CreateAccount(acDetails models.AccountDetails) error {
	return service.db.InsertNewAccount(acDetails)
}
