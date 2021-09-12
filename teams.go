package discordtypes

// https://discord.com/developers/docs/topics/teams#data-models-team-object
type Team struct {
	Icon        string         `json:"icon"`
	ID          string         `json:"id"`
	Members     []*TeamMembers `json:"members"`
	Name        string         `json:"name"`
	OwnerUserID string         `json:"owner_user_id"`
}

// https://discord.com/developers/docs/topics/teams#data-models-team-members-object
type TeamMembers struct {
	User            *User    `json:"user"`
	TeamID          string   `json:"team_id"`
	MembershipState int      `json:"membership_state"`
	Permissions     []string `json:"permissions"`
}
