package light

import (
	"flag"
	"path"

	"cici/andromda/libs/log"
	"cici/andromeda/libs/confutil"
)

var (
	root  string
	index = new(Index)
)

type Index struct {
	Inuse bool
}

func init() {
	flag.StringVar(&root, "cfg", "/etc/gocfg", "配置文件根目录")
}

func Inuse() bool {
	return index.Inuse
}

func Root() string {
	return root
}

func LoadIndex() (err error) {
	confutil.SetRoot(root)
	var file = path.Join(root, "index.json")
	err = confutil.LoadExtendConf(file, index)
	if err != nil {
		log.Errorf("loadConfig error :%s", err.Error())
	}
	if Inuse() {
		log.Debug("load config from :", root)
	}
	return
}

func LoadConfig(cfg interface{}, local string, files ...string) (err error) {
	var file string
	if Inuse() {
		file = path.Join(root, path.Join(files...))
	} else {
		file = local
	}
	log.Debug("load config from ", file)
	err = confutil.LoadExtendConf(file, cfg)
	if err != nil {
		log.Errorf("loadConfig error :%s", err.Error())
	}
	return
}
