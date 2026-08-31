package system

import (
	"context"
	"github.com/WJX2001/gin-vue-admin-server/global"
	"github.com/WJX2001/gin-vue-admin-server/model/system"
)

type LoginLogService struct{}

func (loginLogService *LoginLogService) CreateLoginLog(ctx context.Context, loginLog system.SysLoginLog) (err error) {
	// 系统未初始化 (GVA_DB==nil)时，/base/login 仍可被未认证请求触达
	// 此处静默跳过登录日志 避免对 nil *gorm.DB 解引用 panic(与 CreateSysError 守卫一致)
	if global.GVA_DB == nil {
		return nil
	}
	err = global.GVA_DB.WithContext(ctx).Create(&loginLog).Error
	return err
}
