//Package rest provides a RocketChat rest client.
package rest

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

var (
	// ErrResponse for general responice errors
	ErrResponse = fmt.Errorf("got false response")
)

// Response is everythuing is ok ?
type Response interface {
	OK() error
}

// Status ...
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

// OK ...
func (s Status) OK() error {
	if s.Success {
		return nil
	}

	if len(s.Error) > 0 {
		return fmt.Errorf(s.Error)
	}

	if s.Status == "success" {
		return nil
	}

	if len(s.Message) > 0 {
		return fmt.Errorf("status: %s, message: %s", s.Status, s.Message)
	}
	return ErrResponse
}

// StatusResponse The base for the most of the json responses
type StatusResponse struct {
	Status
	Channel string `json:"channel"`
}

func (rest *RestService) getURL() string {
	if len(rest.client.Version) == 0 {
		rest.client.Version = "v1"
	}
	return fmt.Sprintf("%v://%v:%v/api/%s", rest.client.Protocol, rest.client.Host, rest.client.Port, rest.client.Version)
}

// Get call Get
func (rest *RestService) Get(api string, params url.Values, response Response) error {
	return rest.doRequest(http.MethodGet, api, params, nil, response)
}

// Post call as JSON
func (rest *RestService) Post(api string, body io.Reader, response Response) error {
	return rest.doRequest(http.MethodPost, api, nil, body, response)
}

// PostForm call as Form Data
func (rest *RestService) PostForm(api string, params url.Values, response Response) error {
	return rest.doRequest(http.MethodPost, api, params, nil, response)
}

func (rest *RestService) doRequest(method, api string, params url.Values, body io.Reader, response Response) error {
	contentType := "application/x-www-form-urlencoded"
	if method == http.MethodPost {
		if body != nil {
			contentType = "application/json"
		} else if len(params) > 0 {
			body = bytes.NewBufferString(params.Encode())
		}
	}

	request, err := http.NewRequest(method, rest.getURL()+"/"+api, body)
	if err != nil {
		return err
	}

	if method == http.MethodGet {
		if len(params) > 0 {
			request.URL.RawQuery = params.Encode()
		}
	} else {
		request.Header.Set("Content-Type", contentType)
	}

	if rest.client.auth != nil {
		request.Header.Set("X-Auth-Token", rest.client.auth.token)
		request.Header.Set("X-User-Id", rest.client.auth.id)
	}

	if rest.client.Debug {
		log.Println(request)
	}

	resp, err := rest.client.myDoer.Do(request)

	if err != nil {
		return err
	}

	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)

	if rest.client.Debug {
		log.Println(string(bodyBytes))
	}

	var parse bool
	if err == nil {
		if e := json.Unmarshal(bodyBytes, response); e == nil {
			parse = true
		}
	}
	if resp.StatusCode != http.StatusOK {
		if parse {
			return response.OK()
		}
		return errors.New("Request error: " + resp.Status)
	}

	if err != nil {
		return err
	}

	return response.OK()
}
