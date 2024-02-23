package task

import (
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/pkg/core/cron"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"strconv"
)

type Task interface {
	Run(options map[string]interface{}) (f func() error)
	CallBack(options map[string]interface{}, err error) func()
}

type Proxy struct {
	source     string
	cron       string
	task       Task
	taskId     string
	ServerCron *cron.ServerCron `inject:""`
	Factory    *Factory         `inject:""`
}

func (p *Proxy) GetTaskId(taskIdPostFix string) (taskId string) {
	taskId = p.source + "_" + taskIdPostFix
	return
}

func NewProxy(source, cron string) (proxy Proxy) {
	proxy = Proxy{
		source: source,
		cron:   cron,
	}

	return
}

func (p *Proxy) Init(source, cron string) {
	p.source = source
	p.cron = cron

	p.Factory.name = source
	p.task = p.Factory.Create()
}

func (p *Proxy) Add(options map[string]interface{}) (err error) {
	taskFunc := p.getTaskFunc(options)

	var taskId string
	if v, ok := options["taskId"]; ok {
		taskId = strconv.Itoa(int(v.(uint)))
	}
	err = p.ServerCron.AddCommonTask(p.GetTaskId(taskId), p.cron, taskFunc)

	return
}

func (p Proxy) getTaskFunc(options map[string]interface{}) (taskFunc func()) {
	taskFunc = func() {
		defer func() {
			if err := recover(); err != nil {
				logUtils.Errorf(fmt.Sprintf("%v", err))
			}
		}()

		runFunc := p.task.Run(options)
		//if runFunc() == nil {
		//	return
		//}

		err := runFunc()
		callBackFunc := p.task.CallBack(options, err)
		//if callBackFunc == nil {
		//	return
		//}

		callBackFunc()
	}

	return
}

func Test() {
	options := make(map[string]interface{})
	options["swagger_1"] = 1
	proxy := NewProxy("swagger", "*****")
	proxy.Add(options)
}
