package system

import (
	"context"
	"github.com/WJX2001/gin-vue-admin-server/global"
	"github.com/WJX2001/gin-vue-admin-server/model/system"
	"github.com/WJX2001/gin-vue-admin-server/utils/logger"
)

type JwtService struct{}

func LoadAll(ctx context.Context) {
	var data []string
	err := global.GVA_DB.WithContext(ctx).Model(&system.JwtBlacklist{}).Select("jwt").Find(&data).Error
	if err != nil {
		logger.WithCtx(ctx).Mod("biz").Err(err).Error("加载数据库jwt黑名单失败！")
		return
	}
	for i := 0; i < len(data); i++ {
		global.GVA_CACHE.SetDefault(data[i], "1")
	} // jwt 黑名单 加入 GVA_CACHE 中
}
