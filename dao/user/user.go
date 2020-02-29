package user

import (
	"github.com/DualVectorFoil/stem/model"
	"github.com/DualVectorFoil/stem/util/jwt"
	"github.com/gin-gonic/gin"
	"time"
)

const MAX_REQUEST_TIME = time.Second * 6

type UserDao struct {
	// TODO client
}

func NewUserDao() *UserDao {
	return &UserDao{}
}

func (u *UserDao) LoginWithPwd(ctx *gin.Context, userName string, pwd string) (*model.ProfileModel, error) {
	token, err := jwt.CreateToken(0, userName)
	if err != nil {
		return nil, err
	}

	return &model.ProfileModel{
		UserName:     userName,
		PWD:          pwd,
		AvatarUrl:    "http://139.155.46.62:4869/90405e5063c8166c1f5ae29c746daeff",
		RegisteredAt: time.Now().Unix(),
		LastLoginAt:  time.Now().Unix(),
		Token:        token,
	}, nil
}

func (u *UserDao) LoginWithToken(ctx *gin.Context, userName string, token string) (*model.ProfileModel, error) {
	jwtToken, err := jwt.CreateToken(0, userName)
	if err != nil {
		return nil, err
	}

	return &model.ProfileModel{
		UserName:     userName,
		PWD:          token,
		AvatarUrl:    "http://139.155.46.62:4869/812ac00b18c08db8e2773a63f933dd6f",
		RegisteredAt: time.Now().Unix(),
		LastLoginAt:  time.Now().Unix(),
		Token:        jwtToken,
	}, nil
}

func (u *UserDao) Register(ctx *gin.Context, userName string, pwd string, avatarUrl string) (*model.ProfileModel, error) {
	return &model.ProfileModel{
		UserName:     userName,
		PWD:          pwd,
		AvatarUrl:    avatarUrl,
		RegisteredAt: time.Now().Unix(),
		LastLoginAt:  0,
		Token:        "ffffffffff",
	}, nil
}
