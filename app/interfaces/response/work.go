package response

type (
	WorkID struct {
		WorkID string `json:"workID"`
	}
	Tag struct {
		Tag string `json:"tag"`
	}
	Image struct {
		Image string `json:"image"`
	}
	ReadWork struct {
		Title       string  `json:"title"`
		Description string  `json:"description"`
		Thumbnail   string  `json:"thumbnail"`
		UserIcon    string  `json:"user_icon"`
		UserName    string  `json:"user_name"`
		Images      []Image `json:"images"`
		WorkUrl     string  `json:"work_url"`
		MovieUrl    string  `json:"movie_url"`
		Tags        []Tag   `json:"tags"`
		Security    int     `json:"security"`
	}
	ReadWorks struct {
		WorkID      string `json:"workID"`
		Title       string `json:"title"`
		Thumbnail   string `json:"thumbnail"`
		Image       string `json:"image"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	}
	ReadWorksList struct {
		Works []ReadWorks `json:"works"`
	}
)
