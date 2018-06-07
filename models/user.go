package models

type User struct {
	Id           string `json:"_id"`
	UserName     string `json:"username"`
	Token        string `json:"token"`
	TokenExpires int64  `json:"tokenExpires"`
}
