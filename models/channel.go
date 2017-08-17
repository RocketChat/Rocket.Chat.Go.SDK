package models

type Channel struct {
	Id           string   `json:"_id"`
	Name         string   `json:"name"`
	MessageCount int      `json:"msgs"`
	UserNames    []string `json:"usernames"`
	Default      bool     `json:"default"`

	User User `json:"u"`

	ReadOnly  bool   `json:"ro"`
	Timestamp string `json:"ts"`
	Type      string `json:"t"`
	UpdatedAt string `json:"_updatedAt"`
	SysMes    bool   `json:"sysMes"`
}

type ChannelSubscription struct {
	Id          string   `json:"_id"`
	Alert       bool     `json:"alert"`
	Name        string   `json:"name"`
	DisplayName string   `json:"fname"`
	Open        bool     `json:"open"`
	RoomId      string   `json:"rid"`
	Type        string   `json:"c"`
	User        User     `json:"u"`
	Roles       []string `json:"roles"`
	Unread      float64      `json:"unread"`
}
