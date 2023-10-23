package request

type CreateWorkRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Thumbnail   string `json:"thumbnail"`
	Images      []struct {
		Image string `json:"image"`
	} `json:"images"`
	WorkUrl  string `json:"work_url"`
	MovieUrl string `json:"movie_url"`
	Tags     []struct {
		Tag string `json:"tag"`
	} `json:"tags"`
	GroupID  string `json:"group_id"`
	Security int    `json:"security"`
}

type UpdateWorkRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Thumbnail   string `json:"thumbnail"`
	Images      []struct {
		Image string `json:"image"`
	} `json:"images"`
	WorkUrl  string `json:"work_url"`
	MovieUrl string `json:"movie_url"`
	Tags     []struct {
		Tag string `json:"tag"`
	} `json:"tags"`
	GroupID  string `json:"group_id"`
	Security int    `json:"security"`
}
