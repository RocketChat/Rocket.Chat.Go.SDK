package rest

import (
	"github.com/RocketChat/Rocket.Chat.Go.SDK/common_testing"
	"github.com/RocketChat/Rocket.Chat.Go.SDK/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

// you have to set access-permissions on role "user" to run this test successfully!
func TestRocket_SetPermissions(t *testing.T) {
	client := getAuthenticatedClient(t, common_testing.GetRandomString(), common_testing.GetRandomEmail(), common_testing.GetRandomString())

	request := UpdatePermissionsRequest{
		Permissions: []models.Permission{{ID: "add-user-to-any-p-room", Roles: []string{"admin"}}},
	}
	response, err := client.UpdatePermissions(&request)

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.NotEmpty(t, response.Permissions)
}
