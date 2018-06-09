package hand

import (
	"path"
	"runtime"

	"cici/andromeda/libs/confutil"
	"cici/andromeda/libs/log"
)

func LoadConfig(file string, cfg interface{}) (err error) {
	jsonPath := path.Join(ConfRoot, file+".json")
	err = confutil.LoadExtendConf(jsonPath, cfg)
	if err != nil {
		log.Errorf("loadConfig error :%s", err.Error())
	}
	return
}

var Root string
var ConfRoot string

func init() {
	_, filename, _, _ := runtime.Caller(0)
	Root = path.Dir(filename)
	ConfRoot = path.Join(Root, "..", "..", "..", "config", "files")
}
