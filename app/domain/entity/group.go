package entity

import "github.com/google/uuid"

type (
	Group struct {
		GroupID     string `db:"id"`
		Name        string `db:"name"`
		Description string `db:"description"`
		LeaderEmail string `db:"leader_email"`
		Icon        string `db:"icon"`
	}

	GroupSkill struct {
		GroupID   string `db:"group_id"`
		SkillName string `db:"skill_name"`
	}

	GroupMember struct {
		GroupID   string `db:"group_id"`
		GroupName string `db:"name"`
		UserID    string `db:"user_id"`
		Role      string `db:"role"`
		Status    bool   `db:"status"`
	}
)

func NewGroup(name string, description string, leaderEmail string, icon string) *Group {
	return &Group{
		GroupID:     uuid.NewString(),
		Name:        name,
		Description: description,
		LeaderEmail: leaderEmail,
		Icon:        icon,
	}
}

func NewGroupSkill(groupId string, skillName string) *GroupSkill {
	return &GroupSkill{
		GroupID:   groupId,
		SkillName: skillName,
	}
}
