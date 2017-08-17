package models

type Permission struct {
	Id        string   `json:"_id"`
	UpdatedAt string   `json:"_updatedAt.$date"`
	Roles     []string `json:"roles"`
}
