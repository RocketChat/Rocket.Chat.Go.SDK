package models

import "time"

type Team struct {
	ID        string    `json:"_id"`
	Name      string    `json:"name"`
	Type      int       `json:"type"`
	CreatedAt time.Time `json:"createdAt"`
	CreatedBy struct {
		ID       string `json:"_id"`
		Username string `json:"username"`
	} `json:"createdBy"`
	UpdatedAt time.Time `json:"_updatedAt"`
	RoomID    string    `json:"roomId"`
	IsOwner   bool      `json:"isOwner"`
}
