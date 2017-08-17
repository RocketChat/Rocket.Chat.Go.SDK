package rest

import (
	"testing"

	"github.com/RocketChat/Rocket.Chat.Go.SDK/common_testing"
	"github.com/stretchr/testify/assert"
)

func TestRocket_GetServerInfo(t *testing.T) {
	rocket := Client{Protocol: common_testing.Protocol, Host: common_testing.Host, Port: common_testing.Port}

	info, err := rocket.GetServerInfo()

	assert.Nil(t, err)
	assert.NotNil(t, info)

	assert.NotEmpty(t, info.Version)

	assert.NotEmpty(t, info.Build.Arch)
	assert.NotZero(t, info.Build.CpuCount)
	assert.NotEmpty(t, info.Build.Platform)
	assert.NotEmpty(t, info.Build.Date)
	assert.NotZero(t, info.Build.FreeMemory)
	assert.NotZero(t, info.Build.TotalMemory)
	assert.NotEmpty(t, info.Build.NodeVersion)
	assert.NotEmpty(t, info.Build.OsRelease)

	assert.NotEmpty(t, info.Travis.Branch)
	assert.NotEmpty(t, info.Travis.BuildNumber)
	assert.NotEmpty(t, info.Travis.Tag)

	assert.NotEmpty(t, info.Commit.Author)
	assert.NotEmpty(t, info.Commit.Branch)
	assert.NotEmpty(t, info.Commit.Date)
	assert.NotEmpty(t, info.Commit.Hash)
	assert.NotEmpty(t, info.Commit.Subject)
	assert.NotEmpty(t, info.Commit.Tag)

	assert.NotEmpty(t, info.ImageMagick.Version)
	assert.NotNil(t, info.GraphicsMagick)
}
