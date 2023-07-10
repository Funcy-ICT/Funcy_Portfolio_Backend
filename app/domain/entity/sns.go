package entity

type SNS struct {
	UserID string `db:"user_id"`
	SnsURL string `db:"sns"`
}
