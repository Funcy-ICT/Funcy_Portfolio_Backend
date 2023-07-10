package entity

type Skill struct {
	SkillName string `db:"skill_name"`
	UserID    string `db:"user_id"`
}
