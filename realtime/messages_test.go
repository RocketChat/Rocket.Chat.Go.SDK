package realtime

import (
	"testing"

	"github.com/RocketChat/Rocket.Chat.Go.SDK/models"
	"github.com/stretchr/testify/assert"
)

func TestClient_SubscribeToMessageStream(t *testing.T) {
	c := getLoggedInClient(t)

	general := models.Channel{ID: "GENERAL"}
	textToSend := "RealtimeTest"

	messageChannel, err := c.SubscribeToMessageStream(&general)

	assert.Nil(t, err, "Function returned error")
	assert.NotNil(t, messageChannel, "Function didn't returned general")

	go func() {
		sendAndAssertNoError(t, c, &general, textToSend)
		// sendAndAssertNoError(t, c, &general, textToSend)
		// sendAndAssertNoError(t, c, &general, textToSend)
	}()

	receivedMessage1 := <-messageChannel
	// receivedMessage2 := <-messageChannel
	// receivedMessage3 := <-messageChannel

	assertMessage(t, receivedMessage1)
	// assertMessage(t, receivedMessage2)
	// assertMessage(t, receivedMessage3)
}

func assertMessage(t *testing.T, message models.Message) {
	assert.NotNil(t, message.ID, "Id was not set")
	assert.Equal(t, "GENERAL", message.RoomID, "Wrong channel id")
	assert.NotNil(t, message.Timestamp, "Timestamp was not set")
	assert.NotNil(t, message.User.ID, "UserId was not set")
	assert.NotNil(t, message.User.UserName, "Username was not set")
}

func sendAndAssertNoError(t *testing.T, c *Client, channel *models.Channel, text string) {
	m, err := c.SendMessage(channel, text)
	assert.Nil(t, err, "Error while sending message")
	assert.NotNil(t, m, "SendMessage should return a Message object")
}

func TestClient_SubscribeToMessageStream_UnknownChannel(t *testing.T) {

	c := getLoggedInClient(t)
	channel := models.Channel{ID: "unknown"}

	_, err := c.SubscribeToMessageStream(&channel)

	assert.NotNil(t, err, "Function didn't return error")
}
