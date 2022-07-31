package models

type User struct {
	ID           string `json:"_id"`
	Name         string `json:"name"`
	UserName     string `json:"username"`
	Status       string `json:"status"`
	Token        string `json:"token"`
	TokenExpires int64  `json:"tokenExpires"`
}

type UserStatus struct {
	Message          string `json:"message"`
	Status           string `json:"status"`
	ConnectionStatus string `json:"connectionStatus"`
}
