package router

import (
	"github.com/DualVectorFoil/stem/app/handler"
	"github.com/DualVectorFoil/stem/middleware"
	"github.com/DualVectorFoil/stem/util/jsonUtil"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Init() {
	h := handler.NewHandlerInstance()

	router := gin.Default()
	router.NoRoute()

	userRouter := router.Group("/user")
	userRouter.POST("/login", h.UserCtrl.Login)
	userRouter.POST("register", h.UserCtrl.Register)
	//userRouter.POST("modify_pwd", )

	//privateRouter := router.Group("/private", middleware.AuthMiddleware())
	imageRouter := router.Group("/images", middleware.AuthMiddleware())
	imageRouter.POST("get_images", h.ImageCtrl.GetImages)

	router.Run(":30534")
}

func notFound(c *gin.Context) {
	c.String(http.StatusNotFound, jsonUtil.JsonResp(http.StatusNotFound, nil, "404 router not found"))
}
