package entity

type Userinfo struct {
	Profile      *Profile
	JoinedGroups *[]*GroupMember
	Skills       *[]*Skill
	SNS          *[]*SNS
}

type UpdateUserinfo struct {
	Profile *Profile
	Skills  *[]*Skill
	SNS     *[]*SNS
}
