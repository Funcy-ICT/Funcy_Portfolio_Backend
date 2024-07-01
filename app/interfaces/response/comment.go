package response

type CreateCommentResponse struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	ID         string `json:"id,omitempty"`
}
