package entity

type (
	WorkTable struct {
		ID          string `db:"id"`
		Title       string `db:"title"`
		Description string `db:"description"`
		URL         string `db:"url"`
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
		ImageURLs   []string `db:"image_url"`
		Tags        []string `db:"tag"`
		WorkURL     string   `db:"url"`
		MovieUrl    string   `db:"movie_url"`
		Security    int      `db:"security"`
	}
	ReadWorksList struct {
		WorkID      string `db:"id"`
		Title       string `db:"title"`
		Images      string `db:"image_url"`
		Description string `db:"description"`
		Icon        string `db:"icon"`
	}
)

func NewReadWorksList(workID string, title string, images string, description string, icon string) *ReadWorksList {
	return &ReadWorksList{
		WorkID:      workID,
		Title:       title,
		Images:      images,
		Description: description,
		Icon:        icon,
	}
}
