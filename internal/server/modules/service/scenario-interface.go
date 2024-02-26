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

func (s *ScenarioInterfaceService) GetDebugDataFromScenarioInterface(tenantId consts.TenantId, scenarioInterfaceId uint) (req domain.DebugData, err error) {
	scenarioInterfacePo, _ := s.ScenarioInterfaceRepo.GetDetail(tenantId, scenarioInterfaceId)
	if err != nil {
		return
	}

	endpointInterface, _ := s.EndpointInterfaceRepo.Get(tenantId, scenarioInterfacePo.EndpointInterfaceId)

	s.SetProps(tenantId, endpointInterface, &scenarioInterfacePo, &req)

	return
}

func (s *ScenarioInterfaceService) SetProps(tenantId consts.TenantId,
	endpointInterface model.EndpointInterface, scenarioInterfacePo *model.DebugInterface, debugData *domain.DebugData) {

	endpoint, err := s.EndpointRepo.GetAll(tenantId, endpointInterface.EndpointId, "v0.1.0")
	serve, err := s.ServeRepo.Get(tenantId, endpoint.ServeId)
	if err != nil {
		return
	}

	securities, err := s.ServeRepo.ListSecurity(tenantId, serve.ID)
	if err != nil {
		return
	}

	serve.Securities = securities
	debugData.EndpointInterfaceId = endpointInterface.ID

	copier.CopyWithOption(&debugData, scenarioInterfacePo, copier.Option{DeepCopy: true})
	debugData.EndpointInterfaceId = endpointInterface.ID // reset

	if debugData.Headers == nil {
		debugData.Headers = &[]domain.Header{}
	}
	*debugData.Headers = append(*debugData.Headers, domain.Header{Name: "", Value: ""})

	if debugData.QueryParams == nil {
		debugData.QueryParams = &[]domain.Param{}
	}
	*debugData.QueryParams = append(*debugData.QueryParams, domain.Param{Name: "", Value: "", ParamIn: consts.ParamInQuery})

	if debugData.PathParams == nil {
		debugData.PathParams = &[]domain.Param{}
	}
	*debugData.PathParams = append(*debugData.PathParams, domain.Param{Name: "", Value: "", ParamIn: consts.ParamInPath})

	if debugData.BodyFormData == nil {
		debugData.BodyFormData = &[]domain.BodyFormDataItem{}
	}
	*debugData.BodyFormData = append(*debugData.BodyFormData, domain.BodyFormDataItem{
		Name: "", Value: "", Type: consts.FormDataTypeText})

	if debugData.BodyFormUrlencoded == nil {
		debugData.BodyFormUrlencoded = &[]domain.BodyFormUrlEncodedItem{}
	}
	*debugData.BodyFormUrlencoded = append(*debugData.BodyFormUrlencoded, domain.BodyFormUrlEncodedItem{
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

func (s *ScenarioInterfaceService) SaveDebugData(tenantId consts.TenantId, req domain.DebugData) (debug model.DebugInterface, err error) {
	s.DebugInterfaceService.CopyValueFromRequest(tenantId, &debug, req)

	//endpointInterface, _ := s.EndpointInterfaceRepo.Get(req.EndpointInterfaceId)
	//debug.EndpointId = endpointInterface.EndpointId

	if req.DebugInterfaceId > 0 {
		debug.ID = req.DebugInterfaceId
	}

	err = s.ScenarioInterfaceRepo.SaveDebugData(tenantId, &debug)

	//更新执行器method
	s.ScenarioProcessorRepo.UpdateMethod(tenantId, debug.ScenarioProcessorId, debug.Method)

	return
}

func (s *ScenarioInterfaceService) ResetDebugData(tenantId consts.TenantId, scenarioProcessorId int, createBy uint) (
	newProcessor model.Processor, err error) {

	scenarioProcessor, _ := s.ScenarioProcessorRepo.Get(tenantId, uint(scenarioProcessorId))
	parentProcessor, _ := s.ScenarioProcessorRepo.Get(tenantId, scenarioProcessor.ParentId)
	debugInterface, _ := s.DebugInterfaceRepo.Get(tenantId, scenarioProcessor.EntityId)

	if debugInterface.DiagnoseInterfaceId > 0 {
		diagnoseInterface, _ := s.DiagnoseInterfaceRepo.Get(tenantId, debugInterface.DiagnoseInterfaceId)
		if diagnoseInterface.Deleted {
			err = errors.New("interface is deleted")
			return
		}

		diagnoseInterfaceTo := s.DiagnoseInterfaceRepo.ToTo(tenantId, &diagnoseInterface)
		newProcessor, err = s.ScenarioNodeService.createDirOrInterfaceFromDiagnose(tenantId, diagnoseInterfaceTo, parentProcessor, scenarioProcessor.Ordr)

	} else if debugInterface.CaseInterfaceId > 0 {
		endpointCase, _ := s.EndpointCaseRepo.Get(tenantId, debugInterface.CaseInterfaceId)
		if endpointCase.Deleted {
			err = errors.New("interface is deleted")
			return
		}

		interfaceCase := serverDomain.InterfaceCase{}
		copier.CopyWithOption(&interfaceCase, &endpointCase, copier.Option{IgnoreEmpty: true, DeepCopy: true})

		endpointCaseTo := s.EndpointCaseService.EndpointCaseToTo(tenantId, &interfaceCase)
		newProcessor, err = s.ScenarioNodeService.createDirOrInterfaceFromCase(tenantId, endpointCaseTo, parentProcessor, scenarioProcessor.Ordr)

	} else if debugInterface.EndpointInterfaceId > 0 {
		serveId := uint(0)
		newProcessor, err = s.ScenarioNodeService.createInterfaceFromDefine(tenantId, debugInterface.EndpointInterfaceId, &serveId, createBy, parentProcessor, scenarioProcessor.Name, scenarioProcessor.Ordr)
	}

	if err != nil {
		return
	}

	s.DebugInvokeRepo.ChangeProcessorOwner(tenantId, scenarioProcessor.ID, newProcessor.ID, newProcessor.EntityId, newProcessor.EndpointInterfaceId)

	// must put below, since creation will use its DebugInterface
	s.DebugInterfaceRepo.Delete(tenantId, scenarioProcessor.EntityId)
	s.ScenarioProcessorRepo.Delete(tenantId, scenarioProcessor.ID)

	return
}
