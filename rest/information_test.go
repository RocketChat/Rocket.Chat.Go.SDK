package rest

import (
	"testing"

	"github.com/RocketChat/Rocket.Chat.Go.SDK/common_testing"
	"github.com/stretchr/testify/assert"
)

func TestRocket_GetServerInfo(t *testing.T) {
	rocket := Client{Protocol: common_testing.Protocol, Host: common_testing.Host, Port: common_testing.Port}

	// API has changed
	// https://github.com/RocketChat/Rocket.Chat/blob/develop/packages/rocketchat-api/server/v1/misc.js#L13
	info, err := rocket.GetServerInfo()

	assert.Nil(t, err)
	assert.NotNil(t, info)

	assert.NotEmpty(t, info.Version)
}
