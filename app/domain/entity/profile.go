package entity

type Profile struct {
	UserID          string `db:"user_id"`
	DisplayName     string `db:"display_name"`
	Icon            string `db:"icon"`
	HeaderImagePath string `db:"header_image"`
	Biography       string `db:"bio"`
	Mail            string `db:"mail"`
	Course          string `db:"course"`
}
