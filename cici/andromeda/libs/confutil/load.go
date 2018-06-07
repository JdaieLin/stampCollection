// 加载json（可配置扩展字段）配置文件
//
// {
// "Debug": true,
// "Log": "@extend:./log.json"
// }
package confutil

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

import "github.com/json-iterator/go"

var json = jsoniter.ConfigCompatibleWithStandardLibrary

var (
	envTag    = ":@env[K=V1:V2]"
	extendTag = "@extend:"
	pwdTag    = "@pwd@"
	rootTag   = "@root@"
	root      = ""
)

// 设置扩展标识，如果不设置，默认为 '@extend:'
func SetExtendTag(tag string) {
	extendTag = tag
}

func SetRoot(r string) {
	root = r
}

// 设置当前路径标识，如果不设置，默认为 '@pwd@'
// @pwd@ 会被替换成当前文件的路径，
// 至于是绝对路径还是相对路径，取决于读取文件时，传入的是绝对路径还是相对路径
func SetPathTag(tag string) {
	pwdTag = tag
}

//加载json（可配置扩展字段）配置文件
func LoadExtendConf(filePath string, v interface{}) error {
	data, err := extendFile(filePath)
	if err != nil {
		return err
	}

	if data, err = replaceEnv(data); err != nil {
		return err
	}

	return json.Unmarshal(data, v)
}

func extendFile(filePath string) (data []byte, err error) {
	fi, err := os.Stat(filePath)
	if err != nil {
		return
	}
	if fi.IsDir() {
		err = fmt.Errorf(filePath + " is not a file.")
		return
	}

	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		return
	}

	if len(root) != 0 {
		b = bytes.Replace(b, []byte(rootTag), []byte(root), -1)
	}

	dir := filepath.Dir(filePath)
	return extendFileContent(dir, bytes.Replace(b, []byte(pwdTag), []byte(dir), -1))
}

func extendFileContent(dir string, content []byte) (data []byte, err error) {
	//检查是不是规范的json
	test := new(interface{})
	err = json.Unmarshal(content, &test)
	if err != nil {
		return
	}

	// 替换子json文件
	reg := regexp.MustCompile(`"` + extendTag + `.*?"`)
	data = reg.ReplaceAllFunc(content, func(match []byte) []byte {
		match = match[len(extendTag)+1 : len(match)-1]
		sb, e := extendFile(filepath.Join(dir, string(match)))
		if e != nil {
			err = fmt.Errorf("替换json配置[%s]失败：%s\n", match, e.Error())
		}
		return sb
	})

	return
}

func replaceEnv(src []byte) (dest []byte, err error) {
	var m map[string]interface{}
	if err = json.Unmarshal(src, &m); err != nil {
		return
	}
	replaceEnvMap(m)
	dest, err = json.Marshal(m)
	return
}

var envReg = regexp.MustCompile(`.*?:@env\[.*?=.*?\]`)

func matchKey(key string) (prefix string, match, ok bool) {
	if ok = envReg.MatchString(key); ok {
		if beg := strings.Index(key, ":@env["); beg != -1 {
			prefix = key[:beg]
			end := strings.Index(key[beg:], "]")
			buf := key[beg+len(":@env[") : beg+end]
			pair := strings.Split(buf, "=")
			match = MatchEnv(pair[0], pair[1])
		}
	}
	return
}

func assertValue(v interface{}) {
	switch v.(type) {
	case map[string]interface{}:
		replaceEnvMap(v.(map[string]interface{}))
	case []interface{}:
		var slice_v = v.([]interface{})
		for i := 0; i < len(slice_v); i++ {
			assertValue(slice_v[i])
		}
	}
}

func replaceEnvMap(m map[string]interface{}) (err error) {
	for k, v := range m {
		if tag, match, found := matchKey(k); found {
			delete(m, k)
			if match {
				m[tag] = v
			}
		}
	}
	for _, v := range m {
		assertValue(v)
	}
	return
}
