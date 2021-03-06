package dao

import (
	"github.com/DualVectorFoil/stem/model"
	"github.com/gin-gonic/gin"
)

type ImageDao interface {
	GetImages(ctx *gin.Context, itemCount int, userId int64) ([]*model.ImageDetailModel, error)
}
