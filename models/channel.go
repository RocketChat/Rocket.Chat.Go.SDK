package models

import "time"

// Channel ...
type Channel struct {
	ID    string `json:"_id"`
	Name  string `json:"name"`
	Fname string `json:"fname,omitempty"`
	Type  string `json:"t"`
	Msgs  int    `json:"msgs"`

	ReadOnly  bool `json:"ro,omitempty"`
	SysMes    bool `json:"sysMes,omitempty"`
	Default   bool `json:"default"`
	Broadcast bool `json:"broadcast,omitempty"`

	Timestamp *time.Time `json:"ts,omitempty"`
	UpdatedAt *time.Time `json:"_updatedAt,omitempty"`

	User        *User    `json:"u,omitempty"`
	LastMessage *Message `json:"lastMessage,omitempty"`

	// Lm          interface{} `json:"lm"`
	// CustomFields struct {
	// } `json:"customFields,omitempty"`
}

// ChannelSubscription ...
type ChannelSubscription struct {
	ID          string   `json:"_id"`
	Alert       bool     `json:"alert"`
	Name        string   `json:"name"`
	DisplayName string   `json:"fname"`
	Open        bool     `json:"open"`
	RoomID      string   `json:"rid"`
	Type        string   `json:"c"`
	User        User     `json:"u"`
	Roles       []string `json:"roles"`
	Unread      float64  `json:"unread"`
}

//MembersList members in a channel
type MembersList struct {
	Members []Member `json:"members"`
	Count   int      `json:"count"`
	Offset  int      `json:"offset"`
	Total   int      `json:"total"`
	Success bool     `json:"success"`
}

// Member ...
type Member struct {
	ID       string `json:"_id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Status   string `json:"status"`
}
