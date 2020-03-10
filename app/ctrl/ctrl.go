package ctrl

import (
	"github.com/DualVectorFoil/stem/app/dao"
	"github.com/DualVectorFoil/stem/controller"
)

var UserCtrl = controller.NewUserCtrl(dao.UserDao)

var ImageCtrl = controller.NewImageCtrl(dao.ImageDao)
