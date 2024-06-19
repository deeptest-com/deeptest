package agentExec

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"github.com/dop251/goja"
	uuid "github.com/satori/go.uuid"
	"io"
	"math/rand"
	"strings"
	"time"
)

type SysFunc struct {
	Label string
	Name  string
	Func  interface{}
	Value string
	Desc  string
}

var SysFuncList []SysFunc

func init() {
	SysFuncList = []SysFunc{
		{"__uuid()", "__uuid", __uuid, "__uuid()", "唯一id"},
		{"__random(min,max)", "__random", __random, "__random(1,100)", "随机数"},
		{"__timestamp()", "__timestamp", __timestamp, "__timestamp()", "时间戳"},
		{"__base64encode(str)", "__base64encode", __base64encode, "__base64encode('')", "base64编码"},
		{"__base64decode(str)", "__base64decode", __base64decode, "__base64decode('')", "base64解码"},
		{"__md5(str)", "__md5", __md5, "__md5('')", "md5加密"},
		{"__strreplace(s,old,new)", "__strreplace", __strreplace, "__strreplace('abc','b','d')", "字符串替换"},
		{"__strlen(str)", "__strlen", __strlen, "__strlen('')", "字符串长度"},
		{"__strtolower(str)", "__strtolower", __strtolower, "__strtolower('')", "字符串转小写"},
		{"__strtoupper(str)", "__strtoupper", __strtoupper, "__strtoupper('')", "字符串转大写"},
		{"__substr(start,length)", "__substr", __substr, "__substr('abc',1,2)", "字符串截取"},
	}
}

func defineJsSysFunc(runtime *goja.Runtime) {
	for _, sysFunc := range SysFuncList {
		runtime.Set(sysFunc.Name, sysFunc.Func)
	}
}

func __uuid() string {
	uuid := uuid.NewV4()
	return uuid.String()
}

func __random(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	randomInt := rand.Intn(max - min)
	return min + randomInt
}

func __timestamp() int64 {
	timestamp := time.Now().UnixMilli()
	return timestamp
}

func __base64encode(str string) string {
	originalData := []byte(str)
	encodedData := base64.StdEncoding.EncodeToString(originalData)
	return encodedData
}

func __base64decode(str string) string {
	decodedData, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return ""
	}
	return string(decodedData)
}

func __md5(str string) string {
	// 创建一个新的MD5哈希生成器
	hasher := md5.New()
	// 写入要生成哈希的数据
	io.WriteString(hasher, str)

	// 计算最终的哈希值，返回的是一个16字节的数组
	hashBytes := hasher.Sum(nil)

	// 将字节数组转换为16进制的字符串表示
	hashString := hex.EncodeToString(hashBytes)
	return hashString

}

func __strreplace(s, old, new string) string {
	return strings.ReplaceAll(s, old, new)
}

func __substr(s string, args ...int) string {
	runeS := []rune(s)
	lenS := len(runeS)
	if len(args) == 0 {
		return s
	}
	start := args[0]
	if start < 0 {
		start = lenS + start
	}
	if start < 0 {
		start = 0
	}

	if len(args) == 1 {
		return string(runeS[start:])
	}

	length := args[1]

	end := start + length
	if end > lenS {
		end = lenS
	}
	return string(runeS[start:end])
}

func __strlen(str string) int {
	return len(str)
}

func __strtolower(str string) string {
	return strings.ToLower(str)
}

func __strtoupper(str string) string {
	return strings.ToUpper(str)
}
