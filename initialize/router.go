package initialize

import (
	"github.com/WJX2001/gin-vue-admin-server/global"
	"github.com/WJX2001/gin-vue-admin-server/router"
	"github.com/gin-gonic/gin"
)

// 初始化总路由
func Routers() *gin.Engine {
	Router := gin.New()
	systemRouter := router.RouterGroupApp.System

	PrivateGroup := Router.Group(global.GVA_CONFIG.System.RouterPrefix)
	{
		systemRouter.InitUserRouter(PrivateGroup)
	}
	return Router
}
