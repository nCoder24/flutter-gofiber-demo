package models

type UserData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (usrData UserData) IsValid() bool {
	return usrData.Username != "" && usrData.Password != ""
}
