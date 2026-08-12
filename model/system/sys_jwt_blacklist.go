package system

import "github.com/WJX2001/gin-vue-admin-server/global"

type JwtBlacklist struct {
	global.GVA_MODEL
	Jwt string `gorm:"type:text;comment:jwt"`
}
