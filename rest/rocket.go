package rest

import (
	"net/http"
	"net/url"
)

// RocketClient is a generitic client object
//
// Only holds key information that is needed to user Rocket.chat's API's
type RocketClient struct {
	Protocol string
	Host     string
	Port     string
	Version  string

	// Use this switch to see all network communication.
	Debug bool

	auth *authInfo

	myDoer Doer
	service

	//Services
	Rest *RestService
}

type service struct {
	client *RocketClient
}

// RestService rest API service
type RestService service

// Doer to make testing easer !
type Doer interface {
	Do(*http.Request) (*http.Response, error)
}

//NewRestClient RocketClient Client
func NewRestClient(serverURL *url.URL, debug bool) (*RocketClient, error) {

	r := RocketClient{}

	protocol := "http"
	port := "80"

	if serverURL.Scheme == "https" {
		protocol = "https"
		port = "443"
	}

	if len(serverURL.Port()) > 0 {
		port = serverURL.Port()
	}

	r.myDoer = http.DefaultClient
	r.Protocol = protocol
	r.Port = port
	r.service.client = &r
	r.Version = "v1"
	r.Debug = debug
	r.Host = serverURL.Hostname()

	r.Rest = (*RestService)(&r.service)

	return &r, nil
}
