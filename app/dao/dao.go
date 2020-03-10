package dao

import (
	"github.com/DualVectorFoil/stem/dao/image"
	"github.com/DualVectorFoil/stem/dao/user"
)

var UserDao = user.NewUserDao()

var ImageDao = image.NewImageDao()
