package entity

import "github.com/google/uuid"

type (
	Group struct {
		GroupID string `db:"id"`
		Name    string `db:"name"`
	}

	GroupMember struct {
		GroupID string `db:"group_id"`
		UserID  string `db:"user_id"`
		Role    string `db:"role"`
	}
)

func NewGroup(name string) *Group {
	return &Group{
		GroupID: uuid.NewString(),
		Name:    name,
	}
}

func NewGroupMember(groupID, userID, role string) *GroupMember {
	return &GroupMember{
		GroupID: groupID,
		UserID:  userID,
		Role:    role,
	}
}
