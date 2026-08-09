package system

import (
	"context"
	"errors"
	"github.com/WJX2001/gin-vue-admin-server/global"
	"github.com/WJX2001/gin-vue-admin-server/model/system"
	"github.com/WJX2001/gin-vue-admin-server/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type UserService struct{}

func (userService *UserService) Register(ctx context.Context, u system.SysUser) (userInter system.SysUser, err error) {
	var user system.SysUser
	if !errors.Is(global.GVA_DB.WithContext(ctx).Where("username = ?", u.Username).First(&user).Error, gorm.ErrRecordNotFound) {
		return userInter, errors.New("用户名已注册")
	}
	// 否则 附加 uuid 密码 hash 加密 注册
	cfg := (&SecurityConfigService{}).Current(ctx)
	u.MustChangePassword = cfg.ForceNewUserChangePassword
	u.Password = utils.BcryptHash(u.Password)
	u.UUID = uuid.New()
	now := time.Now()
	u.PasswordUpdatedAt = &now
	err = global.GVA_DB.WithContext(ctx).Create(&u).Error
	return u, err
}
