package rest

import (
	"bytes"
	"encoding/json"

	"github.com/RocketChat/Rocket.Chat.Go.SDK/models"
)

type updatePermissionsRequest struct {
	Permissions []models.Permission `json:"permissions"`
}

type updatePermissionsResponse struct {
	Status
	Permissions []models.Permission `json:"permissions"`
}

// UpdatePermissions updates permissions
//
// https://rocket.chat/docs/developer-guides/rest-api/permissions/update/
func (c *Client) UpdatePermissions(permissions []models.Permission) ([]models.Permission, error) {
	req := updatePermissionsRequest{Permissions: permissions}
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	response := new(updatePermissionsResponse)
	if err := c.Post("permissions.update", bytes.NewBuffer(body), response); err != nil {
		return nil, err
	}
	return response.Permissions, nil
}
