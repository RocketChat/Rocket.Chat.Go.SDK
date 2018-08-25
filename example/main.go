package main

import (
	"flag"
	"fmt"
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

	flag.Parse()

	required := []string{"host", "user", "pass"}
	flag.Parse()

	seen := make(map[string]bool)
	flag.Visit(func(f *flag.Flag) { seen[f.Name] = true })
	for _, req := range required {
		if !seen[req] {
			fmt.Fprintf(os.Stderr, "missing required -%s argument/flag\n", req)
			os.Exit(2)
		}
	}

	rockerServer := &url.URL{Host: *hostPtr, Scheme: *schemePtr}
	debug := false

	credentials := &models.UserCredentials{Email: *userPtr, Password: *passPtr}
	rc := rest.NewClient(rockerServer, debug)

	rc.Login(credentials)

	rcChannels, _ := rc.GetPublicChannels()

	for _, channel := range rcChannels.Channels {
		fmt.Printf("channel\n\tName:\t%s\n\tID:\t%s\n", channel.Name, channel.ID)
	}

	general := &models.Channel{ID: "GENERAL", Name: "general"}

	//rc.Send(eegabeevas, "Hello -  using Send ")

	messages, _ := rc.GetMessages(general, &models.Pagination{Count: 5})

	for _, message := range messages {
		fmt.Printf("%s %s\n", message.Timestamp, message.User.UserName)
		fmt.Printf("%s\n", message.Msg)
		fmt.Println("")
	}

}
