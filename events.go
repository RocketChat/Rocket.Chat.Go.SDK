package goRocket

import "fmt"

// StartTyping ...
func (c *LiveService) StartTyping(roomID string, username string) error {
	_, err := c.client.ddp.Call("stream-notify-room", fmt.Sprintf("%s/typing", roomID), username, true)
	if err != nil {
		return err
	}

	return nil
}

// StopTyping ...
func (c *LiveService) StopTyping(roomID string, username string) error {
	_, err := c.client.ddp.Call("stream-notify-room", fmt.Sprintf("%s/typing", roomID), username, false)
	if err != nil {
		return err
	}

	return nil
}
