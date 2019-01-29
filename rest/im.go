package rest

import (
	"bytes"
	"encoding/json"

	"github.com/RocketChat/Rocket.Chat.Go.SDK/models"
)

type CreateDirectMessageRequest struct {
	Username string `json:"username"`
}

type CreateDirectMessageResponse struct {
	Status
	Room    models.Room `json:"room"`
	Success bool        `json:"success"`
}

// CreateDirectMessage (IM Create) Create a direct message session with another user.
//
// https://rocket.chat/docs/developer-guides/rest-api/im/create
func (c *Client) CreateDirectMessage(username string) (*CreateDirectMessageResponse, error) {
	body, err := json.Marshal(map[string]string{"username": username})
	if err != nil {
		return nil, err
	}

	response := new(CreateDirectMessageResponse)
	err = c.Post("im.create", bytes.NewBuffer(body), response)
	return response, err
}
