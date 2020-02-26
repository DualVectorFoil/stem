package model

// TODO should transform userModel to profileModel when response
type ProfileModel struct {
	UserName     string `json:"user_name"`
	PWD          string `json:"pwd"`
	AvatarUrl    string `json:"avatar_url"`
	RegisteredAt int64  `json:"registered_at"`
	LastLoginAt  int64  `json:"last_login_at"`
	Token        string `json:"token"`
}
