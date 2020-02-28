package router

import (
	"github.com/DualVectorFoil/stem/app/handler"
	"github.com/DualVectorFoil/stem/util"
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

	router.Run(":30534")
}

func notFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, util.JsonResp(http.StatusNotFound, nil, "404 router not found"))
}
