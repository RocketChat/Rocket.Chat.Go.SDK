package rest

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestRocket_SendAndReceive(t *testing.T) {
	rocket := getDefaultClient(t)

	rooms, err := rocket.GetPublicChannels()
	assert.Nil(t, err)

	general := getChannel(rooms, "general")

	err = rocket.Send(general, "Test")
	assert.Nil(t, err)

	messages, err := rocket.GetMessages(general, &Page{Count: 10})
	assert.Nil(t, err)

	message := findMessage(messages, testUserName, "Test")
	assert.NotNil(t, message)
}