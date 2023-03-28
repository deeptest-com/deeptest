package openapi

import (
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	_commUtils "github.com/aaronchen2k/deeptest/pkg/lib/comm"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/jinzhu/copier"
)

type interfaces2debug struct {
	Inter model.EndpointInterface
}

func NewInterfaces2debug(inter model.EndpointInterface) *interfaces2debug {
	return &interfaces2debug{Inter: inter}
}

func (i *interfaces2debug) Convert() (debugInterface *model.DebugInterface) {
	debugInterface = new(model.DebugInterface)
	copier.CopyWithOption(debugInterface, &i.Inter, copier.Option{DeepCopy: true})
	debugInterface.ID = 0
	debugInterface.BodyFormData = i.BodyFormData(i.Inter)
	debugInterface.BodyFormUrlencoded = i.BodyFormUrlencoded(i.Inter)
	debugInterface.Body = i.Body(i.Inter)
	return
}

func (i *interfaces2debug) BodyFormData(interf model.EndpointInterface) (bodyFormData []model.DebugInterfaceBodyFormDataItem) {
	return
}

func (i *interfaces2debug) BodyFormUrlencoded(interf model.EndpointInterface) (bodyFormUrlencoded []model.DebugInterfaceBodyFormUrlEncodedItem) {
	return
}

func (i *interfaces2debug) Body(interf model.EndpointInterface) (body string) {
	var examples openapi3.Examples
	_commUtils.JsonDecode(interf.RequestBody.Examples, &examples)
	for _, example := range examples {
		return _commUtils.JsonEncode(example.Value)
	}
	return
}
