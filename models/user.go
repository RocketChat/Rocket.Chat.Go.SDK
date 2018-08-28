package models

// User ...
type User struct {
	ID           string `json:"_id"`
	Name         string `json:"name"`
	UserName     string `json:"username"`
	Status       string `json:"status"`
	Token        string `json:"token"`
	TokenExpires int64  `json:"tokenExpires"`
}
