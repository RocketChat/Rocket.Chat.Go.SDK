package rest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html"
	"net/http"

	"github.com/RocketChat/Rocket.Chat.Go.SDK/models"
)

type MessagesResponse struct {
	StatusResponse
	Messages []models.Message `json:"messages"`
}

type MessageResponse struct {
	StatusResponse
	Message models.Message `json:"message"`
}

type Page struct {
	Count int
}

// Sends a message to a channel. The name of the channel has to be not nil.
// The message will be html escaped.
//
// https://rocket.chat/docs/developer-guides/rest-api/chat/postmessage
func (c *Client) Send(channel *models.Channel, msg string) error {
	body := fmt.Sprintf(`{ "channel": "%s", "text": "%s"}`, channel.Name, html.EscapeString(msg))
	request, _ := http.NewRequest("POST", c.getUrl()+"/api/v1/chat.postMessage", bytes.NewBufferString(body))

	response := new(MessageResponse)

	return c.doRequest(request, response)
}

// PostMessage send a message to a channel. The channel or roomId has to be not nil.
// The message will be json encode.
//
// https://rocket.chat/docs/developer-guides/rest-api/chat/postmessage
func (c *Client) PostMessage(msg *models.PostMessage) (*MessageResponse, error) {
	body, err := json.Marshal(msg)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", c.getUrl()+"/api/v1/chat.postMessage", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	response := new(MessageResponse)
	err = c.doRequest(request, response)
	return response, err
}

// Get messages from a channel. The channel id has to be not nil. Optionally a
// count can be specified to limit the size of the returned messages.
//
// https://rocket.chat/docs/developer-guides/rest-api/channels/history
func (c *Client) GetMessages(channel *models.Channel, page *Page) ([]models.Message, error) {
	u := fmt.Sprintf("%s/api/v1/channels.history?roomId=%s", c.getUrl(), channel.Id)

	if page != nil {
		u = fmt.Sprintf("%s&count=%d", u, page.Count)
	}

	request, _ := http.NewRequest("GET", u, nil)
	response := new(MessagesResponse)

	if err := c.doRequest(request, response); err != nil {
		return nil, err
	}

	return response.Messages, nil
}
