package discordtypes

// https://discord.com/developers/docs/resources/application#application-object
type Application struct {
	ID                  string   `json:"id"`
	Name                string   `json:"name"`
	Icon                string   `json:"icon"`
	Description         string   `json:"description"`
	RPCOrigins          []string `json:"rpc_origins,omitempty"`
	BotPublic           bool     `json:"bot_public"`
	BotRequireCodeGrant bool     `json:"bot_require_code_grant"`
	TermsOfServiceURL   string   `json:"terms_of_service_url,omitempty"`
	PrivacyPolicyURL    string   `json:"privacy_policy_url,omitempty"`
	Owner               *User    `json:"owner,omitempty"`
	Summary             string   `json:"summary"`
	VerifyKey           string   `json:"verify_key"`
	Team                *Team    `json:"team"`
	GuildID             string   `json:"guild_id,omitempty"`
	PrimarySKUID        string   `json:"primary_sku_id,omitempty"`
	Slug                string   `json:"slug,omitempty"`
	CoverImage          string   `json:"cover_image,omitempty"`
	Flags               int      `json:"flags,omitempty"`
}
