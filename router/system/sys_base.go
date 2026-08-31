package system

import (
	"github.com/WJX2001/gin-vue-admin-server/middleware"
	"github.com/gin-gonic/gin"
)

type BaseRouter struct{}

func (s *BaseRouter) InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	baseRouter := Router.Group("base")
	{
		baseRouter.POST("login", middleware.SecurityLimit(), baseApi.Login)
	}
	return baseRouter
}
