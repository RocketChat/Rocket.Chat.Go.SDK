package rest

import (
	"bytes"
	"fmt"
	"log"
)

type directMessageResponse struct {
	Status
	Room Room `json:"room"`
}

type Room struct {
	ID        string   `json:"_id"`
	Rid       string   `json:"rid"`
	Type      string   `json:"t"`
	Usernames []string `json:"usernames"`
}

// Creates a DirectMessage
//
// https://developer.rocket.chat/api/rest-api/methods/im/create
func (c *Client) CreateDirectMessage(username string) (*Room, error) {
	body := fmt.Sprintf(`{ "username": "%s" }`, username)
	resp := new(directMessageResponse)

	if err := c.Post("im.create", bytes.NewBufferString(body), resp); err != nil {
		return nil, err
	}

	log.Println(resp)

	return &resp.Room, nil
}
