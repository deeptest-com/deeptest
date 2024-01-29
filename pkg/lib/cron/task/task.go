package task

type task interface {
	Run() (f func())
	GetTaskId() (taskId interface{})
}

type Proxy struct {
	Name    string                 `json:"name"`
	Cron    string                 `json:"cron"`
	Options map[string]interface{} `json:"options"`
	task    task
}

func NewProxy(name, cron string, options map[string]interface{}) (proxy Proxy) {
	proxy = Proxy{
		Name:    name,
		Cron:    cron,
		Options: options,
	}

	taskEntity := Factory{
		name: name,
	}

	proxy.task = taskEntity.Create()
	return
}

func (p *Proxy) Run() (err error) {
	//function := p.task.Run()
	//taskId := p.task.GetTaskId()
	//cron := p.Cron
	// TODO 调用 AddCommonTask
	return
}

func Test() {
	options := make(map[string]interface{})
	options["swagger_1"] = 1
	proxy := NewProxy("swagger", "*****", options)
	proxy.Run()
}
