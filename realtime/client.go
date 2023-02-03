//Package realtime provides access to Rocket.Chat's realtime API via ddp
package realtime

import (
	"errors"
	"fmt"
	"log"
	"net/url"
	"strconv"

	"github.com/gopackage/ddp"
	"github.com/sony/sonyflake"
)

type Client struct {
	ddp *ddp.Client
	sf  *sonyflake.Sonyflake
}

//NewClient creates a new instance and connects to the websocket.
func NewClient(serverURL *url.URL, debug bool) (*Client, error) {
	sf := sonyflake.NewSonyflake(sonyflake.Settings{})
	if sf == nil {
		return nil, errors.New("random id generator failed to be created")
	}

	wsURL := "ws"
	port := 80

	if serverURL.Scheme == "https" {
		wsURL = "wss"
		port = 443
	}

	if len(serverURL.Port()) > 0 {
		port, _ = strconv.Atoi(serverURL.Port())
	}

	wsURL = fmt.Sprintf("%s://%v:%v%s/websocket", wsURL, serverURL.Hostname(), port, serverURL.Path)

	log.Println("About to connect to:", wsURL, port, serverURL.Scheme)

	c := new(Client)
	c.ddp = ddp.NewClient(wsURL, serverURL.String())
	c.sf = sf

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
	id, err := c.sf.NextID()
	if err != nil {
		log.Fatalf("failed to create a unique id: %v", err.Error())
	}

	return fmt.Sprintf("go%d", id)
}
