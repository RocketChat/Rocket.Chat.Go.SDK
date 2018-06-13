package rest

import (
	"net/url"
	"testing"

	"github.com/RocketChat/Rocket.Chat.Go.SDK/common_testing"
	"github.com/stretchr/testify/assert"
)

func TestRocket_GetServerInfo(t *testing.T) {
	rocket := Client{Protocol: common_testing.Protocol, Host: common_testing.Host, Port: common_testing.Port}

	info, err := rocket.GetServerInfo()

	assert.Nil(t, err)
	assert.NotNil(t, info)
	assert.NotEmpty(t, info.Version)
}

func TestRocket_GetDirectory(t *testing.T) {
	rocket := getDefaultClient(t)

	directory, err := rocket.GetDirectory(url.Values{"query": []string{`{"text": "gene", "type": "channels"}`}})

	assert.Nil(t, err)
	assert.NotNil(t, directory)
}

func TestRocket_GetSpotlight(t *testing.T) {
	rocket := getDefaultClient(t)

	spotlight, err := rocket.GetSpotlight(url.Values{"query": []string{`#foobar`}})

	assert.Nil(t, err)
	assert.NotNil(t, spotlight)
}

func TestRocket_GetStatistics(t *testing.T) {
	rocket := getDefaultClient(t)

	_, err := rocket.GetStatistics()
	assert.NotNil(t, err)

	// TODO admin user
	rocket.auth.id = "4SicoW2wDjcAaRh4M"
	rocket.auth.token = "4nI24JkcHTsqTtOUeJYbbrhum5T_Y6IdjAwyj72qDBu"
	statistics, err := rocket.GetStatistics()
	assert.Nil(t, err)
	assert.NotNil(t, statistics)
}
