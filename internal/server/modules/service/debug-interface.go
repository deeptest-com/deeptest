package service

import (
	"fmt"
	agentExec "github.com/aaronchen2k/deeptest/internal/agent/exec"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
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

	DebugSceneService  *DebugSceneService  `inject:""`
	SceneService       *SceneService       `inject:""`
	EnvironmentService *EnvironmentService `inject:""`
	DatapoolService    *DatapoolService    `inject:""`
}

func (s *DebugInterfaceService) Load(loadReq domain.DebugReq) (ret agentExec.InterfaceExecObj, err error) {
	if loadReq.ScenarioProcessorId > 0 {
		processor, _ := s.ScenarioProcessorRepo.Get(loadReq.ScenarioProcessorId)
		loadReq.EndpointInterfaceId = processor.EndpointInterfaceId
	}

	if loadReq.EndpointInterfaceId == 0 {
		return
	}

	// gen debug data
	debugData, _ := s.GetDebugInterface(loadReq.EndpointInterfaceId)

	debugData.UsedBy = loadReq.UsedBy
	if loadReq.ScenarioProcessorId > 0 {
		debugData.ScenarioProcessorId = loadReq.ScenarioProcessorId
	}

	debugData.BaseUrl, debugData.ShareVars, debugData.EnvVars, debugData.GlobalEnvVars, debugData.GlobalParamVars =
		s.DebugSceneService.LoadScene(debugData.EndpointInterfaceId, debugData.ScenarioProcessorId, debugData.UsedBy)

	debugData.ScenarioProcessorId = loadReq.ScenarioProcessorId
	debugData.UsedBy = loadReq.UsedBy

	ret.DebugData = debugData

	// get variables
	var projectId uint
	ret.EnvToVariables, ret.InterfaceToEnvMap, projectId, _ = s.SceneService.LoadEnvVarMapByEndpointInterface(debugData.EndpointInterfaceId)
	ret.GlobalVars, _ = s.EnvironmentService.GetGlobalVars(projectId)
	ret.GlobalParams, _ = s.EnvironmentService.GetGlobalParams(projectId)

	ret.Datapools, _ = s.DatapoolService.ListForExec(projectId)

	return
}

func (s *DebugInterfaceService) GetDebugInterface(endpointInterfaceId uint) (ret domain.DebugData, err error) {
	debugInterfaceId, _ := s.DebugInterfaceRepo.HasDebugInterfaceRecord(endpointInterfaceId)

	if debugInterfaceId > 0 {
		ret, err = s.GetDebugDataFromDebugInterface(debugInterfaceId)
	} else {
		ret, err = s.ConvertDebugDataFromEndpointInterface(endpointInterfaceId)
	}

	return
}

func (s *DebugInterfaceService) Save(req domain.DebugData) (debug model.DebugInterface, err error) {
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

func (s *DebugInterfaceService) GetDebugDataFromDebugInterface(debugInterfaceId uint) (req domain.DebugData, err error) {
	debugInterface, _ := s.DebugInterfaceRepo.GetDetail(debugInterfaceId)
	if err != nil {
		return
	}

	endpointInterface, _ := s.EndpointInterfaceRepo.Get(debugInterface.EndpointInterfaceId)

	s.SetProps(endpointInterface, &debugInterface, &req)

	return
}

func (s *DebugInterfaceService) ConvertDebugDataFromEndpointInterface(endpointInterfaceId uint) (req domain.DebugData, err error) {
	endpointInterface, err := s.EndpointInterfaceRepo.GetDetail(endpointInterfaceId)
	if err != nil {
		return
	}

	s.SetProps(endpointInterface, nil, &req)

	req.UsedBy = consts.InterfaceDebug

	return
}

func (s *DebugInterfaceService) SetProps(
	endpointInterface model.EndpointInterface, debugInterface *model.DebugInterface, req *domain.DebugData) {

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
		interfaces2debug := openapi.NewInterfaces2debug(endpointInterface, serve)
		debugInterface = interfaces2debug.Convert()

		debugInterface.Name = fmt.Sprintf("%s - %s", endpoint.Title, debugInterface.Method)
	}

	copier.CopyWithOption(&req, &debugInterface, copier.Option{DeepCopy: true})
	req.EndpointInterfaceId = endpointInterface.ID // reset

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

func (s *DebugInterfaceService) CopyValueFromRequest(interf *model.DebugInterface, req domain.DebugData) (err error) {
	copier.CopyWithOption(interf, req, copier.Option{DeepCopy: true})

	return
}
