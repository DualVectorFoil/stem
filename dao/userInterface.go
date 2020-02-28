package dao

import (
	"github.com/DualVectorFoil/stem/model"
	"github.com/gin-gonic/gin"
)

type IUserDao interface {
	LoginWithPwd(ctx *gin.Context, userName string, pwd string) (*model.ProfileModel, error)
	LoginWithToken(ctx *gin.Context, userName string, token string) (*model.ProfileModel, error)

	Register(ctx *gin.Context, userName string, pwd string, avatarUrl string) (*model.ProfileModel, error)
}
