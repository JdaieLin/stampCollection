package confutil

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"path"
	"strconv"
	"time"

	"cici/andromeda/libs/log"
	"cici/andromeda/libs/util"
)

type DbConfig struct {
	NetBase
	// 用户名
	UserName string
	// 密码
	Password string
	// 队列名
	Key string
	// 有效期
	Expire time.Duration
}

type NetBase struct {
	Host string
	Port int
}

type MqConfig struct {
	NetBase
	Topic  string
	Expire int64
}

type HttpServerBase struct {
	NetBase
	DaemonBase
}

type DaemonBase struct {
	Name     string
	DataDir  string
	LogLevel int
}

func (this *DaemonBase) GetDataFile(file string) string {
	return path.Join(this.DataDir, this.Name, file)
}

func (this *DaemonBase) GetDataDir() string {
	return path.Join(this.DataDir, this.Name)
}

func (this *DaemonBase) InitAll() (err error) {
	if err = os.MkdirAll(this.GetDataDir(), os.ModePerm); err != nil {
		return
	}
	// if this.DataDir != "-" && this.DataDir != "" {
	// log.SetRootFileAppender(this.GetDataFile(".log"))
	// }
	log.SetRootLevel(this.LogLevel)

	err = ioutil.WriteFile(this.GetDataFile(this.Name+".pid"), []byte(strconv.Itoa(os.Getpid())), 0666)
	return
}

func (this *NetBase) GetAddr() string {
	return net.JoinHostPort(this.Host, strconv.Itoa(this.Port))
}

func (this *NetBase) Check() error {
	if util.TryDial(this.Host, this.Port) == nil {
		return errors.New(this.GetAddr() + " already inused")
	}
	return nil
}

func (this *NetBase) HttpAddr() string {
	return fmt.Sprintf("http://%s:%d", this.Host, this.Port)
}

func (this *NetBase) StartHttp(handler http.Handler) error {
	ls, err := net.Listen("tcp", this.GetAddr())
	if err != nil {
		return err
	}
	err = util.StartHttp(ls, handler)
	return err
}
