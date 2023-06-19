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
	ScenarioInterfaceRepo *repo.ScenarioInterfaceRepo `inject:""`
	TestInterfaceRepo     *repo.TestInterfaceRepo     `inject:""`

	EndpointRepo          *repo.EndpointRepo          `inject:""`
	ServeRepo             *repo.ServeRepo             `inject:""`
	ScenarioProcessorRepo *repo.ScenarioProcessorRepo `inject:""`
	ServeServerRepo       *repo.ServeServerRepo       `inject:""`

	DebugSceneService        *DebugSceneService        `inject:""`
	ScenarioInterfaceService *ScenarioInterfaceService `inject:""`
	SceneService             *SceneService             `inject:""`
	EnvironmentService       *EnvironmentService       `inject:""`
	DatapoolService          *DatapoolService          `inject:""`
}

func (s *DebugInterfaceService) Load(loadReq domain.DebugReq) (debugData domain.DebugData, err error) {
	if loadReq.ScenarioProcessorId > 0 {
		processor, _ := s.ScenarioProcessorRepo.Get(loadReq.ScenarioProcessorId)
		loadReq.EndpointInterfaceId = processor.EndpointInterfaceId
	}

	if loadReq.EndpointInterfaceId == 0 && loadReq.TestInterfaceId == 0 {
		return
	}

	if loadReq.ScenarioProcessorId > 0 || loadReq.UsedBy == consts.ScenarioDebug {
		debugData, _ = s.ScenarioInterfaceService.GetScenarioInterface(loadReq.EndpointInterfaceId)
	} else if loadReq.EndpointInterfaceId > 0 {
		debugData, _ = s.GetDebugInterfaceByEndpointInterface(loadReq.EndpointInterfaceId)
	} else if loadReq.TestInterfaceId > 0 {
		debugData, _ = s.GetDebugInterfaceByTestInterface(loadReq.TestInterfaceId)
	}

	debugData.UsedBy = loadReq.UsedBy
	if loadReq.ScenarioProcessorId > 0 {
		debugData.ScenarioProcessorId = loadReq.ScenarioProcessorId
	}

	debugData.BaseUrl, debugData.ShareVars, debugData.EnvVars =
		s.DebugSceneService.LoadScene(debugData.EndpointInterfaceId, debugData.ServerId, debugData.ScenarioProcessorId, debugData.UsedBy)

	return
}

func (s *DebugInterfaceService) LoadForExec(loadReq domain.DebugReq) (ret agentExec.InterfaceExecObj, err error) {
	ret.DebugData, _ = s.Load(loadReq)

	ret.ExecScene.ShareVars = ret.DebugData.ShareVars
	ret.DebugData.ShareVars = nil // for display on debug page only

	// get project environment
	projectId, _ := s.SceneService.LoadEnvVarMapByEndpointInterface(&ret.ExecScene,
		ret.DebugData.EndpointInterfaceId, ret.DebugData.ServerId)

	s.SceneService.LoadProjectSettings(&ret.ExecScene, projectId)

	return
}

func (s *DebugInterfaceService) GetDebugInterfaceByEndpointInterface(endpointInterfaceId uint) (ret domain.DebugData, err error) {
	debugInterfaceId, _ := s.DebugInterfaceRepo.HasDebugInterfaceRecord(endpointInterfaceId)

	if debugInterfaceId > 0 {
		ret, err = s.GetDebugDataFromDebugInterface(debugInterfaceId)
	} else {
		ret, err = s.ConvertDebugDataFromEndpointInterface(endpointInterfaceId)
	}

	return
}

func (s *DebugInterfaceService) GetDebugInterfaceByTestInterface(testInterfaceId uint) (ret domain.DebugData, err error) {
	testInterface, err := s.TestInterfaceRepo.Get(testInterfaceId)
	if err != nil {
		return
	}

	copier.CopyWithOption(&ret, testInterface, copier.Option{
		DeepCopy: true,
	})

	if ret.ServerId <= 0 {
		server, _ := s.ServeServerRepo.GetDefaultByServe(testInterface.ServeId)
		ret.ServerId = server.ID
	}

	ret.TestInterfaceId = testInterfaceId

	return
}

func (s *DebugInterfaceService) GetDetail(id uint) (ret model.DebugInterface, err error) {
	ret, err = s.DebugInterfaceRepo.GetDetail(id)

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

func (s *DebugInterfaceService) GetDebugDataFromDebugInterface(debugInterfaceId uint) (ret domain.DebugData, err error) {
	debugInterfacePo, err := s.DebugInterfaceRepo.GetDetail(debugInterfaceId)
	if err != nil {
		return
	}

	endpointInterface, _ := s.EndpointInterfaceRepo.Get(debugInterfacePo.EndpointInterfaceId)

	s.SetProps(endpointInterface, &debugInterfacePo, &ret)

	return
}

func (s *DebugInterfaceService) ConvertDebugDataFromEndpointInterface(endpointInterfaceId uint) (debugData domain.DebugData, err error) {
	endpointInterface, err := s.EndpointInterfaceRepo.GetDetail(endpointInterfaceId)
	if err != nil {
		return
	}

	s.SetProps(endpointInterface, nil, &debugData)

	endpoint, _ := s.EndpointRepo.Get(endpointInterface.EndpointId)
	debugData.ServerId = endpoint.ServerId

	debugData.UsedBy = consts.InterfaceDebug

	return
}

func (s *DebugInterfaceService) SetProps(
	endpointInterface model.EndpointInterface, debugInterfacePo *model.DebugInterface, debugData *domain.DebugData) {

	endpoint, err := s.EndpointRepo.GetAll(endpointInterface.EndpointId, "v0.1.0")
	serve, err := s.ServeRepo.Get(endpoint.ServeId)
	if err != nil {
		return
	}

	securities, err := s.ServeRepo.ListSecurity(serve.ID)
	if err != nil {
		return
	}

	serve.Securities = securities
	debugData.EndpointInterfaceId = endpointInterface.ID

	if debugInterfacePo == nil {
		interfaces2debug := openapi.NewInterfaces2debug(endpointInterface, endpoint, serve)
		debugInterfacePo = interfaces2debug.Convert()

		debugInterfacePo.Name = fmt.Sprintf("%s - %s", endpoint.Title, debugInterfacePo.Method)
	}

	copier.CopyWithOption(&debugData, debugInterfacePo, copier.Option{DeepCopy: true})
	debugData.EndpointInterfaceId = endpointInterface.ID // reset

	debugData.Headers = append(debugData.Headers, domain.Header{Name: "", Value: ""})
	debugData.QueryParams = append(debugData.QueryParams, domain.Param{Name: "", Value: ""})
	debugData.PathParams = append(debugData.PathParams, domain.Param{Name: "", Value: ""})

	debugData.BodyFormData = append(debugData.BodyFormData, domain.BodyFormDataItem{
		Name: "", Value: "", Type: consts.FormDataTypeText})
	debugData.BodyFormUrlencoded = append(debugData.BodyFormUrlencoded, domain.BodyFormUrlEncodedItem{
		Name: "", Value: "",
	})

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

func (s *DebugInterfaceService) GenSample(projectId, serveId uint) (ret *model.DebugInterface, err error) {
	return
}
