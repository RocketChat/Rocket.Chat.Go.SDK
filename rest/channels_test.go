package rest

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRocket_GetPublicChannels(t *testing.T) {
	rocket := getDefaultClient(t)

	channels, err := rocket.GetPublicChannels()
	assert.Nil(t, err)

	assert.Len(t, channels, 1)
	assert.Equal(t, "general", channels[0].Name)
}

func TestRocket_GetJoinedChannels(t *testing.T) {
	rocket := getDefaultClient(t)

	channels, err := rocket.GetPublicChannels()
	assert.Nil(t, err)

	general := getChannel(channels, "general")
	err = rocket.JoinChannel(general)
	assert.Nil(t, err)

	channels, err = rocket.GetJoinedChannels()
	assert.Nil(t, err)

	assert.Len(t, channels, 1)
	assert.Equal(t, "general", channels[0].Name)
}

func TestRocket_LeaveChannel(t *testing.T) {
	rocket := getDefaultClient(t)

	rooms, err := rocket.GetPublicChannels()
	assert.Nil(t, err)

	general := getChannel(rooms, "general")
	err = rocket.JoinChannel(general)
	assert.Nil(t, err)

	err = rocket.LeaveChannel(general)
	assert.Nil(t, err)
}

func TestRocket_GetChannelInfo(t *testing.T) {
	rocket := getDefaultClient(t)

	rooms, err := rocket.GetPublicChannels()
	assert.Nil(t, err)

	general := getChannel(rooms, "general")
	err = rocket.JoinChannel(general)
	assert.Nil(t, err)

	updatedChannelInfo, err := rocket.GetChannelInfo(general)
	assert.Nil(t, err)
	assert.NotNil(t, updatedChannelInfo)

	assert.Equal(t, general.Id, updatedChannelInfo.Id)
	assert.NotEmpty(t, updatedChannelInfo.Name)
	assert.NotEmpty(t, updatedChannelInfo.Type)
	assert.NotEmpty(t, updatedChannelInfo.UpdatedAt)
	assert.NotEmpty(t, updatedChannelInfo.Timestamp)
	// API has changed
	// Please use the `channel/dm/group.members` endpoint. This is disabled for performance reasons
	// https://github.com/RocketChat/Rocket.Chat/blob/develop/packages/rocketchat-api/server/v1/channels.js#L424
	// https://github.com/RocketChat/Rocket.Chat/blob/develop/packages/rocketchat-api/server/api.js#L15
	// assert.NotZero(t, len(updatedChannelInfo.UserNames))
}
