package request

// アカウント作成リクエスト
type SignUpRequest struct {
	Icon        string `json:"icon" validate:"required"`
	FamilyName  string `json:"familyName" validate:"required"`
	FirstName   string `json:"firstName" validate:"required"`
	Mail        string `json:"mail" validate:"required"`
	Password    string `json:"password" validate:"required"`
	Grade       string `json:"grade" validate:"required"`
	Course      string `json:"course" validate:"required"`
	DisplayName string `json:"displayName" validate:"required"`
}

// アカウント認証リクエスト
type SignInRequest struct {
	Mail     string `json:"mail" validate:"required"`
	Password string `json:"password" validate:"required"`
}
