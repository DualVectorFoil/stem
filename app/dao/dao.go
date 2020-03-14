package dao

import (
	"github.com/DualVectorFoil/stem/app/client"
	"github.com/DualVectorFoil/stem/dao/image"
	"github.com/DualVectorFoil/stem/dao/user"
)

var UserDao = user.NewUserDao(client.AccountClient)

var ImageDao = image.NewImageDao()
