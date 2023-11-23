package models

type Account struct {
	Username string `json:"username"`
	password string
}

func AccountWith(acDetails AccountDetails) Account {
	return Account{
		Username: acDetails.Username,
		password: acDetails.Password,
	}
}

func (ac Account) Authenticate(password string) bool {
	return ac.password == password
}
