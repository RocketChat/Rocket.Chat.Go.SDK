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

func TestRocket_CreateChannel(t *testing.T) {
	rocket := getDefaultClient(t)

	// create channel with same name as testUserName so that channels aren't duplicated
	_, err := rocket.CreateChannel(testUserName, nil, false)
	assert.Nil(t, err)
}

func TestRocket_InviteChannel(t *testing.T) {
	rocket := getDefaultClient(t)

	general := &models.Channel{ID: "GENERAL"}
	user, err := rocket.UserInfo(testUserName)
	assert.Nil(t, err)

	invitedUsers := []*models.User{user}
	_, err = rocket.InviteChannel(general, invitedUsers)
	assert.Nil(t, err)
}

func TestRocket_JoinChannel(t *testing.T) {
	rocket := getDefaultClient(t)
	general := &models.Channel{ID: "GENERAL"}
	_, err := rocket.JoinChannel(general, "")
	assert.Nil(t, err)
}

func TestRocket_GetChannelMembers(t *testing.T) {
	rocket := getDefaultClient(t)
	t.Run("With channel name", func(t *testing.T) {
		general := &models.Channel{Name: "general"}
		members, err := rocket.GetChannelMembers(general)
		assert.Nil(t, err)
		// are new users added to general by default?
		assert.GreaterOrEqual(t, len(members), 1)

	})
	t.Run("With channel ID", func(t *testing.T) {
		general := &models.Channel{ID: "GENERAL"}
		members, err := rocket.GetChannelMembers(general)
		assert.Nil(t, err)
		assert.GreaterOrEqual(t, len(members), 1)
	})

}
