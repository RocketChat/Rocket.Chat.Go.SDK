package goRocket

// nolint - is unused
func (c *LiveService) getCustomEmoji() error {
	_, err := c.client.ddp.Call("listEmojiCustom")
	if err != nil {
		return err
	}

	return nil
}
