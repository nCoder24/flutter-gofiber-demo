package models

type Account struct {
	Username string `json:"username"`
	password string
}

func CreateAccount(username string, password string) Account {
	return Account{
		Username: username,
		password: password,
	}
}

func (ac Account) Authenticate(password string) bool {
	return ac.password == password
}
