package initialize

import (
	"github.com/WJX2001/gin-vue-admin-server/global"
	"github.com/WJX2001/gin-vue-admin-server/utils"
)

func OtherInit() {
	// 提前解析下jwt 的配置文件，防止等到发token 的时候再报错，提前将错误暴露出来
	_, err := utils.ParseDuration(global.GVA_CONFIG.JWT.ExpiresTime)
	if err != nil {
		panic(err)
	}

	_, err = utils.ParseDuration(global.GVA_CONFIG.JWT.BufferTime)
	if err != nil {
		panic(err)
	}
}
