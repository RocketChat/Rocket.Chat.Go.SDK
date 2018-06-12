// This package provides a RocketChat rest client.
package rest

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

var (
	ResponseErr = fmt.Errorf("got false response")
)

type Client struct {
	Protocol string
	Host     string
	Port     string

	// Use this switch to see all network communication.
	Debug bool

	auth *authInfo
}

type Status struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

type authInfo struct {
	token string
	id    string
}

func (s Status) OK() error {
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

// The base for the most of the json responses
type StatusResponse struct {
	Status
	Channel string `json:"channel"`
}

func NewClient(serverUrl *url.URL, debug bool) *Client {
	protocol := "http"
	port := "80"

	if serverUrl.Scheme == "https" {
		protocol = "https"
		port = "443"
	}

	if len(serverUrl.Port()) > 0 {
		port = serverUrl.Port()
	}

	return &Client{Host: serverUrl.Hostname(), Port: port, Protocol: protocol, Debug: debug}
}

func (c *Client) getUrl() string {
	return fmt.Sprintf("%v://%v:%v", c.Protocol, c.Host, c.Port)
}

func (c *Client) doRequest(request *http.Request, responseBody interface{}) error {

	request.Header.Set("Content-Type", "application/json")
	if c.auth != nil {
		request.Header.Set("X-Auth-Token", c.auth.token)
		request.Header.Set("X-User-Id", c.auth.id)
	}

	if c.Debug {
		log.Println(request)
	}

	response, err := http.DefaultClient.Do(request)

	if err != nil {
		return err
	}

	defer response.Body.Close()
	bodyBytes, err := ioutil.ReadAll(response.Body)

	if c.Debug {
		log.Println(string(bodyBytes))
	}

	if response.StatusCode != http.StatusOK {
		return errors.New("Request error: " + response.Status)
	}

	if err != nil {
		return err
	}

	return json.Unmarshal(bodyBytes, responseBody)
}
