package models

type User struct {
	Username string `json:"username"`
	password string
}

func UserWith(acDetails UserDetails) User {
	return User{
		Username: acDetails.Username,
		password: acDetails.Password,
	}
}

func (ac User) Authenticate(password string) bool {
	return ac.password == password
}
