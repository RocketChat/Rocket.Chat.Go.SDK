package rest

import (
	"bytes"
	"encoding/json"

	"github.com/RocketChat/Rocket.Chat.Go.SDK/models"
)

type logoutResponse struct {
	Status
	Data struct {
		Message string `json:"message"`
	} `json:"data"`
}

type logonResponse struct {
	Status
	Data struct {
		Token  string `json:"authToken"`
		UserID string `json:"userID"`
	} `json:"data"`
}

// Login a user. The Email and the Password are mandatory. The auth token of the user is stored in the Client instance.
//
// https://rocket.chat/docs/developer-guides/rest-api/authentication/login
func (c *RestService) Login(credentials *models.UserCredentials) error {
	if c.client.auth != nil {
		return nil
	}

	if credentials.ID != "" && credentials.Token != "" {
		c.client.auth = &authInfo{id: credentials.ID, token: credentials.Token}
		return nil
	}

	response := new(logonResponse)

	type userlogin struct {
		UserName string `json:"username"`
		Password string `json:"password"`
	}

	data, _ := json.Marshal(userlogin{UserName: credentials.Email, Password: credentials.Password})

	if err := c.Post("login", bytes.NewBuffer(data), response); err != nil {
		return err
	}

	c.client.auth = &authInfo{id: response.Data.UserID, token: response.Data.Token}
	credentials.ID, credentials.Token = response.Data.UserID, response.Data.Token
	return nil
}

// Logout a user. The function returns the response message of the server.
//
// https://rocket.chat/docs/developer-guides/rest-api/authentication/logout
func (c *RestService) Logout() (string, error) {

	if c.client.auth == nil {
		return "Was not logged in", nil
	}

	response := new(logoutResponse)
	if err := c.Get("logout", nil, response); err != nil {
		return "", err
	}

	return response.Data.Message, nil
}
