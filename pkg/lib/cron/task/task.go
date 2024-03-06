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
	source             string
	cron               string
	task               Task
	taskId             string
	tenantId           consts.TenantId
	ServerCron         *cron.ServerCron            `inject:""`
	Factory            *Factory                    `inject:""`
	ProjectCronService *service.ProjectCronService `inject:""`
}

func (p *Proxy) GetTaskId() (taskId string) {
	taskId = fmt.Sprintf("%s_%s_%s", p.source, p.tenantId, p.taskId)
	return
}

func NewProxy(source, cron string) (proxy Proxy) {
	proxy = Proxy{
		source: source,
		cron:   cron,
	}

	return
}

func (p *Proxy) Init(tenantId consts.TenantId, source, taskId, cron string) {
	p.tenantId = tenantId
	p.source = source
	p.cron = cron
	p.taskId = taskId
	p.Factory.name = source
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
