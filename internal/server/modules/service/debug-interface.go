package service

import (
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/helper/openapi"
	model "github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	_httpUtils "github.com/aaronchen2k/deeptest/pkg/lib/http"
	"github.com/jinzhu/copier"
)

type DebugInterfaceService struct {
	EndpointInterfaceRepo *repo.EndpointInterfaceRepo `inject:""`
	DebugInterfaceRepo    *repo.DebugInterfaceRepo    `inject:""`
	EndpointRepo          *repo.EndpointRepo          `inject:""`
	ServeRepo             *repo.ServeRepo             `inject:""`
	ScenarioProcessorRepo *repo.ScenarioProcessorRepo `inject:""`
	ServeServerRepo       *repo.ServeServerRepo       `inject:""`

	DebugSceneService *DebugSceneService `inject:""`
}

func (s *DebugInterfaceService) Load(defineReq v1.DebugReq) (req v1.DebugData, err error) {
	var debugInterfaceId uint

	if defineReq.EndpointInterfaceId > 0 {
		debugInterfaceId, _ = s.DebugInterfaceRepo.HasDebugInterfaceRecord(defineReq.EndpointInterfaceId)

		if debugInterfaceId > 0 {
			req, err = s.GetDebugDataFromDebugInterface(debugInterfaceId)
		} else {
			req, err = s.ConvertDebugDataFromEndpointInterface(defineReq.EndpointInterfaceId, 0)
		}
		req.UsedBy = consts.InterfaceDebug

		req.BaseUrl, req.ShareVars, req.EnvVars, req.GlobalEnvVars, req.GlobalParamVars =
			s.DebugSceneService.LoadScene(req.EndpointInterfaceId, req.ScenarioProcessorId, req.UsedBy)
	} else if defineReq.ScenarioProcessorId > 0 {
		// TODO:
		req.UsedBy = consts.ScenarioDebug
	}

	return
}
func (s *DebugInterfaceService) Save(req v1.DebugData) (err error) {
	debug := model.DebugInterface{}
	s.CopyValueFromRequest(&debug, req)

	debugInterfaceId, _ := s.DebugInterfaceRepo.HasDebugInterfaceRecord(debug.EndpointInterfaceId)
	if debugInterfaceId > 0 {
		debug.ID = debugInterfaceId
	}

	err = s.DebugInterfaceRepo.Save(debug)

	return
}

func (s *DebugInterfaceService) GetDebugDataFromDebugInterface(debugInterfaceId uint) (req v1.DebugData, err error) {
	debugInterface, _ := s.DebugInterfaceRepo.GetDetail(debugInterfaceId)
	if err != nil {
		return
	}

	endpointInterface, _ := s.EndpointInterfaceRepo.Get(debugInterface.EndpointInterfaceId)

	s.SetProps(&endpointInterface, &debugInterface, &req)

	return
}

func (s *DebugInterfaceService) ConvertDebugDataFromEndpointInterface(endpointInterfaceId, endpointId uint) (req v1.DebugData, err error) {
	var endpointInterface model.EndpointInterface

	if endpointInterfaceId != 0 {
		endpointInterface, err = s.EndpointInterfaceRepo.GetDetail(endpointInterfaceId)
	} else if endpointId != 0 {
		endpointInterface, err = s.EndpointRepo.GetFirstMethod(endpointId)
	} else {
		return
	}

	if err != nil {
		return
	}

	s.SetProps(&endpointInterface, nil, &req)

	req.UsedBy = consts.InterfaceDebug

	return
}

func (s *DebugInterfaceService) SetProps(
	endpointInterface *model.EndpointInterface, debugInterface *model.DebugInterface, req *v1.DebugData) {

	endpoint, err := s.EndpointRepo.Get(endpointInterface.EndpointId)
	serve, err := s.ServeRepo.Get(endpoint.ServeId)
	if err != nil {
		return
	}

	Securities, err := s.ServeRepo.ListSecurity(serve.ID)
	if err != nil {
		return
	}

	serve.Securities = Securities
	req.EndpointInterfaceId = endpointInterface.ID

	if debugInterface == nil {
		interfaces2debug := openapi.NewInterfaces2debug(*endpointInterface, serve)
		debugInterface = interfaces2debug.Convert()
	}

	copier.CopyWithOption(&req, &debugInterface, copier.Option{DeepCopy: true})

	if debugInterface == nil {
		serveServer, _ := s.ServeServerRepo.GetByEndpoint(endpointInterface.EndpointId)
		req.Url = _httpUtils.AddSepIfNeeded(serveServer.Url) + req.Url
	}

	return
}

func (s *DebugInterfaceService) GetEndpointAndServeIdForEndpointInterface(endpointInterfaceId uint) (
	endpointId, serveId uint) {

	endpointInterface, _ := s.EndpointInterfaceRepo.Get(endpointInterfaceId)

	endpointId = endpointInterface.EndpointId
	endpoint, _ := s.EndpointRepo.Get(endpointId)

	serveId = endpoint.ServeId

	return
}
func (s *DebugInterfaceService) GetEndpointAndServeIdForDebugInterface(debugInterfaceId uint) (
	endpointId, serveId uint) {

	debugInterface, _ := s.DebugInterfaceRepo.Get(debugInterfaceId)

	endpointId = debugInterface.EndpointId
	endpoint, _ := s.EndpointRepo.Get(endpointId)

	serveId = endpoint.ServeId

	return
}
func (s *DebugInterfaceService) GetScenarioIdForDebugInterface(processorId uint) (
	scenarioId uint) {

	processor, _ := s.ScenarioProcessorRepo.Get(processorId)
	scenarioId = processor.ScenarioId

	return
}

func (s *DebugInterfaceService) CopyValueFromRequest(interf *model.DebugInterface, req v1.DebugData) (err error) {
	copier.CopyWithOption(interf, req, copier.Option{DeepCopy: true})

	return
}
