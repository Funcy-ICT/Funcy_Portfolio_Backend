package entity

type Userinfo struct {
	Profile      *Profile
	JoinedGroups *[]*GroupMember
	Skills       *[]*Skill
	SNS          *[]*SNS
	Course       string
}

type UpdateUserinfo struct {
	Profile *Profile
	Skills  *[]*Skill
	SNS     *[]*SNS
}

type UserSearchResult struct {
	UserID      string   `db:"user_id" json:"userID"`
	DisplayName string   `db:"display_name" json:"displayName"`
	Icon        string   `db:"icon" json:"icon"`
	Course      string   `db:"course" json:"course"`
	Skills      []string `json:"skills"`
}
