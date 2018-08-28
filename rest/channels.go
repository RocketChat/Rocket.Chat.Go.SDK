package rest

import (
	"bytes"
	"fmt"
	"net/url"

	"github.com/RocketChat/Rocket.Chat.Go.SDK/models"
)

// ChannelsResponse used when returning channel lists
type ChannelsResponse struct {
	Status
	models.Pagination
	Channels []models.Channel `json:"channels"`
}

// ChannelResponse on a single channel
type ChannelResponse struct {
	Status
	Channel models.Channel `json:"channel"`
}

// MembersResponse on a single channel
type MembersResponse struct {
	Status
	models.Pagination
	MembersList []models.Member `json:"members"`
}

// GetPublicChannels returns all channels that can be seen by the logged in user.
//
// https://rocket.chat/docs/developer-guides/rest-api/channels/list
func (c *RestService) GetPublicChannels() (*ChannelsResponse, error) {
	response := new(ChannelsResponse)
	if err := c.Get("channels.list", nil, response); err != nil {
		return nil, err
	}

	return response, nil
}

// GetJoinedChannels returns all channels that the user has joined.
//
// https://rocket.chat/docs/developer-guides/rest-api/channels/list-joined
func (c *RestService) GetJoinedChannels(params url.Values) (*ChannelsResponse, error) {
	response := new(ChannelsResponse)
	if err := c.Get("channels.list.joined", params, response); err != nil {
		return nil, err
	}

	return response, nil
}

// KickChannel Removes a user from the channel.
//
// https://rocket.chat/docs/developer-guides/rest-api/channels/kic
func (c *RestService) KickChannel(channel *models.Channel) error {
	var body = fmt.Sprintf(`{ "roomId": "%s", "userId": "%s"}`, channel.ID, channel.User.ID)
	return c.Post("channels.kick", bytes.NewBufferString(body), new(ChannelResponse))
}

// LeaveChannel leaves a channel. The id of the channel has to be not nil.
//
// https://rocket.chat/docs/developer-guides/rest-api/channels/leave
func (c *RestService) LeaveChannel(channel *models.Channel) error {
	var body = fmt.Sprintf(`{ "roomId": "%s"}`, channel.ID)
	return c.Post("channels.leave", bytes.NewBufferString(body), new(ChannelResponse))
}

// GetMembersList returns all channels that the user has joined.
//
// https://rocket.chat/docs/developer-guides/rest-api/channels/members/
func (c *RestService) GetMembersList(roomID string) (*MembersResponse, error) {

	params := url.Values{}
	params.Add("roomId", roomID)

	response := new(MembersResponse)

	if err := c.Get("channels.members", params, response); err != nil {
		return nil, err
	}

	return response, nil
}

// GetChannelInfo get information about a channel. That might be useful to update the usernames.
//
// https://rocket.chat/docs/developer-guides/rest-api/channels/info
func (c *RestService) GetChannelInfo(channel *models.Channel) (*models.Channel, error) {
	response := new(ChannelResponse)
	if err := c.Get("channels.info", url.Values{"roomId": []string{channel.ID}}, response); err != nil {
		return nil, err
	}
	return &response.Channel, nil
}