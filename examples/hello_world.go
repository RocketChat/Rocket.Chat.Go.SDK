package main

import (
	"log"
	"net/url"

	"github.com/RocketChat/Rocket.Chat.Go.SDK/models"
	"github.com/RocketChat/Rocket.Chat.Go.SDK/rest"
)

// Create a rocket chat client, login, and send a message to general channel
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

	message := models.PostMessage{
		Channel: "general",
		Text:    "Hello World!",
	}

	if _, err := rc_client.PostMessage(&message); err != nil {
		log.Fatal(err)
	}
}
