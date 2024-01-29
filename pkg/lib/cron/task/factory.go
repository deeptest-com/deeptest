package task

import "github.com/aaronchen2k/deeptest/internal/server/modules/service"

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

func newFactory(name string, task2 task) (factory *Factory) {
	factory = &Factory{
		name: name,
	}

	return
}

func (e *Factory) Create() (res task) {
	switch e.name {
	case "swagger":
		res = new(service.Swagger)
	case "lecang":
		res = new(service.Lecang)
	default:
		res = new(service.Swagger)
	}

	return
}
