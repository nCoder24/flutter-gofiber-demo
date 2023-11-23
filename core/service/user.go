package service

import (
	"demo/core/models"
	"demo/database"
)

type UserService struct {
	db database.UserDB
}

func initUserService(db database.DB) UserService {
	return UserService{db.UserDB}
}

func (service *UserService) GetAccount(username string) (models.Account, error) {
	data, err := service.db.GetAccountByUsername(username)

	if err != nil {
		return models.Account{}, err
	}

	return models.AccountWith(models.AccountDetails(data)), nil
}

func (service *UserService) CreateAccount(acDetails models.AccountDetails) error {
	return service.db.InsertNewAccount(acDetails)
}

func (service *UserService) AuthenticateUser(username, password string) (bool, error) {
	ac, err := service.GetAccount(username)

	if err != nil {
		return false, err
	}

	return ac.Authenticate(password), nil
}
