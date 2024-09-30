package task

import (
	"fmt"
	"github.com/deeptest-com/deeptest/internal/pkg/consts"
	"github.com/deeptest-com/deeptest/internal/pkg/core/cron"
	"github.com/deeptest-com/deeptest/internal/server/modules/service"
	logUtils "github.com/deeptest-com/deeptest/pkg/lib/log"
)

type Task interface {
	Run(options map[string]interface{}) (f func() error)
	CallBack(options map[string]interface{}, err error) func()
}

type Proxy struct {
	source             consts.CronSource
	cron               string
	task               Task
	taskId             string
	beforeCall         func(tenantId consts.TenantId, taskId string, source consts.CronSource) error
	callBack           func(tenantId consts.TenantId, taskId string, source consts.CronSource, err error) error
	tenantId           consts.TenantId
	ServerCron         *cron.ServerCron            `inject:""`
	Factory            *Factory                    `inject:""`
	ProjectCronService *service.ProjectCronService `inject:""`
}

func (p *Proxy) GetTaskId() (taskId string) {
	taskId = fmt.Sprintf("%s_%s_%s", p.source, p.tenantId, p.taskId)
	return
}

//func NewProxy(source, cron string, f func(tenantId consts.TenantId, taskId string, source string, err error)) (proxy Proxy) {
//	proxy = Proxy{
//		source:   source,
//		cron:     cron,
//		callBack: f,
//	}
//
//	return
//}

func (p *Proxy) Init(tenantId consts.TenantId, source consts.CronSource, beforeCall func(tenantId consts.TenantId, taskId string, source consts.CronSource) error, callBack func(tenantId consts.TenantId, taskId string, source consts.CronSource, err error) error, taskId, cron string) {
	p.tenantId = tenantId
	p.source = source
	p.cron = cron
	p.taskId = taskId
	p.Factory.name = source
	p.callBack = callBack
	p.beforeCall = beforeCall
	p.task = p.Factory.Create()
}

func (p *Proxy) Add(options map[string]interface{}) (err error) {
	taskFunc := p.getTaskFunc(options)

	err = p.ServerCron.AddCommonTask(p.GetTaskId(), p.cron, taskFunc)

	return
}

func (p *Proxy) Remove() {
	taskId := p.GetTaskId()
	p.ServerCron.RemoveTask(taskId)
}

func (p Proxy) getTaskFunc(options map[string]interface{}) (taskFunc func()) {
	taskFunc = func() {
		defer func() {
			var ret error
			err := recover()
			if err != nil {
				logUtils.Errorf(fmt.Sprintf("%v", err))
				ret = fmt.Errorf("%v", err)
			}

			if p.callBack != nil {
				p.callBack(p.tenantId, p.taskId, p.source, ret) //后置处理任务，处理调用结果
			}

		}()

		if p.beforeCall != nil { // 前置处理任务，初始化任务状态为执行中，防止任务重复执行
			p.beforeCall(p.tenantId, p.taskId, p.source)
		}

		runFunc := p.task.Run(options)

		runFunc()

	}

	return
}

func (p Proxy) Run(options map[string]interface{}) {
	f := p.getTaskFunc(options)
	//异步执行
	go f()
}

func Test() {
	options := make(map[string]interface{})
	options["swagger_1"] = 1
	//proxy := NewProxy("swagger", "*****")
	//proxy.Add(options)
}
