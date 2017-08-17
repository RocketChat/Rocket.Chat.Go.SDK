// Provides access to Rocket.Chat's realtime API via ddp
package realtime

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/gopackage/ddp"
)

type Client struct {
	ddp *ddp.Client
}

// Creates a new instance and connects to the websocket.
func NewClient(host string, ssl bool, debug bool) (*Client, error) {
	rand.Seed(time.Now().UTC().UnixNano())

	wsUrl := "ws"
	origin := "http"

	if ssl {
		wsUrl = "wss"
		origin = "https"
	}

	wsUrl = fmt.Sprintf("%s://%v/websocket", wsUrl, host)
	origin = fmt.Sprintf("%s://%v", origin, host)

	c := new(Client)
	c.ddp = ddp.NewClient(wsUrl, origin)

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
