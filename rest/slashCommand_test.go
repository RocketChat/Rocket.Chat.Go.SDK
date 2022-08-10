package rest

import (
	"net/url"
	"testing"

	"github.com/RocketChat/Rocket.Chat.Go.SDK/common_testing"
	"github.com/RocketChat/Rocket.Chat.Go.SDK/models"
	"github.com/stretchr/testify/assert"
)

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
