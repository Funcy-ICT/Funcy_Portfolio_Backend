package entity

import "time"

type Comment struct {
	ID              string    `db:"id"`
	UserID          string    `db:"user_id"`
	WorksID         string    `db:"works_id"`
	Content         string    `db:"content"`
	CreatedAt       time.Time `db:"created_at"`
	UpdatedAt       time.Time `db:"updated_at"`
	UserDisplayName string    `db:"user_display_name"`
	UserIcon        string    `db:"user_icon"`
}
