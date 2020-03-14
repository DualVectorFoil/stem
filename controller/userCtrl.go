package controller

import (
	"github.com/DualVectorFoil/stem/app/conf"
	"github.com/DualVectorFoil/stem/dao"
	"github.com/DualVectorFoil/stem/formatter"
	"github.com/DualVectorFoil/stem/util/jsonUtil"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type UserCtrl struct {
	UserDao dao.UserDao
}

func NewUserCtrl(userDao dao.UserDao) *UserCtrl {
	return &UserCtrl{
		UserDao: userDao,
	}
}

func (ctrl *UserCtrl) Login(ctx *gin.Context) {
	phoneNum := ctx.PostForm("phone_num")
	userName := ctx.PostForm("user_name")
	password := ctx.PostForm("password")
	token := ctx.GetHeader(conf.AUTHORIZATION_KEY)
	if phoneNum == "" && userName == "" {
		logrus.WithFields(logrus.Fields{
			"phone_num": phoneNum,
			"user_name": userName,
		}).Error("Uncorrected username")
		ctx.String(http.StatusInternalServerError, jsonUtil.JsonResp(http.StatusInternalServerError, nil, "Uncorrected username"))
		return
	}
	if password == "" && token == "" {
		logrus.WithFields(logrus.Fields{
			"phone_num": phoneNum,
			"user_name": userName,
			"password":  "*",
			"token":     token,
		}).Error("Uncorrected login info")
		ctx.String(http.StatusInternalServerError, jsonUtil.JsonResp(http.StatusInternalServerError, nil, "Uncorrected login info"))
		return
	}

	if token != "" && ctrl.loginWithToken(ctx, userName, phoneNum, token) || password != "" && ctrl.loginWithPwd(ctx, userName, phoneNum, password) {
		return
	}
	ctx.String(http.StatusNonAuthoritativeInfo, jsonUtil.JsonResp(http.StatusNonAuthoritativeInfo, nil, "Uncorrect login info"))
}

func (ctrl *UserCtrl) Register(ctx *gin.Context) {
	phoneNum := ctx.PostForm("phone_num")
	userName := ctx.PostForm("user_name")
	password := ctx.PostForm("password")
	avatarUrl := ctx.PostForm("avatar_url")
	if phoneNum == "" || userName == "" || password == "" {
		logrus.WithFields(logrus.Fields{
			"phone_num": phoneNum,
			"user_name": userName,
			"password":  password,
		}).Error("Uncorrected register info")
		ctx.String(http.StatusInternalServerError, jsonUtil.JsonResp(http.StatusInternalServerError, nil, "Uncorrected register info"))
		return
	}
	if avatarUrl == "" {
		logrus.WithFields(logrus.Fields{
			"avatar_url": "",
		}).Warn("avatar_url is empty")
	}

	profileModel, err := ctrl.UserDao.Register(ctx, userName, phoneNum, password, avatarUrl)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"phone_num":  phoneNum,
			"user_name":  userName,
			"password":   password,
			"avatar_url": avatarUrl,
			"err":        err.Error(),
		}).Error("login failed")
		ctx.String(http.StatusInternalServerError, jsonUtil.JsonResp(http.StatusInternalServerError, nil, err.Error()))
		return
	}

	ctx.String(http.StatusOK, jsonUtil.JsonResp(http.StatusOK, formatter.ProfileFormat(profileModel), ""))
}

func (ctrl *UserCtrl) loginWithToken(ctx *gin.Context, userName string, phoneNum string, token string) bool {
	profileModel, err := ctrl.UserDao.LoginWithToken(ctx, userName, phoneNum, token)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"phone_num": phoneNum,
			"user_name": userName,
			"token":     token,
		}).Error("login failed, err: " + err.Error())
		return false
	}

	ctx.String(http.StatusOK, jsonUtil.JsonResp(http.StatusOK, formatter.ProfileFormat(profileModel), ""))
	return true
}

func (ctrl *UserCtrl) loginWithPwd(ctx *gin.Context, userName string, phoneNum string, pwd string) bool {
	profileModel, err := ctrl.UserDao.LoginWithPwd(ctx, userName, phoneNum, pwd)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"phone_num": phoneNum,
			"user_name": userName,
			"pwd":       pwd,
		}).Error("login failed, err: " + err.Error())
		return false
	}

	ctx.String(http.StatusOK, jsonUtil.JsonResp(http.StatusOK, formatter.ProfileFormat(profileModel), ""))
	return true
}
