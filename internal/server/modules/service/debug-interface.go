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
	EndpointRepo          *repo.EndpointRepo          `inject:""`
	ServeRepo             *repo.ServeRepo             `inject:""`
	ScenarioProcessorRepo *repo.ScenarioProcessorRepo `inject:""`
	ServeServerRepo       *repo.ServeServerRepo       `inject:""`

	DebugSceneService *DebugSceneService `inject:""`
}

func (s *DebugInterfaceService) Load(loadReq v1.DebugReq) (req v1.DebugData, err error) {
	if loadReq.ScenarioProcessorId > 0 {
		processor, _ := s.ScenarioProcessorRepo.Get(loadReq.ScenarioProcessorId)
		loadReq.EndpointInterfaceId = processor.EndpointInterfaceId
	}

	if loadReq.EndpointInterfaceId == 0 {
		return
	}

	debugInterfaceId, _ := s.DebugInterfaceRepo.HasDebugInterfaceRecord(loadReq.EndpointInterfaceId)

	if debugInterfaceId > 0 {
		req, err = s.GetDebugDataFromDebugInterface(debugInterfaceId)
	} else {
		req, err = s.ConvertDebugDataFromEndpointInterface(loadReq.EndpointInterfaceId)
	}

	req.BaseUrl, req.ShareVars, req.EnvVars, req.GlobalEnvVars, req.GlobalParamVars =
		s.DebugSceneService.LoadScene(req.EndpointInterfaceId, req.ScenarioProcessorId, req.UsedBy)

	req.ScenarioProcessorId = loadReq.ScenarioProcessorId
	req.UsedBy = loadReq.UsedBy

	return
}
func (s *DebugInterfaceService) Save(req v1.DebugData) (debug model.DebugInterface, err error) {
	s.CopyValueFromRequest(&debug, req)

	endpointInterface, _ := s.EndpointInterfaceRepo.Get(req.EndpointInterfaceId)
	debug.EndpointId = endpointInterface.EndpointId

	debugInterfaceId, _ := s.DebugInterfaceRepo.HasDebugInterfaceRecord(debug.EndpointInterfaceId)
	if debugInterfaceId > 0 {
		debug.ID = debugInterfaceId
	}

	err = s.DebugInterfaceRepo.Save(&debug)

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

func (s *DebugInterfaceService) ConvertDebugDataFromEndpointInterface(endpointInterfaceId uint) (req v1.DebugData, err error) {
	endpointInterface, err := s.EndpointInterfaceRepo.GetDetail(endpointInterfaceId)
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
	req.EndpointInterfaceId = endpointInterface.EndpointId // reset

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
