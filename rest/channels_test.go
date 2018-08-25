package rest

import (
	"testing"

	"github.com/RocketChat/Rocket.Chat.Go.SDK/models"
	"github.com/stretchr/testify/assert"
)

func TestRocket_GetPublicChannels(t *testing.T) {
	rocket := getDefaultClient(t)

	resp, err := rocket.GetPublicChannels()
	assert.Nil(t, err)
	assert.NotZero(t, len(resp.Channels))
}

func TestRocket_GetJoinedChannels(t *testing.T) {
	rocket := getDefaultClient(t)

	resp, err := rocket.GetJoinedChannels(nil)
	assert.Nil(t, err)
	assert.NotZero(t, len(resp.Channels))
}

func TestRocket_LeaveChannel(t *testing.T) {
	rocket := getDefaultClient(t)

	general := &models.Channel{ID: "GENERAL"}
	err := rocket.LeaveChannel(general)
	assert.Nil(t, err)
}

func TestRocket_GetChannelInfo(t *testing.T) {
	rocket := getDefaultClient(t)

	general := &models.Channel{ID: "GENERAL"}
	updatedChannelInfo, err := rocket.GetChannelInfo(general)
	assert.Nil(t, err)
	assert.NotNil(t, updatedChannelInfo)

	assert.Equal(t, general.ID, updatedChannelInfo.ID)
	assert.NotEmpty(t, updatedChannelInfo.Name)
	assert.NotEmpty(t, updatedChannelInfo.Type)
	assert.NotEmpty(t, updatedChannelInfo.UpdatedAt)
	assert.NotEmpty(t, updatedChannelInfo.Timestamp)
}
