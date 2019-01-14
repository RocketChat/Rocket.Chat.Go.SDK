package rest

import (
	"net/url"
	"testing"

	"github.com/facundomedica/Rocket.Chat.Go.SDK/common_testing"
	"github.com/facundomedica/Rocket.Chat.Go.SDK/models"
	"github.com/facundomedica/Rocket.Chat.Go.SDK/realtime"
	"github.com/stretchr/testify/assert"
)

var (
	testUserName  string
	testUserEmail string
	testPassword  = "test"
	rocketClient  *Client
)

func getDefaultClient(t *testing.T) *Client {

	if rocketClient == nil {
		testUserEmail = common_testing.GetRandomEmail()
		testUserName = common_testing.GetRandomString()
		rocketClient = getAuthenticatedClient(t, testUserName, testUserEmail, testPassword)
	}

	return rocketClient
}

func getAuthenticatedClient(t *testing.T, name, email, password string) *Client {
	client := Client{Protocol: common_testing.Protocol, Host: common_testing.Host, Port: common_testing.Port}
	credentials := &models.UserCredentials{Name: name, Email: email, Password: password}

	rtClient, err := realtime.NewClient(&url.URL{Host: common_testing.Host + ":" + common_testing.Port}, true)
	assert.Nil(t, err)
	_, regErr := rtClient.RegisterUser(credentials)
	assert.Nil(t, regErr)

	loginErr := client.Login(credentials)
	assert.Nil(t, loginErr)

	return &client
}

func findMessage(messages []models.Message, user string, msg string) *models.Message {
	var m *models.Message
	for i := range messages {
		m = &messages[i]
		if m.User.UserName == user && m.Msg == msg {
			return m
		}
	}

	return nil
}

func getChannel(channels []models.Channel, name string) *models.Channel {
	for _, r := range channels {
		if r.Name == name {
			return &r
		}
	}

	return nil
}
