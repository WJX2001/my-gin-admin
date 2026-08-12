package v1

import "github.com/WJX2001/gin-vue-admin-server/api/v1/system"

var ApiGroupApp = new(ApiGroup)

type ApiGroup struct {
	SystemApiGroup system.ApiGroup
}
