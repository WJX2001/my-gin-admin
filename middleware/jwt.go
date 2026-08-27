package middleware

import (
	"errors"
	"github.com/WJX2001/gin-vue-admin-server/global"
	"github.com/WJX2001/gin-vue-admin-server/model/common/response"
	"github.com/WJX2001/gin-vue-admin-server/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"strconv"
	"time"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 我们这里 jwt 鉴权取头部信息 x-token 登录时返回 token 信息，这里前端需要把 token 存储到 cookie 或者本地 localStorage中 不过需要跟后端协商过期时间，可以约定刷新令牌或者重新登陆
		token := utils.GetToken(c)
		if token == "" {
			response.NoAuth("未登陆或非法访问，请登录", c)
			c.Abort()
			return
		}
		if isBlacklist(token) {
			response.NoAuth("您的账户异地登陆或令牌失效", c)
			utils.ClearToken(c)
			c.Abort()
			return
		}
		j := utils.NewJWT()
		// parseToken 解析 token 包含的信息
		claims, err := j.ParseToken(token)
		if err != nil {
			if errors.Is(err, utils.TokenExpired) {
				response.NoAuth("登录已过期，请重新登录", c)
				utils.ClearToken(c)
				c.Abort()
				return
			}
			response.NoAuth(err.Error(), c)
			utils.ClearToken(c)
			c.Abort()
			return
		}

		c.Set("claims", claims)
		if claims.ExpiresAt.Unix()-time.Now().Unix() < claims.BufferTime {
			dr, _ := utils.ParseDuration(global.GVA_CONFIG.JWT.ExpiresTime)
			claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(dr))
			newToken, _ := j.CreateTokenByOldToken(token, *claims)
			newClaims, _ := j.ParseToken(newToken)
			c.Header("new-token", newToken)
			c.Header("new-expires-at", strconv.FormatInt(newClaims.ExpiresAt.Unix(), 10))
			utils.SetToken(c, newToken, int(dr.Seconds()))
			if global.GVA_CONFIG.System.UseMultipoint {
				// 记录新的活跃jwt
				//_ = utils.SetR
			}
		}
		c.Next()

		if newToken, exists := c.Get("new-token"); exists {
			c.Header("new-token", newToken.(string))
		}
		if newExpiresAt, exists := c.Get("new-expires-at"); exists {
			c.Header("new-expires-at", newExpiresAt.(string))
		}
	}
}

// @author: [WJX2001](https://github.com/WJX2001)
// @function: isBlacklist
// @description:
// @param: jwt string
// @return:
func isBlacklist(jwt string) bool {
	_, ok := global.GVA_CACHE.Get(jwt)
	return ok
}
