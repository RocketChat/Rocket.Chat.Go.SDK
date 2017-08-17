package models

type UserCredentials struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"pass"`
}
