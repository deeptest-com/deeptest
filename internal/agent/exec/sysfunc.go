package agentExec

import (
	"github.com/dop251/goja"
	uuid "github.com/satori/go.uuid"
	"math/rand"
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
