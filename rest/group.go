package rest

import (
	"net/url"
	"strconv"
	"time"

	"github.com/RocketChat/Rocket.Chat.Go.SDK/models"
)

// Get messages from a dm. The channel id has to be not nil. Optionally a
// count can be specified to limit the size of the returned messages.
//
// https://rocket.chat/docs/developer-guides/rest-api/channels/history
func (c *Client) GroupHistory(channel *models.Channel, inclusive bool, fromDate time.Time, page *models.Pagination) ([]models.Message, error) {
	params := url.Values{
		"roomId": []string{channel.ID},
	}

	if page != nil {
		params.Add("count", strconv.Itoa(page.Count))
	}

	params.Add("inclusive", "true")
	params.Add("oldest", fromDate.String())

	response := new(MessagesResponse)
	if err := c.Get("groups.history", params, response); err != nil {
		return nil, err
	}

	return response.Messages, nil
}
