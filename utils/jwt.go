package utils

import (
	"context"
	"errors"
	"github.com/WJX2001/gin-vue-admin-server/global"
	"github.com/WJX2001/gin-vue-admin-server/model/system/request"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type JWT struct {
	SigningKey []byte
}

var (
	TokenExpired          = errors.New("token已过期")
	TokenNotValidYet      = errors.New("token尚未激活")
	TokenMalformed        = errors.New("这不是一个token")
	TokenSignatureInvalid = errors.New("无效签名")
	TokenInvalid          = errors.New("无法处理此token")
)

func NewJWT() *JWT {
	return &JWT{
		[]byte(global.GVA_CONFIG.JWT.SigningKey),
	}
}

func (j *JWT) CreateClaims(baseClaims request.BaseClaims) request.CustomClaims {
	bf, _ := ParseDuration(global.GVA_CONFIG.JWT.BufferTime)
	ep, _ := ParseDuration(global.GVA_CONFIG.JWT.ExpiresTime)
	claims := request.CustomClaims{
		BaseClaims: baseClaims,
		BufferTime: int64(bf / time.Second), // 缓冲时间1天 缓冲时间内会获得新的 token 刷新令牌，此时一个用户会存在两个有效令牌，但是前端只留下一个有效令牌 另一个会丢失
		RegisteredClaims: jwt.RegisteredClaims{
			Audience:  jwt.ClaimStrings{"GVA"},                   // 受众
			NotBefore: jwt.NewNumericDate(time.Now().Add(-1000)), // 签名生效时间
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(ep)),    // 过期时间 7天 配置文件
			Issuer:    global.GVA_CONFIG.JWT.Issuer,              // 签名的发行者
		},
	}
	return claims
}

// CreateToken 创建一个token
func (j *JWT) CreateToken(claims request.CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// CreateTokenByOldToken 旧token 换新token 使用归并回源避免并发问题
func (j *JWT) CreateTokenByOldToken(oldToken string, claims request.CustomClaims) (string, error) {
	/*
	* 这里使用 singleflight 防并发，同一个 oldToken，同一时刻只执行一次 CreateToken
	* 这个函数主要在 JWT 中间件里使用，用户快过期时，前端可能并发发多个请求（多个tab、页面同时请求）。如果不合并：
	* 	- 每个请求都会签一个新 token
	* 	- 响应头里 new-token 各不相同
	*   - 前端/Redis 状态容易乱
	* 有了 singleflight:
	* 	- 第一个请求真正执行 CreateToken
	*   - 其他并发请求 等同一个结果
	*   - 大家拿到 同一个新 token
	 */
	v, err, _ := global.GVA_Concurrency_Control.Do("JWT:"+oldToken, func() (interface{}, error) {
		return j.CreateToken(claims)
	})
	return v.(string), err
}

// ParseToken 解析 token
func (j *JWT) ParseToken(tokenString string) (*request.CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &request.CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})

	if err != nil {
		switch {
		case errors.Is(err, jwt.ErrTokenExpired):
			return nil, TokenExpired
		case errors.Is(err, jwt.ErrTokenMalformed):
			return nil, TokenMalformed
		case errors.Is(err, jwt.ErrTokenSignatureInvalid):
			return nil, TokenSignatureInvalid
		case errors.Is(err, jwt.ErrTokenNotValidYet):
			return nil, TokenNotValidYet
		default:
			return nil, TokenInvalid
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*request.CustomClaims); ok && token.Valid {
			return claims, nil
		}
	}
	return nil, TokenInvalid
}

func SetRedisJWT(jwt string, userName string) (err error) {
	// 此处过期时间等于 jwt 过期时间
	dr, err := ParseDuration(global.GVA_CONFIG.JWT.ExpiresTime)
	if err != nil {
		return err
	}
	timer := dr
	err = global.GVA_REDIS.Set(context.Background(), userName, jwt, timer).Err()
	return err
}
