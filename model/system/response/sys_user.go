package response

import "github.com/WJX2001/gin-vue-admin-server/model/system"

type SysUserResponse struct {
	User system.SysUser `json:"user"`
}
