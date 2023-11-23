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

func (service *UserOperator) GetUser(username string) (models.User, error) {
	data, err := service.db.GetUserByUsername(username)

	if err != nil {
		return models.User{}, err
	}

	return models.UserWith(models.UserDetails(data)), nil
}

func (service *UserOperator) CreateUser(acDetails models.UserDetails) error {
	return service.db.InsertNewUser(acDetails)
}

func (service *UserOperator) AuthenticateUser(username, password string) (bool, error) {
	ac, err := service.GetUser(username)

	if err != nil {
		return false, err
	}

	return ac.Authenticate(password), nil
}
