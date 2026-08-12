package core

import (
	"context"
	"fmt"
	"github.com/WJX2001/gin-vue-admin-server/global"
	"github.com/WJX2001/gin-vue-admin-server/initialize"
	"github.com/WJX2001/gin-vue-admin-server/service/system"
	"time"
)

func RunServer() {

	// 初始化通用缓存（必须在 Redis 之后：有 Redis 用 Redis，否则用内存）
	initialize.InitGvaCache()

	if global.GVA_DB != nil {
		system.LoadAll(context.Background())
		(&system.SecurityConfigService{}).LoadAll(context.Background())
	}

	Router := initialize.Routers()
	address := fmt.Sprintf(":%d", global.GVA_CONFIG.System.Addr)

	fmt.Printf(`
	欢迎使用 gin-vue-admin
	当前版本:%s
	项目地址:https://github.com/flipped-aurora/gin-vue-admin
	插件市场:https://plugin.gin-vue-admin.com
	默认自动化文档地址:http://127.0.0.1%s/swagger/index.html
	MCP 独立服务请手动启动: go run ./cmd/mcp -config ./cmd/mcp/config.yaml
	默认MCP StreamHTTP地址:%s
	默认前端文件运行地址:http://127.0.0.1:8080
`, global.Version, address)
	//r := gin.Default()
	////
	////// 简单GET接口
	////r.GET("/hello", func(c *gin.Context) {
	////	c.JSON(200, gin.H{
	////		"msg":  "hello gin!",
	////		"code": 0,
	////	})
	////})
	////tmp := &system.BaseApi{}
	////r.POST("/admin_register", tmp.Register)
	////
	////// 启动服务，监听0.0.0.0:8080
	//err := r.Run(":8080")
	//if err != nil {
	//	panic("服务启动失败:" + err.Error())
	//}

	initServer(address, Router, 10*time.Minute, 10*time.Minute)
}
