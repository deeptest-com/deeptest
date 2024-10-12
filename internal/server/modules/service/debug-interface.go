package service

import (
	"fmt"
	serverDomain "github.com/deeptest-com/deeptest/cmd/server/v1/domain"
	agentExec "github.com/deeptest-com/deeptest/internal/agent/exec"
	agentService "github.com/deeptest-com/deeptest/internal/agent/service"
	"github.com/deeptest-com/deeptest/internal/pkg/consts"
	"github.com/deeptest-com/deeptest/internal/pkg/domain"
	"github.com/deeptest-com/deeptest/internal/pkg/helper/openapi"
	schemaHelper "github.com/deeptest-com/deeptest/internal/pkg/helper/schema"
	model "github.com/deeptest-com/deeptest/internal/server/modules/model"
	"github.com/deeptest-com/deeptest/internal/server/modules/repo"
	_stringUtils "github.com/deeptest-com/deeptest/pkg/lib/string"
	"github.com/jinzhu/copier"
	"strings"
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

	ConditionRepo *repo.ConditionRepo `inject:""`
	MetricsRepo   *repo.MetricsRepo   `inject:""`
	LlmToolRepo   *repo.LlmToolRepo   `inject:""`
}

func (s *DebugInterfaceService) Load(tenantId consts.TenantId, loadReq domain.DebugInfo) (debugData domain.DebugData, err error) {
	if loadReq.DebugInterfaceId > 0 {
		debugData, _ = s.GetDebugDataFromDebugInterface(tenantId, loadReq.DebugInterfaceId)
	} else {
		if loadReq.ScenarioProcessorId > 0 {
			debugData, _ = s.GetDebugInterfaceByScenarioInterface(tenantId, loadReq.ScenarioProcessorId)
		} else if loadReq.DiagnoseInterfaceId > 0 {
			debugData, _ = s.GetDebugInterfaceByDiagnoseInterface(tenantId, loadReq.DiagnoseInterfaceId)
		} else if loadReq.CaseInterfaceId > 0 {
			debugData, _ = s.GetDebugInterfaceByEndpointCase(tenantId, loadReq.CaseInterfaceId)
		} else if loadReq.EndpointInterfaceId > 0 {
			debugData, _ = s.GetDebugInterfaceByEndpointInterface(tenantId, loadReq.EndpointInterfaceId, loadReq.FromDefine)
		}
	}

	if debugData.QueryParams == nil {
		debugData.QueryParams = &[]domain.Param{}
	}
	if debugData.PathParams == nil {
		debugData.PathParams = &[]domain.Param{}
	}
	if debugData.Headers == nil {
		debugData.Headers = &[]domain.Header{}
	}
	if debugData.Cookies == nil {
		debugData.Cookies = &[]domain.ExecCookie{}
	}
	if debugData.GlobalParams == nil {
		debugData.GlobalParams = &[]domain.GlobalParam{}
	}

	if debugData.BodyFormData == nil {
		debugData.BodyFormData = &[]domain.BodyFormDataItem{}
	}
	if debugData.BodyFormUrlencoded == nil {
		debugData.BodyFormUrlencoded = &[]domain.BodyFormUrlEncodedItem{}
	}

	debugData.UsedBy = loadReq.UsedBy
	debugData.EnvDataToView = &domain.EnvDataToView{}

	debugData.BaseUrl, debugData.EnvDataToView.ShareVars,
		debugData.EnvDataToView.EnvVars, debugData.EnvDataToView.GlobalVars,
		*debugData.GlobalParams =
		s.DebugSceneService.LoadScene(tenantId, &debugData, loadReq.UserId, loadReq.EnvironmentId)

	debugData.ResponseDefine = s.ConditionRepo.CreateDefaultResponseDefine(tenantId, debugData.DebugInterfaceId, debugData.EndpointInterfaceId, loadReq.UsedBy)

	return
}

func (s *DebugInterfaceService) LoadForExec(tenantId consts.TenantId, loadReq domain.DebugInfo) (ret agentExec.InterfaceExecObj, err error) {
	ret, err = s.loadDetail(tenantId, loadReq, true)

	return
}

func (s *DebugInterfaceService) loadDetail(tenantId consts.TenantId, loadReq domain.DebugInfo, withConditions bool) (ret agentExec.InterfaceExecObj, err error) {
	ret.DebugData, _ = s.Load(tenantId, loadReq)

	// load default environment for user
	env, _ := s.EnvironmentRepo.GetByUserAndProject(tenantId, loadReq.UserId, uint(loadReq.ProjectId))
	if env.ID > 0 {
		server, _ := s.ServeServerRepo.FindByServeAndExecEnv(tenantId, ret.DebugData.ServeId, env.ID)
		if server.ID > 0 {
			ret.DebugData.ServerId = server.ID
		}
	}

	if withConditions {
		ret.PreConditions, _ = s.ConditionRepo.ListTo(tenantId,
			ret.DebugData.DebugInterfaceId, ret.DebugData.EndpointInterfaceId, loadReq.UsedBy, "false", consts.ConditionSrcPre)
		ret.PostConditions, _ = s.ConditionRepo.ListTo(tenantId,
			ret.DebugData.DebugInterfaceId, ret.DebugData.EndpointInterfaceId, loadReq.UsedBy, "false", consts.ConditionSrcPost)

		ret.Metrics, _ = s.MetricsRepo.ListTo(tenantId,
			ret.DebugData.DebugInterfaceId, ret.DebugData.EndpointInterfaceId)
		ret.ToolLlm, _ = s.LlmToolRepo.GetDefault(tenantId, uint(loadReq.ProjectId))
	}

	ret.ExecScene.ShareVars = ret.DebugData.EnvDataToView.ShareVars // for execution
	ret.DebugData.EnvDataToView = nil

	// get environment and settings on project level
	s.SceneService.LoadEnvVars(tenantId, &ret.ExecScene, ret.DebugData.ServerId, ret.DebugData.DebugInterfaceId)
	s.SceneService.LoadProjectSettings(tenantId, &ret.ExecScene, ret.DebugData.ProjectId)

	//ret.ExecScene.GlobalParams = s.DebugSceneService.MergeGlobalParams(ret.ExecScene.GlobalParams, ret.DebugData.GlobalParams)

	return
}

func (s *DebugInterfaceService) GetDetail(tenantId consts.TenantId, id uint) (ret model.DebugInterface, err error) {
	ret, err = s.DebugInterfaceRepo.GetDetail(tenantId, id)

	return
}

func (s *DebugInterfaceService) CreateOrUpdate(tenantId consts.TenantId, req domain.DebugData) (debugInterface model.DebugInterface, err error) {
	if req.DebugInterfaceId > 0 {
		debugInterface, err = s.Update(tenantId, req, req.DebugInterfaceId)
	} else {
		debugInterface, err = s.Create(tenantId, req)
	}

	return
}

func (s *DebugInterfaceService) Create(tenantId consts.TenantId, req domain.DebugData) (debugInterface model.DebugInterface, err error) {
	s.CopyValueFromRequest(tenantId, &debugInterface, req)
	debugInterface.ServeId = req.ServeId
	debugInterface.ID = 0

	err = s.DebugInterfaceRepo.Save(tenantId, &debugInterface)

	// first time to save a debug interface that convert from endpoint interface, clone conditions
	// it's different from cloning data between two debug interfaces when do importing

	s.ConditionRepo.CloneAll(tenantId, 0, req.EndpointInterfaceId, debugInterface.ID, req.UsedBy, "", "false")

	s.EndpointInterfaceRepo.SetDebugInterfaceId(tenantId, req.EndpointInterfaceId, debugInterface.ID)

	return
}

func (s *DebugInterfaceService) Update(tenantId consts.TenantId, req domain.DebugData, debugInterfaceId uint) (debugInterface model.DebugInterface, err error) {
	s.CopyValueFromRequest(tenantId, &debugInterface, req)
	debugInterface.ID = debugInterfaceId

	err = s.DebugInterfaceRepo.Save(tenantId, &debugInterface)

	// 更新method
	s.DiagnoseInterfaceRepo.UpdateMethod(tenantId, debugInterface.DiagnoseInterfaceId, debugInterface.Method)

	return
}

func (s *DebugInterfaceService) SaveAs(tenantId consts.TenantId, debugData domain.DebugData, srcDebugInterfaceId uint, srcUsedBy consts.UsedBy) (debugInterface model.DebugInterface, err error) {
	s.CopyValueFromRequest(tenantId, &debugInterface, debugData)
	debugInterface.ServeId = debugData.ServeId
	debugData.DebugInterfaceId = 0
	debugInterface.ID = 0

	err = s.DebugInterfaceRepo.Save(tenantId, &debugInterface)

	s.ConditionRepo.CloneAll(tenantId, srcDebugInterfaceId, debugData.EndpointInterfaceId, debugInterface.ID, debugData.UsedBy, srcUsedBy, "false")

	return
}

func (s *DebugInterfaceService) GetDebugDataFromDebugInterface(tenantId consts.TenantId, debugInterfaceId uint) (debugData domain.DebugData, err error) {
	debugInterfacePo, err := s.DebugInterfaceRepo.GetDetail(tenantId, debugInterfaceId)
	if err != nil {
		return
	}

	endpointInterface, _ := s.EndpointInterfaceRepo.Get(tenantId, debugInterfacePo.EndpointInterfaceId)

	s.CopyDebugDataPropsFromPo(tenantId, &debugData, &debugInterfacePo, endpointInterface)

	return
}
func (s *DebugInterfaceService) GetDebugDataFromEndpointInterface(tenantId consts.TenantId, endpointInterfaceId uint) (ret domain.DebugData, err error) {
	endpointInterface, _ := s.EndpointInterfaceRepo.Get(tenantId, endpointInterfaceId)

	if endpointInterface.DebugInterfaceId > 0 {
		ret, err = s.GetDebugDataFromDebugInterface(tenantId, endpointInterface.DebugInterfaceId)
	} else {
		ret, err = s.ConvertDebugDataFromEndpointInterface(tenantId, endpointInterfaceId)
	}

	return
}

func (s *DebugInterfaceService) ConvertDebugDataFromEndpointInterface(tenantId consts.TenantId, endpointInterfaceId uint) (debugData domain.DebugData, err error) {
	endpointInterface, err := s.EndpointInterfaceRepo.GetDetail(tenantId, endpointInterfaceId)
	if err != nil {
		return
	}

	s.CopyDebugDataPropsFromPo(tenantId, &debugData, nil, endpointInterface)

	// init ServeId and ServerId id by endpoint
	endpoint, _ := s.EndpointRepo.Get(tenantId, endpointInterface.EndpointId)
	debugData.ServeId = endpoint.ServeId
	debugData.ServerId = endpoint.ServerId
	debugData.ProjectId = endpoint.ProjectId
	server, err := s.ServeRepo.GetDefaultServer(tenantId, endpoint.ServeId)
	if err != nil {
		debugData.BaseUrl = server.Url
	}

	debugData.UsedBy = consts.InterfaceDebug

	return
}

func (s *DebugInterfaceService) CopyDebugDataPropsFromPo(tenantId consts.TenantId, debugData *domain.DebugData,
	debugInterfacePo *model.DebugInterface, endpointInterface model.EndpointInterface) {

	endpoint, err := s.EndpointRepo.GetAll(tenantId, endpointInterface.EndpointId, "v0.1.0")
	serve, err := s.ServeRepo.Get(tenantId, endpoint.ServeId)
	if err != nil {
		//	return
	}

	securities, err := s.ServeRepo.ListSecurity(tenantId, serve.ID)
	if err != nil {
		//	return
	}

	serve.Securities = securities
	debugData.EndpointInterfaceId = endpointInterface.ID

	if debugInterfacePo == nil { // is null when converting from EndpointInterface
		schema2conv := schemaHelper.NewSchema2conv()

		schema2conv.Components = s.ServeService.Components(tenantId, endpoint.ProjectId)

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

func (s *DebugInterfaceService) GetEndpointAndServeIdForEndpointInterface(tenantId consts.TenantId, endpointInterfaceId uint) (
	endpointId, serveId uint) {

	endpointInterface, _ := s.EndpointInterfaceRepo.Get(tenantId, endpointInterfaceId)

	endpointId = endpointInterface.EndpointId
	endpoint, _ := s.EndpointRepo.Get(tenantId, endpointId)

	serveId = endpoint.ServeId

	return
}

func (s *DebugInterfaceService) GetScenarioIdForDebugInterface(tenantId consts.TenantId, processorId uint) (
	scenarioId uint) {

	processor, _ := s.ScenarioProcessorRepo.Get(tenantId, processorId)
	scenarioId = processor.ScenarioId

	return
}

func (s *DebugInterfaceService) GenSample(projectId, serveId uint) (ret *model.DebugInterface, err error) {
	return
}

func (s *DebugInterfaceService) GetDebugInterfaceByEndpointInterface(tenantId consts.TenantId, endpointInterfaceId uint, fromDefine bool) (ret domain.DebugData, err error) {
	endpointInterface, _ := s.EndpointInterfaceRepo.Get(tenantId, endpointInterfaceId)

	if endpointInterface.DebugInterfaceId > 0 && !fromDefine {
		ret, err = s.GetDebugDataFromDebugInterface(tenantId, endpointInterface.DebugInterfaceId)
	} else {
		ret, err = s.ConvertDebugDataFromEndpointInterface(tenantId, endpointInterfaceId)
	}

	return
}

func (s *DebugInterfaceService) GetDebugInterfaceByScenarioInterface(tenantId consts.TenantId, scenarioProcessorId uint) (ret domain.DebugData, err error) {
	processor, err := s.ScenarioProcessorRepo.Get(tenantId, scenarioProcessorId)
	if err != nil {
		return
	}

	debugInterfaceId := processor.EntityId
	debugData, _ := s.DebugInterfaceRepo.GetDetail(tenantId, debugInterfaceId)

	copier.CopyWithOption(&ret, debugData, copier.Option{
		DeepCopy: true,
	})

	if ret.ServerId <= 0 {
		endpointInterface, _ := s.EndpointInterfaceRepo.Get(tenantId, processor.EndpointInterfaceId)
		server, _ := s.ServeServerRepo.GetByEndpoint(tenantId, endpointInterface.EndpointId)
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

func (s *DebugInterfaceService) GetDebugInterfaceByDiagnoseInterface(tenantId consts.TenantId, diagnoseInterfaceId uint) (ret domain.DebugData, err error) {
	diagnoseInterface, err := s.DiagnoseInterfaceRepo.GetDetail(tenantId, diagnoseInterfaceId)
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

func (s *DebugInterfaceService) GetDebugInterfaceByEndpointCase(tenantId consts.TenantId, endpointCaseId uint) (ret domain.DebugData, err error) {
	endpointCase, err := s.EndpointCaseRepo.GetDetail(tenantId, endpointCaseId)
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

func (s *DebugInterfaceService) CreateDefault(tenantId consts.TenantId, src consts.ProcessorInterfaceSrc, projectId uint) (id uint, err error) {
	id, err = s.DebugInterfaceRepo.CreateDefault(tenantId, src, projectId)

	return
}

func (s *DebugInterfaceService) CopyValueFromRequest(tenantId consts.TenantId, interf *model.DebugInterface, req domain.DebugData) (err error) {
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

func (s *DebugInterfaceService) LoadCurl(tenantId consts.TenantId, req serverDomain.DiagnoseCurlLoadReq) (ret string, err error) {
	if req.EndpointId > 0 {
		_, req.EndpointInterfaceId = s.EndpointInterfaceRepo.GetByMethod(tenantId, req.EndpointId, req.InterfaceMethod)
	}

	loadReq := domain.DebugInfo{
		DebugInterfaceId:    req.DebugInterfaceId,
		EndpointInterfaceId: req.EndpointInterfaceId,
		CaseInterfaceId:     req.CaseId,
		DiagnoseInterfaceId: req.DiagnoseId,

		EnvironmentId: req.EnvironmentId,
		ProjectId:     req.ProjectId,
		UserId:        req.UserId,
		UsedBy:        req.UsedBy,
		FromDefine:    req.FromDefine,
	}

	if req.EndpointId != 0 {
		loadReq.FromDefine = true
	}

	execObj, err := s.loadDetail(tenantId, loadReq, false)

	// replace variables
	uuid := fmt.Sprintf("load_curl_on_server_side_user%d_%s", req.UserId, _stringUtils.Uuid())
	session := agentExec.ExecSession{
		ExecUuid:       uuid,
		ExecScene:      execObj.ExecScene,
		InterfaceDebug: new(agentExec.InterfaceDebugSession),
		ScenarioDebug:  new(agentExec.ScenarioDebugSession),
	}
	agentExec.ReplaceVariables(&execObj.DebugData.BaseRequest, &session)

	// gen url
	execObj.DebugData.BaseRequest.Url, _ = agentService.UpdateUrl(execObj.DebugData)

	// gen bytes for form file item
	//if execObj.DebugData.BodyFormData != nil {
	//	for index, item := range *execObj.DebugData.BodyFormData {
	//		if item.Type == consts.FormDataTypeFile {
	//			(*execObj.DebugData.BodyFormData)[index].Value = filepath.Join(consts.WorkDir, item.Value)
	//		}
	//	}
	//}

	// generate curl command
	ret = s.genCurlCommand(execObj, uuid)

	return
}

func (s *DebugInterfaceService) genCurlCommand(execObj agentExec.InterfaceExecObj, execUuid string) (ret string) {
	debugData := execObj.DebugData

	command := "curl -i "

	// basic auth
	if debugData.BasicAuth.Username != "" {
		command += fmt.Sprintf("-u '%s:%s'",
			debugData.BasicAuth.Username, debugData.BasicAuth.Password)
	}

	// method
	command += fmt.Sprintf("-X %s ", debugData.Method)

	// url param
	arr := []string{}
	for _, param := range *debugData.QueryParams {
		if param.Name == "" {
			continue
		}
		str := fmt.Sprintf("%s=%s", param.Name, param.Value)
		//str = url.QueryEscape(str)
		arr = append(arr, str)
	}
	if len(arr) > 0 {
		command += fmt.Sprintf("'%s?%s' ", debugData.Url, strings.Join(arr, "&"))
	} else {
		command += fmt.Sprintf("'%s' ", debugData.Url)
	}

	// header
	for _, header := range *debugData.Headers {
		if header.Name == "" {
			continue
		}
		command += fmt.Sprintf("-H '%s: %s' ", header.Name, header.Value)
	}

	// cookie
	for _, cookie := range *debugData.Cookies {
		if cookie.Name == "" {
			continue
		}
		command += fmt.Sprintf("-b '%s=%s' ", cookie.Name, cookie.Value)
	}

	// body
	if debugData.BodyType == consts.ContentTypeFormData {
		arr := []string{}
		for _, item := range *debugData.BodyFormData {
			if item.Name == "" || item.Value == "" {
				continue
			}
			str := fmt.Sprintf("-F %s=%s", item.Name, item.Value)
			arr = append(arr, str)
		}

		command += fmt.Sprintf("%s", strings.Join(arr, " "))

	} else if debugData.BodyType == consts.ContentTypeFormUrlencoded {
		arr := []string{}
		for _, item := range *debugData.BodyFormUrlencoded {
			if item.Name == "" || item.Value == "" {
				continue
			}
			str := fmt.Sprintf("--data-urlencode %s=%s", item.Name, item.Value)
			arr = append(arr, str)
		}

		command += fmt.Sprintf("%s", strings.Join(arr, " "))

	} else if debugData.Method != consts.GET &&
		debugData.Method != consts.DELETE &&
		debugData.Method != consts.TRACE &&
		debugData.Method != consts.OPTIONS &&
		debugData.Method != consts.HEAD {

		body := strings.ReplaceAll(debugData.Body, "\n", "")
		command += fmt.Sprintf("-H 'Content-Type: %s' -d '%s' ", debugData.BodyType, body)

	}

	ret = command
	return
}
