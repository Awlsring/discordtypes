package discordtypes

// https://discord.com/developers/docs/interactions/receiving-and-responding#interaction-object-interaction-structure
type Interaction struct {
	ID            string           `json:"id"`
	ApplicationID string           `json:"application_id"`
	Type          int              `json:"type"` // https://discord.com/developers/docs/interactions/receiving-and-responding#interaction-object-interaction-type
	Data          *InteractionData `json:"data,omitempty"`
	GuildID       string           `json:"guild_id,omitempty"`
	ChannelID     string           `json:"channel_id,omitempty"`
	Member        *GuildMember     `json:"member,omitempty"`
	User          *User            `json:"user,omitempty"`
	Token         string           `json:"token"`
	Version       int              `json:"version"`
	Message       *Message         `json:"message,omitempty"`
}

// https://discord.com/developers/docs/interactions/receiving-and-responding#interaction-object-interaction-data-structure
type InteractionData struct {
	ID            string                                   `json:"id"`
	Name          string                                   `json:"name"`
	Type          int                                      `json:"type"` // https://discord.com/developers/docs/interactions/receiving-and-responding#interaction-object-interaction-type
	Resolved      *ResolvedData                            `json:"resolved,omitempty`
	Options       *ApplicationCommandInteractionDataOption `json:"options,omitempty"`
	CustomID      string                                   `json:"custom_id,omitempty"`
	ComponentType int                                      `json:"component_type,omitempty"`
	Values        []*SelectOption                          `json:"values,omitempty"`
	TargetID      string                                   `json:"target_id,omitempty"`
}

//https://discord.com/developers/docs/interactions/application-commands#application-command-object-application-command-interaction-data-option-structure
type ApplicationCommandInteractionDataOption struct {
	Name    string                                     `json:"name"`
	Type    string                                     `json:"type"`
	Value   int                                        `json:"value,omitempty"`
	Options []*ApplicationCommandInteractionDataOption `json:"options,omitempty"`
}

// https://discord.com/developers/docs/interactions/receiving-and-responding#interaction-object-resolved-data-structure
type ResolvedData struct {
	Users    map[string]*User        `json:"users,omitempty"`
	Members  map[string]*GuildMember `json:"members,omitempty"`
	Roles    map[string]*Role        `json:"roles,omitempty"`
	Channels map[string]*Channel     `json:"channels,omitempty"`
	Messages map[string]*Message     `json:"messages,omitempty"`
}

// https://discord.com/developers/docs/interactions/receiving-and-responding#message-interaction-object-message-interaction-structure
type MessageInteraction struct {
	ID   string `json:"id"`
	Type int    `json:"type"` //https://discord.com/developers/docs/interactions/receiving-and-responding#interaction-object-interaction-type
	Name string `json:"name"`
	User *User  `json:"user"`
}

// https://discord.com/developers/docs/interactions/message-components#component-object
type Component struct {
	Type        int             `json:"type"`
	CustomID    string          `json:"custom_id,omitempty"`
	Disabled    bool            `json:"disabled,omitempty"`
	Style       int             `json:"style,omitempty"`
	Label       string          `json:"label,omitempty"`
	Emoji       *Emoji          `json:"emoji,omitempty"`
	Url         string          `json:"url,omitempty"`
	Options     []*SelectOption `json:"options,omitempty"`
	Placeholder string          `json:"placeholder,omitempty"`
	MinValues   int             `json:"min_values,omitempty"`
	MaxValues   int             `json:"max_values,omitempty"`
	Components  []*Component    `json:"components,omitempty"`
}

// https://discord.com/developers/docs/interactions/message-components#select-menu-object-select-option-structure
type SelectOption struct {
	Label       string `json:"label"`
	Value       string `json:"value"`
	Description string `json:"description,omitempty"`
	Emoji       Emoji  `json:"emoji,omitempty"`
	Default     bool   `json:"default,omitempty"`
}

// https://discord.com/developers/docs/interactions/receiving-and-responding#interaction-response-object-interaction-response-structure
type InteractionResponse struct {
	Type int                            `json:"type"` // https://discord.com/developers/docs/interactions/receiving-and-responding#interaction-response-object-interaction-callback-type
	Data DiscordInteractionResponseData `json:"data,omitempty"`
}

// https://discord.com/developers/docs/interactions/receiving-and-responding#interaction-response-object-interaction-callback-data-structure
type InteractionCallbackData struct {
	TTS             bool             `json:"tts,omitempty"`
	Content         string           `json:"content,omitempty"`
	Embeds          []*Embed         `json:"embeds,omitempty"`
	AllowedMentions *AllowedMentions `json:"allowed_mentions,omitempty"`
	Flags           int              `json:"flags,omitempty"` // https://discord.com/developers/docs/interactions/receiving-and-responding#interaction-response-object-interaction-callback-data-flags
	Components      []*Component     `json:"components,omitempty"`
}
