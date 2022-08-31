package main

import (
	"fmt"
	"log"
	"net/url"

	"github.com/RocketChat/Rocket.Chat.Go.SDK/models"
	"github.com/RocketChat/Rocket.Chat.Go.SDK/rest"
)

// Create a rocket chat client, login, and send a message to all public channels.
func main() {
	serverURL := url.URL{
		Scheme: "http",
		Host:   "localhost:3000",
	}
	authInfo := models.UserCredentials{
		ID:    "example_id",
		Token: "example_token",
	}
	debug := false

	rc_client := rest.NewClient(&serverURL, debug)

	if err := rc_client.Login(&authInfo); err != nil {
		log.Fatal(err)
	}

	channelResp, err := rc_client.GetPublicChannels()
	if err != nil {
		log.Fatal(err)
	}
	channels := channelResp.Channels

	for _, channel := range channels {
		greeting := fmt.Sprintf("Hello %s", channel.Name)
		message := models.PostMessage{
			Channel: channel.Name,
			Text:    greeting,
		}
		if _, err := rc_client.PostMessage(&message); err != nil {
			log.Fatal(err)
		}
	}
}
