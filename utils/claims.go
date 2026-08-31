package utils

import (
	"github.com/WJX2001/gin-vue-admin-server/model/system"
	systemReq "github.com/WJX2001/gin-vue-admin-server/model/system/request"
	"github.com/WJX2001/gin-vue-admin-server/utils/logger"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

func ClearToken(c *gin.Context) {
	setTokenCookie(c, "", -1, time.Unix(1, 0))
}

func SetToken(c *gin.Context, token string, maxAge int) {
	setTokenCookie(c, token, maxAge, time.Time{})
}

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

// GetUserID 从Gin的Context中获取从 jwt 解析出来用户ID
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

func setTokenCookie(c *gin.Context, value string, maxAge int, expires time.Time) {
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "x-token",
		Value:    value,
		Path:     "/",
		Expires:  expires,
		MaxAge:   maxAge,
		Secure:   requestUsesHTTPS(c.Request),
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})
}

func requestUsesHTTPS(r *http.Request) bool {
	// 如果客户端直连Go服务走HTTPS，TLS握手完成后，r.TLS 就不会 nil，直接返回 true
	// 注意：经过 Nginx 反向代理后，Go 后端收到的是 http，此时 r.TLS = nil，这也就是为什么要处理 X-Forwarded-Proto 请求头
	if r.TLS != nil {
		return true
	}
	// X-Forwarded-Proto 是反向代理标准头，当 Nginx 对外接受https 请求，转发给后端 Go 服务时，Nginx 会加这个Header
	// X-Forwarded-Proto: https 告诉后端原始客户端用的是什么协议（http/https）
	// 有些多层代理环境，这个头会带多个值，逗号分隔，例如： X-Forwarded-Proto：https,http
	forwardedProto := strings.SplitN(r.Header.Get("X-Forwarded-Proto"), ",", 2)[0]
	return strings.EqualFold(strings.TrimSpace(forwardedProto), "https")
}

// LoginTokenWithExpire 签发登录 token 可携带 MustChangePwd 强制改密标记
func LoginTokenWithExpire(user system.Login, mustChangePwd bool) (token string, claims systemReq.CustomClaims, err error) {
	j := NewJWT()
	claims = j.CreateClaims(systemReq.BaseClaims{
		UUID:     user.GetUUID(),
		ID:       user.GetUserId(),
		Nickname: user.GetNickname(),
		Username: user.GetUsername(),
		UserType: user.GetUserType(),
	})
	claims.MustChangePwd = mustChangePwd
	token, err = j.CreateToken(claims)
	return
}
