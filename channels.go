package goRocket

import (
	"bytes"
	"fmt"
	"log"
	"net/url"

	"github.com/Jeffail/gabs"
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

// LeaveChannel leaves a channel. The id of the channel has to be not nil.
//
// https://rocket.chat/docs/developer-guides/rest-api/channels/leave
func (c *RestService) LeaveChannel(channel *models.Channel) error {
	var body = fmt.Sprintf(`{ "roomId": "%s"}`, channel.ID)
	return c.Post("channels.leave", bytes.NewBufferString(body), new(ChannelResponse))
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

// GetChannelID ..
func (c *LiveService) GetChannelID(name string) (string, error) {
	rawResponse, err := c.client.ddp.Call("getRoomIdByNameOrId", name)
	if err != nil {
		return "", err
	}

	log.Println(rawResponse)

	return rawResponse.(string), nil
}

// GetChannelsIn returns list of channels
// Optionally includes date to get all since last check or 0 to get all
//
// https://rocket.chat/docs/developer-guides/realtime-api/method-calls/get-rooms/
func (c *LiveService) GetChannelsIn() ([]models.Channel, error) {
	rawResponse, err := c.client.ddp.Call("rooms/get", map[string]int{
		"$date": 0,
	})
	if err != nil {
		return nil, err
	}

	document, _ := gabs.Consume(rawResponse.(map[string]interface{})["update"])

	chans, err := document.Children()
	if err != nil {
		log.Println(err)
	}

	var channels []models.Channel

	for _, i := range chans {
		channels = append(channels, models.Channel{
			ID: stringOrZero(i.Path("_id").Data()),
			//Default: stringOrZero(i.Path("default").Data()),
			Name: stringOrZero(i.Path("name").Data()),
			Type: stringOrZero(i.Path("t").Data()),
		})
	}

	return channels, nil
}

// GetChannelSubscriptions gets users channel subscriptions
// Optionally includes date to get all since last check or 0 to get all
//
// https://rocket.chat/docs/developer-guides/realtime-api/method-calls/get-subscriptions
func (c *LiveService) GetChannelSubscriptions() ([]models.ChannelSubscription, error) {
	rawResponse, err := c.client.ddp.Call("subscriptions/get", map[string]int{
		"$date": 0,
	})
	if err != nil {
		return nil, err
	}

	document, err := gabs.Consume(rawResponse.(map[string]interface{})["update"])
	if err != nil {
		log.Println(err)
	}

	channelSubs, err := document.Children()
	if err != nil {
		log.Println(err)
	}

	var channelSubscriptions []models.ChannelSubscription

	for _, sub := range channelSubs {
		channelSubscription := models.ChannelSubscription{
			ID:          stringOrZero(sub.Path("_id").Data()),
			Alert:       sub.Path("alert").Data().(bool),
			Name:        stringOrZero(sub.Path("name").Data()),
			DisplayName: stringOrZero(sub.Path("fname").Data()),
			Open:        sub.Path("open").Data().(bool),
			Type:        stringOrZero(sub.Path("t").Data()),
			User: models.User{
				ID:       stringOrZero(sub.Path("u._id").Data()),
				UserName: stringOrZero(sub.Path("u.username").Data()),
			},
			Unread: sub.Path("unread").Data().(float64),
		}

		if sub.Path("roles").Data() != nil {
			var roles []string
			for _, role := range sub.Path("roles").Data().([]interface{}) {
				roles = append(roles, role.(string))
			}

			channelSubscription.Roles = roles
		}

		channelSubscriptions = append(channelSubscriptions, channelSubscription)
	}

	return channelSubscriptions, nil
}

// GetChannelRoles returns room roles
//
// https://rocket.chat/docs/developer-guides/realtime-api/method-calls/get-room-roles
func (c *LiveService) GetChannelRoles(roomID string) error {
	_, err := c.client.ddp.Call("getRoomRoles", roomID)
	if err != nil {
		return err
	}

	return nil
}

// CreateChannel creates a channel
// Takes name and users array
//
// https://rocket.chat/docs/developer-guides/realtime-api/method-calls/create-channels
func (c *LiveService) CreateChannel(name string, users []string) error {
	_, err := c.client.ddp.Call("createChannel", name, users)
	if err != nil {
		return err
	}

	return nil
}

// CreateGroup creates a private group
// Takes group name and array of users
//
// https://rocket.chat/docs/developer-guides/realtime-api/method-calls/create-private-groups
func (c *LiveService) CreateGroup(name string, users []string) error {
	_, err := c.client.ddp.Call("createPrivateGroup", name, users)
	if err != nil {
		return err
	}

	return nil
}

// JoinChannel joins a channel
// Takes roomID
//
// https://rocket.chat/docs/developer-guides/realtime-api/method-calls/joining-channels
func (c *LiveService) JoinChannel(roomID string) error {
	_, err := c.client.ddp.Call("joinRoom", roomID)
	if err != nil {
		return err
	}

	return nil
}

// LeaveChannel leaves a channel
// Takes roomID
//
// https://rocket.chat/docs/developer-guides/realtime-api/method-calls/leaving-rooms
func (c *LiveService) LeaveChannel(roomID string) error {
	_, err := c.client.ddp.Call("leaveRoom", roomID)
	if err != nil {
		return err
	}

	return nil
}

// ArchiveChannel archives the channel
// Takes roomID
//
// https://rocket.chat/docs/developer-guides/realtime-api/method-calls/archive-rooms
func (c *LiveService) ArchiveChannel(roomID string) error {
	_, err := c.client.ddp.Call("archiveRoom", roomID)
	if err != nil {
		return err
	}

	return nil
}

// UnArchiveChannel unarchives the channel
// Takes roomID
//
// https://rocket.chat/docs/developer-guides/realtime-api/method-calls/unarchive-rooms
func (c *LiveService) UnArchiveChannel(roomID string) error {
	_, err := c.client.ddp.Call("unarchiveRoom", roomID)
	if err != nil {
		return err
	}

	return nil
}

// DeleteChannel deletes the channel
// Takes roomID
//
// https://rocket.chat/docs/developer-guides/realtime-api/method-calls/delete-rooms
func (c *LiveService) DeleteChannel(roomID string) error {
	_, err := c.client.ddp.Call("eraseRoom", roomID)
	if err != nil {
		return err
	}

	return nil
}

// SetChannelTopic sets channel topic
// takes roomID and topic
//
// https://rocket.chat/docs/developer-guides/realtime-api/method-calls/save-room-settings
func (c *LiveService) SetChannelTopic(roomID string, topic string) error {
	_, err := c.client.ddp.Call("saveRoomSettings", roomID, "roomTopic", topic)
	if err != nil {
		return err
	}

	return nil
}

// SetChannelType sets the channel type
// takes roomID and roomType
//
// https://rocket.chat/docs/developer-guides/realtime-api/method-calls/save-room-settings
func (c *LiveService) SetChannelType(roomID string, roomType string) error {
	_, err := c.client.ddp.Call("saveRoomSettings", roomID, "roomType", roomType)
	if err != nil {
		return err
	}

	return nil
}

// SetChannelJoinCode sets channel join code
// takes roomID and joinCode
//
// https://rocket.chat/docs/developer-guides/realtime-api/method-calls/save-room-settings
func (c *LiveService) SetChannelJoinCode(roomID string, joinCode string) error {
	_, err := c.client.ddp.Call("saveRoomSettings", roomID, "joinCode", joinCode)
	if err != nil {
		return err
	}

	return nil
}

// SetChannelReadOnly sets channel as read only
// takes roomID and boolean
//
// https://rocket.chat/docs/developer-guides/realtime-api/method-calls/save-room-settings
func (c *LiveService) SetChannelReadOnly(roomID string, readOnly bool) error {
	_, err := c.client.ddp.Call("saveRoomSettings", roomID, "readOnly", readOnly)
	if err != nil {
		return err
	}

	return nil
}

// SetChannelDescription sets channels description
// takes roomID and description
//
// https://rocket.chat/docs/developer-guides/realtime-api/method-calls/save-room-settings
func (c *LiveService) SetChannelDescription(roomID string, description string) error {
	_, err := c.client.ddp.Call("saveRoomSettings", roomID, "roomDescription", description)
	if err != nil {
		return err
	}

	return nil
}
