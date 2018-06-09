package conf

import (
	"cici/andromeda/libs/confutil"
	"cici/andromeda/libs/db/mongo/conf"
	"cici/andromeda/libs/log"
	"path"
	"runtime"
)

type Config struct {
	Debug   bool
	Network int // 0.windows本地, 1. macos本地 2. sunteng服务器
	Mongo   conf.MultiDbConfig
	Log     log.Config
}

var defaultConfig = new(Config)

func App() *Config {
	return defaultConfig
}

func init() {
	_, file, _, _ := runtime.Caller(0)
	configFile := path.Join(path.Dir(file), "config.json")
	if err := confutil.LoadExtendConf(configFile, defaultConfig); err != nil {
		panic("load config faild : " + err.Error())
	}
}
