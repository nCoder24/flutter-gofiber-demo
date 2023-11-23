package errors

import "errors"

var (
	ErrUserDoesNotExists = errors.New("user does not exists")
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrCouldNotAddUser   = errors.New("could not add user")
	ErrCouldNotGetUser   = errors.New("could not get user")
)
