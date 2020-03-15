package user

import (
	"context"
	"errors"
	"github.com/DualVectorFoil/stem/account"
	"github.com/DualVectorFoil/stem/app/conf"
	"github.com/DualVectorFoil/stem/model"
	"github.com/DualVectorFoil/stem/pb"
	"github.com/DualVectorFoil/stem/util/ptr"
	"github.com/gin-gonic/gin"
)

type UserDaoImpl struct {
	AccountClient pb.AccountServiceClient
}

func NewUserDao(accountClient pb.AccountServiceClient) *UserDaoImpl {
	return &UserDaoImpl{
		AccountClient: accountClient,
	}
}

// TODO login with phoneCode

func (u *UserDaoImpl) LoginWithPwd(ctx *gin.Context, userName string, phoneNum string, pwd string) (*model.ProfileModel, error) {
	loginMutation, err := account.LoginWithPwd(phoneNum, userName, pwd, "0", "")
	if err != nil {
		return nil, err
	}

	timeoutCtx, _ := context.WithTimeout(ctx.Request.Context(), conf.USER_INFO_TTL)
	resp, err := u.AccountClient.Login(timeoutCtx, &pb.LoginRequest{
		Id:        ptr.StringPtr(string(loginMutation.Login.ID)),
		PhoneNum:  ptr.StringPtr(phoneNum),
		UserName:  ptr.StringPtr(userName),
		Email:     ptr.StringPtr(string(loginMutation.Login.Email)),
		Pwd:       ptr.StringPtr(pwd),
		PhoneCode: ptr.Int32Ptr(0),
	})
	if err != nil {
		return nil, err
	}

	profile := resp.GetProfile()
	if !resp.GetIsSuccess() || profile == nil {
		return nil, errors.New("gprc error, LoginWithPwd failed")
	}

	return transformProfile(profile, string(loginMutation.Login.Token)), nil
}

func (u *UserDaoImpl) LoginWithToken(ctx *gin.Context, userName string, phoneNum string, token string) (*model.ProfileModel, error) {
	id, err := account.VerifyToken(token)
	if err != nil {
		return nil, err
	}
	if len(id) == 0 {
		return nil, errors.New("verify token failed")
	}

	timeoutCtx, _ := context.WithTimeout(ctx.Request.Context(), conf.USER_INFO_TTL)
	resp, err := u.AccountClient.Login(timeoutCtx, &pb.LoginRequest{
		Id:        ptr.StringPtr(id),
		PhoneNum:  ptr.StringPtr(phoneNum),
		UserName:  ptr.StringPtr(userName),
		Email:     ptr.StringPtr(""),
		Pwd:       ptr.StringPtr(""),
		PhoneCode: ptr.Int32Ptr(0),
	})
	if err != nil {
		return nil, err
	}

	profile := resp.GetProfile()
	if !resp.GetIsSuccess() || profile == nil {
		return nil, errors.New("gprc error, LoginWithToken failed")
	}

	refreshTokenProfileMutation, err := account.LoginWithPwd(profile.GetPhoneNum(), profile.GetUserName(), profile.GetPwd(), "0", profile.GetEmail())
	if err != nil {
		return nil, err
	}

	resp, err = u.AccountClient.Login(timeoutCtx, &pb.LoginRequest{
		Id:        ptr.StringPtr(string(refreshTokenProfileMutation.Login.ID)),
		PhoneNum:  ptr.StringPtr(phoneNum),
		UserName:  ptr.StringPtr(string(refreshTokenProfileMutation.Login.Username)),
		Email:     ptr.StringPtr(string(refreshTokenProfileMutation.Login.Email)),
		Pwd:       ptr.StringPtr(""),
		PhoneCode: ptr.Int32Ptr(0),
	})
	if err != nil {
		return nil, err
	}

	profile = resp.GetProfile()
	if !resp.GetIsSuccess() || profile == nil {
		return nil, errors.New("gprc error, LoginWithToken failed")
	}

	return transformProfile(profile, string(refreshTokenProfileMutation.Login.Token)), nil
}

func (u *UserDaoImpl) Register(ctx *gin.Context, userName string, phoneNum string, pwd string, avatarUrl string) (*model.ProfileModel, error) {
	registerMutation, err := account.Register(phoneNum, pwd)
	if err != nil {
		return nil, err
	}

	timeoutCtx, _ := context.WithTimeout(ctx.Request.Context(), conf.USER_INFO_TTL)
	resp, err := u.AccountClient.Register(timeoutCtx, &pb.RegisterRequest{
		Id:        ptr.StringPtr(string(registerMutation.Register.ID)),
		PhoneNum:  ptr.StringPtr(string(registerMutation.Register.Phone)),
		UserName:  ptr.StringPtr(userName),
		Email:     ptr.StringPtr(string(registerMutation.Register.Email)),
		Pwd:       ptr.StringPtr(pwd),
		PhoneCode: ptr.Int32Ptr(0),
		AvatarUrl: ptr.StringPtr(avatarUrl),
	})
	if err != nil {
		account.DeleteAccounts([]string{string(registerMutation.Register.ID)})
		return nil, err
	}

	profile := resp.GetProfile()
	if !resp.GetIsSuccess() || profile == nil {
		account.DeleteAccounts([]string{string(registerMutation.Register.ID)})
		return nil, errors.New("gprc error, register failed")
	}

	return u.LoginWithPwd(ctx, profile.GetUserName(), profile.GetPhoneNum(), profile.GetPwd())
}

func transformProfile(profile *pb.Profile, token string) *model.ProfileModel {
	return &model.ProfileModel{
		ID:           profile.GetId(),
		PhoneNum:     profile.GetPhoneNum(),
		UserName:     profile.GetUserName(),
		PWD:          profile.GetPwd(),
		AvatarUrl:    profile.GetAvatarUrl(),
		RegisteredAt: profile.GetRegisteredAt(),
		LastLoginAt:  profile.GetLastLoginAt(),
		Token:        token,
	}
}
