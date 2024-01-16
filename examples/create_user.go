package main

import (
	"fmt"
	"log"
	"net/url"

	"github.com/RocketChat/Rocket.Chat.Go.SDK/models"
	"github.com/RocketChat/Rocket.Chat.Go.SDK/rest"
)

// Create a rocket chat client, login,
// create a new user and send them a direct message
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

	createUserRequest := models.CreateUserRequest{
		Name:     "john doe",
		Email:    "johndoe@example.com",
		Password: "password",
		Username: "john_doe",
	}

	createUserResp, err := rc_client.CreateUser(&createUserRequest)
	if err != nil {
		log.Fatal(err)
	}
	newUser := createUserResp.User

	room, err := rc_client.CreateDirectMessage(newUser.Username)
	if err != nil {
		log.Fatal(err)
	}

	greeting := fmt.Sprintf("Hello %s, welcome to rocketchat", newUser.Name)
	message := models.PostMessage{
		RoomID: room.ID,
		Text:   greeting,
	}
	if _, err := rc_client.PostMessage(&message); err != nil {
		log.Fatal(err)
	}
}
