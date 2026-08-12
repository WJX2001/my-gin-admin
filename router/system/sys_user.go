package system

import (
	//"github.com/WJX2001/gin-vue-admin-server/api/v1/system"
	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (s *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("user")
	{
		userRouter.POST("admin_register", baseApi.Register)
	}
}
