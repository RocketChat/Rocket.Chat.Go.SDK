package goRocket

import (
	"log"

	"github.com/Jeffail/gabs"
)

// GetPermissions gets permissions
//
// https://rocket.chat/docs/developer-guides/realtime-api/method-calls/get-permissions
func (c *LiveService) GetPermissions() ([]Permission, error) {
	rawResponse, err := c.client.ddp.Call("permissions/get")
	if err != nil {
		return nil, err
	}

	document, _ := gabs.Consume(rawResponse)

	perms, _ := document.Children()

	var permissions []Permission

	for _, permission := range perms {
		var roles []string
		for _, role := range permission.Path("roles").Data().([]interface{}) {
			roles = append(roles, role.(string))
		}

		permissions = append(permissions, Permission{
			ID:    stringOrZero(permission.Path("_id").Data()),
			Roles: roles,
		})
	}

	return permissions, nil
}

// GetUserRoles gets current users roles
//
// https://rocket.chat/docs/developer-guides/realtime-api/method-calls/get-user-roles
func (c *LiveService) GetUserRoles() error {
	rawResponse, err := c.client.ddp.Call("getUserRoles")
	if err != nil {
		return err
	}

	document, _ := gabs.Consume(rawResponse)

	roles, err := document.Children()
	// TODO: Figure out if this function is even useful if so return it
	log.Println(roles)

	return nil
}
