package models

type UserDetails struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (details UserDetails) IsValid() bool {
	return details.Username != "" && details.Password != ""
}
