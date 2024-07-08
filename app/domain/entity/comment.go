package entity

import "time"

type Comment struct {
	ID                 string             `db:"id"`
	UserID             string             `db:"user_id"`
	WorksID            string             `db:"works_id"`
	Content            string             `db:"content"`
	CreatedAt          time.Time          `db:"created_at"`
	UpdatedAt          time.Time          `db:"updated_at"`
	CommentUserProfile CommentUserProfile `db:"-"`
}

type CommentUserProfile struct {
	DisplayName string `db:"display_name"`
	Icon        string `db:"icon"`
}
