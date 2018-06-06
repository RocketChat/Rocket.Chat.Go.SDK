// This package provides a RocketChat rest client.
package rest

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Client struct {
	Protocol string
	Host     string
	Port     string

	// Use this switch to see all network communication.
	Debug bool

	auth *authInfo
}

type authInfo struct {
	token string
	id    string
}

// The base for the most of the json responses
type StatusResponse struct {
	Success bool   `json:"success"`
	Channel string `json:"channel"`
}

func NewClient(host, port string, tls, debug bool) *Client {
	var protocol string

	if tls {
		protocol = "https"
	} else {
		protocol = "http"
	}

	return &Client{Host: host, Port: port, Protocol: protocol, Debug: debug}
}

func (c *Client) getUrl() string {
	return fmt.Sprintf("%v://%v:%v", c.Protocol, c.Host, c.Port)
}

func (c *Client) doRequest(request *http.Request, responseBody interface{}) error {

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
