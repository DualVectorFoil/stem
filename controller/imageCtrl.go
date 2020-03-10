package controller

import (
	"github.com/DualVectorFoil/stem/dao"
	"github.com/DualVectorFoil/stem/formatter"
	"github.com/DualVectorFoil/stem/util/jsonUtil"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type ImageCtrl struct {
	ImageDao dao.IImageDao
}

func NewImageCtrl(imageDao dao.IImageDao) *ImageCtrl {
	return &ImageCtrl{
		ImageDao: imageDao,
	}
}

func (i *ImageCtrl) GetImages(ctx *gin.Context) {
	images, err := i.ImageDao.GetImages(ctx, 0, 0)
	if err != nil {
		logrus.Error("something wrong to get images")
		return
	}
	ctx.String(http.StatusOK, jsonUtil.JsonResp(http.StatusOK, formatter.ImageFormat(images), ""))
}
