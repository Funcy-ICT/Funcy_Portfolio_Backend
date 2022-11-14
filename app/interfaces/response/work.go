package response

type WorkID struct {
	WorkID string `json:"workID"`
}

type Work struct {
	WorkID      string `json:"workID"`
	Title       string `json:"title"`
	Image       string `json:"image"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type Works struct {
	Works []Work `json:"works"`
}
