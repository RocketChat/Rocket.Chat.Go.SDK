package rest

import (
	"bytes"
	"errors"
	"net/http"
	"net/url"

	"github.com/RocketChat/Rocket.Chat.Go.SDK/models"
)

type logoutResponse struct {
	Status string `json:"status"`
	Data   struct {
		Message string `json:"message"`
	} `json:"data"`
}

type logonResponse struct {
	Status string `json:"status"`
	Data   struct {
		Token  string `json:"authToken"`
		UserId string `json:"userId"`
	} `json:"data"`
}

// Login a user. The Email and the Password are mandatory. The auth token of the user is stored in the Client instance.
//
// https://rocket.chat/docs/developer-guides/rest-api/authentication/login
func (c *Client) Login(credentials *models.UserCredentials) error {
	if c.auth != nil {
		return nil
	}

	if credentials.ID != "" && credentials.Token != "" {
		c.auth = &authInfo{id: credentials.ID, token: credentials.Token}
		return nil
	}

	data := url.Values{"user": {credentials.Email}, "password": {credentials.Password}}
	request, _ := http.NewRequest("POST", c.getUrl()+"/api/v1/login", bytes.NewBufferString(data.Encode()))
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	response := new(logonResponse)

	if err := c.doRequest(request, response); err != nil {
		return err
	}

	if response.Status == "success" {
		c.auth = &authInfo{id: response.Data.UserId, token: response.Data.Token}
		credentials.ID, credentials.Token = response.Data.UserId, response.Data.Token
		return nil
	} else {
		return errors.New("Response status: " + response.Status)
	}
}

// Logout a user. The function returns the response message of the server.
//
// https://rocket.chat/docs/developer-guides/rest-api/authentication/logout
func (c *Client) Logout() (string, error) {

	if c.auth == nil {
		return "Was not logged in", nil
	}

	request, _ := http.NewRequest("POST", c.getUrl()+"/api/v1/logout", nil)

	response := new(logoutResponse)

	if err := c.doRequest(request, response); err != nil {
		return "", err
	}

	if response.Status == "success" {
		return response.Data.Message, nil
	} else {
		return "", errors.New("Response status: " + response.Status)
	}
}
