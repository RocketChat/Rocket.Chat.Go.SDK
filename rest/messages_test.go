package rest

import (
	"testing"

	"github.com/facundomedica/Rocket.Chat.Go.SDK/models"
	"github.com/stretchr/testify/assert"
)

func TestRocket_SendAndReceive(t *testing.T) {
	rocket := getDefaultClient(t)
	general := &models.Channel{ID: "GENERAL", Name: "general"}

	err := rocket.Send(general, "Test")
	assert.Nil(t, err)

	messages, err := rocket.GetMessages(general, &models.Pagination{Count: 10})
	assert.Nil(t, err)

	message := findMessage(messages, testUserName, "Test")
	assert.NotNil(t, message)
}
