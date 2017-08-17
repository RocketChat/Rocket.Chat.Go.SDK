package realtime

import (
	"crypto/sha256"
	"encoding/hex"

	"github.com/Jeffail/gabs"
	"github.com/RocketChat/Rocket.Chat.Go.SDK/models"
)

type ddpLoginRequest struct {
	User     ddpUser     `json:"user"`
	Password ddpPassword `json:"password"`
}

type ddpTokenLoginRequest struct {
	Token string `json:"resume"`
}

type ddpUser struct {
	Email string `json:"email"`
}

type ddpPassword struct {
	Digest    string `json:"digest"`
	Algorithm string `json:"algorithm"`
}

// RegisterUser a new user on the server. This function does not need a logged in user. The registered user gets logged in
// to set its username.
func (c *Client) RegisterUser(credentials *models.UserCredentials) (*models.User, error) {

	if _, err := c.ddp.Call("registerUser", credentials); err != nil {
		return nil, err
	}

	user, err := c.Login(credentials)
	if err != nil {
		return nil, err
	}

	if _, err := c.ddp.Call("setUsername", credentials.Name); err != nil {
		return nil, err
	}

	return user, nil
}

// Login a user. The password and the email are not allowed to be nil.
//
// https://rocket.chat/docs/developer-guides/realtime-api/method-calls/login/
func (c *Client) Login(credentials *models.UserCredentials) (*models.User, error) {

	digest := sha256.Sum256([]byte(credentials.Password))

	rawResponse, err := c.ddp.Call("login", ddpLoginRequest{
		User:     ddpUser{Email: credentials.Email},
		Password: ddpPassword{Digest: hex.EncodeToString(digest[:]), Algorithm: "sha-256"}})

	return getUserFromData(rawResponse.(map[string]interface{})), err
}

func getUserFromData(data interface{}) *models.User {
	document, _ := gabs.Consume(data)

	return &models.User{
		Id: stringOrZero(document.Path("id").Data()),
	}
}

// SetPresence set user presence
func (c *Client) SetPresence(status string) error {
	_, err := c.ddp.Call("UserPresence:setDefaultStatus", status)
	if err != nil {
		return err
	}

	return nil
}
