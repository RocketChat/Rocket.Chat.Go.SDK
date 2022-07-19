package rest

import (
	"testing"

	"github.com/RocketChat/Rocket.Chat.Go.SDK/common_testing"
	"github.com/RocketChat/Rocket.Chat.Go.SDK/models"
	"github.com/stretchr/testify/assert"
)

func TestRocket_LoginLogout(t *testing.T) {
	client := getAuthenticatedClient(t, common_testing.GetRandomString(), common_testing.GetRandomEmail(), common_testing.GetRandomString())
	_, logoutErr := client.Logout()
	assert.Nil(t, logoutErr)

	// channels, err := client.GetJoinedChannels()
	// assert.Nil(t, channels)
	// assert.NotNil(t, err)
}

func TestRocket_UserInfo(t *testing.T) {
	rocket := getDefaultClient(t)
	user, err := rocket.UserInfo(testUserName)
	assert.Nil(t, err)
	assert.NotNil(t, user)
}

func TestRocket_Me(t *testing.T) {
	rocket := getDefaultClient(t)
	me, err := rocket.Me()
	assert.Nil(t, err)
	assert.NotNil(t, me)
	assert.Equal(t, testUserEmail, me.Emails[0].Address)
	assert.Equal(t, testUserName, me.UserName)
}

func TestRocket_GetAvatar(t *testing.T) {
	rocket := getDefaultClient(t)
	url, err := rocket.GetAvatar(testUserName)
	assert.Nil(t, err)
	assert.NotNil(t, url)
}

func TestRocket_GetUsers(t *testing.T) {
	rocket := getDefaultClient(t)
	users, err := rocket.GetUsers(nil)
	assert.Nil(t, err)
	assert.NotNil(t, users)
	assert.GreaterOrEqual(t, len(users), 1)
}

func TestRocket_GetPreferences(t *testing.T) {
	rocket := getDefaultClient(t)
	pref, err := rocket.GetPreferences()
	assert.Nil(t, err)
	assert.NotNil(t, pref)
}

func TestRocket_SetPreferences(t *testing.T) {
	rocket := getDefaultClient(t)
	preferences := &models.Preferences{User: testUserName}
	updatedPreferences, err := rocket.SetPreferences(preferences)
	assert.Nil(t, err)
	assert.NotNil(t, updatedPreferences)
}

func TestRocket_ListTeams(t *testing.T) {
	rocket := getDefaultClient(t)
	user, err := rocket.UserInfo(testUserName)
	assert.Nil(t, err)
	teams, err := rocket.ListTeams(user)
	assert.Nil(t, err)
	assert.NotNil(t, teams)
}
