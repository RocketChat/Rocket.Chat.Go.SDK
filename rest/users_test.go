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
}

func TestRocket_CreateUser(t *testing.T) {
	client := getDefaultClient(t)

	newUser := &models.CreateUserRequest{
		Name:     "Test User",
		Email:    "testUser@test.com",
		Password: "testPassword",
		Username: "TestUser",
	}

	createdUser, err := client.CreateUser(newUser)
	assert.Nil(t, err)
	assert.NotNil(t, createdUser)
}
