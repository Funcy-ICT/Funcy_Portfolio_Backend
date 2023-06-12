package entity

type (
	WorkTable struct {
		ID          string `db:"id"`
		Title       string `db:"title"`
		Description string `db:"description"`
		Thumbnail   string `db:"thumbnail"`
		WorkUrl     string `db:"work_url"`
		MovieUrl    string `db:"movie_url"`
		Security    int    `db:"security"`
	}
	Tag struct {
		ID     string `db:"id"`
		WorkID string `db:"work_id"`
		Tag    string `db:"tag"`
	}
	Image struct {
		ID     string `db:"id"`
		WorkID string `db:"work_id"`
		Image  string `db:"image_url"`
	}
	ReadWork struct {
		Title       string   `db:"title"`
		Description string   `db:"description"`
		Thumbnail   string   `db:"thumbnail"`
		UserId      string   `db:"user_id"`
		ImageURLs   []string `db:"image_url"`
		Tags        []string `db:"tag"`
		WorkUrl     string   `db:"url"`
		MovieUrl    string   `db:"movie_url"`
		Security    int      `db:"security"`
	}
	ReadWorksList struct {
		WorkID      string `db:"id"`
		Title       string `db:"title"`
		Thumbnail   string `db:"thumbnail"`
		Images      string `db:"image_url"`
		Description string `db:"description"`
		Icon        string `db:"icon"`
	}
	UpdateWork struct {
		Title       string   `db:"title"`
		Description string   `db:"description"`
		Thumbnail   string   `db:"thumbnail"`
		ImageURLs   []string `db:"image_url"`
		Tags        []string `db:"tag"`
		WorkUrl     string   `db:"work_url"`
		MovieUrl    string   `db:"movie_url"`
		Security    int      `db:"security"`
	}
)
