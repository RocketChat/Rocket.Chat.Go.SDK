package goRocket

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"strconv"

	"github.com/Jeffail/gabs"
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
func (c *RestService) Login(credentials *UserCredentials) error {
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
func (c *LiveService) RegisterUser(credentials *UserCredentials) (*User, error) {

	if _, err := c.client.ddp.Call("registerUser", credentials); err != nil {
		return nil, err
	}

	user, err := c.Login(credentials)
	if err != nil {
		return nil, err
	}

	if _, err := c.client.ddp.Call("setUsername", credentials.Name); err != nil {
		return nil, err
	}

	return user, nil
}

// Login a user.
// token shouldn't be nil, otherwise the password and the email are not allowed to be nil.
//
// https://rocket.chat/docs/developer-guides/realtime-api/method-calls/login/
func (c *LiveService) Login(credentials *UserCredentials) (*User, error) {
	var request interface{}
	if credentials.Token != "" {
		request = ddpTokenLoginRequest{
			Token: credentials.Token,
		}
	} else {
		digest := sha256.Sum256([]byte(credentials.Password))
		request = ddpLoginRequest{
			User: ddpUser{Email: credentials.Email},
			Password: ddpPassword{
				Digest:    hex.EncodeToString(digest[:]),
				Algorithm: "sha-256",
			},
		}
	}

	rawResponse, err := c.client.ddp.Call("login", request)
	if err != nil {
		return nil, err
	}

	user := getUserFromData(rawResponse.(map[string]interface{}))
	if credentials.Token == "" {
		credentials.ID, credentials.Token = user.ID, user.Token
	}

	return user, nil
}

func getUserFromData(data interface{}) *User {
	document, _ := gabs.Consume(data)

	expires, _ := strconv.ParseFloat(stringOrZero(document.Path("tokenExpires.$date").Data()), 64)
	return &User{
		ID:           stringOrZero(document.Path("id").Data()),
		Token:        stringOrZero(document.Path("token").Data()),
		TokenExpires: int64(expires),
	}
}

// SetPresence set user presence
func (c *LiveService) SetPresence(status string) error {
	_, err := c.client.ddp.Call("UserPresence:setDefaultStatus", status)
	if err != nil {
		return err
	}

	return nil
}
