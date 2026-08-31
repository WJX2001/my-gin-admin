package system

import (
	"context"
	"github.com/WJX2001/gin-vue-admin-server/global"
	"github.com/WJX2001/gin-vue-admin-server/model/system"
	"github.com/WJX2001/gin-vue-admin-server/utils/logger"
)

type JwtService struct{}

// @author: [WJX2001](https://github.com/WJX2001)
// @function: GetRedisJWT
// @description: 从 redis 取 jwt
// @param ctx
// @param userName
// @return redisJWT
// @return err
func (jwtService *JwtService) GetRedisJWT(ctx context.Context, userName string) (redisJWT string, err error) {
	redisJWT, err = global.GVA_REDIS.Get(ctx, userName).Result()
	return redisJWT, err
}

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

func (JwtService *JwtService) JsonInBlacklist(ctx context.Context, jwtList system.JwtBlacklist) (err error) {
	err = global.GVA_DB.WithContext(ctx).Create(&jwtList).Error
	if err != nil {
		return
	}
	global.GVA_CACHE.SetDefault(jwtList.Jwt, "1")
	return
}
