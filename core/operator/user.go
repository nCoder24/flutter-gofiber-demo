package operator

import (
	internalErrs "demo/core/errors"
	"demo/core/models"
	"demo/database"
	dbErrs "demo/database/errors"
	"errors"
	"log"
)

type UserOperator struct {
	db database.UserDB
}

func initUserOperator(db database.DB) UserOperator {
	return UserOperator{db.UserDB}
}

func (op *UserOperator) GetUser(username string) (models.UserData, error) {
	data, err := op.db.FindUserByUsername(username)

	if err == nil {
		return models.UserData(data), nil
	}

	log.Println(err)
	if errors.Is(err, dbErrs.ErrDocumentNotFound) {
		return models.UserData{}, internalErrs.ErrUserDoesNotExists
	}

	return models.UserData{}, internalErrs.ErrCouldNotFetchUsers
}

func (op *UserOperator) AddNewUser(usrData models.UserData) error {
	userExists, err := op.db.CheckUserExists(usrData.Username)

	if err != nil {
		log.Println(err)
		return internalErrs.ErrCouldNotAddUser
	}

	if userExists {
		return internalErrs.ErrUserAlreadyExists
	}

	return op.db.InsertUser(usrData)
}
