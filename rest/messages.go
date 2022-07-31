package rest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html"
	"net/url"
	"strconv"

	"github.com/RocketChat/Rocket.Chat.Go.SDK/models"
)

type messagesResponse struct {
	Status
	Messages []models.Message `json:"messages"`
}

type messageResponse struct {
	Status
	Message models.Message `json:"message"`
}

type deleteMessageResponse struct {
	Status
	Message models.Message
}

// Sends a message to a channel. The name of the channel has to be not nil.
// The message will be html escaped.
//
// https://rocket.chat/docs/developer-guides/rest-api/chat/postmessage
func (c *Client) Send(channel *models.Channel, msg string) error {
	body := fmt.Sprintf(`{ "channel": "%s", "text": "%s"}`, channel.Name, html.EscapeString(msg))
	return c.Post("chat.postMessage", bytes.NewBufferString(body), new(messageResponse))
}

// PostMessage send a message to a channel. The channel or roomId has to be not nil.
// The message will be json encode.
//
// https://rocket.chat/docs/developer-guides/rest-api/chat/postmessage
func (c *Client) PostMessage(msg *models.PostMessage) (*models.Message, error) {
	body, err := json.Marshal(msg)
	if err != nil {
		return nil, err
	}

	response := new(messageResponse)
	err = c.Post("chat.postMessage", bytes.NewBuffer(body), response)
	return &response.Message, err
}

// GetMessage retrieves a single chat message by the provided id.
// Callee must have permission to access the room where the message resides.
//
// https://developer.rocket.chat/reference/api/rest-api/endpoints/core-endpoints/chat-endpoints/getmessage

func (c *Client) GetMessage(msgId string) (models.Message, error) {
	params := url.Values{
		"msgId": []string{msgId},
	}
	response := new(messageResponse)
	if err := c.Get("chat.getMessage", params, response); err != nil {
		return models.Message{}, err
	}
	return response.Message, nil
}

// Get messages from a channel. The channel id has to be not nil. Optionally a
// count can be specified to limit the size of the returned messages.
//
// https://rocket.chat/docs/developer-guides/rest-api/channels/history
func (c *Client) GetMessages(channel *models.Channel, page *models.Pagination) ([]models.Message, error) {
	params := url.Values{
		"roomId": []string{channel.ID},
	}

	if page != nil {
		params.Add("count", strconv.Itoa(page.Count))
		params.Add("offset", strconv.Itoa(page.Offset))
	}

	response := new(messagesResponse)
	if err := c.Get("channels.history", params, response); err != nil {
		return nil, err
	}

	return response.Messages, nil
}

// GetMentionedMessages retrieves mentioned messages.
// It supports the Offset and Count Query Parameters.
//
// https://developer.rocket.chat/reference/api/rest-api/endpoints/core-endpoints/chat-endpoints/getmentionedmessages
func (c *Client) GetMentionedMessages(channel *models.Channel, page *models.Pagination) ([]models.Message, error) {
	params := url.Values{
		"roomId": []string{channel.ID},
	}

	if page != nil {
		params.Add("count", strconv.Itoa(page.Count))
		params.Add("offset", strconv.Itoa(page.Offset))
	}

	response := new(messagesResponse)
	if err := c.Get("chat.getMentionedMessages", params, response); err != nil {
		return nil, err
	}

	return response.Messages, nil
}

// UpdateMessage updates a specific message.
//
// https://developer.rocket.chat/reference/api/rest-api/endpoints/core-endpoints/chat-endpoints/message-update
func (c *Client) UpdateMessage(msg *models.UpdateMessage) (*models.Message, error) {
	body, err := json.Marshal(msg)
	if err != nil {
		return nil, err
	}

	response := new(messageResponse)
	err = c.Post("chat.update", bytes.NewBuffer(body), response)
	return &response.Message, err

}

// DeleteMessage deletes a specific message.
//
// https://developer.rocket.chat/reference/api/rest-api/endpoints/core-endpoints/chat-endpoints/delete
func (c *Client) DeleteMessage(msg *models.DeleteMessage) (*models.Message, error) {
	body, err := json.Marshal(msg)
	if err != nil {
		return nil, err
	}

	response := new(deleteMessageResponse)
	err = c.Post("chat.delete", bytes.NewBuffer(body), response)
	return &response.Message, err
}

// SearchMessages searches for messages in a channel by id and text message
//
// https://developer.rocket.chat/reference/api/rest-api/endpoints/core-endpoints/chat-endpoints/search-message

func (c *Client) SearchMessages(channel *models.Channel, searchText string) ([]models.Message, error) {
	params := url.Values{
		"roomId":     []string{channel.ID},
		"searchText": []string{searchText},
	}
	response := new(messagesResponse)
	if err := c.Get("chat.search", params, response); err != nil {
		return nil, err
	}
	return response.Messages, nil
}
