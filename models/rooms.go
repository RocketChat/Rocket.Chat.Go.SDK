package models

type RoomInfo struct {
	RoomParticipant bool   `json:"roomParticipant", mapstructure:"roomParticipant"`
	RoomType        string `json:"roomType", mapstructure:"roomType"`
}
