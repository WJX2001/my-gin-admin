package utils

import (
	systemReq "github.com/WJX2001/gin-vue-admin-server/model/system/request"
	"github.com/WJX2001/gin-vue-admin-server/utils/logger"
	"github.com/gin-gonic/gin"
)

func GetToken(c *gin.Context) string {
	token := c.Request.Header.Get("x-token")
	if token != "" {
		return token
	}
	token, _ = c.Cookie("x-token")
	return token
}

func GetClaims(c *gin.Context) (*systemReq.CustomClaims, error) {
	token := GetToken(c)
	j := NewJWT()
	claims, err := j.ParseToken(token)
	if err != nil {
		logger.WithCtx(c.Request.Context()).Mod("system").Error("从Gin的Context中获取从jwt解析信息失败, 请检查请求头是否存在x-token且claims是否为规定结构")
	}
	return claims, err
}

// GetUserId 从Gin的Context中获取从 jwt 解析出来用户ID
func GetUserID(c *gin.Context) uint {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(c); err != nil {
			return 0
		} else {
			return cl.BaseClaims.ID
		}
	} else {
		waitUse := claims.(*systemReq.CustomClaims)
		return waitUse.BaseClaims.ID
	}
}
