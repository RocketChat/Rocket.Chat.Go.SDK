package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/RocketChat/Rocket.Chat.Go.SDK/models"
	"github.com/RocketChat/Rocket.Chat.Go.SDK/rest"
)

func main() {

	hostPtr := flag.String("host", "", "Rocket.chat server host")
	schemePtr := flag.String("scheme", "https", "http/https")
	userPtr := flag.String("user", "", "Rocket.chat user")
	passPtr := flag.String("pass", "", "Rocket.chat password")

	required := []string{"host", "user", "pass"}
	flag.Parse()

	seen := make(map[string]bool)
	flag.Visit(func(f *flag.Flag) { seen[f.Name] = true })
	for _, req := range required {
		if !seen[req] {
			fmt.Fprintf(os.Stderr, "Missing required argument -%s\n", req)
			os.Exit(2)
		}
	}

	rockerServer := &url.URL{Host: *hostPtr, Scheme: *schemePtr}
	debug := false

	credentials := &models.UserCredentials{Email: *userPtr, Password: *passPtr}
	rc, err := rest.NewRestClient(rockerServer, debug)
	if err != nil {
		log.Fatal(err)
	}
	err = rc.Rest.Login(credentials)
	if err != nil {
		log.Fatal(err)
	}

	myroom := &models.Channel{Name: "myroom"}

	err = rc.Rest.ChannelsCreate(myroom)
	if err != nil {
		log.Println(err)
	}

	rcChannels, err := rc.Rest.GetPublicChannels()
	if err != nil {
		log.Println(err)
	}

	for _, channel := range rcChannels.Channels {
		fmt.Printf("channel\n\tName:\t%s\n\tID:\t%s\n", channel.Name, channel.ID)

		if channel.Name == "myroom" {
			closeThisRoom := &models.Channel{ID: channel.ID}
			err = rc.Rest.ChannelClose(closeThisRoom)
			if err != nil {
				log.Println(err)
			}
		}
	}

	general := &models.Channel{ID: "GENERAL", Name: "general"}

	members, _ := rc.Rest.GetMembersList("GENERAL")
	fmt.Println(members)

	messages, err := rc.Rest.GetMessages(general, &models.Pagination{Count: 5})
	if err != nil {
		log.Println(err)
	}

	for _, message := range messages {
		fmt.Printf("%s %s\n", message.Timestamp, message.User.UserName)
		fmt.Printf("%s\n", message.Msg)
		fmt.Println("")
	}

	msgOBJ := models.Attachment{
		Color:    "#00ff00",
		Text:     "Yay for the gopher!",
		ImageURL: "https://ih1.redbubble.net/image.377846240.0222/ap,550x550,12x16,1,transparent,t.png",
		Title:    "PostMessage Example for Go",
		Fields: []models.AttachmentField{
			models.AttachmentField{Short: true, Title: "Get the package", Value: "[Link](https://github.com/RocketChat/Rocket.Chat.Go.SDK) Rocket.Chat.Go.SDK"},
		},
	}

	msgPOST := models.PostMessage{
		RoomID:  "GENERAL",
		Channel: "general",
		Emoji:   ":smirk:",
		Text:    "PostMessage API using GoLang works ok",
		Attachments: []models.Attachment{
			msgOBJ,
		},
	}

	message, err := rc.Rest.PostMessage(&msgPOST)
	log.Println(message)

	if err != nil {
		log.Println(err)
	}
}

