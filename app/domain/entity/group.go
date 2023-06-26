package entity

import "github.com/google/uuid"

type (
	Group struct {
		GroupID string `db:"id"`
		Name    string `db:"name"`
	}

	GroupMember struct {
		GroupID   string `db:"group_id"`
		GroupName string `db:"name"`
		UserID    string `db:"user_id"`
		Role      string `db:"role"`
		Status    bool   `db:"status"`
	}
)

func NewGroup(name string) *Group {
	return &Group{
		GroupID: uuid.NewString(),
		Name:    name,
	}
}
