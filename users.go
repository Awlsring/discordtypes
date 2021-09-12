package discordtypes

// https://discord.com/developers/docs/resources/user#user-object
type User struct {
	// the user's id. oauth2: identity
	ID string `json:"id"`

	// the user's username, not unique across the platform. oauth2: identity
	Username string `json:"username"`

	// the user's 4-digit discord-tag. oauth2: identity
	Discriminator string `json:"discriminator"`

	// the user's avatar hash. https://discord.com/developers/docs/reference#image-formatting oauth2: identity
	Avatar string `json:"avatar"`

	// whether the user belongs to an OAuth2 application. oauth2: identity
	Bot bool `json:"bot,omitempty"`

	// whether the user is an Official Discord System user (part of the urgent message system). oauth2: identity
	System bool `json:"system,omitempty"`

	// Whether the user has multi-factor authentication enabled.
	MFAEnabled bool `json:"mfa_enabled,omitempty"`

	// the user's banner hash. https://discord.com/developers/docs/reference#image-formatting oauth2: identity
	Banner string `json:"banner,omitempty"`

	// the user's banner color encoded as an integer representation of hexadecimal color code. oauth2: identity
	AccentColor int `json:"accent_color,omitempty"`

	// the user's chosen language option. oauth2: identity
	Locale string `json:"locale,omitempty"`

	// whether the email on this account has been verified. oauth2: email
	Verified bool `json:"verified,omitempty"`

	// the user's email. oauth2: email
	Email string `json:"email,omitempty"`

	// the flags on a user's account. oauth2: identity
	Flags int `json:"flags,omitempty"`

	// the type of Nitro subscription on a user's account. oauth2: identity
	PremiumType int `json:"premium_type,omitempty"`

	// the public flags on a user's account. oauth2: identity
	PublicFlags int `json:"public_flags,omitempty"`
}
