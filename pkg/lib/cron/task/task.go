package task

import (
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/core/cron"
	"github.com/aaronchen2k/deeptest/internal/server/modules/service"

	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
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

func (p *Proxy) Init(tenantId consts.TenantId, source consts.CronSource, callBack func(tenantId consts.TenantId, taskId string, source consts.CronSource, err error) error, taskId, cron string) {
	p.tenantId = tenantId
	p.source = source
	p.cron = cron
	p.taskId = taskId
	p.Factory.name = source
	p.callBack = callBack
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

			p.callBack(p.tenantId, p.taskId, p.source, ret)

		}()

		runFunc := p.task.Run(options)
		//if runFunc() == nil {
		//	return
		//}

		runFunc()
		//callBackFunc := p.task.CallBack(options, err)
		//if callBackFunc == nil {
		//	return
		//}

		//callBackFunc()
	}

	return
}

func Test() {
	options := make(map[string]interface{})
	options["swagger_1"] = 1
	//proxy := NewProxy("swagger", "*****")
	//proxy.Add(options)
}
