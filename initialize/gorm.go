package initialize

import (
	"fmt"
	"github.com/WJX2001/gin-vue-admin-server/global"
	"github.com/WJX2001/gin-vue-admin-server/model/system"
	"gorm.io/gorm"
	"os"
)

func Gorm() *gorm.DB {
	switch global.GVA_CONFIG.System.DbType {
	case "mysql":
		global.GVA_ACTIVE_DBNAME = &global.GVA_CONFIG.Mysql.Dbname
		return GormMysql()
	case "pgsql":
		global.GVA_ACTIVE_DBNAME = &global.GVA_CONFIG.Pgsql.Dbname
		return nil
	default:
		global.GVA_ACTIVE_DBNAME = &global.GVA_CONFIG.Mysql.Dbname
		return GormMysql()
	}
}

func RegisterTables() {
	if global.GVA_CONFIG.System.DisableAutoMigrate {
		fmt.Println("auto-migrate is disabled,skipping table registration")
		return
	}

	db := global.GVA_DB
	err := db.AutoMigrate(
		system.SysUser{},
		system.SysSecurityConfig{},
	)
	if err != nil {
		// TODO: 这里将来用 zap 日志
		fmt.Println("register table failed")
		os.Exit(1)
	}
}
