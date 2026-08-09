package system

import (
	"github.com/WJX2001/gin-vue-admin-server/model/common/response"
	"github.com/WJX2001/gin-vue-admin-server/model/system"
	systemReq "github.com/WJX2001/gin-vue-admin-server/model/system/request"
	systemRes "github.com/WJX2001/gin-vue-admin-server/model/system/response"
	"github.com/WJX2001/gin-vue-admin-server/utils"
	"github.com/gin-gonic/gin"
)

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
		//logger.WithCtx(c.Request.Context()).Mod("biz").E
		// TODO: 临时写法 后续接入zap 日志
		println("注册失败", err.Error())
		response.FailWithDetailed(systemRes.SysUserResponse{User: userReturn}, "注册失败", c)
		return
	}
	response.OkWithDetailed(systemRes.SysUserResponse{User: userReturn}, "注册成功", c)
}
