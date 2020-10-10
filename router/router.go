package router

import (
	"github.com/bluebell/controller"
	"github.com/bluebell/middlewares"
	"github.com/bluebell/settings"

	"github.com/bluebell/logger"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	//如果配置是release，就使用release模式启动，否则使用debug模式（默认）
	//注意这一行不能写在gin.New()之后
	if settings.Conf.Mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	// 使用token校验插件
	r.Use(middlewares.JWTAuthMiddleware())

	r.GET("/index", controller.IndexHandler)

	return r
}
