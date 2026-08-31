package system

import "github.com/WJX2001/gin-vue-admin-server/service"

type ApiGroup struct {
	BaseApi
}

var (
	userService           = service.ServiceGroupApp.SystemServiceGroup.UserService
	jwtService            = service.ServiceGroupApp.SystemServiceGroup.JwtService
	securityConfigService = service.ServiceGroupApp.SystemServiceGroup.SecurityConfigService
	loginLogService       = service.ServiceGroupApp.SystemServiceGroup.LoginLogService
)
