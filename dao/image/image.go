package image

import (
	"github.com/DualVectorFoil/stem/model"
	"github.com/gin-gonic/gin"
	"time"
)

const MAX_REQUEST_TIME = time.Second * 6

type ImageDao struct {
	// TODO account
}

func NewImageDao() *ImageDao {
	return &ImageDao{}
}

// TODO
func (i *ImageDao) GetImages(ctx *gin.Context, itemCount int, userId int64) ([]*model.ImageDetailModel, error) {
	return []*model.ImageDetailModel{
		&model.ImageDetailModel{
			Url:    "httpManager://139.155.46.62:4869/90405e5063c8166c1f5ae29c746daeff?w=524&h=320",
			Width:  530,
			Height: 320,
		},
		&model.ImageDetailModel{
			Url:    "httpManager://139.155.46.62:4869/90405e5063c8166c1f5ae29c746daeff?w=524&h=320",
			Width:  530,
			Height: 320,
		},
		&model.ImageDetailModel{
			Url:    "httpManager://139.155.46.62:4869/90405e5063c8166c1f5ae29c746daeff?w=524&h=320",
			Width:  530,
			Height: 320,
		},
		&model.ImageDetailModel{
			Url:    "httpManager://139.155.46.62:4869/90405e5063c8166c1f5ae29c746daeff?w=524&h=320",
			Width:  530,
			Height: 320,
		},
		&model.ImageDetailModel{
			Url:    "httpManager://139.155.46.62:4869/90405e5063c8166c1f5ae29c746daeff?w=524&h=320",
			Width:  530,
			Height: 320,
		},
		&model.ImageDetailModel{
			Url:    "httpManager://139.155.46.62:4869/90405e5063c8166c1f5ae29c746daeff?w=524&h=320",
			Width:  530,
			Height: 320,
		},
		&model.ImageDetailModel{
			Url:    "httpManager://139.155.46.62:4869/90405e5063c8166c1f5ae29c746daeff?w=524&h=320",
			Width:  530,
			Height: 320,
		},
		&model.ImageDetailModel{
			Url:    "httpManager://139.155.46.62:4869/90405e5063c8166c1f5ae29c746daeff?w=524&h=320",
			Width:  530,
			Height: 320,
		},
		&model.ImageDetailModel{
			Url:    "httpManager://139.155.46.62:4869/90405e5063c8166c1f5ae29c746daeff?w=524&h=320",
			Width:  530,
			Height: 320,
		},
		&model.ImageDetailModel{
			Url:    "httpManager://139.155.46.62:4869/90405e5063c8166c1f5ae29c746daeff?w=524&h=320",
			Width:  530,
			Height: 320,
		},
		&model.ImageDetailModel{
			Url:    "httpManager://139.155.46.62:4869/90405e5063c8166c1f5ae29c746daeff?w=524&h=320",
			Width:  530,
			Height: 320,
		},
		&model.ImageDetailModel{
			Url:    "httpManager://139.155.46.62:4869/90405e5063c8166c1f5ae29c746daeff?w=524&h=320",
			Width:  530,
			Height: 320,
		},
		&model.ImageDetailModel{
			Url:    "httpManager://139.155.46.62:4869/90405e5063c8166c1f5ae29c746daeff?w=524&h=320",
			Width:  530,
			Height: 320,
		},
		&model.ImageDetailModel{
			Url:    "httpManager://139.155.46.62:4869/90405e5063c8166c1f5ae29c746daeff?w=524&h=320",
			Width:  530,
			Height: 320,
		},
		&model.ImageDetailModel{
			Url:    "httpManager://139.155.46.62:4869/90405e5063c8166c1f5ae29c746daeff?w=524&h=320",
			Width:  530,
			Height: 320,
		},
		&model.ImageDetailModel{
			Url:    "httpManager://139.155.46.62:4869/90405e5063c8166c1f5ae29c746daeff?w=524&h=320",
			Width:  530,
			Height: 320,
		},
		&model.ImageDetailModel{
			Url:    "httpManager://139.155.46.62:4869/90405e5063c8166c1f5ae29c746daeff?w=524&h=320",
			Width:  530,
			Height: 320,
		},
		&model.ImageDetailModel{
			Url:    "httpManager://139.155.46.62:4869/90405e5063c8166c1f5ae29c746daeff?w=524&h=320",
			Width:  530,
			Height: 320,
		},
		&model.ImageDetailModel{
			Url:    "httpManager://139.155.46.62:4869/90405e5063c8166c1f5ae29c746daeff?w=524&h=320",
			Width:  530,
			Height: 320,
		},
		&model.ImageDetailModel{
			Url:    "httpManager://139.155.46.62:4869/90405e5063c8166c1f5ae29c746daeff?w=524&h=320",
			Width:  530,
			Height: 320,
		},
		&model.ImageDetailModel{
			Url:    "httpManager://139.155.46.62:4869/90405e5063c8166c1f5ae29c746daeff?w=524&h=320",
			Width:  530,
			Height: 320,
		},
		&model.ImageDetailModel{
			Url:    "httpManager://139.155.46.62:4869/90405e5063c8166c1f5ae29c746daeff?w=524&h=320",
			Width:  530,
			Height: 320,
		},
		&model.ImageDetailModel{
			Url:    "httpManager://139.155.46.62:4869/90405e5063c8166c1f5ae29c746daeff?w=524&h=320",
			Width:  530,
			Height: 320,
		},
		&model.ImageDetailModel{
			Url:    "httpManager://139.155.46.62:4869/90405e5063c8166c1f5ae29c746daeff?w=524&h=320",
			Width:  530,
			Height: 320,
		},
		&model.ImageDetailModel{
			Url:    "httpManager://139.155.46.62:4869/90405e5063c8166c1f5ae29c746daeff?w=524&h=320",
			Width:  530,
			Height: 320,
		},
		&model.ImageDetailModel{
			Url:    "httpManager://139.155.46.62:4869/90405e5063c8166c1f5ae29c746daeff?w=524&h=320",
			Width:  530,
			Height: 320,
		},
		&model.ImageDetailModel{
			Url:    "httpManager://139.155.46.62:4869/90405e5063c8166c1f5ae29c746daeff?w=524&h=320",
			Width:  530,
			Height: 320,
		},
	}, nil
}
