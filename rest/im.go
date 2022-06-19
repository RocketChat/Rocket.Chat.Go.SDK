package rest

import (
	"bytes"
	"fmt"
	"log"
)

type DirectMessageResponse struct {
	Status
	Room Room `json:"room"`
}

type Room struct {
	ID        string   `json:"_id"`
	Rid       string   `json:"rid"`
	Type      string   `json:"t"`
	Usernames []string `json:"usernames"`
}

// CreateDirectMessage creates a direct message session with another user.
//
// https://developer.rocket.chat/reference/api/rest-api/endpoints/team-collaboration-endpoints/im-endpoints/create
func (c *Client) CreateDirectMessage(username string) (*Room, error) {
	body := fmt.Sprintf(`{ "username": "%s" }`, username)
	resp := new(DirectMessageResponse)

	if err := c.Post("im.create", bytes.NewBufferString(body), resp); err != nil {
		return nil, err
	}

	log.Println(resp)

	return &resp.Room, nil
}
