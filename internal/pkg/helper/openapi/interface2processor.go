package openapi

import (
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	_commUtils "github.com/aaronchen2k/deeptest/pkg/lib/comm"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/jinzhu/copier"
)

type interfaces2Processor struct {
	Inter model.Interface
}

func NewInterfaces2Processor(inter model.Interface) *interfaces2Processor {
	return &interfaces2Processor{Inter: inter}
}

func (i *interfaces2Processor) Convert() (processorInterface *model.ProcessorInterface) {
	processorInterface = new(model.ProcessorInterface)
	copier.CopyWithOption(processorInterface, i.Inter, copier.Option{DeepCopy: true})
	processorInterface.ID = 0
	processorInterface.BodyFormData = i.BodyFormData(i.Inter)
	processorInterface.BodyFormUrlencoded = i.BodyFormUrlencoded(i.Inter)
	processorInterface.Body = i.Body(i.Inter)
	return
}

func (i *interfaces2Processor) BodyFormData(interf model.Interface) (bodyFormData []model.ProcessorInterfaceBodyFormDataItem) {
	return
}

func (i *interfaces2Processor) BodyFormUrlencoded(interf model.Interface) (bodyFormUrlencoded []model.ProcessorInterfaceBodyFormUrlEncodedItem) {
	return
}

func (i *interfaces2Processor) Body(interf model.Interface) (body string) {
	var examples openapi3.Examples
	_commUtils.JsonDecode(interf.RequestBody.Examples, &examples)
	for _, example := range examples {
		return _commUtils.JsonEncode(example.Value)
	}
	return
}
