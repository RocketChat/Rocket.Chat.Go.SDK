package rest

import (
	"net/http"
	"net/url"

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

	if err = response.OK(); err != nil {
		return nil, err
	}

	return &response.Info, nil
}

type DirectoryResponse struct {
	Status
	models.Directory
}

// GetDirectory
// A method, that searches by users or channels on all users and channels available on server.
// It supports the Offset, Count, and Sort Query Parameters along with Query and Fields Query Parameters.
//
// https://rocket.chat/docs/developer-guides/rest-api/miscellaneous/directory
func (c *Client) GetDirectory(params url.Values) (*models.Directory, error) {
	request, err := http.NewRequest("GET", c.getUrl()+"/api/v1/directory", nil)
	if err != nil {
		return nil, err
	}
	if len(params) > 0 {
		request.URL.RawQuery = params.Encode()
	}

	response := new(DirectoryResponse)
	if err = c.doRequest(request, response); err != nil {
		return nil, err
	}

	if err = response.OK(); err != nil {
		return nil, err
	}

	return &response.Directory, nil
}
