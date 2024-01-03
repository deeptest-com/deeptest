package service

import (
	agentExec "github.com/aaronchen2k/deeptest/internal/agent/exec"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/helper/openapi"
	schemaHelper "github.com/aaronchen2k/deeptest/internal/pkg/helper/schema"
	model "github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	"github.com/jinzhu/copier"
)

type DebugInterfaceService struct {
	EndpointInterfaceRepo *repo.EndpointInterfaceRepo `inject:""`
	DebugInterfaceRepo    *repo.DebugInterfaceRepo    `inject:""`
	ScenarioInterfaceRepo *repo.ScenarioInterfaceRepo `inject:""`
	DiagnoseInterfaceRepo *repo.DiagnoseInterfaceRepo `inject:""`
	EndpointCaseRepo      *repo.EndpointCaseRepo      `inject:""`

	EndpointRepo          *repo.EndpointRepo          `inject:""`
	ServeRepo             *repo.ServeRepo             `inject:""`
	ScenarioProcessorRepo *repo.ScenarioProcessorRepo `inject:""`
	ServeServerRepo       *repo.ServeServerRepo       `inject:""`
	ExtractorRepo         *repo.ExtractorRepo         `inject:""`
	CheckpointRepo        *repo.CheckpointRepo        `inject:""`
	EnvironmentRepo       *repo.EnvironmentRepo       `inject:""`

	DebugSceneService        *DebugSceneService        `inject:""`
	ScenarioInterfaceService *ScenarioInterfaceService `inject:""`
	SceneService             *SceneService             `inject:""`
	EnvironmentService       *EnvironmentService       `inject:""`
	DatapoolService          *DatapoolService          `inject:""`
	ServeService             *ServeService             `inject:""`

	PreConditionRepo  *repo.PreConditionRepo  `inject:""`
	PostConditionRepo *repo.PostConditionRepo `inject:""`
}

func (s *DebugInterfaceService) Load(loadReq domain.DebugInfo) (debugData domain.DebugData, err error) {
	if loadReq.DebugInterfaceId > 0 {
		debugData, _ = s.GetDebugDataFromDebugInterface(loadReq.DebugInterfaceId)
	} else {
		if loadReq.ScenarioProcessorId > 0 {
			debugData, _ = s.GetDebugInterfaceByScenarioInterface(loadReq.ScenarioProcessorId)
		} else if loadReq.DiagnoseInterfaceId > 0 {
			debugData, _ = s.GetDebugInterfaceByDiagnoseInterface(loadReq.DiagnoseInterfaceId)
		} else if loadReq.CaseInterfaceId > 0 {
			debugData, _ = s.GetDebugInterfaceByEndpointCase(loadReq.CaseInterfaceId)
		} else if loadReq.EndpointInterfaceId > 0 {
			debugData, _ = s.GetDebugInterfaceByEndpointInterface(loadReq.EndpointInterfaceId)
		}
	}

	debugData.UsedBy = loadReq.UsedBy
	debugData.EnvDataToView = &domain.EnvDataToView{}

	debugData.BaseUrl, debugData.EnvDataToView.ShareVars,
		debugData.EnvDataToView.EnvVars, debugData.EnvDataToView.GlobalVars,
		*debugData.GlobalParams =
		s.DebugSceneService.LoadScene(&debugData, loadReq.UserId, loadReq.EnvironmentId)

	debugData.ResponseDefine = s.PostConditionRepo.CreateDefaultResponseDefine(debugData.DebugInterfaceId, debugData.EndpointInterfaceId, loadReq.UsedBy)

	return
}

func (s *DebugInterfaceService) LoadForExec(loadReq domain.DebugInfo) (ret agentExec.InterfaceExecObj, err error) {
	ret.DebugData, _ = s.Load(loadReq)

	// load default environment for user
	env, _ := s.EnvironmentRepo.GetByUserAndProject(loadReq.UserId, uint(loadReq.ProjectId))
	if env.ID > 0 {
		server, _ := s.ServeServerRepo.FindByServeAndExecEnv(ret.DebugData.ServeId, env.ID)
		if server.ID > 0 {
			ret.DebugData.ServerId = server.ID
		}
	}

	ret.PreConditions, _ = s.PreConditionRepo.ListTo(
		ret.DebugData.DebugInterfaceId, ret.DebugData.EndpointInterfaceId, loadReq.UsedBy, "false")
	ret.PostConditions, _ = s.PostConditionRepo.ListTo(
		ret.DebugData.DebugInterfaceId, ret.DebugData.EndpointInterfaceId, loadReq.UsedBy, "false")

	ret.ExecScene.ShareVars = ret.DebugData.EnvDataToView.ShareVars // for execution
	ret.DebugData.EnvDataToView = nil

	// get environment and settings on project level
	s.SceneService.LoadEnvVars(&ret.ExecScene, ret.DebugData.ServerId, ret.DebugData.DebugInterfaceId)
	s.SceneService.LoadProjectSettings(&ret.ExecScene, ret.DebugData.ProjectId)

	//ret.ExecScene.GlobalParams = s.DebugSceneService.MergeGlobalParams(ret.ExecScene.GlobalParams, ret.DebugData.GlobalParams)

	return
}

func (s *DebugInterfaceService) GetDetail(id uint) (ret model.DebugInterface, err error) {
	ret, err = s.DebugInterfaceRepo.GetDetail(id)

	return
}

func (s *DebugInterfaceService) CreateOrUpdate(req domain.DebugData) (debugInterface model.DebugInterface, err error) {
	if req.DebugInterfaceId > 0 {
		debugInterface, err = s.Update(req, req.DebugInterfaceId)
	} else {
		debugInterface, err = s.Create(req)
	}

	return
}

func (s *DebugInterfaceService) Create(req domain.DebugData) (debugInterface model.DebugInterface, err error) {
	s.CopyValueFromRequest(&debugInterface, req)
	debugInterface.ServeId = req.ServeId
	debugInterface.ID = 0

	err = s.DebugInterfaceRepo.Save(&debugInterface)

	// first time to save a debug interface that convert from endpoint interface, clone conditions
	// it's different from cloning data between two debug interfaces when do importing
	s.PreConditionRepo.CloneAll(0, req.EndpointInterfaceId, debugInterface.ID, req.UsedBy, "", "false")
	s.PostConditionRepo.CloneAll(0, req.EndpointInterfaceId, debugInterface.ID, req.UsedBy, "", "false")

	s.EndpointInterfaceRepo.SetDebugInterfaceId(req.EndpointInterfaceId, debugInterface.ID)

	return
}

func (s *DebugInterfaceService) Update(req domain.DebugData, debugInterfaceId uint) (debugInterface model.DebugInterface, err error) {
	s.CopyValueFromRequest(&debugInterface, req)
	debugInterface.ID = debugInterfaceId

	err = s.DebugInterfaceRepo.Save(&debugInterface)

	// 更新method
	s.DiagnoseInterfaceRepo.UpdateMethod(debugInterface.DiagnoseInterfaceId, debugInterface.Method)

	return
}

func (s *DebugInterfaceService) SaveAs(debugData domain.DebugData, srcDebugInterfaceId uint, srcUsedBy consts.UsedBy) (debugInterface model.DebugInterface, err error) {
	s.CopyValueFromRequest(&debugInterface, debugData)
	debugInterface.ServeId = debugData.ServeId
	debugData.DebugInterfaceId = 0
	debugInterface.ID = 0

	err = s.DebugInterfaceRepo.Save(&debugInterface)

	s.PreConditionRepo.CloneAll(srcDebugInterfaceId, debugData.EndpointInterfaceId, debugInterface.ID, debugData.UsedBy, srcUsedBy, "false")
	s.PostConditionRepo.CloneAll(srcDebugInterfaceId, debugData.EndpointInterfaceId, debugInterface.ID, debugData.UsedBy, srcUsedBy, "false")

	return
}

func (s *DebugInterfaceService) GetDebugDataFromDebugInterface(debugInterfaceId uint) (debugData domain.DebugData, err error) {
	debugInterfacePo, err := s.DebugInterfaceRepo.GetDetail(debugInterfaceId)
	if err != nil {
		return
	}

	endpointInterface, _ := s.EndpointInterfaceRepo.Get(debugInterfacePo.EndpointInterfaceId)

	s.CopyDebugDataPropsFromPo(&debugData, &debugInterfacePo, endpointInterface)

	return
}
func (s *DebugInterfaceService) GetDebugDataFromEndpointInterface(endpointInterfaceId uint) (ret domain.DebugData, err error) {
	endpointInterface, _ := s.EndpointInterfaceRepo.Get(endpointInterfaceId)

	if endpointInterface.DebugInterfaceId > 0 {
		ret, err = s.GetDebugDataFromDebugInterface(endpointInterface.DebugInterfaceId)
	} else {
		ret, err = s.ConvertDebugDataFromEndpointInterface(endpointInterfaceId)
	}

	return
}

func (s *DebugInterfaceService) ConvertDebugDataFromEndpointInterface(endpointInterfaceId uint) (debugData domain.DebugData, err error) {
	endpointInterface, err := s.EndpointInterfaceRepo.GetDetail(endpointInterfaceId)
	if err != nil {
		return
	}

	s.CopyDebugDataPropsFromPo(&debugData, nil, endpointInterface)

	// init ServeId and ServerId id by endpoint
	endpoint, _ := s.EndpointRepo.Get(endpointInterface.EndpointId)
	debugData.ServeId = endpoint.ServeId
	debugData.ServerId = endpoint.ServerId
	debugData.ProjectId = endpoint.ProjectId

	debugData.UsedBy = consts.InterfaceDebug

	return
}

func (s *DebugInterfaceService) CopyDebugDataPropsFromPo(debugData *domain.DebugData,
	debugInterfacePo *model.DebugInterface, endpointInterface model.EndpointInterface) {

	endpoint, err := s.EndpointRepo.GetAll(endpointInterface.EndpointId, "v0.1.0")
	serve, err := s.ServeRepo.Get(endpoint.ServeId)
	if err != nil {
		//	return
	}

	securities, err := s.ServeRepo.ListSecurity(serve.ID)
	if err != nil {
		//	return
	}

	serve.Securities = securities
	debugData.EndpointInterfaceId = endpointInterface.ID

	if debugInterfacePo == nil { // is null when converting from EndpointInterface
		schema2conv := schemaHelper.NewSchema2conv()
		schema2conv.Components = s.ServeService.Components(serve.ID)
		interfaces2debug := openapi.NewInterfaces2debug(endpointInterface, endpoint, serve, schema2conv)
		debugInterfacePo = interfaces2debug.Convert()

		//debugInterfacePo.Name = fmt.Sprintf("%s - %s", endpoint.Title, debugInterfacePo.Method)
		debugInterfacePo.Name = endpoint.Title
	}

	copier.CopyWithOption(&debugData, debugInterfacePo, copier.Option{DeepCopy: true})
	debugData.EndpointInterfaceId = endpointInterface.ID // reset

	debugData.DebugInterfaceId = debugInterfacePo.ID
	debugData.ServeId = debugInterfacePo.ServeId

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

	if debugData.Cookies == nil {
		debugData.Cookies = &[]domain.ExecCookie{}
	}
	*debugData.Cookies = append(*debugData.Cookies, domain.ExecCookie{Name: "", Value: ""})

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

func (s *DebugInterfaceService) GetEndpointAndServeIdForEndpointInterface(endpointInterfaceId uint) (
	endpointId, serveId uint) {

	endpointInterface, _ := s.EndpointInterfaceRepo.Get(endpointInterfaceId)

	endpointId = endpointInterface.EndpointId
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

func (s *DebugInterfaceService) GenSample(projectId, serveId uint) (ret *model.DebugInterface, err error) {
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
		endpointInterface, _ := s.EndpointInterfaceRepo.Get(processor.EndpointInterfaceId)
		server, _ := s.ServeServerRepo.GetByEndpoint(endpointInterface.EndpointId)
		ret.ServerId = server.ID
	}

	ret.DebugInterfaceId = debugInterfaceId
	ret.ScenarioProcessorId = scenarioProcessorId
	ret.ProcessorInterfaceSrc = debugData.ProcessorInterfaceSrc

	if ret.Headers == nil {
		ret.Headers = &[]domain.Header{}
	}
	*ret.Headers = append(*ret.Headers, domain.Header{Name: "", Value: ""})

	if ret.QueryParams == nil {
		ret.QueryParams = &[]domain.Param{}
	}
	*ret.QueryParams = append(*ret.QueryParams, domain.Param{Name: "", Value: "", ParamIn: consts.ParamInQuery})

	if ret.PathParams == nil {
		ret.PathParams = &[]domain.Param{}
	}
	*ret.PathParams = append(*ret.PathParams, domain.Param{Name: "", Value: "", ParamIn: consts.ParamInPath})

	if ret.Cookies == nil {
		ret.Cookies = &[]domain.ExecCookie{}
	}
	*ret.Cookies = append(*ret.Cookies, domain.ExecCookie{Name: "", Value: ""})

	if ret.BodyFormData == nil {
		ret.BodyFormData = &[]domain.BodyFormDataItem{}
	}
	*ret.BodyFormData = append(*ret.BodyFormData, domain.BodyFormDataItem{
		Name: "", Value: "", Type: consts.FormDataTypeText})

	if ret.BodyFormUrlencoded == nil {
		ret.BodyFormUrlencoded = &[]domain.BodyFormUrlEncodedItem{}
	}
	*ret.BodyFormUrlencoded = append(*ret.BodyFormUrlencoded, domain.BodyFormUrlEncodedItem{
		Name: "", Value: "",
	})

	return
}

func (s *DebugInterfaceService) GetDebugInterfaceByDiagnoseInterface(diagnoseInterfaceId uint) (ret domain.DebugData, err error) {
	diagnoseInterface, err := s.DiagnoseInterfaceRepo.GetDetail(diagnoseInterfaceId)
	if err != nil {
		return
	}

	copier.CopyWithOption(&ret, diagnoseInterface.DebugData, copier.Option{
		DeepCopy: true,
	})

	ret.ServerId = diagnoseInterface.DebugData.ServerId
	ret.DebugInterfaceId = diagnoseInterface.DebugInterfaceId
	ret.DiagnoseInterfaceId = diagnoseInterfaceId

	if ret.Headers == nil {
		ret.Headers = &[]domain.Header{}
	}
	*ret.Headers = append(*ret.Headers, domain.Header{Name: "", Value: ""})

	if ret.QueryParams == nil {
		ret.QueryParams = &[]domain.Param{}
	}
	*ret.QueryParams = append(*ret.QueryParams, domain.Param{Name: "", Value: "", ParamIn: consts.ParamInQuery})

	if ret.PathParams == nil {
		ret.PathParams = &[]domain.Param{}
	}
	*ret.PathParams = append(*ret.PathParams, domain.Param{Name: "", Value: "", ParamIn: consts.ParamInPath})

	if ret.Cookies == nil {
		ret.Cookies = &[]domain.ExecCookie{}
	}
	*ret.Cookies = append(*ret.Cookies, domain.ExecCookie{Name: "", Value: ""})

	if ret.BodyFormData == nil {
		ret.BodyFormData = &[]domain.BodyFormDataItem{}
	}
	*ret.BodyFormData = append(*ret.BodyFormData, domain.BodyFormDataItem{
		Name: "", Value: "", Type: consts.FormDataTypeText})

	if ret.BodyFormUrlencoded == nil {
		ret.BodyFormUrlencoded = &[]domain.BodyFormUrlEncodedItem{}
	}
	*ret.BodyFormUrlencoded = append(*ret.BodyFormUrlencoded, domain.BodyFormUrlEncodedItem{
		Name: "", Value: "",
	})

	return
}

func (s *DebugInterfaceService) GetDebugInterfaceByEndpointCase(endpointCaseId uint) (ret domain.DebugData, err error) {
	endpointCase, err := s.EndpointCaseRepo.GetDetail(endpointCaseId)
	if err != nil {
		return
	}

	copier.CopyWithOption(&ret, endpointCase.DebugData, copier.Option{
		DeepCopy: true,
	})

	ret.ServerId = endpointCase.DebugData.ServerId
	ret.DebugInterfaceId = endpointCase.DebugInterfaceId
	ret.CaseInterfaceId = endpointCaseId

	if ret.Headers == nil {
		ret.Headers = &[]domain.Header{}
	}
	*ret.Headers = append(*ret.Headers, domain.Header{Name: "", Value: ""})

	if ret.QueryParams == nil {
		ret.QueryParams = &[]domain.Param{}
	}
	*ret.QueryParams = append(*ret.QueryParams, domain.Param{Name: "", Value: "", ParamIn: consts.ParamInQuery})

	if ret.PathParams == nil {
		ret.PathParams = &[]domain.Param{}
	}
	*ret.PathParams = append(*ret.PathParams, domain.Param{Name: "", Value: "", ParamIn: consts.ParamInPath})

	if ret.Cookies == nil {
		ret.Cookies = &[]domain.ExecCookie{}
	}
	*ret.Cookies = append(*ret.Cookies, domain.ExecCookie{Name: "", Value: ""})

	if ret.BodyFormData == nil {
		ret.BodyFormData = &[]domain.BodyFormDataItem{}
	}
	*ret.BodyFormData = append(*ret.BodyFormData, domain.BodyFormDataItem{
		Name: "", Value: "", Type: consts.FormDataTypeText})

	if ret.BodyFormUrlencoded == nil {
		ret.BodyFormUrlencoded = &[]domain.BodyFormUrlEncodedItem{}
	}
	*ret.BodyFormUrlencoded = append(*ret.BodyFormUrlencoded, domain.BodyFormUrlEncodedItem{
		Name: "", Value: "",
	})

	return
}

func (s *DebugInterfaceService) CreateDefault(src consts.ProcessorInterfaceSrc, projectId uint) (id uint, err error) {
	id, err = s.DebugInterfaceRepo.CreateDefault(src, projectId)

	return
}

func (s *DebugInterfaceService) CopyValueFromRequest(interf *model.DebugInterface, req domain.DebugData) (err error) {
	copier.CopyWithOption(interf, &req, copier.Option{IgnoreEmpty: true, DeepCopy: true})

	return
}

func (s *DebugInterfaceService) MergeGlobalParams(globalParams []domain.GlobalParam, selfGlobalParam []model.DebugInterfaceGlobalParam) (ret []domain.GlobalParam) {

	ret = globalParams
	for key, globalParam := range ret {
		for _, param := range selfGlobalParam {
			if param.Name == globalParam.Name && param.In == globalParam.In {
				ret[key].Disabled = param.Disabled
			}
		}
	}

	return
}
