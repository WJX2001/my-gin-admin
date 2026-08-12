package main

import (
	"github.com/WJX2001/gin-vue-admin-server/core"
	"github.com/WJX2001/gin-vue-admin-server/global"
	"github.com/WJX2001/gin-vue-admin-server/initialize"
)

func main() {
	initializeSystem()
	core.RunServer()

}

func initializeSystem() {
	global.GVA_VP = core.Viper() // 初始化Viper
	initialize.OtherInit()
	global.GVA_LOG = core.Zap()       // 初始化zap日志库
	global.GVA_DB = initialize.Gorm() // gorm连接数据库
	//fmt.Println("wjx-test。。。。。", global.GVA_VP.GetString("jwt.wjx-test"))
	if global.GVA_DB != nil {
		initialize.RegisterTables() // 初始化表
	}
}
