package response

type UserInfo struct {
	Icon            string      `json:"icon"`
	HeaderImagePath string      `json:"header"`
	UserDescription string      `json:"user_description"`
	SNS             []string    `json:"sns"`
	Group           []string    `json:"group"`
	Skills          []string    `json:"skills"`
	DisplayName     string      `json:"displayName"`
	Works           []ReadWorks `json:"works"`
}
