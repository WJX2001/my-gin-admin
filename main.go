package main

import (
	"fmt"
	"github.com/WJX2001/gin-vue-admin-server/core"
	"github.com/WJX2001/gin-vue-admin-server/global"
	"github.com/gin-gonic/gin"
)

func main() {
	initializeSystem()
	RunServer()

}

func initializeSystem() {
	global.GVA_VP = core.Viper()
	fmt.Println("wjx-test222:", global.GVA_VP.GetString("jwt.wjx-test"))
	//fmt.Println("wjxTest", global.GVA_VP.GetString("jwt.signing-key"))
}

func RunServer() {

	// 创建默认路由引擎（包含日志+Recovery崩溃恢复中间件）
	r := gin.Default()

	// 简单GET接口
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg":  "hello gin!",
			"code": 0,
		})
	})

	// 启动服务，监听0.0.0.0:8080
	err := r.Run(":8080")
	if err != nil {
		panic("服务启动失败:" + err.Error())
	}
}
