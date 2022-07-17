package rest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
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

type CreateUserResponse struct {
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

type UserStatusResponse struct {
	ID               string `json:"_id"`
	ConnectionStatus string `json:"connectionStatus"`
	Message          string `json:"message"`
	Status           string `json:"status"`
	Error            string `json:"error"`
	Success          bool   `json:"success"`
}

type meResponse struct {
	models.Me
	Status
}

type avatarResponse struct {
	Url string `json:"url"`
	Status
}

type usersResponse struct {
	Status
	Users []models.User `json:"users"`
}

func (s UserStatusResponse) OK() error {
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
func (c *Client) CreateUser(req *models.CreateUserRequest) (*CreateUserResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	response := new(CreateUserResponse)
	err = c.Post("users.create", bytes.NewBuffer(body), response)
	return response, err
}

// UpdateUser updates a user's data being logged in with a user that has permission to do so.
//
// https://rocket.chat/docs/developer-guides/rest-api/users/update/
func (c *Client) UpdateUser(req *models.UpdateUserRequest) (*CreateUserResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	response := new(CreateUserResponse)
	err = c.Post("users.update", bytes.NewBuffer(body), response)
	return response, err
}

// SetUserAvatar updates a user's avatar being logged in with a user that has permission to do so.
// Currently only passing an URL is possible.
//
// https://rocket.chat/docs/developer-guides/rest-api/users/setavatar/
func (c *Client) SetUserAvatar(userID, username, avatarURL string) (*Status, error) {
	body := fmt.Sprintf(`{ "userId": "%s","username": "%s","avatarUrl":"%s"}`, userID, username, avatarURL)
	response := new(Status)
	err := c.Post("users.setAvatar", bytes.NewBufferString(body), response)
	return response, err
}

func (c *Client) GetUserStatus(username string) (*UserStatusResponse, error) {
	params := url.Values{
		"username": []string{username},
	}

	response := new(UserStatusResponse)
	err := c.Get("users.getStatus", params, response)
	return response, err
}

// Me returns quick information about the authenticated user.
//
// https://developer.rocket.chat/reference/api/rest-api/endpoints/other-important-endpoints/authentication-endpoints/me
func (c *Client) Me() (*models.Me, error) {
	response := new(meResponse)
	err := c.Get("me", nil, response)
	return &response.Me, err
}

// GetAvatar gets the URL for a user’s avatar.
//
// https://developer.rocket.chat/reference/api/rest-api/endpoints/core-endpoints/users-endpoints/get-avatar
func (c *Client) GetAvatar(username string) (*avatarResponse, error) {
	params := url.Values{
		"username": []string{username},
	}
	response := new(avatarResponse)
	err := c.Get("users.getAvatar", params, response)
	fmt.Println(response)
	return response, err
}

// GetUsers gets all of the users in the system and their information.
// The result is only limited to what the callee has access to view.
// It supports the Offset, Count, and Sort Query Parameters.
//
// https://developer.rocket.chat/reference/api/rest-api/endpoints/core-endpoints/users-endpoints/get-users-list
func (c *Client) GetUsers(page *models.Pagination) ([]models.User, error) {
	params := url.Values{}
	if page != nil {
		params.Add("count", strconv.Itoa(page.Count))
		params.Add("offset", strconv.Itoa(page.Offset))
		params.Add("total", strconv.Itoa(page.Total))
	}
	response := new(usersResponse)
	if err := c.Get("users.list", params, response); err != nil {
		return nil, err
	}
	return response.Users, nil
}
