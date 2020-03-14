package model

// for remote account service
type VerifyTokenRespModel struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Code    int    `json:"code"`
	Token   struct {
		Data struct {
			Email    string `json:"email"`
			ID       string `json:"id"`
			ClientID string `json:"clientId"`
		} `json:"data"`
		Iat int64 `json:"iat"`
		Exp int64 `json:"exp"`
	} `json:"token"`
}
