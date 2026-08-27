package initialize

import (
	"github.com/WJX2001/gin-vue-admin-server/global"
	"github.com/WJX2001/gin-vue-admin-server/middleware"
	"github.com/WJX2001/gin-vue-admin-server/router"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 初始化总路由
func Routers() *gin.Engine {
	Router := gin.New()
	systemRouter := router.RouterGroupApp.System

	PublicGroup := Router.Group(global.GVA_CONFIG.System.RouterPrefix)
	PrivateGroup := Router.Group(global.GVA_CONFIG.System.RouterPrefix)

	PrivateGroup.Use(middleware.JWTAuth()).Use(middleware.MustChangePwdGuard())
	{
		// 健康检测
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, "ok")
		})
	}
	{
		systemRouter.InitUserRouter(PrivateGroup)
	}
	return Router
}
