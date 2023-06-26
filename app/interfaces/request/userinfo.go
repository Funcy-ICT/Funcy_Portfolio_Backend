package request

type UserInfo struct {
	Icon            string   `json:"icon" validate:"required"`
	HeaderImagePath string   `json:"header" validate:"required"`
	Bio             string   `json:"user_description" validate:"required"`
	SNS             []string `json:"sns" validate:"required"`
	Group           []string `json:"group" validate:"required"`
	Skills          []string `json:"skills" validate:"required"`
	DisplayName     string   `json:"displayName" validate:"required"`
	Works           []struct {
		WorkID      string `json:"workID" validate:"required"`
		Title       string `json:"title" validate:"required"`
		Image       string `json:"image" validate:"required"`
		Description string `json:"description" validate:"required"`
		Icon        string `json:"icon" validate:"required"`
	}
}
