package services

type Services struct {
	UserService UserService
}

func InitServices(user UserService) Services {
	return Services{user}
}
