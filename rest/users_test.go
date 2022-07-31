package rest

import (
	"testing"

	"github.com/RocketChat/Rocket.Chat.Go.SDK/common_testing"
	"github.com/stretchr/testify/assert"
)

func TestRocket_LoginLogout(t *testing.T) {
	client := getAuthenticatedClient(t, common_testing.GetRandomString(), common_testing.GetRandomEmail(), common_testing.GetRandomString())
	_, logoutErr := client.Logout()
	assert.Nil(t, logoutErr)
}

func TestRocket_CreateUser(t *testing.T) {
	rocket := getDefaultClient(t)

	newUser := &createUserRequest{
		Name:     "Test User",
		Email:    "testUser@test.com",
		Password: "testPassword",
		Username: "TestUser",
	}

	createdUser, err := rocket.CreateUser(newUser)
	assert.Nil(t, err)
	assert.NotNil(t, createdUser)
}

func TestRocket_GetUserStatus(t *testing.T) {
	rocket := getDefaultClient(t)

	status, err := rocket.GetUserStatus(testUserName)
	assert.Nil(t, err)
	assert.NotNil(t, status)
	assert.Equal(t, "online", status.Status)
	assert.Equal(t, "online", status.ConnectionStatus)
}
