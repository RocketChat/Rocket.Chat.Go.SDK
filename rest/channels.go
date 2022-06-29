package rest

import (
	"bytes"
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

type GroupsResponse struct {
	Status
	models.Pagination
	Groups []models.Channel `json:"groups"`
}

type GroupResponse struct {
	Status
	Group models.Channel `json:"group"`
}

// GetPublicChannels lists all of the channels on the server.
//
// https://developer.rocket.chat/reference/api/rest-api/endpoints/core-endpoints/channels-endpoints/list
func (c *Client) GetPublicChannels() (*ChannelsResponse, error) {
	response := new(ChannelsResponse)
	if err := c.Get("channels.list", nil, response); err != nil {
		return nil, err
	}

	return response, nil
}

// GetPrivateGroups lists all of the private groups the calling user has joined.
//
// https://developer.rocket.chat/reference/api/rest-api/endpoints/core-endpoints/groups-endpoints/list
func (c *Client) GetPrivateGroups() (*GroupsResponse, error) {
	response := new(GroupsResponse)
	if err := c.Get("groups.list", nil, response); err != nil {
		return nil, err
	}

	return response, nil
}

// GetJoinedChannels lists all of the channels the calling user has joined.
//
// https://developer.rocket.chat/reference/api/rest-api/endpoints/core-endpoints/channels-endpoints/list-joined
func (c *Client) GetJoinedChannels(params url.Values) (*ChannelsResponse, error) {
	response := new(ChannelsResponse)
	if err := c.Get("channels.list.joined", params, response); err != nil {
		return nil, err
	}

	return response, nil
}

// LeaveChannel causes the callee to be removed from the channel.
// The id of the channel must not be nil.
//
// https://developer.rocket.chat/reference/api/rest-api/endpoints/core-endpoints/channels-endpoints/leave
func (c *Client) LeaveChannel(channel *models.Channel) error {
	var body = fmt.Sprintf(`{ "roomId": "%s"}`, channel.ID)
	return c.Post("channels.leave", bytes.NewBufferString(body), new(ChannelResponse))
}

// GetChannelInfo retrieves the information about the channel.
//
// https://developer.rocket.chat/reference/api/rest-api/endpoints/core-endpoints/channels-endpoints/info
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

// GetGroupInfo retrieves the information about the private group, only if you're part of the group.
//
// https://developer.rocket.chat/reference/api/rest-api/endpoints/core-endpoints/groups-endpoints/info
func (c *Client) GetGroupInfo(channel *models.Channel) (*models.Channel, error) {
	response := new(GroupResponse)
	switch {
	case channel.Name != "" && channel.ID == "":
		if err := c.Get("groups.info", url.Values{"roomName": []string{channel.Name}}, response); err != nil {
			return nil, err
		}
	default:
		if err := c.Get("groups.info", url.Values{"roomId": []string{channel.ID}}, response); err != nil {
			return nil, err
		}
	}
	return &response.Group, nil
}
