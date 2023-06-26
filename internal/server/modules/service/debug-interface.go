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
	ExtractorRepo         *repo.ExtractorRepo         `inject:""`
	CheckpointRepo        *repo.CheckpointRepo        `inject:""`

	DebugSceneService        *DebugSceneService        `inject:""`
	ScenarioInterfaceService *ScenarioInterfaceService `inject:""`
	SceneService             *SceneService             `inject:""`
	EnvironmentService       *EnvironmentService       `inject:""`
	DatapoolService          *DatapoolService          `inject:""`
}

func (s *DebugInterfaceService) Load(loadReq domain.DebugReq) (debugData domain.DebugData, err error) {
	if loadReq.DebugInterfaceId > 0 {
		debugData, _ = s.GetDebugDataFromDebugInterface(loadReq.DebugInterfaceId)
	} else {
		if loadReq.ScenarioProcessorId > 0 {
			debugData, _ = s.GetDebugInterfaceByScenarioInterface(loadReq.ScenarioProcessorId)
		} else if loadReq.TestInterfaceId > 0 {
			debugData, _ = s.GetDebugInterfaceByTestInterface(loadReq.TestInterfaceId)
		} else if loadReq.EndpointInterfaceId > 0 {
			debugData, _ = s.GetDebugInterfaceByEndpointInterface(loadReq.EndpointInterfaceId)
		}
	}

	debugData.UsedBy = loadReq.UsedBy

	debugData.BaseUrl, debugData.ShareVars, debugData.EnvVars, debugData.GlobalVars, debugData.GlobalParams =
		s.DebugSceneService.LoadScene(debugData, debugData.UsedBy)

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
	endpointInterface, _ := s.EndpointInterfaceRepo.Get(endpointInterfaceId)

	if endpointInterface.DebugInterfaceId > 0 {
		ret, err = s.GetDebugDataFromDebugInterface(endpointInterface.DebugInterfaceId)
	} else {
		ret, err = s.ConvertDebugDataFromEndpointInterface(endpointInterfaceId)
	}

	return
}

func (s *DebugInterfaceService) GetDebugInterfaceByScenarioInterface(scenarioProcessorId uint) (ret domain.DebugData, err error) {
	processor, err := s.ScenarioProcessorRepo.Get(scenarioProcessorId)
	if err != nil {
		return
	}

	debugInterfaceId := processor.EntityId
	debugData, _ := s.DebugInterfaceRepo.GetDetail(debugInterfaceId)

	copier.CopyWithOption(&ret, debugData, copier.Option{
		DeepCopy: true,
	})

	if ret.ServerId <= 0 {
		server, _ := s.ServeServerRepo.GetByEndpoint(processor.EndpointInterfaceId)
		ret.ServerId = server.ID
	}

	ret.DebugInterfaceId = debugInterfaceId
	ret.ScenarioProcessorId = scenarioProcessorId

	ret.Headers = append(ret.Headers, domain.Header{Name: "", Value: ""})
	ret.QueryParams = append(ret.QueryParams, domain.Param{Name: "", Value: ""})
	ret.PathParams = append(ret.PathParams, domain.Param{Name: "", Value: ""})

	ret.BodyFormData = append(ret.BodyFormData, domain.BodyFormDataItem{
		Name: "", Value: "", Type: consts.FormDataTypeText})
	ret.BodyFormUrlencoded = append(ret.BodyFormUrlencoded, domain.BodyFormUrlEncodedItem{
		Name: "", Value: "",
	})

	return
}

func (s *DebugInterfaceService) GetDebugInterfaceByTestInterface(testInterfaceId uint) (ret domain.DebugData, err error) {
	testInterface, err := s.TestInterfaceRepo.GetDetail(testInterfaceId)
	if err != nil {
		return
	}

	copier.CopyWithOption(&ret, testInterface.DebugData, copier.Option{
		DeepCopy: true,
	})

	if ret.ServerId <= 0 {
		server, _ := s.ServeServerRepo.GetDefaultByServe(testInterface.ServeId)
		ret.ServerId = server.ID
	}

	ret.DebugInterfaceId = testInterface.DebugInterfaceId
	ret.TestInterfaceId = testInterfaceId

	ret.Headers = append(ret.Headers, domain.Header{Name: "", Value: ""})
	ret.QueryParams = append(ret.QueryParams, domain.Param{Name: "", Value: ""})
	ret.PathParams = append(ret.PathParams, domain.Param{Name: "", Value: ""})

	ret.BodyFormData = append(ret.BodyFormData, domain.BodyFormDataItem{
		Name: "", Value: "", Type: consts.FormDataTypeText})
	ret.BodyFormUrlencoded = append(ret.BodyFormUrlencoded, domain.BodyFormUrlEncodedItem{
		Name: "", Value: "",
	})

	return
}

func (s *DebugInterfaceService) GetDetail(id uint) (ret model.DebugInterface, err error) {
	ret, err = s.DebugInterfaceRepo.GetDetail(id)

	return
}

func (s *DebugInterfaceService) Save(req domain.DebugData) (debugInterface model.DebugInterface, err error) {
	s.CopyValueFromRequest(&debugInterface, req)

	if req.DebugInterfaceId > 0 { // to update
		debugInterface.ID = req.DebugInterfaceId
	} else {
		debugInterface.ServeId = req.ServeId
	}

	err = s.DebugInterfaceRepo.Save(&debugInterface)

	if req.DebugInterfaceId <= 0 { // first time to save
		// clone extractors and checkpoints if needed
		s.ExtractorRepo.CloneFromEndpointInterfaceToDebugInterface(req.EndpointInterfaceId, debugInterface.ID, req.UsedBy)
		s.CheckpointRepo.CloneFromEndpointInterfaceToDebugInterface(req.EndpointInterfaceId, debugInterface.ID, req.UsedBy)

		s.EndpointInterfaceRepo.SetDebugInterfaceId(req.EndpointInterfaceId, debugInterface.ID)
	}

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
	debugData.ServeId = endpoint.ServeId
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

	debugData.DebugInterfaceId = debugInterfacePo.ID

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
