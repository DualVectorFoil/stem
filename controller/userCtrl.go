package controller

import (
	"github.com/DualVectorFoil/stem/dao"
	"github.com/gin-gonic/gin"
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

}

func (ctrl *UserCtrl) Register(ctx *gin.Context) {

}
