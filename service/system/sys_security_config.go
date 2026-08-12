package system

import (
	"context"
	"errors"
	"github.com/WJX2001/gin-vue-admin-server/utils/logger"
	"sync/atomic"

	"github.com/WJX2001/gin-vue-admin-server/global"
	"github.com/WJX2001/gin-vue-admin-server/model/system"
	"gorm.io/gorm"
)

type SecurityConfigService struct{}

// securityConfigCache 进程内当前生效配置 热读
var securityConfigCache atomic.Value

func setSecurityConfigCache(cfg system.SysSecurityConfig) {
	securityConfigCache.Store(cfg)
}

func getSecurityConfigCache() system.SysSecurityConfig {
	if v := securityConfigCache.Load(); v != nil {
		return v.(system.SysSecurityConfig)
	}
	return system.SysSecurityConfig{}
}

func (s *SecurityConfigService) Get(ctx context.Context) (system.SysSecurityConfig, error) {
	var cfg system.SysSecurityConfig
	// 系统尚未初始化(未走 init 向导)或连库失败时 global.GVA_DB 为 nil
	// 此时返回代码默认配置并带错误: 调用方 Current 据此不写缓存
	// 待数据库就绪后再惰性加载真实行 同时避免对 nil 的 *gorm.DB 解引用导致 panic
	if global.GVA_DB == nil {
		return system.DefaultSecurityConfig(), errors.New("数据库未初始化")
	}
	err := global.GVA_DB.WithContext(ctx).Where("id = ?", 1).First(&cfg).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		cfg = system.DefaultSecurityConfig()
		cfg.ID = 1
		if err = global.GVA_DB.WithContext(ctx).Create(&cfg).Error; err != nil {
			return cfg, err
		}
		return cfg, nil
	}
	return cfg, err
}

// Current 返回内存缓存当前配置 未加载则惰性 Get
func (s *SecurityConfigService) Current(ctx context.Context) system.SysSecurityConfig {
	if v := securityConfigCache.Load(); v != nil {
		return v.(system.SysSecurityConfig)
	}
	cfg, err := s.Get(ctx)
	if err == nil {
		setSecurityConfigCache(cfg)
	}
	return cfg
}

// LoadAll 启动时加载配置入内存缓存
func (s *SecurityConfigService) LoadAll(ctx context.Context) {
	cfg, err := s.Get(ctx)
	if err != nil {
		logger.WithCtx(ctx).Mod("biz").Error("加载安全配置失效!")
		return
	}
	setSecurityConfigCache(cfg)
}
