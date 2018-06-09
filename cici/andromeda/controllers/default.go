package controllers

import (
	"cici/andromeda/constant"
	"encoding/json"
	"fmt"
	"math/rand"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"

	"sunteng/commons/jsonutil"
)

func init() {
	rand.Seed(time.Now().Unix())
}

type Controller struct {
	beego.Controller
}

func (this *Controller) StrQuery(name string) string {
	return this.Ctx.Input.Query(name)
}

func (this *Controller) IntQuery(name string) int {
	s := this.Ctx.Input.Query(name)
	i, _ := strconv.Atoi(s)
	return i
}

func (this *Controller) Int64Query(name string) int64 {
	s := this.StrQuery(name)
	i, _ := strconv.ParseInt(s, 10, 64)
	return i
}

func (this *Controller) Int64Param(name string) int64 {
	s := this.Ctx.Input.Param(name)
	i, _ := strconv.ParseInt(s, 10, 64)
	return i
}

func (this *Controller) Int64SliceQuery(name string) []int64 {
	s := this.StrQuery(name)
	if s == "" {
		return []int64{}
	}
	tmp := strings.Split(s, ",")
	res := make([]int64, len(tmp))
	for i, str := range tmp {
		res[i], _ = strconv.ParseInt(str, 10, 64)
	}
	return res
}

func (this *Controller) BodyItem(inst interface{}) error {
	return json.Unmarshal(this.Ctx.Input.RequestBody, inst)
}

func (this *Controller) BodyQuery() map[string]string {
	a := jsonutil.Parse(string(this.Ctx.Input.RequestBody)).All()
	return a
}

func (this *Controller) PostBody() []byte {
	return this.Ctx.Input.RequestBody
}

func (this *Controller) Page() int {
	ret := this.IntQuery("page")
	if ret == 0 {
		ret = 1
	}
	return ret
}
func (this *Controller) Offset() int {
	return (this.Page() - 1) * this.Limit()
}

func (this *Controller) Limit() int {
	ret := this.IntQuery("limit")
	if ret == 0 {
		ret = 20
	}
	return ret
}

func (this *Controller) WriteString(str string) {
	this.Ctx.WriteString(str)
}

func (this *Controller) WriteJsonEncodeError(err error) {
	this.WriteError(constant.ERR_REQ_JSON_ENCODE, err.Error())
}

func (this *Controller) WriteError(code int, params ...interface{}) {
	msg := fmt.Sprintf(constant.GetErrorInfo(code), params...)
	msg = fmt.Sprintf(`{"success":false,"error_code":%d, "error_info":"%s"}`, code, msg)
	this.Ctx.WriteString(msg)
}

func (this *Controller) WriteErrorResult(result interface{}, err error, code int, params ...interface{}) {
	var ret = map[string]interface{}{
		"success":    false,
		"debug_info": err.Error(),
		"error_code": code,
		"error_info": fmt.Sprintf(constant.GetErrorInfo(code), params...),
		"result":     result,
	}
	bs, err2 := json.Marshal(ret)
	if err2 != nil {
		this.WriteJsonEncodeError(err2)
		return
	}

	this.Ctx.WriteString(string(bs))
}

func (this *Controller) WriteSuccess(msg string) {
	this.Ctx.WriteString(`{"success":true,"result":` + msg + `}`)
}

func (this *Controller) WriteSuccessResult(result interface{}) {
	var ret = map[string]interface{}{
		"success": true,
		"result":  result,
	}
	bs, err2 := json.Marshal(ret)
	if err2 != nil {
		this.WriteJsonEncodeError(err2)
		return
	}

	this.Ctx.WriteString(string(bs))
}

func (this *Controller) WriteResult(result interface{}, err ...error) {
	if len(err) > 0 && err[0] != nil {
		this.WriteError(constant.ERR_REQ_SERVER, err[0].Error())
		return
	}
	var ret = map[string]interface{}{
		"success": true,
		"result":  result,
	}
	bs, err2 := json.Marshal(ret)
	if err2 != nil {
		this.WriteJsonEncodeError(err2)
		return
	}
	this.Ctx.WriteString(string(bs))
}

func (this *Controller) WriteItems(items interface{}, amount interface{}, total int) {
	this.WriteResult(map[string]interface{}{
		"items":  items,
		"amount": amount,
		"total":  total,
		"page":   this.Page(),
		"size":   reflect.ValueOf(items).Len(),
	})
}
