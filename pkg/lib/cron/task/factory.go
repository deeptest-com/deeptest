package task

import (
	"github.com/aaronchen2k/deeptest/internal/server/modules/service"
)

type Factory struct {
	name              string
	taskId            uint
	SwaggerCron       *service.SwaggerCron       `inject:""`
	LecangCronService *service.LecangCronService `inject:""`
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
		res = e.SwaggerCron
	case "lecang":
		res = e.LecangCronService
	default:
		res = e.SwaggerCron
	}

	return
}
