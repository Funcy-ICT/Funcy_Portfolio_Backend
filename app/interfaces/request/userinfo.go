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

type UpdateUserInfo struct {
	Icon            string   `json:"icon"`
	HeaderImagePath string   `json:"header"`
	Bio             string   `json:"user_description"`
	SNS             []string `json:"sns"`
	Skills          []string `json:"skills"`
	DisplayName     string   `json:"displayName"`
	FamilyName      string   `json:"familyName"`
	FirstName       string   `json:"firstName"`
	Grade           string   `json:"grade"`
	Course          string   `json:"course"`
}
