package controller

import (
	"github.com/DualVectorFoil/stem/dao"
	"github.com/DualVectorFoil/stem/formatter"
	"github.com/DualVectorFoil/stem/util"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type UserCtrl struct {
	UserDao dao.IUserDao
}

func NewUserCtrl(userDao dao.IUserDao) *UserCtrl {
	return &UserCtrl{
		UserDao: userDao,
	}
}

func (ctrl *UserCtrl) Login(ctx *gin.Context) {
	userName := ctx.PostForm("user_name")
	password := ctx.PostForm("password")
	token := ctx.PostForm("token")
	if userName == "" {
		logrus.WithFields(logrus.Fields{
			"user_name": userName,
		}).Error("Uncorrected username")
		ctx.JSON(http.StatusInternalServerError, util.JsonResp(http.StatusInternalServerError, nil, "Uncorrected username"))
		return
	}
	if password == "" && token == "" {
		logrus.WithFields(logrus.Fields{
			"user_name": userName,
			"password":  "*",
			"token": token,
		}).Error("Uncorrected login info")
		ctx.JSON(http.StatusInternalServerError, util.JsonResp(http.StatusInternalServerError, nil, "Uncorrected login info"))
		return
	}

	if token != "" {
		ctrl.loginWithToken(ctx, userName, token)
	} else {
		ctrl.loginWithPwd(ctx, userName, password)
	}
}

func (ctrl *UserCtrl) Register(ctx *gin.Context) {
	userName := ctx.PostForm("user_name")
	password := ctx.PostForm("password")
	avatarUrl := ctx.PostForm("avatar_url")
	if userName == "" || password == "" {
		logrus.WithFields(logrus.Fields{
			"user_name": userName,
			"password":  "*",
		}).Error("Uncorrected register info")
		ctx.JSON(http.StatusInternalServerError, util.JsonResp(http.StatusInternalServerError, nil, "Uncorrected register info"))
		return
	}
	if avatarUrl == "" {
		logrus.WithFields(logrus.Fields{
			"avatar_url": "",
		}).Warn("avatar_url is empty")
	}

	profileModel, err := ctrl.UserDao.Register(ctx, userName, password, avatarUrl)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"user_name":  userName,
			"password":   "*",
			"avatar_url": "",
			"err":        err.Error(),
		}).Error("login failed")
		ctx.JSON(http.StatusInternalServerError, util.JsonResp(http.StatusInternalServerError, nil, err.Error()))
		return
	}

	ctx.String(http.StatusOK, util.JsonResp(http.StatusOK, formatter.ProfileFormat(profileModel)))
}

func (ctrl *UserCtrl) loginWithToken(ctx *gin.Context, userName string, token string) {
	profileModel, err := ctrl.UserDao.LoginWithToken(ctx, userName, token)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"user_name":  userName,
			"token":   token,
			"err":        err.Error(),
		}).Error("login failed")
		ctx.JSON(http.StatusNonAuthoritativeInfo, util.JsonResp(http.StatusNonAuthoritativeInfo, nil, err.Error()))
		return
	}

	ctx.String(http.StatusOK, util.JsonResp(http.StatusOK, formatter.ProfileFormat(profileModel)))
}

func (ctrl *UserCtrl) loginWithPwd(ctx *gin.Context, userName string, pwd string) {
	profileModel, err := ctrl.UserDao.LoginWithPwd(ctx, userName, pwd)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"user_name":  userName,
			"pwd":   pwd,
			"err":        err.Error(),
		}).Error("login failed")
		ctx.JSON(http.StatusNonAuthoritativeInfo, util.JsonResp(http.StatusNonAuthoritativeInfo, nil, err.Error()))
		return
	}

	ctx.String(http.StatusOK, util.JsonResp(http.StatusOK, formatter.ProfileFormat(profileModel)))
}