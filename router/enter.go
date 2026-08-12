package router

import "github.com/WJX2001/gin-vue-admin-server/router/system"

var RouterGroupApp = new(RouterGroup)

type RouterGroup struct {
	System system.RouterGroup
}
