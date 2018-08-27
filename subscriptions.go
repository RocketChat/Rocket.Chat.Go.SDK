package goRocket

import (
	"fmt"
	"log"

	"github.com/gopackage/ddp"
)

// Sub subscribes to stream-notify-logged
// Returns a buffered channel
//
// https://rocket.chat/docs/developer-guides/realtime-api/subscriptions/stream-room-messages/
func (c *LiveService) Sub(name string, args ...interface{}) (chan string, error) {

	if args == nil {
		log.Println("no args passed")
		if err := c.client.ddp.Sub(name); err != nil {
			return nil, err
		}
	} else {
		if err := c.client.ddp.Sub(name, args[0], false); err != nil {
			return nil, err
		}
	}

	msgChannel := make(chan string, defaultBufferSize)
	c.client.ddp.CollectionByName("stream-room-messages").AddUpdateListener(genericExtractor{msgChannel, "update"})

	return msgChannel, nil
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
