package rest

import (
	"github.com/RocketChat/Rocket.Chat.Go.SDK/common_testing"
	"github.com/RocketChat/Rocket.Chat.Go.SDK/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRocket_SetPermissions(t *testing.T) {
	rocket := Client{Protocol: common_testing.Protocol, Host: common_testing.Host, Port: common_testing.Port}

	err := rocket.Login(&models.UserCredentials{Email: "chat@catalysts.cc", Password: "test002"})
	if err != nil {
		t.Error(err)
	}

	request := UpdatePermissionsRequest{
		Permissions: []models.Permission{{ID: "add-user-to-any-p-room", Roles: []string{"admin"}}},
	}
	response, err := rocket.UpdatePermissions(&request)

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.NotEmpty(t, response.Permissions)
}
