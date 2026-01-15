package entity

type Userinfo struct {
	Profile      *Profile
	JoinedGroups *[]*GroupMember
	Skills       *[]*Skill
	SNS          *[]*SNS
	Course       string
	FamilyName   string
	FirstName    string
	Grade        string
}

type UpdateUserinfo struct {
	Profile *Profile
	Skills  *[]*Skill
	SNS     *[]*SNS
}

type UserSearchResult struct {
	UserID      string   `json:"user_id"`
	DisplayName string   `json:"display_name"`
	Icon        string   `json:"icon"`
	Course      string   `json:"course"`
	Skills      []string `json:"skills"`
}
