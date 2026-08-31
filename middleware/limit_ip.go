package middleware

import (
	"errors"
	"github.com/WJX2001/gin-vue-admin-server/global"
	"github.com/WJX2001/gin-vue-admin-server/model/common/response"
	"github.com/WJX2001/gin-vue-admin-server/service"
	"github.com/WJX2001/gin-vue-admin-server/utils/logger"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func CacheCheckOrMark(key string, expire int, limit int) error {
	if global.GVA_CACHE == nil {
		return nil
	}
	n, err := global.GVA_CACHE.IncrementWithExpire(key, 1, time.Duration(expire)*time.Second)
	if err != nil {
		logger.Bg().Mod("system").Err(err).Error("limit")
		return nil // fail-open
	}
	if int(n) > limit {
		return errors.New("请求太过频繁，请稍后再试")
	}
	return nil
}

// SecurityLimit 按安全配置对登录/敏感接口限流 未开启则放行
func SecurityLimit() gin.HandlerFunc {
	return func(c *gin.Context) {
		enable, window, count := service.ServiceGroupApp.SystemServiceGroup.SecurityConfigService.CurrentLimit(c.Request.Context())
		if !enable {
			c.Next()
			return
		}
		// 拼限流用的缓存 key，用来区分 谁+哪个接口的请求次数
		key := "GVA_SecLimit" + c.ClientIP() + c.FullPath()
		if err := CacheCheckOrMark(key, window, count); err != nil {
			c.JSON(http.StatusOK, gin.H{"code": response.ERROR, "msg": err.Error()})
			c.Abort()
			return
		}
		c.Next()
	}
}
