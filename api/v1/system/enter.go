package system

import "github.com/WJX2001/gin-vue-admin-server/service"

var (
	userService           = service.ServiceGroupApp.SystemServiceGroup.UserService
	securityConfigService = service.ServiceGroupApp.SystemServiceGroup.SecurityConfigService
)
