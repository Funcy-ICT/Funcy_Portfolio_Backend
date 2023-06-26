package request

type UserInfo struct {
	Icon            string   `json:"icon"`
	HeaderImagePath string   `json:"header"`
	Bio             string   `json:"bio"`
	SNS             []string `json:"sns"`
	Group           []string `json:"group"`
	Skills          []string `json:"skills"`
	DisplayName     string   `json:"displayName"`
	Works           []struct {
		WorkID      string `json:"workID"`
		Title       string `json:"title"`
		Image       string `json:"image"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	}
}
