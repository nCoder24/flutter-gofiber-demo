package operator

import (
	internalErrs "demo/core/errors"
	"demo/core/models"
	"demo/database"
	dbErs "demo/database/errors"
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

	if err != nil {
		log.Println(err)

		if err == dbErs.ErrDocumentNotFound {
			return models.UserData{}, internalErrs.ErrUserDoesNotExists
		}

		return models.UserData{}, internalErrs.ErrCouldNotGetUser
	}

	return models.UserData(data), nil
}

func (op *UserOperator) AddNewUser(usrData models.UserData) error {
	return op.db.InsertUser(usrData)
}
