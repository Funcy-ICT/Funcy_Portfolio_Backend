package entity

import "time"

type Comment struct {
	ID        string    `db:"id"`
	UserID    string    `db:"user_id"`
	WorksID   string    `db:"works_id"`
	Text      string    `db:"text"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
