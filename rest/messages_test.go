package rest

import (
	"fmt"
	"testing"

	"github.com/RocketChat/Rocket.Chat.Go.SDK/models"
	"github.com/stretchr/testify/assert"
)

func TestRocket_PostMessage(t *testing.T) {
	rocket := getDefaultClient(t)
	postMessage := &models.PostMessage{
		Channel: "general",
		Text:    "TestRocket_PostMessage",
	}
	resp, err := rocket.PostMessage(postMessage)
	assert.Nil(t, err)
	assert.NotNil(t, resp)
}

func TestRocket_GetMessage(t *testing.T) {
	rocket := getDefaultClient(t)
	text := "TestRocket_GetMessage"
	postMessage := &models.PostMessage{
		Channel: "general",
		Text:    text,
	}
	postResp, err := rocket.PostMessage(postMessage)
	assert.Nil(t, err)
	assert.NotNil(t, postResp)

	msgId := postResp.Message.ID
	msg, err := rocket.GetMessage(msgId)
	assert.Nil(t, err)
	assert.Equal(t, text, msg.Msg)
}

func TestRocket_GetMessages(t *testing.T) {
	rocket := getDefaultClient(t)
	texts := [2]string{"TestRocket_GetMessages_1", "TestRocket_GetMessages_2"}
	for _, text := range texts {
		postMessage := &models.PostMessage{
			Channel: "general",
			Text:    text,
		}
		postResp, err := rocket.PostMessage(postMessage)
		assert.Nil(t, err)
		assert.NotNil(t, postResp)
	}

	general := &models.Channel{ID: "GENERAL", Name: "general"}
	messages, err := rocket.GetMessages(general, nil)
	assert.Nil(t, err)

	for _, text := range texts {
		found := findMessage(messages, testUserName, text)
		assert.NotNil(t, found)
	}
}

func TestRocket_GetMentionedMessages(t *testing.T) {
	rocket := getDefaultClient(t)
	text := fmt.Sprint("TestRocket_GetMentionedMessages @", testUserName)
	postMessage := &models.PostMessage{
		Channel: "general",
		Text:    text,
	}
	postResp, err := rocket.PostMessage(postMessage)
	assert.Nil(t, err)
	assert.NotNil(t, postResp)

	general := &models.Channel{ID: "GENERAL", Name: "general"}
	mentionedMessages, err := rocket.GetMentionedMessages(general, nil)
	assert.Nil(t, err)
	found := findMessage(mentionedMessages, testUserName, text)
	assert.NotNil(t, found)
}

func TestRocket_UpdateMessage(t *testing.T) {
	rocket := getDefaultClient(t)
	textOriginal := "TestRocket_UpdateMessageOriginal"
	textUpdated := "TestRocket_UpdateMessageUpdated"
	postMessage := &models.PostMessage{
		Channel: "general",
		Text:    textOriginal,
	}
	postResp, err := rocket.PostMessage(postMessage)
	assert.Nil(t, err)
	assert.NotNil(t, postResp)

	roomId := postResp.Message.RoomID
	msgId := postResp.Message.ID
	updateMessage := &models.UpdateMessage{
		RoomID: roomId,
		MsgID:  msgId,
		Text:   textUpdated,
	}
	resp, err := rocket.UpdateMessage(updateMessage)
	assert.Nil(t, err)
	assert.Equal(t, textUpdated, resp.Message.Msg)
}

func TestRocket_DeleteMessage(t *testing.T) {
	rocket := getDefaultClient(t)
	t.Run("asUser = true", func(t *testing.T) {
		text := "TestRocket_DeleteMessageAsUser"
		postMessage := &models.PostMessage{
			Channel: "general",
			Text:    text,
		}
		postResp, err := rocket.PostMessage(postMessage)
		assert.Nil(t, err)
		assert.NotNil(t, postResp)

		roomId := postResp.Message.RoomID
		msgId := postResp.Message.ID
		deleteMessage := &models.DeleteMessage{
			RoomID: roomId,
			MsgID:  msgId,
			AsUser: true,
		}
		resp, err := rocket.DeleteMessage(deleteMessage)
		assert.Nil(t, err)
		assert.NotNil(t, resp)
	})
	t.Run("asUser = false", func(t *testing.T) {
		text := "TestRocket_DeleteMessageNotAsUser"
		postMessage := &models.PostMessage{
			Channel: "general",
			Text:    text,
		}
		postResp, err := rocket.PostMessage(postMessage)
		assert.Nil(t, err)
		assert.NotNil(t, postResp)

		roomId := postResp.Message.RoomID
		msgId := postResp.Message.ID
		deleteMessage := &models.DeleteMessage{
			RoomID: roomId,
			MsgID:  msgId,
			AsUser: false,
		}
		resp, err := rocket.DeleteMessage(deleteMessage)
		assert.Nil(t, err)
		assert.NotNil(t, resp)
	})
}

func TestRocket_SearchMessages(t *testing.T) {
	rocket := getDefaultClient(t)
	text := "TestRocket_SearchMessages"
	postMessage := &models.PostMessage{
		Channel: "general",
		Text:    text,
	}
	postResp, err := rocket.PostMessage(postMessage)
	assert.Nil(t, err)
	assert.NotNil(t, postResp)

	general := &models.Channel{ID: "GENERAL", Name: "general"}
	searchMessages, err := rocket.SearchMessages(general, text)
	assert.Nil(t, err)
	found := findMessage(searchMessages, testUserName, text)
	assert.NotNil(t, found)
}

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
