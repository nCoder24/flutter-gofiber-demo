package operator

import (
	"demo/database"
)

type Services struct {
	UserService UserOperator
}

func InitOperators(db database.DB) Services {
	return Services{initUserOperator(db)}
}
