package realtime

import (
	"fmt"
	"log"

	"github.com/gopackage/ddp"
)

// Subscribes to stream-notify-logged
// Returns a buffered channel
//
// https://rocket.chat/docs/developer-guides/realtime-api/subscriptions/stream-room-messages/
func (c *Client) Sub(name string, args ...interface{}) (chan string, error) {

	if args == nil {
		log.Println("no args passed")
		if err := c.ddp.Sub(name); err != nil {
			return nil, err
		}
	} else {
		if err := c.ddp.Sub(name, args[0], false); err != nil {
			return nil, err
		}
	}

	msgChannel := make(chan string, default_buffer_size)
	c.ddp.CollectionByName("stream-room-messages").AddUpdateListener(genericExtractor{msgChannel, "update"})

	return msgChannel, nil
}

func (c *Client) SubNotifyUser(userID, event string) (chan map[string]interface{}, error) {

	if err := c.ddp.Sub("stream-notify-user", userID+"/"+event, true); err != nil {
		return nil, err
	}

	msgChannel := make(chan map[string]interface{}, 1000)
	c.ddp.CollectionByName("stream-notify-user").AddUpdateListener(notificationExtractor{msgChannel, "update"})

	return msgChannel, nil
}

type notificationExtractor struct {
	messageChannel chan map[string]interface{}
	operation      string
}

func (u notificationExtractor) CollectionUpdate(collection, operation, id string, doc ddp.Update) {
	if operation == u.operation {
		u.messageChannel <- doc
	}
}

type genericExtractor struct {
	messageChannel chan string
	operation      string
}

func (u genericExtractor) CollectionUpdate(collection, operation, id string, doc ddp.Update) {
	if operation == u.operation {
		u.messageChannel <- fmt.Sprintf("%s -> update", collection)
	}
}
