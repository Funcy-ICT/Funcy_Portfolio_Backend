package entity

type Userinfo struct {
	Profile      *Profile
	JoinedGroups *[]*GroupMember
	Skills       *[]*Skill
	SNS          *[]*SNS
}
