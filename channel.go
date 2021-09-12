package discordtypes

// https://discord.com/developers/docs/resources/channel#channel-object
type Channel struct {
	ID                         string          `json:"id"`
	Type                       int             `json:"type"`
	GuildID                    string          `json:"guild_id,omitempty"`
	Position                   int             `json:"position,omitempty"`
	PermissionOverwrites       []*Overwrite    `json:"permission_overwrites,omitempty"`
	Name                       string          `json:"name,omitempty"`
	Topic                      string          `json:"topic,omitempty"`
	NSFW                       bool            `json:"nsfw,omitempty"`
	LastMessageID              string          `json:"last_message_id,omitempty"`
	Bitrate                    int             `json:"bitrate,omitempty"`
	UserLimit                  int             `json:"user_limit,omitempty"`
	RateLimitPerUser           int             `json:"rate_limit_per_user,omitempty"`
	Recipients                 []*User         `json:"recipients,omitempty"`
	Icon                       string          `json:"icon,omitempty"`
	OwnerID                    string          `json:"owner_id,omitempty"`
	ApplicationID              string          `json:"application_id,omitempty"`
	ParentID                   string          `json:"parent_id,omitempty"`
	LastPinTimestamp           string          `json:"last_pin_timestamp,omitempty"`
	RTCRegion                  string          `json:"rtc_region,omitempty"`
	VideoQualityMode           int             `json:"video_quality_mode,omitempty"`
	MessageCount               int             `json:"message_count,omitempty"`
	MemberCount                int             `json:"member_count,omitempty"`
	TheadMetadata              *ThreadMetadata `json:"thread_metadata,omitempty"`
	ThreadMember               *ThreadMember   `json:"thread_member,omitempty"`
	DefaultAutoArchiveDuration int             `json:"default_auto_archive_duration,omitempty"`
	Permissions                string          `json:"permissions,omitempty"`
}

// https://discord.com/developers/docs/resources/channel#overwrite-object
type Overwrite struct {
	ID    string `json:"id"`
	Type  int    `json:"type"`
	Allow string `json:"allow"`
	Deny  string `json:"deny"`
}

// https://discord.com/developers/docs/resources/channel#thread-metadata-object
type ThreadMetadata struct {
	Archived            bool   `json:"archived"`
	AutoArchiveDuration int    `json:"auto_archive_duration"`
	ArchiveTimestamp    string `json:"archive_timestamp"`
	Locked              bool   `json:"locked"`
	Invitable           bool   `json:"invitable,omitempty"`
}

// https://discord.com/developers/docs/resources/channel#thread-member-object
type ThreadMember struct {
	ID            string `json:"id,omitempty"`
	UserID        string `json:"user_id,omitempty"`
	JoinTimestamp string `json:"join_timestamp"`
	Flags         int    `json:"flags"`
}

// https://discord.com/developers/docs/resources/channel#message-object
type Message struct {
	ID                 string              `json:"id"`
	ChannelID          string              `json:"channel_id"`
	GuildID            string              `json:"guild_id,omitempty"`
	Author             *User               `json:"author"`
	Member             *GuildMember        `json:"member,omitempty"`
	Content            string              `json:"content"`
	Timestamp          string              `json:"timestamp"`
	EditedTimestamp    string              `json:"edited_timestamp"`
	TTS                bool                `json:"tts"`
	MentionEveryone    bool                `json:"mention_everyone"`
	Mentions           []*User             `json:"mentions"`
	MentionRoles       []string            `json:"mention_roles"`
	MentionChannels    []*Channel          `json:"mention_channels,omitempty"`
	Attachments        []*Attachment       `json:"attachments"`
	Embeds             []*Embed            `json:"embeds"`
	Reactions          []*Reaction         `json:"reactions"`
	Nonce              string              `json:"nonce,omitempty"` // Int or String?
	Pinned             bool                `json:"pinned"`
	WebhookID          string              `json:"webhook_id,omitempty"`
	Type               int                 `json:"type"` // https://discord.com/developers/docs/resources/channel#message-object-message-types
	Activity           *MessageActivity    `json:"activity,omitempty"`
	Application        *Application        `json:"application,omitempty"`
	ApplicationID      string              `json:"application_id,omitempty"`
	MessageReference   *MessageReference   `json:"message_reference"`
	Flags              int                 `json:"flags,omitempty"`
	ReferencedMessages *Message            `json:"referenced_messages,omitempty"`
	Interaction        *MessageInteraction `json:"interaction,omitempty"`
	Thread             *Channel            `json:"channel,omitempty"`
	Components         []*Component        `json:"components,omitempty"`
	StickerItems       []*StickerItem      `json:"sticker_items,omitempty"`
	Sticker            *Sticker            `json:"sticker,omitempty"`
}

// https://discord.com/developers/docs/resources/channel#message-reference-object-message-reference-structure
type MessageReference struct {
	MessageID       string `json:"message_id,omitempty"`
	ChannelID       string `json:"channel_id,omitempty"`
	GuildID         string `json:"guild_id,omitempty"`
	FailIfNotExists bool   `json:"fail_if_not_exists,omitempty"`
}

type MessageActivity struct {
	Type    int    `json:"type"` // https://discord.com/developers/docs/resources/channel#message-object-message-activity-types
	PartyID string `json:"party_id,omitempty"`
}

// https://discord.com/developers/docs/resources/channel#attachment-object
type Attachment struct {
	ID          string `json:"id"`
	Filename    string `json:"filename"`
	ContentType string `json:"content_type,omitempty"`
	Size        int    `json:"size"`
	URL         string `json:"url"`
	ProxyURL    string `json:"proxy_url"`
	Height      int    `json:"height,omitempty"`
	Width       int    `json:"width,omitempty"`
}

// https://discord.com/developers/docs/resources/channel#reaction-object
type Reaction struct {
	Count string `json:"user_id"`
	Me    string `json:"message_id"`
	Emoji Emoji  `json:"emoji"`
}

// https://discord.com/developers/docs/resources/channel#embed-object
type Embed struct {
	Title       string          `json:"title,omitempty"`
	Type        string          `json:"type,omitempty"`
	Description string          `json:"description,omitempty"`
	URL         string          `json:"url,omitempty"`
	Timestamp   string          `json:"timestamp,omitempty"`
	Color       int             `json:"color,omitempty"`
	Footer      *EmbedFooter    `json:"footer,omitempty"`
	Image       *EmbedImage     `json:"image,omitempty"`
	Thumbnail   *EmbedThumbnail `json:"thumbnail,omitempty"`
	Video       *EmbedVideo     `json:"video,omitempty"`
	Provider    *EmbedProvider  `json:"provider,omitempty"`
	Author      *EmbedAuthor    `json:"author,omitempty"`
	Fields      []*EmbedField   `json:"fields,omitempty"`
}

// https://discord.com/developers/docs/resources/channel#embed-object-embed-footer-structure
type EmbedFooter struct {
	Text         string `json:"text"`
	IconURL      string `json:"icon_url,omitempty"`
	ProxyIconURL string `json:"proxy_icon_url,omitempty"`
}

// https://discord.com/developers/docs/resources/channel#embed-object-embed-image-structure
type EmbedImage struct {
	URL      string `json:"url"`
	ProxyURL string `json:"proxy_url,omitempty"`
	Width    int    `json:"width,omitempty"`
	Height   int    `json:"height,omitempty"`
}

// https://discord.com/developers/docs/resources/channel#embed-object-embed-field-structure
type EmbedField struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline,omitempty"`
}

// https://discord.com/developers/docs/resources/channel#embed-object-embed-thumbnail-structure
type EmbedThumbnail struct {
	URL      string `json:"url"`
	ProxyURL string `json:"proxy_url,omitempty"`
	Width    int    `json:"width,omitempty"`
	Height   int    `json:"height,omitempty"`
}

// https://discord.com/developers/docs/resources/channel#embed-object-embed-video-structure
type EmbedVideo struct {
	URL      string `json:"url"`
	ProxyURL string `json:"proxy_url,omitempty"`
	Width    int    `json:"width,omitempty"`
	Height   int    `json:"height,omitempty"`
}

// https://discord.com/developers/docs/resources/channel#embed-object-embed-provider-structure
type EmbedProvider struct {
	Name string `json:"name,omitempty"`
	URL  string `json:"url,omitempty"`
}

// https://discord.com/developers/docs/resources/channel#embed-object-embed-author-structure
type EmbedAuthor struct {
	Name         string `json:"name"`
	URL          string `json:"url,omitempty"`
	IconURL      string `json:"icon_url,omitempty"`
	ProxyIconURL string `json:"proxy_icon_url,omitempty"`
}

// https://discord.com/developers/docs/resources/channel#allowed-mentions-object-allowed-mentions-structure
type AllowedMentions struct {
	Parse       []string `json:"parse,omitempty"` // https://discord.com/developers/docs/resources/channel#allowed-mentions-object-allowed-mention-types
	Roles       []string `json:"roles,omitempty"`
	Users       []string `json:"users,omitempty"`
	RepliedUser bool     `json:"replied_user,omitempty"`
}
