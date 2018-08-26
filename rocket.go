package goRocket

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"time"

	"github.com/gopackage/ddp"
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

	ddp    *ddp.Client
	myDoer Doer

	service

	//Services
	Rest *RestService
	Live *LiveService
}

type service struct {
	client *RocketClient
}

// RestService rest API service
type RestService service

// LiveService live API service
type LiveService service

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

// NewLiveClient a new instance and connects to the websocket.
func NewLiveClient(serverURL *url.URL, debug bool) (*RocketClient, error) {
	rand.Seed(time.Now().UTC().UnixNano())

	r := RocketClient{}
	wsURL := "ws"
	port := "80"

	if serverURL.Scheme == "https" {
		port = "443"
	}

	if len(serverURL.Port()) > 0 {
		port = serverURL.Port()
	}

	wsURL = fmt.Sprintf("%s://%v:%v/websocket", wsURL, serverURL.Hostname(), port)

	r.Live = (*LiveService)(&r.service)

	log.Println("About to connect to:", wsURL, port, serverURL.Scheme)

	r.ddp = ddp.NewClient(wsURL, serverURL.String())

	if debug {
		r.ddp.SetSocketLogActive(true)
	}

	if err := r.ddp.Connect(); err != nil {
		return nil, err
	}

	return &r, nil
}
