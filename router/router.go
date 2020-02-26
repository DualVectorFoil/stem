package router

import (
	"github.com/DualVectorFoil/stem/app/handler"
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
}

func notFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"status": 404,
		"error":  "404 page not found",
	})
}
