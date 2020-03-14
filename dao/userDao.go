package dao

import (
	"github.com/DualVectorFoil/stem/model"
	"github.com/gin-gonic/gin"
)

type UserDao interface {
	// TODO login with phone verifycode
	LoginWithPwd(ctx *gin.Context, userName string, phoneNum string, pwd string) (*model.ProfileModel, error)
	LoginWithToken(ctx *gin.Context, userName string, phoneNum string, token string) (*model.ProfileModel, error)
	Register(ctx *gin.Context, userName string, phoneNum string, pwd string, avatarUrl string) (*model.ProfileModel, error)
}
