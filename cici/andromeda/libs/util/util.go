//一些基本的辅助类
package util

import (
	"crypto/md5"
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	typeTime = reflect.TypeOf(time.Time{})
)

// 对字符串进行md5哈希,
// 返回32位小写md5结果
func MD5(s string) string {
	h := md5.New()
	io.WriteString(h, s)
	return fmt.Sprintf("%x", h.Sum(nil))
}

// 对字符串进行md5哈希,
// 返回32位大写写md5结果
func MD5ToUpper(s string) string {
	h := md5.New()
	io.WriteString(h, s)
	return fmt.Sprintf("%X", h.Sum(nil))
}

// 对字符串进行md5哈希,
// 返回16位小写md5结果
func MD5_16(s string) string {
	return MD5(s)[8:24]
}

// 对字符串进行签名，
// 返回int64的签名结果。
func CreateSign_8_23(s string) int64 {
	i, _ := strconv.ParseInt(MD5(s)[8:23], 16, 64)
	return i
}

func CreateSignInt32(s string) int {
	i, _ := strconv.ParseInt(MD5(s)[8:15], 16, 64)
	return int(i)
}

func StringIsIn(id string, ids []string) bool {
	for i := 0; i < len(ids); i++ {
		if id == ids[i] {
			return true
		}
	}
	return false
}

func MatchArrayFunc(n int, match func(int) bool) bool {
	if n == 0 {
		return false
	}
	for i := 0; i < n; i++ {
		if match(i) {
			return true
		}
	}
	return false
}

func MatchArrayString(id string, ids []string) bool {
	return StringIsIn(id, ids)
}

func MatchArrayInt(id int, ids []int) bool {
	for i := 0; i < len(ids); i++ {
		if id == ids[i] {
			return true
		}
	}
	return false
}

func MatchArrayInt64(id int64, ids []int64) bool {
	for i := 0; i < len(ids); i++ {
		if id == ids[i] {
			return true
		}
	}
	return false
}

func MatchArrayInt32(id int32, ids []int32) bool {
	for i := 0; i < len(ids); i++ {
		if id == ids[i] {
			return true
		}
	}
	return false
}

func MatchArray(id string, ids []string) bool {
	if id == "" {
		return false
	}
	for i := 0; i < len(ids); i++ {
		if strings.Contains(id, ids[i]) {
			return true
		}
	}
	return false
}

func Sort(ids []int) []int {
	is := sort.IntSlice(ids)
	is.Sort()
	return is
}

var _level = -1
var once sync.Once

func Ip2long(ipstr string) uint32 {
	defer func() {
		if r := recover(); r != nil {
			// return uint32(0)
		}
	}()
	ip := net.ParseIP(ipstr)
	if ip == nil {
		return 0
	}
	ip = ip.To4()
	return binary.BigEndian.Uint32(ip)
}

func Long2ip(ip uint32) string {
	return fmt.Sprintf("%d.%d.%d.%d", ip>>24, ip<<8>>24, ip<<16>>24, ip<<24>>24)
	// ipByte := make([]byte, 4)
	// binary.BigEndian.PutUint32(ipByte, ipLong)
	// ip := net.IP(ipByte)
	// return ip.String()
}
