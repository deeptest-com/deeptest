package service

import (
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/helper/openapi"
	model "github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	"github.com/jinzhu/copier"
)

type DebugInterfaceService struct {
	EndpointInterfaceRepo *repo.EndpointInterfaceRepo `inject:""`
	DebugInterfaceRepo    *repo.DebugInterfaceRepo    `inject:""`
	ServeRepo             *repo.ServeRepo             `inject:""`
	EndpointRepo          *repo.EndpointRepo          `inject:""`

	DebugSceneService *DebugSceneService `inject:""`
}

func (s *DebugInterfaceService) Load(call v1.DebugCall) (req v1.DebugRequest, err error) {
	hasDebugInterfaceRecord, err := s.DebugInterfaceRepo.HasDebugInterfaceRecord(call.InterfaceId)

	if hasDebugInterfaceRecord {
		req, err = s.GetDebugRequestFromPo(call.InterfaceId)
	} else {
		req, err = s.ConvertFromEndpointInterface(call.InterfaceId, call.EndpointId)
	}

	req.BaseUrl, req.ShareVars, req.EnvVars, req.GlobalEnvVars, req.GlobalParamVars =
		s.DebugSceneService.LoadScene(req.InterfaceId, req.EndpointId, req.ProcessorId, req.UsedBy)

	return
}
func (s *DebugInterfaceService) Save(req v1.DebugRequest) (err error) {
	debug := model.DebugInterface{}
	s.CopyValueFromRequest(&debug, req)

	err = s.DebugInterfaceRepo.Update(debug)

	return
}

func (s *DebugInterfaceService) GetDebugRequestFromPo(endpointInterfaceId uint) (req v1.DebugRequest, err error) {
	debugInterface, err := s.DebugInterfaceRepo.GetByEndpointInterfaceId(endpointInterfaceId)
	if err != nil {
		return
	}

	copier.CopyWithOption(&req, &debugInterface, copier.Option{DeepCopy: true})

	return
}

func (s *DebugInterfaceService) ConvertFromEndpointInterface(interfaceId, endpointId uint) (req v1.DebugRequest, err error) {
	var interf model.EndpointInterface

	if interfaceId != 0 {
		interf, err = s.EndpointInterfaceRepo.GetDetail(interfaceId)
	} else if endpointId != 0 {
		interf, err = s.EndpointRepo.GetFirstMethod(endpointId)
	} else {
		return
	}

	if err != nil {
		return
	}

	var endpoint model.Endpoint
	var serve model.Serve

	endpoint, err = s.EndpointRepo.Get(interf.EndpointId)
	serve, err = s.ServeRepo.Get(endpoint.ServeId)
	if err != nil {
		return
	}

	Securities, err := s.ServeRepo.ListSecurity(serve.ID)
	if err != nil {
		return
	}

	serve.Securities = Securities
	interfaces2debug := openapi.NewInterfaces2debug(interf, serve)
	debugInterface := interfaces2debug.Convert()

	copier.CopyWithOption(&req, &debugInterface, copier.Option{DeepCopy: true})

	req.InterfaceId = interfaceId
	req.UsedBy = consts.InterfaceDebug

	return
}

func (s *DebugInterfaceService) CopyValueFromRequest(interf *model.DebugInterface, req v1.DebugRequest) (err error) {
	copier.CopyWithOption(interf, req, copier.Option{DeepCopy: true})

	return
}
