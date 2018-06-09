package confutil

import (
	"os"
	"strings"

	"cici/andromeda/libs/util"
)

var envs = map[string]string{
	"IDC":  "bj",
	"test": "a",
}

func init() {
	for _, senv := range os.Environ() {
		segs := strings.Split(senv, "=")
		if len(segs) == 2 {
			envs[segs[0]] = segs[1]
		}
	}
}

func MergeEnv(es map[string]string) {
	for key, value := range es {
		envs[key] = value
	}
}

func SetEnv(key, value string) {
	envs[key] = value
}

func MatchEnv(key, value string) bool {
	v, ok := envs[key]
	if ok {
		if util.StringIsIn(v, strings.Split(value, ":")) {
			return true
		}
	}
	return false
}
