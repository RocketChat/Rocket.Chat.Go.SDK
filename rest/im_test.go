package rest

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRocket_CreateDirectMessage(t *testing.T) {
	rocket := getDefaultClient(t)

	room, err := rocket.CreateDirectMessage(testUserName)
	assert.Nil(t, err)
	assert.NotNil(t, room)
}
