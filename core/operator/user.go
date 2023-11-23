package operator

import (
	"demo/core/models"
	"demo/database"
)

type UserOperator struct {
	db database.UserDB
}

func initUserService(db database.DB) UserOperator {
	return UserOperator{db.UserDB}
}

func (service *UserOperator) GetAccount(username string) (models.User, error) {
	data, err := service.db.GetAccountByUsername(username)

	if err != nil {
		return models.User{}, err
	}

	return models.AccountWith(models.AccountDetails(data)), nil
}

func (service *UserOperator) CreateAccount(acDetails models.AccountDetails) error {
	return service.db.InsertNewAccount(acDetails)
}

func (service *UserOperator) AuthenticateUser(username, password string) (bool, error) {
	ac, err := service.GetAccount(username)

	if err != nil {
		return false, err
	}

	return ac.Authenticate(password), nil
}
