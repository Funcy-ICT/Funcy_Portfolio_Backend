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
	ReadWorkResponse struct {
		Title       string  `json:"title"`
		Description string  `json:"description"`
		Images      []Image `json:"images"`
		WorkURL     string  `json:"work_url"`
		MovieUrl    string  `json:"movie_url"`
		Tags        []Tag   `json:"tags"`
		Security    int     `json:"security"`
	}
	Work struct {
		WorkID      string `json:"workID"`
		Title       string `json:"title"`
		Image       string `json:"image"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	}
	Works struct {
		Works []Work `json:"works"`
	}
)
