package service

import (
	"errors"
	serverDomain "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	"github.com/jinzhu/copier"
)

type ScenarioInterfaceService struct {
	EndpointInterfaceRepo *repo.EndpointInterfaceRepo `inject:""`
	DebugInterfaceRepo    *repo.DebugInterfaceRepo    `inject:""`
	ScenarioInterfaceRepo *repo.ScenarioInterfaceRepo `inject:""`
	EndpointRepo          *repo.EndpointRepo          `inject:""`
	ServeRepo             *repo.ServeRepo             `inject:""`
	ScenarioProcessorRepo *repo.ScenarioProcessorRepo `inject:""`
	ServeServerRepo       *repo.ServeServerRepo       `inject:""`
	DiagnoseInterfaceRepo *repo.DiagnoseInterfaceRepo `inject:""`
	EndpointCaseRepo      *repo.EndpointCaseRepo      `inject:""`
	DebugInvokeRepo       *repo.DebugInvokeRepo       `inject:""`

	ScenarioNodeService   *ScenarioNodeService   `inject:""`
	DebugSceneService     *DebugSceneService     `inject:""`
	DebugInterfaceService *DebugInterfaceService `inject:""`
	SceneService          *SceneService          `inject:""`
	EnvironmentService    *EnvironmentService    `inject:""`
	DatapoolService       *DatapoolService       `inject:""`
	EndpointCaseService   *EndpointCaseService   `inject:""`
}

func (s *ScenarioInterfaceService) GetDebugDataFromScenarioInterface(scenarioInterfaceId uint) (req domain.DebugData, err error) {
	scenarioInterfacePo, _ := s.ScenarioInterfaceRepo.GetDetail(scenarioInterfaceId)
	if err != nil {
		return
	}

	endpointInterface, _ := s.EndpointInterfaceRepo.Get(scenarioInterfacePo.EndpointInterfaceId)

	s.SetProps(endpointInterface, &scenarioInterfacePo, &req)

	return
}

func (s *ScenarioInterfaceService) SetProps(
	endpointInterface model.EndpointInterface, scenarioInterfacePo *model.DebugInterface, debugData *domain.DebugData) {

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

	copier.CopyWithOption(&debugData, scenarioInterfacePo, copier.Option{DeepCopy: true})
	debugData.EndpointInterfaceId = endpointInterface.ID // reset

	debugData.Headers = append(debugData.Headers, domain.Header{Name: "", Value: ""})
	debugData.QueryParams = append(debugData.QueryParams, domain.Param{Name: "", Value: "", ParamIn: consts.ParamInQuery})
	debugData.PathParams = append(debugData.PathParams, domain.Param{Name: "", Value: "", ParamIn: consts.ParamInPath})

	debugData.BodyFormData = append(debugData.BodyFormData, domain.BodyFormDataItem{
		Name: "", Value: "", Type: consts.FormDataTypeText})
	debugData.BodyFormUrlencoded = append(debugData.BodyFormUrlencoded, domain.BodyFormUrlEncodedItem{
		Name: "", Value: "",
	})

	return
}

//func (s *ScenarioInterfaceService) GetScenarioInterface(endpointInterfaceId uint) (ret domain.DebugData, err error) {
//	scenarioInterfaceId, _ := s.ScenarioInterfaceRepo.HasScenarioInterfaceRecord(endpointInterfaceId)
//
//	if scenarioInterfaceId > 0 {
//		ret, err = s.GetDebugDataFromScenarioInterface(scenarioInterfaceId)
//	} else {
//		ret, err = s.DebugInterfaceService.GetDebugDataFromEndpointInterface(endpointInterfaceId)
//		if err != nil || ret.EndpointInterfaceId == 0 {
//			return domain.DebugData{}, err
//		}
//		_, err = s.SaveDebugData(ret)
//	}
//
//	return
//}

func (s *ScenarioInterfaceService) SaveDebugData(req domain.DebugData) (debug model.DebugInterface, err error) {
	s.DebugInterfaceService.CopyValueFromRequest(&debug, req)

	//endpointInterface, _ := s.EndpointInterfaceRepo.Get(req.EndpointInterfaceId)
	//debug.EndpointId = endpointInterface.EndpointId

	if req.DebugInterfaceId > 0 {
		debug.ID = req.DebugInterfaceId
	}

	err = s.ScenarioInterfaceRepo.SaveDebugData(&debug)

	//更新执行器method
	s.ScenarioProcessorRepo.UpdateMethod(debug.ScenarioProcessorId, debug.Method)

	return
}

func (s *ScenarioInterfaceService) ResetDebugData(scenarioProcessorId int, createBy uint) (
	newProcessor model.Processor, err error) {

	scenarioProcessor, _ := s.ScenarioProcessorRepo.Get(uint(scenarioProcessorId))
	parentProcessor, _ := s.ScenarioProcessorRepo.Get(scenarioProcessor.ParentId)
	debugInterface, _ := s.DebugInterfaceRepo.Get(scenarioProcessor.EntityId)

	if debugInterface.DiagnoseInterfaceId > 0 {
		diagnoseInterface, _ := s.DiagnoseInterfaceRepo.Get(debugInterface.DiagnoseInterfaceId)
		if diagnoseInterface.Deleted {
			err = errors.New("interface is deleted")
			return
		}

		diagnoseInterfaceTo := s.DiagnoseInterfaceRepo.ToTo(&diagnoseInterface)
		newProcessor, err = s.ScenarioNodeService.createDirOrInterfaceFromDiagnose(diagnoseInterfaceTo, parentProcessor, scenarioProcessor.Ordr)

	} else if debugInterface.CaseInterfaceId > 0 {
		endpointCase, _ := s.EndpointCaseRepo.Get(debugInterface.CaseInterfaceId)
		if endpointCase.Deleted {
			err = errors.New("interface is deleted")
			return
		}

		interfaceCase := serverDomain.InterfaceCase{}
		copier.CopyWithOption(&interfaceCase, &endpointCase, copier.Option{IgnoreEmpty: true, DeepCopy: true})

		endpointCaseTo := s.EndpointCaseService.EndpointCaseToTo(&interfaceCase)
		newProcessor, err = s.ScenarioNodeService.createDirOrInterfaceFromCase(endpointCaseTo, parentProcessor, scenarioProcessor.Ordr)

	} else if debugInterface.EndpointInterfaceId > 0 {
		serveId := uint(0)
		newProcessor, err = s.ScenarioNodeService.createInterfaceFromDefine(debugInterface.EndpointInterfaceId, &serveId, createBy, parentProcessor, scenarioProcessor.Name, scenarioProcessor.Ordr)
	}

	if err != nil {
		return
	}

	s.DebugInvokeRepo.ChangeProcessorOwner(scenarioProcessor.ID, newProcessor.ID, newProcessor.EntityId, newProcessor.EndpointInterfaceId)

	// must put below, since creation will use its DebugInterface
	s.DebugInterfaceRepo.Delete(scenarioProcessor.EntityId)
	s.ScenarioProcessorRepo.Delete(scenarioProcessor.ID)

	return
}
