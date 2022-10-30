package request

type CreateWorkRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Images      []struct {
		Image string `json:"image"`
	} `json:"images"`
	Work_URL  string `json:"work_url"`
	Movie_url string `json:"movie_url"`
	Tags      []struct {
		Tag string `json:"tag"`
	} `json:"tags"`
	Group    string `json:"group"`
	Security int    `json:"security"`
}
