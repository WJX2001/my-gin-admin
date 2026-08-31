package system

import (
	"github.com/WJX2001/gin-vue-admin-server/global"
	"github.com/WJX2001/gin-vue-admin-server/model/common/response"
	"github.com/WJX2001/gin-vue-admin-server/model/system"
	systemReq "github.com/WJX2001/gin-vue-admin-server/model/system/request"
	systemRes "github.com/WJX2001/gin-vue-admin-server/model/system/response"
	systemSvc "github.com/WJX2001/gin-vue-admin-server/service/system"
	"github.com/WJX2001/gin-vue-admin-server/utils"
	"github.com/WJX2001/gin-vue-admin-server/utils/logger"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func (b *BaseApi) Login(c *gin.Context) {
	var l systemReq.Login
	err := c.ShouldBindJSON(&l)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(l, utils.LoginVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	cfg := securityConfigService.Current(c.Request.Context())

	// 1. 帐号锁定检查
	if cfg.LockEnable && systemSvc.IsAccountLocked(c.Request.Context(), l.Username) {
		response.FailWithMessage("帐号已锁定，请 "+strconv.Itoa(cfg.LockDuration)+" 分钟后再试", c)
		loginLogService.CreateLoginLog(c.Request.Context(), system.SysLoginLog{
			Username:     l.Username,
			Ip:           c.ClientIP(),
			Status:       false,
			ErrorMessage: "帐号已锁定",
			Agent:        c.Request.UserAgent(),
		})
		return
	}

	// 2. 验证码检查
	key := c.ClientIP()
	openCaptcha := cfg.CaptchaOpen
	openCaptchaTimeOut := cfg.CaptchaTimeout
	v, ok := global.GVA_CACHE.Get(key)
	if !ok {
		global.GVA_CACHE.Set(key, int64(1), time.Second*time.Duration(openCaptchaTimeOut))
	}
	var oc = openCaptcha == 0 || openCaptcha < interfaceToInt(v)
	if oc && (l.Captcha == "" || l.CaptchaId == "" || !store.Verify(l.CaptchaId, l.Captcha, true)) {
		global.GVA_CACHE.Increment(key, 1)
		response.FailWithMessage("验证码错误", c)
		loginLogService.CreateLoginLog(c.Request.Context(), system.SysLoginLog{
			Username:     l.Username,
			Ip:           c.ClientIP(),
			Status:       false,
			ErrorMessage: "验证码错误",
			Agent:        c.Request.UserAgent(),
		})
		return
	}

	// 3.凭证校验
	u := &system.SysUser{Username: l.Username, Password: l.Password}
	user, err := userService.Login(c.Request.Context(), u)
	if err != nil {
		logger.WithCtx(c.Request.Context()).Mod("biz").Err(err).Error("登录失败！用户名不存在或者密码错误！")
		global.GVA_CACHE.Increment(key, 1)
		systemSvc.RecordLoginFail(c.Request.Context(), l.Username, cfg)
		response.FailWithMessage("用户名不存在或者密码错误", c)
		loginLogService.CreateLoginLog(c.Request.Context(), system.SysLoginLog{
			Username:     l.Username,
			Ip:           c.ClientIP(),
			Status:       false,
			ErrorMessage: "用户名不存在或者密码错误",
			Agent:        c.Request.UserAgent(),
		})
		return
	}
	if user.Enable != 1 {
		logger.WithCtx(c.Request.Context()).Mod("biz").Error("登录失败！用户被禁止登录！")
		global.GVA_CACHE.Increment(key, 1)
		response.FailWithMessage("用户被禁止登录", c)
		loginLogService.CreateLoginLog(c.Request.Context(), system.SysLoginLog{
			Username:     l.Username,
			Ip:           c.ClientIP(),
			Status:       false,
			ErrorMessage: "用户被禁止登录",
			Agent:        c.Request.UserAgent(),
			UserID:       user.ID,
		})
		return
	}

	// 4. 登录成功 清除失败计数与锁
	systemSvc.ClearLoginFail(c.Request.Context(), l.Username)

	// 5. 初始密码或密码过期检查
	needChange := systemSvc.ShouldForcePasswordChange(c.Request.Context(), *user, cfg, time.Now())
	b.TokenNext(c, *user, needChange)
}

// TokenNext 登录以后签发 Jwt
func (b *BaseApi) TokenNext(c *gin.Context, user system.SysUser, mustChangePwd bool) {
	token, claims, err := utils.LoginTokenWithExpire(&user, mustChangePwd)
	if err != nil {
		logger.WithCtx(c.Request.Context()).Mod("biz").Err(err).Error("获取token失败！")
		response.FailWithMessage("获取token失败", c)
		return
	}
	// 记录登录成功日志
	loginLogService.CreateLoginLog(c.Request.Context(), system.SysLoginLog{
		Username: user.Username, Ip: c.ClientIP(), Agent: c.Request.UserAgent(),
		Status: true, UserID: user.ID, ErrorMessage: "登录成功",
	})
	if !global.GVA_CONFIG.System.UseMultipoint {
		utils.SetToken(c, token, int(claims.RegisteredClaims.ExpiresAt.Unix()-time.Now().Unix()))
		response.OkWithDetailed(systemRes.LoginResponse{
			User:               user,
			Token:              token,
			ExpiresAt:          claims.RegisteredClaims.ExpiresAt.Unix() * 1000,
			NeedChangePassword: mustChangePwd,
		}, "登录成功", c)
		return
	}

	//if jwtStr,err := jwtService.Get
}

func (b *BaseApi) Register(c *gin.Context) {
	var r systemReq.Register
	err := c.ShouldBindJSON(&r)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(r, utils.RegisterVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err = utils.ValidatePasswordComplexity(r.Password, securityConfigService.Current(c.Request.Context())); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// TODO: 这里还需要加入角色权限控制
	user := &system.SysUser{
		Username:  r.Username,
		NickName:  r.NickName,
		Password:  r.Password,
		HeaderImg: r.HeaderImg,
		Enable:    r.Enable,
		Phone:     r.Phone,
		Email:     r.Email,
	}
	userReturn, err := userService.Register(c.Request.Context(), *user)
	if err != nil {
		logger.WithCtx(c.Request.Context()).Mod("biz").Err(err).Error("注册失败!")
		response.FailWithDetailed(systemRes.SysUserResponse{User: userReturn}, "注册失败", c)
		return
	}
	response.OkWithDetailed(systemRes.SysUserResponse{User: userReturn}, "注册成功", c)
}
