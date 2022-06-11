package dto

//Request struct

//アカウント作成リクエスト
type SignUpRequest struct {
	Icon        string `json:"icon"`
	FamilyName  string `json:"FamilyName"`
	FirstName   string `json:"firstName"`
	Mail        string `json:"mail"`
	Password    string `json:"password"`
	Grade       string `json:"grade"`
	Course      string `json:"course"`
	DisplayName string `json:"displayName"`
}

//アカウント認証リクエスト
type SignInRequest struct {
	Mail     string `json:"mail"`
	Password string `json:"password"`
}
