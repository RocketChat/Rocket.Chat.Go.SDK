package rest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/RocketChat/Rocket.Chat.Go.SDK/models"
)

type ChannelsResponse struct {
	Status
	models.Pagination
	Channels []models.Channel `json:"channels"`
}

type ChannelResponse struct {
	Status
	Channel models.Channel `json:"channel"`
}

// CreateChannel is payload for channels.create
//
// https://developer.rocket.chat/reference/api/rest-api/endpoints/core-endpoints/channels-endpoints/create
type CreateChannel struct {
	ChannelName string   `json:"name"`
	Members     []string `json:"members"`
	ReadOnly    bool     `json:"readOnly"`
}

type channelInvite struct {
	RoomId  string   `json:"roomId"`
	UserIds []string `json:"userIds"`
}

// GetPublicChannels returns all channels that can be seen by the logged in user.
//

// https://rocket.chat/docs/developer-guides/rest-api/channels/list
func (c *Client) GetPublicChannels() (*ChannelsResponse, error) {
	response := new(ChannelsResponse)
	if err := c.Get("channels.list", nil, response); err != nil {
		return nil, err
	}

	return response, nil
}

// GetJoinedChannels returns all channels that the user has joined.
//
// https://rocket.chat/docs/developer-guides/rest-api/channels/list-joined
func (c *Client) GetJoinedChannels(params url.Values) (*ChannelsResponse, error) {
	response := new(ChannelsResponse)
	if err := c.Get("channels.list.joined", params, response); err != nil {
		return nil, err
	}

	return response, nil
}

// LeaveChannel leaves a channel. The id of the channel has to be not nil.
//
// https://rocket.chat/docs/developer-guides/rest-api/channels/leave
func (c *Client) LeaveChannel(channel *models.Channel) error {
	var body = fmt.Sprintf(`{ "roomId": "%s"}`, channel.ID)
	return c.Post("channels.leave", bytes.NewBufferString(body), new(ChannelResponse))
}

// GetChannelInfo get information about a channel. That might be useful to update the usernames.
//
// https://rocket.chat/docs/developer-guides/rest-api/channels/info
func (c *Client) GetChannelInfo(channel *models.Channel) (*models.Channel, error) {
	response := new(ChannelResponse)
	switch {
	case channel.Name != "" && channel.ID == "":
		if err := c.Get("channels.info", url.Values{"roomName": []string{channel.Name}}, response); err != nil {
			return nil, err
		}
	default:
		if err := c.Get("channels.info", url.Values{"roomId": []string{channel.ID}}, response); err != nil {
			return nil, err
		}
	}

	return &response.Channel, nil
}

// CreateChannel creates a new public channel, optionally including specified users.
// The channel creator is always included.
//
// https://developer.rocket.chat/reference/api/rest-api/endpoints/core-endpoints/channels-endpoints/create
func (c *Client) CreateChannel(channel *CreateChannel) (*ChannelResponse, error) {
	body, err := json.Marshal(channel)
	if err != nil {
		return nil, err
	}

	response := new(ChannelResponse)
	err = c.Post("channels.create", bytes.NewBuffer(body), response)
	return response, err
}

// InviteChannel adds users to the channel.
//
// https://developer.rocket.chat/reference/api/rest-api/endpoints/core-endpoints/channels-endpoints/invite
func (c *Client) InviteChannel(channel *models.Channel, users []*models.User) (*ChannelResponse, error) {
	var ids []string
	for _, user := range users {
		ids = append(ids, user.ID)
	}
	invite := &channelInvite{RoomId: channel.ID, UserIds: ids}
	body, err := json.Marshal(invite)
	if err != nil {
		return nil, err
	}

	response := new(ChannelResponse)
	err = c.Post("channels.invite", bytes.NewBuffer(body), response)
	return response, err
}
