package config

type System struct {
	DbType             string `mapstructure:"db-type" json:"db-type" yaml:"db-type"`                                        // 数据库类型:mysql(默认)|sqlite|sqlserver|postgresql
	DisableAutoMigrate bool   `mapstructure:"disable-auto-migrate" json:"disable-auto-migrate" yaml:"disable-auto-migrate"` // 自动迁移数据库表结构，生产环境建议设为false，手动迁移
	RouterPrefix       string `mapstructure:"router-prefix" json:"router-prefix" yaml:"router-prefix"`
	Addr               int    `mapstructure:"addr" json:"addr" yaml:"addr"` // 端口值
}
