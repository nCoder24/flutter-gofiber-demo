package service

import (
	"demo/database"
)

type Services struct {
	UserService UserService
}

func InitServices(db database.DB) Services {
	return Services{initUserService(db)}
}
