package auth

type SendCodeRequest struct {
	Phone string `json:"phone"`
}

type SendCodeResponse struct {
	Code      string `json:"code"`
	SessionID string `json:"session_id"`
}

type VerifyCodeRequest struct {
	SessionID string `json:"session_id"`
	Code      string `json:"code"`
}

type VerifyCodeResponse struct {
	Token string `json:"token"`
}
