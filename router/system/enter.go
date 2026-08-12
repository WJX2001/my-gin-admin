package system

import api "github.com/WJX2001/gin-vue-admin-server/api/v1"

type RouterGroup struct {
	UserRouter
}

var (
	baseApi = api.ApiGroupApp.SystemApiGroup.BaseApi
)
