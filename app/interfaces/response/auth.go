package response

type Token struct {
	Token string `json:"token"`
}

type UserID struct {
	UserID string `json:"userID"`
}

type SignInResponse struct {
	UserID string `json:"userID"`
	Token  string `json:"token"`
}
