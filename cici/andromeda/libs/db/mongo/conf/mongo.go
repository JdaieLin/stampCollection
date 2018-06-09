package conf

import (
	"cici/andromeda/libs/confutil"
	"cici/andromeda/libs/confutil/hand"
)

// 数据库配置信息
type MultiDbConfig struct {
	Servers []confutil.NetBase
	// 用户名
	UserName string
	// 密码
	Password string
	// 数据库
	Database string
}

type MultiDbsConfig map[string]MultiDbConfig

func DefaultMultiDbConfig() (cfg *MultiDbConfig, err error) {
	cfg = new(MultiDbConfig)
	err = hand.LoadConfig("mgdb", cfg)
	return
}
