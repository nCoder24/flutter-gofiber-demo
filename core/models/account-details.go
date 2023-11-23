package models

type AccountDetails struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (details AccountDetails) IsValid() bool {
	return details.Username != "" && details.Password != ""
}
