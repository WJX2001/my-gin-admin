package config

type Server struct {
	JWT       JWT     `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Zap       Zap     `mapstructure:"zap" json:"zap" yaml:"zap"`
	Redis     Redis   `mapstructure:"redis" json:"redis" yaml:"redis"`
	RedisList []Redis `mapstructure:"redis-list" json:"redis-list" yaml:"redis-list"`
	System    System  `mapstructure:"system" json:"system" yaml:"system"`

	// gorm
	Mysql Mysql `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Pgsql Pgsql `mapstructure:"pgsql" json:"pgsql" yaml:"pgsql"`

	// 应用身份（日志静态字段 node/app_id/env）
	App App `mapstructure:"app" json:"app" yaml:"app"`
}
