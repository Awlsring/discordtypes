package discordtypes

// https://discord.com/developers/docs/resources/sticker#sticker-item-object
type StickerItem struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	FormatType int    `json:"format_type"`
}

// https://discord.com/developers/docs/resources/sticker#sticker-object
type Sticker struct {
	ID          string `json:"id"`
	PackID      string `json:"pack_id,omitempty"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Tags        string `json:"tags"`
	Asset       string `json:"asset"`
	Type        int    `json:"type"`
	FormatType  int    `json:"format_type"`
	Available   bool   `json:"available,omitempty"`
	GuildID     string `json:"guild_id,omitempty"`
	User        *User  `json:"user"`
	SortValue   int    `json:"sort_value"`
}
