package goRocket

import (
	"bytes"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"time"
)

// nolint
type wonkyReader struct{}

func (wr wonkyReader) Read(p []byte) (n int, err error) {
	return 0, nil
}

type testDoer struct {
	response     string
	responseCode int
	http.Header
}

func (nd testDoer) Do(*http.Request) (*http.Response, error) {
	return &http.Response{
		Body:       ioutil.NopCloser(bytes.NewReader([]byte(nd.response))),
		StatusCode: nd.responseCode,
		Header:     nd.Header,
	}, nil
}

func CreateTestRestClient(d Doer) *RocketClient {
	rockerServer := &url.URL{Host: "rocketchat.localhost", Scheme: "https"}
	client, _ := NewRestClient(rockerServer, false)
	client.myDoer = d
	return client
}

const (
	chars    = "abcdefghijklmnopqrstuvwxyz0123456789" // nolint
	Protocol = "http"                                 // nolint
	Host     = "localhost"                            // nolint
	Port     = "3000"                                 // nolint
)

// nolint
func GetRandomString() string {
	length := 6
	rand.Seed(time.Now().UTC().UnixNano())
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result)
}

// nolint
func GetRandomEmail() string {
	return GetRandomString() + "@localhost.com"
}
