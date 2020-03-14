package model

type RefreshTokenModel struct {
	Data struct {
		RefreshToken struct {
			Token string `json:"token"`
			Iat   int    `json:"iat"`
			Exp   int    `json:"exp"`
		} `json:"refreshToken"`
	} `json:"data"`
}
