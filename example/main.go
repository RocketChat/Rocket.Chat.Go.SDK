package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/RocketChat/Rocket.Chat.Go.SDK"
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
			fmt.Fprintf(os.Stderr, "Missing required argument -%s\n", req)
			os.Exit(2)
		}
	}

	rockerServer := &url.URL{Host: *hostPtr, Scheme: *schemePtr}
	debug := false

	credentials := &goRocket.UserCredentials{Email: *userPtr, Password: *passPtr}
	rc, err := goRocket.NewRestClient(rockerServer, debug)
	if err != nil {
		log.Fatal(err)
	}
	err = rc.Rest.Login(credentials)
	if err != nil {
		log.Fatal(err)
	}

	rcChannels, err := rc.Rest.GetPublicChannels()
	if err != nil {
		log.Println(err)
	}

	for _, channel := range rcChannels.Channels {
		fmt.Printf("channel\n\tName:\t%s\n\tID:\t%s\n", channel.Name, channel.ID)
	}

	general := &goRocket.Channel{ID: "GENERAL", Name: "general"}

	messages, err := rc.Rest.GetMessages(general, &goRocket.Pagination{Count: 5})
	if err != nil {
		log.Println(err)
	}

	for _, message := range messages {
		fmt.Printf("%s %s\n", message.Timestamp, message.User.UserName)
		fmt.Printf("%s\n", message.Msg)
		fmt.Println("")
	}

}
