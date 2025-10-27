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
