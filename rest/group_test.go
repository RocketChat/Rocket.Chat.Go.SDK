package rest

import (
	"testing"

	"github.com/RocketChat/Rocket.Chat.Go.SDK/models"
	"github.com/stretchr/testify/assert"
)

func TestRocket_GetGroupMembers(t *testing.T) {
	rocket := getDefaultClient(t)
	t.Run("With channel name", func(t *testing.T) {
		general := &models.Channel{Name: "general"}
		members, err := rocket.GetGroupMembers(general)
		assert.Nil(t, err)
		// are new users added to general by default?
		assert.GreaterOrEqual(t, len(members), 1)

	})

	t.Run("With channel ID", func(t *testing.T) {
		general := &models.Channel{ID: "GENERAL"}
		members, err := rocket.GetGroupMembers(general)
		assert.Nil(t, err)
		assert.GreaterOrEqual(t, len(members), 1)
	})
}
