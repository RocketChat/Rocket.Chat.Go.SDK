// Provides access to Rocket.Chat's realtime API via ddp
package realtime

import (
	"fmt"
	"log"
	"math/rand"
	"net/url"
	"strconv"
	"time"

	"github.com/gopackage/ddp"
)

type Client struct {
	ddp *ddp.Client
}

// Creates a new instance and connects to the websocket.
func NewClient(serverUrl *url.URL, debug bool) (*Client, error) {
	rand.Seed(time.Now().UTC().UnixNano())

	wsURL := "ws"
	port := 80

	if serverUrl.Scheme == "https" {
		wsURL = "wss"
		port = 443
	}

	if len(serverUrl.Port()) > 0 {
		port, _ = strconv.Atoi(serverUrl.Port())
	}

	wsURL = fmt.Sprintf("%s://%v:%v/websocket", wsURL, serverUrl.Hostname(), port)

	log.Println("About to connect to:", wsURL, port, serverUrl.Scheme)

	c := new(Client)
	c.ddp = ddp.NewClient(wsURL, serverUrl.String())

	if debug {
		c.ddp.SetSocketLogActive(true)
	}

	if err := c.ddp.Connect(); err != nil {
		return nil, err
	}

	return c, nil
}

type statusListener struct {
	listener func(int)
}

func (s statusListener) Status(status int) {
	s.listener(status)
}

func (c *Client) AddStatusListener(listener func(int)) {
	c.ddp.AddStatusListener(statusListener{listener: listener})
}

func (c *Client) Reconnect() {
	c.ddp.Reconnect()
}

// ConnectionAway sets connection status to away
func (c *Client) ConnectionAway() error {
	_, err := c.ddp.Call("UserPresence:away")
	if err != nil {
		return err
	}

	return nil
}

// ConnectionOnline sets connection status to online
func (c *Client) ConnectionOnline() error {
	_, err := c.ddp.Call("UserPresence:online")
	if err != nil {
		return err
	}

	return nil
}

// Close closes the ddp session
func (c *Client) Close() {
	c.ddp.Close()
}

// Some of the rocketchat objects need unique IDs specified by the client
func (c *Client) newRandomId() string {
	return fmt.Sprintf("%f", rand.Float64())
}
