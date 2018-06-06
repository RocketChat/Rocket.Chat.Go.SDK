package models

type Message struct {
	Id        string `json:"_id"`
	ChannelId string `json:"rid"`
	Text      string `json:"msg"`
	Timestamp string `json:"ts"`
	User      User   `json:"u"`
}

// PostMessage Payload for postmessage rest API
//
// https://rocket.chat/docs/developer-guides/rest-api/chat/postmessage/
type PostMessage struct {
	RoomID      string       `json:"roomId,omitempty"`
	Channel     string       `json:"channel,omitempty"`
	Text        string       `json:"text,omitempty"`
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
	Ts          string `json:"ts,omitempty"`
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
	Short bool `json:"short"`
	Title bool `json:"title"`
	Value bool `json:"value"`
}
