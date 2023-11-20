package service

import "demo/db"

type Services struct {
	UserService UserService
}

func InitServices(db db.DB) Services {
	return Services{initUserService(db)}
}
