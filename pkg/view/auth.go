package view

type SignResponse struct {
	Token string `json:"token"`
}

func ReturnSignResponse(token string) SignResponse {
	body := SignResponse{
		Token: token,
	}
	return body
}
