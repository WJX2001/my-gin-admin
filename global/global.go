package global

import (
	"github.com/WJX2001/gin-vue-admin-server/config"
	"github.com/spf13/viper"
)

var (
	GVA_CONFIG config.Server
	GVA_VP     *viper.Viper
)
