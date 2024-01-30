package task

import (
	"github.com/aaronchen2k/deeptest/internal/server/modules/service"
	third_party "github.com/aaronchen2k/deeptest/internal/server/modules/service/third-party"
)

type Factory struct {
	name   string
	taskId uint
}

//func c Create() (res CronRun) {
//	switch e.name {
//	case "swagger":
//		res = new(template.Swagger)
//	case "lecang":
//		res = new(template.Lecang)
//	default:
//		res = new(template.Swagger)
//	}
//
//	return
//}

func newFactory(name string, task2 Task) (factory *Factory) {
	factory = &Factory{
		name: name,
	}

	return
}

func (e *Factory) Create() (res Task) {
	switch e.name {
	case "swagger":
		res = new(service.SwaggerCron)
	case "lecang":
		res = new(third_party.LecangCron)
	default:
		res = new(service.SwaggerCron)
	}

	return
}
