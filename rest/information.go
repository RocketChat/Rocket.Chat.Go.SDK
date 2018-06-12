package rest

import (
	"net/http"

	"github.com/RocketChat/Rocket.Chat.Go.SDK/models"
)

type InfoResponse struct {
	Status
	Info models.Info `json:"info"`
}

// GetServerInfo
// A simple method, requires no authentication,
// that returns information about the server including version information.
//
// https://rocket.chat/docs/developer-guides/rest-api/miscellaneous/info
func (c *Client) GetServerInfo() (*models.Info, error) {
	request, err := http.NewRequest("GET", c.getUrl()+"/api/v1/info", nil)
	if err != nil {
		return nil, err
	}

	response := new(InfoResponse)
	if err = c.doRequest(request, response); err != nil {
		return nil, err
	}

	if !response.OK() {
		return nil, ResponseErr
	}

	return &response.Info, nil
}
