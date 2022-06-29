package rest

import (
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

func TestRocket_GetMentionedMessages(t *testing.T) {
	rocket := getDefaultClient(t)
	text := "TestRocket_GetMentionedMessages @all"
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
	assert.Equal(t, 1, len(mentionedMessages))

	// delete mentioned messages so tests are idempotent
	msgId := mentionedMessages[0].ID
	roomId := mentionedMessages[0].RoomID
	deleteMessage := &models.DeleteMessage{MsgID: msgId, RoomID: roomId, AsUser: true}
	msg, err := rocket.DeleteMessage(deleteMessage)
	assert.Nil(t, err)
	assert.NotNil(t, msg)
}

func TestRocket_UpdateUser(t *testing.T) {
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
	assert.NotNil(t, resp)
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
	assert.Equal(t, 1, len(searchMessages))
	assert.Equal(t, text, searchMessages[0].Msg)

	// delete mentioned messages so tests are idempotent
	msgId := searchMessages[0].ID
	roomId := searchMessages[0].RoomID
	deleteMessage := &models.DeleteMessage{MsgID: msgId, RoomID: roomId, AsUser: true}
	msg, err := rocket.DeleteMessage(deleteMessage)
	assert.Nil(t, err)
	assert.NotNil(t, msg)
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
