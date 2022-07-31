package rest

import (
	"testing"

	"github.com/RocketChat/Rocket.Chat.Go.SDK/models"
	"github.com/stretchr/testify/assert"
)

// you have to set access-permissions on role "user" to run this test successfully!
func TestRocket_SetPermissions(t *testing.T) {
	rocket := getDefaultClient(t)

	permissions := []models.Permission{{ID: "add-user-to-any-p-room", Roles: []string{"admin"}}}

	returnPermissions, err := rocket.UpdatePermissions(permissions)

	assert.Nil(t, err)
	assert.NotNil(t, returnPermissions)
	assert.NotEmpty(t, returnPermissions)
}
