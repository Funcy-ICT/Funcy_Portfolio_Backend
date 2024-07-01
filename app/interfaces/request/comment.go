package request

type CreateCommentRequest struct {
	UserID  string `json:"user_id"`
	WorksID string `json:"works_id"`
	Text    string `json:"text"`
}
