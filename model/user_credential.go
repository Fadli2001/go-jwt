package model

type UserCredential struct {
	Username string `json:"username"`
	Password string `json:"userPassword"`
	Email    string
}
