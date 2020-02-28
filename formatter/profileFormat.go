package formatter

import "github.com/DualVectorFoil/stem/model"

func ProfileFormat(model *model.ProfileModel) *ProfileResp {
	if model == nil {
		return nil
	}

	return &ProfileResp{
		UserName:     model.UserName,
		PWD:          model.PWD,
		AvatarUrl:    model.AvatarUrl,
		RegisteredAt: model.RegisteredAt,
		LastLoginAt:  model.LastLoginAt,
		Token:        model.Token,
	}
}

type ProfileResp struct {
	UserName     string `json:"user_name"`
	PWD          string `json:"pwd"`
	AvatarUrl    string `json:"avatar_url"`
	RegisteredAt int64  `json:"registered_at"`
	LastLoginAt  int64  `json:"last_login_at"`
	Token        string `json:"token"`
}
