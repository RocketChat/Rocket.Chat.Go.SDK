package goRocket

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/url"
	"strconv"
	"time"

	"github.com/Jeffail/gabs"
	"github.com/gopackage/ddp"
)

const (
	// RocketChat doesn't send the `added` event for new messages by default, only `changed`.
	sendAddedEvent    = true
	defaultBufferSize = 100
)

// MessagesResponse hold messages
type MessagesResponse struct {
	Status
	Messages []Message `json:"messages"`
}

// MessageResponse a message
type MessageResponse struct {
	Status
	Message Message `json:"message"`
}

// Send a message to a channel. The name of the channel has to be not nil.
// The message will be html escaped.
//
// https://rocket.chat/docs/developer-guides/rest-api/chat/postmessage
func (c *RestService) Send(channel *Channel, msg string) error {
	body := fmt.Sprintf(`{ "channel": "%s", "text": "%s"}`, channel.Name, html.EscapeString(msg))
	return c.Post("chat.postMessage", bytes.NewBufferString(body), new(MessageResponse))
}

// PostMessage send a message to a channel. The channel or roomId has to be not nil.
// The message will be json encode.
//
// https://rocket.chat/docs/developer-guides/rest-api/chat/postmessage
func (c *RestService) PostMessage(msg *PostMessage) (*MessageResponse, error) {
	body, err := json.Marshal(msg)
	if err != nil {
		return nil, err
	}

	response := new(MessageResponse)
	err = c.Post("chat.postMessage", bytes.NewBuffer(body), response)
	return response, err
}

// GetMessages from a channel. The channel id has to be not nil. Optionally a
// count can be specified to limit the size of the returned messages.
//
// https://rocket.chat/docs/developer-guides/rest-api/channels/history
func (c *RestService) GetMessages(channel *Channel, page *Pagination) ([]Message, error) {
	params := url.Values{
		"roomId": []string{channel.ID},
	}

	if page != nil {
		params.Add("count", strconv.Itoa(page.Count))
	}

	response := new(MessagesResponse)
	if err := c.Get("channels.history", params, response); err != nil {
		return nil, err
	}

	return response.Messages, nil
}

// LoadHistory loads history
// Takes roomID
//
// https://rocket.chat/docs/developer-guides/realtime-api/method-calls/load-history
func (c *LiveService) LoadHistory(roomID string) error {
	_, err := c.client.ddp.Call("loadHistory", roomID)
	if err != nil {
		return err
	}

	return nil
}

// SendMessage sends message to channel
// takes channel and message
//
// https://rocket.chat/docs/developer-guides/realtime-api/method-calls/send-message
func (c *LiveService) SendMessage(channel *Channel, text string) (*Message, error) {
	m := Message{
		ID:     c.newRandomID(),
		RoomID: channel.ID,
		Msg:    text,
	}

	rawResponse, err := c.client.ddp.Call("sendMessage", m)
	if err != nil {
		return nil, err
	}

	return getMessageFromData(rawResponse.(map[string]interface{})), nil
}

// EditMessage edits a message
// takes message object
//
// https://rocket.chat/docs/developer-guides/realtime-api/method-calls/update-message
func (c *LiveService) EditMessage(message *Message) error {
	_, err := c.client.ddp.Call("updateMessage", message)
	if err != nil {
		return err
	}

	return nil
}

// DeleteMessage deletes a message
// takes a message object
//
// https://rocket.chat/docs/developer-guides/realtime-api/method-calls/delete-message
func (c *LiveService) DeleteMessage(message *Message) error {
	_, err := c.client.ddp.Call("deleteMessage", map[string]string{
		"_id": message.ID,
	})
	if err != nil {
		return err
	}

	return nil
}

// ReactToMessage adds a reaction to a message
// takes a message and emoji
//
// https://rocket.chat/docs/developer-guides/realtime-api/method-calls/set-reaction
func (c *LiveService) ReactToMessage(message *Message, reaction string) error {
	_, err := c.client.ddp.Call("setReaction", reaction, message.ID)
	if err != nil {
		return err
	}

	return nil
}

// StarMessage stars message
// takes a message object
//
// https://rocket.chat/docs/developer-guides/realtime-api/method-calls/star-message
func (c *LiveService) StarMessage(message *Message) error {
	_, err := c.client.ddp.Call("starMessage", map[string]interface{}{
		"_id":     message.ID,
		"rid":     message.RoomID,
		"starred": true,
	})

	if err != nil {
		return err
	}

	return nil
}

// UnStarMessage unstars message
// takes message object
//
// https://rocket.chat/docs/developer-guides/realtime-api/method-calls/star-message
func (c *LiveService) UnStarMessage(message *Message) error {
	_, err := c.client.ddp.Call("starMessage", map[string]interface{}{
		"_id":     message.ID,
		"rid":     message.RoomID,
		"starred": false,
	})

	if err != nil {
		return err
	}

	return nil
}

// PinMessage pins a message
// takes a message object
//
// https://rocket.chat/docs/developer-guides/realtime-api/method-calls/pin-message
func (c *LiveService) PinMessage(message *Message) error {
	_, err := c.client.ddp.Call("pinMessage", message)

	if err != nil {
		return err
	}

	return nil
}

// UnPinMessage unpins message
// takes a message object
//
// https://rocket.chat/docs/developer-guides/realtime-api/method-calls/unpin-messages
func (c *LiveService) UnPinMessage(message *Message) error {
	_, err := c.client.ddp.Call("unpinMessage", message)

	if err != nil {
		return err
	}

	return nil
}

// SubscribeToMessageStream Subscribes to the message updates of a channel
// Returns a buffered channel
//
// https://rocket.chat/docs/developer-guides/realtime-api/subscriptions/stream-room-messages/
func (c *LiveService) SubscribeToMessageStream(channel *Channel, msgChannel chan Message) error {

	if err := c.client.ddp.Sub("stream-room-messages", channel.ID, sendAddedEvent); err != nil {
		return err
	}

	//msgChannel := make(chan Message, default_buffer_size)
	c.client.ddp.CollectionByName("stream-room-messages").AddUpdateListener(messageExtractor{msgChannel, "update"})

	return nil
}

func getMessagesFromUpdateEvent(update ddp.Update) []Message {
	document, _ := gabs.Consume(update["args"])
	args, err := document.Children()

	if err != nil {
		log.Printf("Event arguments are in an unexpected format: %v", err)
		return make([]Message, 0)
	}

	messages := make([]Message, len(args))

	for i, arg := range args {
		messages[i] = *getMessageFromDocument(arg)
	}

	return messages
}

func getMessageFromData(data interface{}) *Message {
	// TODO: We should know what this will look like, we shouldn't need to use gabs
	document, _ := gabs.Consume(data)
	return getMessageFromDocument(document)
}

func getMessageFromDocument(arg *gabs.Container) *Message {
	var ts *time.Time
	date := stringOrZero(arg.Path("ts.$date").Data())
	if len(date) > 0 {
		if ti, err := strconv.ParseFloat(date, 64); err == nil {
			t := time.Unix(int64(ti)/1e3, int64(ti)%1e3)
			ts = &t
		}
	}
	return &Message{
		ID:        stringOrZero(arg.Path("_id").Data()),
		RoomID:    stringOrZero(arg.Path("rid").Data()),
		Msg:       stringOrZero(arg.Path("msg").Data()),
		Timestamp: ts,
		User: &User{
			ID:       stringOrZero(arg.Path("u._id").Data()),
			UserName: stringOrZero(arg.Path("u.username").Data()),
		},
	}
}

func stringOrZero(i interface{}) string {
	if i == nil {
		return ""
	}

	switch i.(type) {
	case string:
		return i.(string)
	case float64:
		return fmt.Sprintf("%f", i.(float64))
	default:
		return ""
	}
}

type messageExtractor struct {
	messageChannel chan Message
	operation      string
}

func (u messageExtractor) CollectionUpdate(collection, operation, id string, doc ddp.Update) {
	if operation == u.operation {
		msgs := getMessagesFromUpdateEvent(doc)
		for _, m := range msgs {
			u.messageChannel <- m
		}
	}
}
