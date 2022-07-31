package rest

import (
	"bytes"
	"fmt"

	"github.com/RocketChat/Rocket.Chat.Go.SDK/models"
)

type directMessageResponse struct {
	Status
	Room models.Room `json:"room"`
}

// Creates a DirectMessage
//
// https://developer.rocket.chat/api/rest-api/methods/im/create
func (c *Client) CreateDirectMessage(username string) (*models.Room, error) {
	body := fmt.Sprintf(`{ "username": "%s" }`, username)
	resp := new(directMessageResponse)

	if err := c.Post("im.create", bytes.NewBufferString(body), resp); err != nil {
		return nil, err
	}

	return &resp.Room, nil
}
