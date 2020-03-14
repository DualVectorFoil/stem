package user

import (
	"errors"
	"github.com/DualVectorFoil/stem/account"
	"github.com/DualVectorFoil/stem/model"
	"github.com/gin-gonic/gin"
	"time"
)

type UserDaoImpl struct {
	// TODO account
}

func NewUserDao() *UserDaoImpl {
	return &UserDaoImpl{}
}

func (u *UserDaoImpl) LoginWithPwd(ctx *gin.Context, userName string, phoneNum string, pwd string) (*model.ProfileModel, error) {
	loginMutation, err := account.LoginWithPwd(phoneNum, userName, pwd, "0", "")
	if err != nil {
		return nil, err
	}

	// TODO get real profile model in mysql with grpc
	// TODO remove
	return &model.ProfileModel{
		ID:           string(loginMutation.Login.ID),
		PhoneNum:     phoneNum,
		UserName:     userName,
		PWD:          pwd,
		AvatarUrl:    "http://139.155.46.62:4869/90405e5063c8166c1f5ae29c746daeff",
		RegisteredAt: time.Now().Unix(),
		LastLoginAt:  time.Now().Unix(),
		Token:        string(loginMutation.Login.Token),
	}, nil
}

func (u *UserDaoImpl) LoginWithToken(ctx *gin.Context, userName string, phoneNum string, token string) (*model.ProfileModel, error) {
	id, err := account.VerifyToken(token)
	if err != nil {
		return nil, err
	}
	if len(id) == 0 {
		return nil, errors.New("verify token failed")
	}

	// TODO get real profile model in mysql with grpc, and login with pwd for new refreshed token
	// TODO remove
	return &model.ProfileModel{
		ID:           id,
		PhoneNum:     phoneNum,
		UserName:     userName,
		PWD:          token,
		AvatarUrl:    "http://139.155.46.62:4869/812ac00b18c08db8e2773a63f933dd6f",
		RegisteredAt: time.Now().Unix(),
		LastLoginAt:  time.Now().Unix(),
		Token:        token,
	}, nil
}

func (u *UserDaoImpl) Register(ctx *gin.Context, userName string, phoneNum string, pwd string, avatarUrl string) (*model.ProfileModel, error) {
	registerMutation, err := account.Register(phoneNum, pwd)
	if err != nil {
		return nil, err
	}

	// TODO grpc to save register profile info and return
	// TODO remove
	return &model.ProfileModel{
		ID:           string(registerMutation.Register.ID),
		PhoneNum:     phoneNum,
		UserName:     userName,
		PWD:          pwd,
		AvatarUrl:    avatarUrl,
		RegisteredAt: time.Now().Unix(),
		LastLoginAt:  0,
		Token:        string(registerMutation.Register.Token),
	}, nil
}
