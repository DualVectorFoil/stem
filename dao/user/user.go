package user

import (
	"context"
	"github.com/DualVectorFoil/stem/model"
	"time"
)

const MAX_REQUEST_TIME = time.Second * 6

type UserDao struct {
	// TODO client
}

func NewUserDao() *UserDao {
	return &UserDao{}
}

func (u *UserDao) LoginWithPwd(ctx context.Context, userName string, pwd string) (*model.ProfileModel, error) {
	// TODO
	return nil, nil
}

func (u *UserDao) LoginWithToken(ctx context.Context, token string) (*model.ProfileModel, error) {
	// TODO
	return nil, nil
}

func (u *UserDao) Register(ctx context.Context, userName string, pwd string, avatarUrl string) error {
	// TODO
	return nil
}
