package task

type Task interface {
	Run(options map[string]interface{}) (f func())
}

type Proxy struct {
	Name   string `json:"name"`
	Cron   string `json:"cron"`
	task   Task
	TaskId string `json:"taskId"`
}

func (p *Proxy) GetTaskId() (taskId string) {
	//自己写
	return
}

func NewProxy(name, cron, taskId string) (proxy Proxy) {
	proxy = Proxy{
		Name: name,
		Cron: cron,
	}

	taskEntity := Factory{
		name: name,
	}

	proxy.task = taskEntity.Create()
	return
}

func (p *Proxy) Run(options map[string]interface{}) (err error) {
	//function := p.task.Run(options)
	//taskId := p.GetTaskId()
	//cron := p.Cron
	// TODO 调用 AddCommonTask
	return
}

func Test() {
	options := make(map[string]interface{})
	options["swagger_1"] = 1
	proxy := NewProxy("swagger", "*****", "taskId")
	proxy.Run(options)
}
