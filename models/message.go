package models

type Message struct {
	Id        string `json:"_id"`
	ChannelId string `json:"rid"`
	Text      string `json:"msg"`
	Timestamp string `json:"ts"`
	User      User   `json:"u"`
}
