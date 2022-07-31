package rest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"
	"time"

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

type createUserRequest struct {
	Name         string            `json:"name"`
	Email        string            `json:"email"`
	Password     string            `json:"password"`
	Username     string            `json:"username"`
	Roles        []string          `json:"roles,omitempty"`
	CustomFields map[string]string `json:"customFields,omitempty"`
}

type createUserResponse struct {
	Status
	User struct {
		ID        string    `json:"_id"`
		CreatedAt time.Time `json:"createdAt"`
		Services  struct {
			Password struct {
				Bcrypt string `json:"bcrypt"`
			} `json:"password"`
		} `json:"services"`
		Username string `json:"username"`
		Emails   []struct {
			Address  string `json:"address"`
			Verified bool   `json:"verified"`
		} `json:"emails"`
		Type         string            `json:"type"`
		Status       string            `json:"status"`
		Active       bool              `json:"active"`
		Roles        []string          `json:"roles"`
		UpdatedAt    time.Time         `json:"_updatedAt"`
		Name         string            `json:"name"`
		CustomFields map[string]string `json:"customFields"`
	} `json:"user"`
}

type updateUserRequest struct {
	UserID string `json:"userId"`
	Data   struct {
		Name         string            `json:"name"`
		Email        string            `json:"email"`
		Password     string            `json:"password"`
		Username     string            `json:"username"`
		CustomFields map[string]string `json:"customFields,omitempty"`
	} `json:"data"`
}

type userStatusResponse struct {
	models.UserStatus
	ID      string `json:"_id"`
	Error   string `json:"error"`
	Success bool   `json:"success"`
}

func (s userStatusResponse) OK() error {
	if s.Success {
		return nil
	}

	if len(s.Error) > 0 {
		return fmt.Errorf(s.Error)
	}

	if len(s.Message) > 0 {
		return fmt.Errorf("status: %s, message: %s", s.Status, s.Message)
	}

	return ResponseErr
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

	response := new(logonResponse)
	data := url.Values{"user": {credentials.Email}, "password": {credentials.Password}}
	if err := c.PostForm("login", data, response); err != nil {
		return err
	}

	c.auth = &authInfo{id: response.Data.UserID, token: response.Data.Token}
	credentials.ID, credentials.Token = response.Data.UserID, response.Data.Token
	return nil
}

// CreateToken creates an access token for a user
//
// https://rocket.chat/docs/developer-guides/rest-api/users/createtoken/
func (c *Client) CreateToken(userID, username string) (*models.UserCredentials, error) {
	response := new(logonResponse)
	data := url.Values{"userId": {userID}, "username": {username}}
	if err := c.PostForm("users.createToken", data, response); err != nil {
		return nil, err
	}
	credentials := &models.UserCredentials{}
	credentials.ID, credentials.Token = response.Data.UserID, response.Data.Token
	return credentials, nil
}

// Logout a user. The function returns the response message of the server.
//
// https://rocket.chat/docs/developer-guides/rest-api/authentication/logout
func (c *Client) Logout() (string, error) {

	if c.auth == nil {
		return "Was not logged in", nil
	}

	response := new(logoutResponse)
	if err := c.Get("logout", nil, response); err != nil {
		return "", err
	}

	return response.Data.Message, nil
}

// CreateUser being logged in with a user that has permission to do so.
//
// https://rocket.chat/docs/developer-guides/rest-api/users/create
func (c *Client) CreateUser(req *createUserRequest) (*models.User, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	resp := new(createUserResponse)
	if err := c.Post("users.create", bytes.NewBuffer(body), resp); err != nil {
		return nil, err
	}

	createdUser := &models.User{
		ID:       resp.User.ID,
		Name:     resp.User.Name,
		UserName: resp.User.Username,
		Status:   resp.User.Status,
	}

	return createdUser, nil
}

// UpdateUser updates a user's data being logged in with a user that has permission to do so.
//
// https://rocket.chat/docs/developer-guides/rest-api/users/update/
func (c *Client) UpdateUser(req *updateUserRequest) (*models.User, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	resp := new(createUserResponse)
	err = c.Post("users.update", bytes.NewBuffer(body), resp)
	if err != nil {
		return nil, err
	}

	updatedUser := &models.User{
		ID:       resp.User.ID,
		Name:     resp.User.Name,
		UserName: resp.User.Username,
		Status:   resp.User.Status,
	}

	return updatedUser, nil
}

// SetUserAvatar updates a user's avatar being logged in with a user that has permission to do so.
// Currently only passing an URL is possible.
//
// https://rocket.chat/docs/developer-guides/rest-api/users/setavatar/
func (c *Client) SetUserAvatar(userID, username, avatarURL string) error {
	body := fmt.Sprintf(`{ "userId": "%s","username": "%s","avatarUrl":"%s"}`, userID, username, avatarURL)
	response := new(Status)
	return c.Post("users.setAvatar", bytes.NewBufferString(body), response)
}

func (c *Client) GetUserStatus(username string) (*models.UserStatus, error) {
	params := url.Values{
		"username": []string{username},
	}

	response := new(userStatusResponse)
	if err := c.Get("users.getStatus", params, response); err != nil {
		return nil, err
	}

	userStatus := &models.UserStatus{
		Message:          response.Message,
		Status:           response.Status,
		ConnectionStatus: response.ConnectionStatus,
	}

	return userStatus, nil
}
