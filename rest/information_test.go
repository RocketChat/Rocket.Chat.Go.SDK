package rest

import (
	"net/url"
	"testing"

	"github.com/RocketChat/Rocket.Chat.Go.SDK/common_testing"
	"github.com/RocketChat/Rocket.Chat.Go.SDK/models"
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
	rocket := Client{Protocol: common_testing.Protocol, Host: common_testing.Host, Port: common_testing.Port}
	// TODO admin user
	rocket.auth = &authInfo{id: "4SicoW2wDjcAaRh4M", token: "4nI24JkcHTsqTtOUeJYbbrhum5T_Y6IdjAwyj72qDBu"}

	statistics, err := rocket.GetStatistics()
	assert.Nil(t, err)
	assert.NotNil(t, statistics)
}

func TestRocket_GetStatisticsList(t *testing.T) {
	rocket := Client{Protocol: common_testing.Protocol, Host: common_testing.Host, Port: common_testing.Port}
	// TODO admin user
	rocket.auth = &authInfo{id: "4SicoW2wDjcAaRh4M", token: "4nI24JkcHTsqTtOUeJYbbrhum5T_Y6IdjAwyj72qDBu"}

	statistics, err := rocket.GetStatisticsList(url.Values{"query": []string{`{"_id" : "zT26ye8RAM7MaEN7S"}`}})
	assert.Nil(t, err)
	assert.NotNil(t, statistics)
}

func TestRocket_GetSlashCommandsList(t *testing.T) {
	rocket := Client{Protocol: common_testing.Protocol, Host: common_testing.Host, Port: common_testing.Port}
	// TODO admin user
	rocket.auth = &authInfo{id: "4SicoW2wDjcAaRh4M", token: "4nI24JkcHTsqTtOUeJYbbrhum5T_Y6IdjAwyj72qDBu"}
	slashCommands, err := rocket.GetSlashCommandsList(url.Values{"offset": []string{"0"}, "count": []string{"50"}})
	assert.Nil(t, err)
	assert.NotNil(t, slashCommands)
}

func TestRocket_ExecuteSlashCommand(t *testing.T) {
	rocket := Client{Protocol: common_testing.Protocol, Host: common_testing.Host, Port: common_testing.Port}
	// TODO admin user
	rocket.auth = &authInfo{id: "4SicoW2wDjcAaRh4M", token: "4nI24JkcHTsqTtOUeJYbbrhum5T_Y6IdjAwyj72qDBu"}
	general := &models.ChannelSubscription{RoomId: "GENERAL"}
	resp, err := rocket.ExecuteSlashCommand(general, "tableflip", "@all")
	assert.Nil(t, err)
	assert.NotNil(t, resp)
}
