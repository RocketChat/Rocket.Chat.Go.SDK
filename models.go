package goRocket

import "time"

// Channel model
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

// ChannelSubscription model
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

// Info model
type Info struct {
	Version string `json:"version"`

	Build struct {
		NodeVersion string `json:"nodeVersion"`
		Arch        string `json:"arch"`
		Platform    string `json:"platform"`
		Cpus        int    `json:"cpus"`
	} `json:"build"`

	Commit struct {
		Hash    string `json:"hash"`
		Date    string `json:"date"`
		Author  string `json:"author"`
		Subject string `json:"subject"`
		Tag     string `json:"tag"`
		Branch  string `json:"branch"`
	} `json:"commit"`
}

// Pagination model
type Pagination struct {
	Count  int `json:"count"`
	Offset int `json:"offset"`
	Total  int `json:"total"`
}

// Directory model
type Directory struct {
	Result []struct {
		ID        string    `json:"_id"`
		CreatedAt time.Time `json:"createdAt"`
		Emails    []struct {
			Address  string `json:"address"`
			Verified bool   `json:"verified"`
		} `json:"emails"`
		Name     string `json:"name"`
		Username string `json:"username"`
	} `json:"result"`

	Pagination
}

// Spotlight model
type Spotlight struct {
	Users []User    `json:"users"`
	Rooms []Channel `json:"rooms"`
}

// Statistics model
type Statistics struct {
	ID       string `json:"_id"`
	UniqueID string `json:"uniqueId"`
	Version  string `json:"version"`

	ActiveUsers    int `json:"activeUsers"`
	NonActiveUsers int `json:"nonActiveUsers"`
	OnlineUsers    int `json:"onlineUsers"`
	AwayUsers      int `json:"awayUsers"`
	OfflineUsers   int `json:"offlineUsers"`
	TotalUsers     int `json:"totalUsers"`

	TotalRooms                int `json:"totalRooms"`
	TotalChannels             int `json:"totalChannels"`
	TotalPrivateGroups        int `json:"totalPrivateGroups"`
	TotalDirect               int `json:"totalDirect"`
	TotlalLivechat            int `json:"totlalLivechat"`
	TotalMessages             int `json:"totalMessages"`
	TotalChannelMessages      int `json:"totalChannelMessages"`
	TotalPrivateGroupMessages int `json:"totalPrivateGroupMessages"`
	TotalDirectMessages       int `json:"totalDirectMessages"`
	TotalLivechatMessages     int `json:"totalLivechatMessages"`

	InstalledAt          time.Time `json:"installedAt"`
	LastLogin            time.Time `json:"lastLogin"`
	LastMessageSentAt    time.Time `json:"lastMessageSentAt"`
	LastSeenSubscription time.Time `json:"lastSeenSubscription"`

	Os struct {
		Type     string    `json:"type"`
		Platform string    `json:"platform"`
		Arch     string    `json:"arch"`
		Release  string    `json:"release"`
		Uptime   int       `json:"uptime"`
		Loadavg  []float64 `json:"loadavg"`
		Totalmem int64     `json:"totalmem"`
		Freemem  int       `json:"freemem"`
		Cpus     []struct {
			Model string `json:"model"`
			Speed int    `json:"speed"`
			Times struct {
				User int `json:"user"`
				Nice int `json:"nice"`
				Sys  int `json:"sys"`
				Idle int `json:"idle"`
				Irq  int `json:"irq"`
			} `json:"times"`
		} `json:"cpus"`
	} `json:"os"`

	Process struct {
		NodeVersion string  `json:"nodeVersion"`
		Pid         int     `json:"pid"`
		Uptime      float64 `json:"uptime"`
	} `json:"process"`

	Deploy struct {
		Method   string `json:"method"`
		Platform string `json:"platform"`
	} `json:"deploy"`

	Migration struct {
		ID       string    `json:"_id"`
		Version  int       `json:"version"`
		Locked   bool      `json:"locked"`
		LockedAt time.Time `json:"lockedAt"`
		BuildAt  time.Time `json:"buildAt"`
	} `json:"migration"`

	InstanceCount int       `json:"instanceCount"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"_updatedAt"`
}

// StatisticsInfo model
type StatisticsInfo struct {
	Statistics Statistics `json:"statistics"`
}

// StatisticsList model
type StatisticsList struct {
	Statistics []Statistics `json:"statistics"`

	Pagination
}

// Message model
type Message struct {
	ID       string `json:"_id"`
	RoomID   string `json:"rid"`
	Msg      string `json:"msg"`
	EditedBy string `json:"editedBy,omitempty"`

	Groupable bool `json:"groupable,omitempty"`

	EditedAt  *time.Time `json:"editedAt,omitempty"`
	Timestamp *time.Time `json:"ts,omitempty"`
	UpdatedAt *time.Time `json:"_updatedAt,omitempty"`

	Mentions []User `json:"mentions,omitempty"`
	User     *User  `json:"u,omitempty"`
	PostMessage

	// Bot         interface{}  `json:"bot"`
	// CustomFields interface{} `json:"customFields"`
	// Channels           []interface{} `json:"channels"`
	// SandstormSessionID interface{} `json:"sandstormSessionId"`
}

// PostMessage Payload for postmessage rest API
//
// https://rocket.chat/docs/developer-guides/rest-api/chat/postmessage/
type PostMessage struct {
	RoomID      string       `json:"roomId,omitempty"`
	Channel     string       `json:"channel,omitempty"`
	Text        string       `json:"text,omitempty"`
	ParseUrls   bool         `json:"parseUrls,omitempty"`
	Alias       string       `json:"alias,omitempty"`
	Emoji       string       `json:"emoji,omitempty"`
	Avatar      string       `json:"avatar,omitempty"`
	Attachments []Attachment `json:"attachments,omitempty"`
}

// Attachment Payload for postmessage rest API
//
// https://rocket.chat/docs/developer-guides/rest-api/chat/postmessage/
type Attachment struct {
	Color       string `json:"color,omitempty"`
	Text        string `json:"text,omitempty"`
	Timestamp   string `json:"ts,omitempty"`
	ThumbURL    string `json:"thumb_url,omitempty"`
	MessageLink string `json:"message_link,omitempty"`
	Collapsed   bool   `json:"collapsed"`

	AuthorName string `json:"author_name,omitempty"`
	AuthorLink string `json:"author_link,omitempty"`
	AuthorIcon string `json:"author_icon,omitempty"`

	Title             string `json:"title,omitempty"`
	TitleLink         string `json:"title_link,omitempty"`
	TitleLinkDownload string `json:"title_link_download,omitempty"`

	ImageURL string `json:"image_url,omitempty"`

	AudioURL string `json:"audio_url,omitempty"`
	VideoURL string `json:"video_url,omitempty"`

	Fields []AttachmentField `json:"fields,omitempty"`
}

// AttachmentField Payload for postmessage rest API
//
// https://rocket.chat/docs/developer-guides/rest-api/chat/postmessage/
type AttachmentField struct {
	Short bool   `json:"short"`
	Title string `json:"title"`
	Value string `json:"value"`
}

// Permission model
type Permission struct {
	ID        string   `json:"_id"`
	UpdatedAt string   `json:"_updatedAt.$date"`
	Roles     []string `json:"roles"`
}

// Setting model
type Setting struct {
	ID           string  `json:"_id"`
	Blocked      bool    `json:"blocked"`
	Group        string  `json:"group"`
	Hidden       bool    `json:"hidden"`
	Public       bool    `json:"public"`
	Type         string  `json:"type"`
	PackageValue string  `json:"packageValue"`
	Sorter       int     `json:"sorter"`
	Value        string  `json:"value"`
	ValueBool    bool    `json:"valueBool"`
	ValueInt     float64 `json:"valueInt"`
	ValueSource  string  `json:"valueSource"`
	ValueAsset   Asset   `json:"asset"`
}

// Asset model
type Asset struct {
	DefaultURL string `json:"defaultUrl"`
}

// User model
type User struct {
	ID           string `json:"_id"`
	Name         string `json:"name"`
	UserName     string `json:"username"`
	Status       string `json:"status"`
	Token        string `json:"token"`
	TokenExpires int64  `json:"tokenExpires"`
}

// UserCredentials model
type UserCredentials struct {
	ID    string `json:"id"`
	Token string `json:"token"`

	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"pass"`
}
