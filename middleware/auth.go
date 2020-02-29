package middleware

import (
	"github.com/DualVectorFoil/stem/app/conf"
	"github.com/DualVectorFoil/stem/util/jsonUtil"
	"github.com/DualVectorFoil/stem/util/jwt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		auth := ctx.Request.Header.Get(conf.AUTHORIZATION_KEY)
		if _, err := jwt.ParseToken(auth); err != nil {
			ctx.Abort()
			logrus.Warn(err.Error())
			ctx.JSON(http.StatusNonAuthoritativeInfo, jsonUtil.JsonResp(http.StatusNonAuthoritativeInfo, nil, "认证信息过期，请重新登陆"))
		}
		ctx.Next()
	}
}
