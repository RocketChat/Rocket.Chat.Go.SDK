package models

type User struct {
	ID           string `json:"_id"`
	Name         string `json:"name"`
	UserName     string `json:"username"`
	Status       string `json:"status"`
	Token        string `json:"token"`
	TokenExpires int64  `json:"tokenExpires"`
}

type CreateUserRequest struct {
	Name         string            `json:"name"`
	Email        string            `json:"email"`
	Password     string            `json:"password"`
	Username     string            `json:"username"`
	Roles        []string          `json:"roles,omitempty"`
	CustomFields map[string]string `json:"customFields,omitempty"`
}

type UpdateUserRequest struct {
	UserID string `json:"userId"`
	Data   struct {
		Name         string            `json:"name"`
		Email        string            `json:"email"`
		Password     string            `json:"password"`
		Username     string            `json:"username"`
		CustomFields map[string]string `json:"customFields,omitempty"`
	} `json:"data"`
}

type Preferences struct {
	EnableAutoAway              bool   `json:"enableAutoAway,omitempty"`
	IdleTimeoutLimit            int    `json:"idleTimeoutLimit,omitempty"`
	DesktopNotificationDuration int    `json:"desktopNotificationDuration,omitempty"`
	AudioNotifications          string `json:"audioNotifications,omitempty"`
	DesktopNotifications        string `json:"desktopNotifications,omitempty"`
	MobileNotifications         string `json:"mobileNotifications,omitempty"`
	UnreadAlert                 bool   `json:"unreadAlert,omitempty"`
	UseEmojis                   bool   `json:"useEmojis,omitempty"`
	ConvertASCIIEmoji           bool   `json:"convertAsciiEmoji,omitempty"`
	AutoImageLoad               bool   `json:"autoImageLoad,omitempty"`
	SaveMobileBandwidth         bool   `json:"saveMobileBandwidth,omitempty"`
	CollapseMediaByDefault      bool   `json:"collapseMediaByDefault,omitempty"`
	HideUsernames               bool   `json:"hideUsernames,omitempty"`
	HideRoles                   bool   `json:"hideRoles,omitempty"`
	HideFlexTab                 bool   `json:"hideFlexTab,omitempty"`
	HideAvatars                 bool   `json:"hideAvatars,omitempty"`
	RoomsListExhibitionMode     string `json:"roomsListExhibitionMode,omitempty"`
	SidebarViewMode             string `json:"sidebarViewMode,omitempty"`
	SidebarHideAvatar           bool   `json:"sidebarHideAvatar,omitempty"`
	SidebarShowUnread           bool   `json:"sidebarShowUnread,omitempty"`
	SidebarShowFavorites        bool   `json:"sidebarShowFavorites,omitempty"`
	SendOnEnter                 string `json:"sendOnEnter,omitempty"`
	MessageViewMode             int    `json:"messageViewMode,omitempty"`
	EmailNotificationMode       string `json:"emailNotificationMode,omitempty"`
	RoomCounterSidebar          bool   `json:"roomCounterSidebar,omitempty"`
	NewRoomNotification         string `json:"newRoomNotification,omitempty"`
	NewMessageNotification      string `json:"newMessageNotification,omitempty"`
	MuteFocusedConversations    bool   `json:"muteFocusedConversations,omitempty"`
	NotificationsSoundVolume    int    `json:"notificationsSoundVolume,omitempty"`
}

type Email struct {
	Address  string `json:"address"`
	Verified bool   `json:"verified"`
}

type Me struct {
	User
	Emails           []Email  `json:"emails"`
	StatusConnection string   `json:"statusConnection"`
	UtcOffset        int      `json:"utcOffset"`
	Active           bool     `json:"active"`
	Roles            []string `json:"roles"`
	Settings         struct {
		Preferences `json:"preferences"`
	} `json:"settings"`
	CustomFields struct {
		Twitter string `json:"twitter,omitempty"`
	} `json:"customFields,omitempty"`
	AvatarURL string `json:"avatarUrl"`
}
