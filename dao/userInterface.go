package dao

import (
	"context"
	"github.com/DualVectorFoil/stem/model"
)

type IUserDao interface {
	LoginWithPwd(ctx context.Context, userName string, pwd string) (*model.ProfileModel, error)
	LoginWithToken(ctx context.Context, token string) (*model.ProfileModel, error)

	Register(ctx context.Context, userName string, pwd string, avatarUrl string) error
}
