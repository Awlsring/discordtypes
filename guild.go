package discordtypes

import (
	"time"
)

// https://discord.com/developers/docs/resources/guild#guild-member-object
type GuildMember struct {
	User         User      `json:"user,omitempty"`
	Nick         string    `json:"nick,omitempty"`
	Roles        []string  `json:"roles"`
	JoinedAt     time.Time `json:"joined_at"`
	PremiumSince time.Time `json:"premium_since,omitempty"`
	Deaf         bool      `json:"deaf"`
	Mute         bool      `json:"mute"`
	Pending      bool      `json:"pending,omitempty"`
	Permissions  string    `json:"permissions,omitempty"`
}

// https://discord.com/developers/docs/topics/permissions#role-object
type Role struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Color       int      `json:"color"`
	Hoist       bool     `json:"hoist"`
	Position    int      `json:"position"`
	Permissions int64    `json:"permissions,string"`
	Managed     bool     `json:"managed"`
	Mentionable bool     `json:"mentionable"`
	Tags        RoleTags `json:"tags,omitempty"`
}

// https://discord.com/developers/docs/topics/permissions#role-object-role-tags-structure
type RoleTags struct {
	BotID             string      `json:"bot_id,omitempty"`
	IntegrationID     string      `json:"integration_id,omitempty"`
	PremiumSubscriber interface{} `json:"premium_subscriber,omitempty"` //Idk what null would be if exists. bool?
}
