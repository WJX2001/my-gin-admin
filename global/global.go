package global

import (
	"github.com/WJX2001/gin-vue-admin-server/config"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	GVA_DB            *gorm.DB
	GVA_CONFIG        config.Server
	GVA_VP            *viper.Viper
	GVA_LOG           *zap.Logger
	GVA_ACTIVE_DBNAME *string
)
