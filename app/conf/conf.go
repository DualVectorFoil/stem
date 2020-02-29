package conf

import "time"

const (
	// jwt token
	JWT_SECRET        = "pith_and_love"
	USER_INFO_TTL     = time.Hour * 240
	AUTHORIZATION_KEY = "Authorization"
)
