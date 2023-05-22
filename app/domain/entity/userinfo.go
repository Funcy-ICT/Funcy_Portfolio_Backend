package entity

type Userinfo struct {
	Profile *Profile
	Groups  *[]*Group
	Skills  *[]*Skill
	SNS     *[]*SNS
	Works   *[]*ReadWorksList
}
