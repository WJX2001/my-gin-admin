package initialize

import (
	"github.com/WJX2001/gin-vue-admin-server/config"
	"github.com/WJX2001/gin-vue-admin-server/global"
	"github.com/WJX2001/gin-vue-admin-server/initialize/internal"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

func GormMysql() *gorm.DB {
	m := global.GVA_CONFIG.Mysql
	return initMysqlDatabase(m)
}

func initMysqlDatabase(m config.Mysql) *gorm.DB {
	if m.Dbname == "" {
		return nil
	}

	mysqlConfig := mysql.Config{
		DSN:                       m.Dsn(), // DSN Data source name
		DefaultStringSize:         191,     // string 类型字段的默认长度
		SkipInitializeWithVersion: false,   // 根据版本自动配置
	}

	// 数据库配置
	general := m.GeneralDB
	if db, err := gorm.Open(mysql.New(mysqlConfig), internal.Gorm.Config(general)); err != nil {
		panic(err)
	} else {
		db.InstanceSet("gorm:table_options", "ENGINE="+m.Engine)
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		sqlDB.SetConnMaxLifetime(time.Duration(m.ConnMaxLifetime) * time.Second)
		return db
	}
}
