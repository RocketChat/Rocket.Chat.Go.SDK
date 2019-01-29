package models

import "time"

type Room struct {
	ID        string    `json:"_id"`
	UpdatedAt time.Time `json:"_updatedAt"`
	Type      string    `json:"t"`
	Messages  int       `json:"msgs"`
	Timestamp time.Time `json:"ts"`
	Meta      struct {
		Revision int   `json:"revision"`
		Created  int64 `json:"created"`
		Version  int   `json:"version"`
	} `json:"meta"`
	Loki      int      `json:"$loki"`
	Usernames []string `json:"usernames"`
}
