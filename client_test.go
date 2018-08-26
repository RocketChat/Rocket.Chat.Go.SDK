package goRocket

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	testUserName  string
	testUserEmail string
	testPassword  = "test"
	rocketClient  *RocketClient
)

// no lint - deadcode !
func getDefaultClient(t *testing.T) *RocketClient {

	if rocketClient == nil {
		testUserEmail = GetRandomEmail()
		testUserName = GetRandomString()
		rocketClient = getAuthenticatedClient(t, testUserName, testUserEmail, testPassword)
	}

	return rocketClient
}

func getAuthenticatedClient(t *testing.T, name, email, password string) *RocketClient {
	client := RocketClient{Protocol: Protocol, Host: Host, Port: Port}
	credentials := &UserCredentials{Name: name, Email: email, Password: password}

	rtClient, err := NewLiveClient(&url.URL{Host: Host + ":" + Port}, true)
	assert.Nil(t, err)
	_, regErr := rtClient.Live.RegisterUser(credentials)
	assert.Nil(t, regErr)

	loginErr := client.Rest.Login(credentials)
	assert.Nil(t, loginErr)

	return &client
}

// no lint - deadcode !
func findMessage(messages []Message, user string, msg string) *Message {
	var m *Message
	for i := range messages {
		m = &messages[i]
		if m.User.UserName == user && m.Msg == msg {
			return m
		}
	}

	return nil
}

// no lint - deadcode !
func getChannel(channels []Channel, name string) *Channel {
	for _, r := range channels {
		if r.Name == name {
			return &r
		}
	}

	return nil
}
