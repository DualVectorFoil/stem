package middleware

import (
	"github.com/DualVectorFoil/stem/account"
	"github.com/DualVectorFoil/stem/app/conf"
	"github.com/DualVectorFoil/stem/util/jsonUtil"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if _, err := account.VerifyToken(ctx.Request.Header.Get(conf.AUTHORIZATION_KEY)); err != nil {
			ctx.Abort()
			logrus.Warn(err.Error())
			ctx.String(http.StatusNonAuthoritativeInfo, jsonUtil.JsonResp(http.StatusNonAuthoritativeInfo, nil, "认证信息过期，请重新登陆"))
		}
		ctx.Next()

	}
}
