package service

import (
	"encoding/json"
	"github.com/deeptest-com/deeptest/internal/pkg/consts"
	"github.com/deeptest-com/deeptest/internal/pkg/domain"
	model "github.com/deeptest-com/deeptest/internal/server/modules/model"
	"github.com/deeptest-com/deeptest/internal/server/modules/repo"
	"github.com/kataras/iris/v12"
	"time"
)

type DebugInvokeService struct {
	DebugRepo          *repo.DebugRepo          `inject:""`
	DebugInterfaceRepo *repo.DebugInterfaceRepo `inject:""`
	DebugInvokeRepo    *repo.DebugInvokeRepo    `inject:""`

	EndpointRepo          *repo.EndpointRepo          `inject:""`
	ScenarioRepo          *repo.ScenarioRepo          `inject:""`
	DiagnoseInterfaceRepo *repo.DiagnoseInterfaceRepo `inject:""`
	EndpointCaseRepo      *repo.EndpointCaseRepo      `inject:""`

	DebugInterfaceService *DebugInterfaceService `inject:""`
	ExecConditionService  *ExecConditionService  `inject:""`

	ConditionRepo      *repo.ConditionRepo      `inject:""`
	ExtractorRepo      *repo.ExtractorRepo      `inject:""`
	CheckpointRepo     *repo.CheckpointRepo     `inject:""`
	ScriptRepo         *repo.ScriptRepo         `inject:""`
	DatabaseOptRepo    *repo.DatabaseOptRepo    `inject:""`
	ResponseDefineRepo *repo.ResponseDefineRepo `inject:""`

	ScenarioInterfaceRepo *repo.ScenarioInterfaceRepo `inject:""`
}

func (s *DebugInvokeService) SubmitResult(tenantId consts.TenantId, req domain.SubmitDebugResultRequest) (invoke model.DebugInvoke, err error) {
	usedBy := req.Request.UsedBy
	var endpointId, serveId, processorId, scenarioId, projectId uint

	if usedBy == consts.InterfaceDebug {
		endpointId, serveId = s.DebugInterfaceService.GetEndpointAndServeIdForEndpointInterface(tenantId, req.Request.EndpointInterfaceId)

		endpoint, _ := s.EndpointRepo.Get(tenantId, endpointId)
		serveId = endpoint.ServeId
		projectId = endpoint.ProjectId

	} else if usedBy == consts.ScenarioDebug {
		processorId = req.Request.ScenarioProcessorId
		scenarioId = s.DebugInterfaceService.GetScenarioIdForDebugInterface(tenantId, req.Request.ScenarioProcessorId)

		scenario, _ := s.ScenarioRepo.Get(tenantId, scenarioId)
		projectId = scenario.ProjectId

	} else if usedBy == consts.DiagnoseDebug {
		diagnoseInterface, _ := s.DiagnoseInterfaceRepo.Get(tenantId, req.Request.DiagnoseInterfaceId)

		serveId = diagnoseInterface.ServeId
		projectId = diagnoseInterface.ProjectId

	} else if usedBy == consts.CaseDebug {
		caseInterface, _ := s.EndpointCaseRepo.Get(tenantId, req.Request.CaseInterfaceId)

		serveId = caseInterface.ServeId
		projectId = caseInterface.ProjectId
	}

	invoke, err = s.Create(tenantId, req, serveId, processorId, scenarioId, projectId)

	s.ExecConditionService.SavePreConditionResult(tenantId, invoke.ID,
		req.Request.DebugInterfaceId, req.Request.CaseInterfaceId, req.Request.EndpointInterfaceId,
		serveId, processorId, scenarioId, usedBy,
		req.PreConditions)

	s.ExecConditionService.SavePostConditionResult(tenantId, invoke.ID,
		req.Request.DebugInterfaceId, req.Request.CaseInterfaceId, req.Request.EndpointInterfaceId,
		serveId, processorId, scenarioId, usedBy,
		req.PostConditions)

	if err != nil {
		return
	}

	return
}

func (s *DebugInvokeService) Create(tenantId consts.TenantId, req domain.SubmitDebugResultRequest,
	serveId, scenarioProcessorId, scenarioId, projectId uint) (po model.DebugInvoke, err error) {

	debugInterface, _ := s.DebugInterfaceRepo.Get(tenantId, req.Request.DebugInterfaceId)

	po = model.DebugInvoke{
		ServeId: serveId,

		ScenarioProcessorId: scenarioProcessorId,
		ScenarioId:          scenarioId,

		InvocationBase: model.InvocationBase{
			ResultStatus: req.ResultStatus,

			Name:                time.Now().Format("01-02 15:04:05"),
			EndpointInterfaceId: req.Request.EndpointInterfaceId,
			DebugInterfaceId:    debugInterface.ID, // may be 0
			ProjectId:           projectId,
		},
	}

	bytesDebugData, _ := json.Marshal(req.Request)
	po.ReqContent = string(bytesDebugData)

	bytesResp, _ := json.Marshal(req.Response)
	po.RespContent = string(bytesResp)

	bytesPreConditions, _ := json.Marshal(req.PreConditions)
	po.PreConditionsContent = string(bytesPreConditions)

	bytesPostConditions, _ := json.Marshal(req.PostConditions)
	po.PostConditionsContent = string(bytesPostConditions)

	err = s.DebugInvokeRepo.Save(tenantId, &po)

	return
}

func (s *DebugInvokeService) ListByInterface(tenantId consts.TenantId, debugInterfaceId, endpointInterfaceId uint) (invocations []model.DebugInvoke, err error) {
	invocations, err = s.DebugRepo.List(tenantId, debugInterfaceId, endpointInterfaceId)

	return
}

func (s *DebugInvokeService) GetLastResp(tenantId consts.TenantId, debugInterfaceId, endpointInterfaceId uint) (ret iris.Map, err error) {
	po, _ := s.DebugRepo.GetLast(tenantId, debugInterfaceId, endpointInterfaceId)

	req := domain.DebugData{}
	resp := domain.DebugResponse{}

	if po.ID > 0 {
		json.Unmarshal([]byte(po.ReqContent), &req)
		json.Unmarshal([]byte(po.RespContent), &resp)

		resp.InvokeId = po.ID

	} else {
		resp = domain.DebugResponse{
			ContentLang: consts.LangHTML,
			Content:     "",
		}
	}

	ret = iris.Map{}
	ret["req"] = req
	ret["resp"] = resp

	return
}

func (s *DebugInvokeService) GetResult(tenantId consts.TenantId, invokeId int) (results []interface{}, err error) {
	invocation, err := s.DebugInvokeRepo.Get(tenantId, uint(invokeId))

	conditions, err := s.ConditionRepo.List(tenantId, invocation.DebugInterfaceId, invocation.EndpointInterfaceId, consts.ConditionCategoryResult, "", "false", "")

	for _, condition := range conditions {
		typ := condition.EntityType
		var log interface{}

		if typ == consts.ConditionTypeCheckpoint {
			log, _ = s.CheckpointRepo.GetLog(tenantId, condition.ID, uint(invokeId))
			results = append(results, log)

		} else if typ == consts.ConditionTypeResponseDefine {
			log, _ = s.ResponseDefineRepo.GetLog(tenantId, condition.ID, uint(invokeId))
			results = append(results, log)

		} else if typ == consts.ConditionTypeScript {
			logs, _ := s.CheckpointRepo.GetLogFromScriptAssert(tenantId, condition.ID, uint(invokeId))
			for _, item := range logs {
				results = append(results, item)
			}
		}
	}

	return
}

func (s *DebugInvokeService) GetLog(tenantId consts.TenantId, invokeId int) (results []interface{}, err error) {
	invocation, err := s.DebugInvokeRepo.Get(tenantId, uint(invokeId))

	preConditions, err := s.ConditionRepo.List(tenantId, invocation.DebugInterfaceId, invocation.EndpointInterfaceId,
		consts.ConditionCategoryConsole, "", "false", consts.ConditionSrcPre)

	for _, condition := range preConditions {
		if condition.Disabled {
			continue
		}

		typ := condition.EntityType
		var log interface{}

		if typ == consts.ConditionTypeScript {
			log, _ = s.ScriptRepo.GetLog(tenantId, condition.ID, uint(invokeId))
		} else if typ == consts.ConditionTypeDatabase {
			log, _ = s.DatabaseOptRepo.GetLog(tenantId, condition.ID, uint(invokeId))
		}

		if log != nil {
			results = append(results, log)
		}
	}

	postConditions, err := s.ConditionRepo.List(tenantId, invocation.DebugInterfaceId, invocation.EndpointInterfaceId,
		consts.ConditionCategoryConsole, "", "false", consts.ConditionSrcPost)

	for _, condition := range postConditions {
		if condition.Disabled {
			continue
		}

		typ := condition.EntityType
		var log interface{}

		if typ == consts.ConditionTypeExtractor {
			log, _ = s.ExtractorRepo.GetLog(tenantId, condition.ID, uint(invokeId))

		} else if typ == consts.ConditionTypeScript {
			log, _ = s.ScriptRepo.GetLog(tenantId, condition.ID, uint(invokeId))

		} else if typ == consts.ConditionTypeDatabase {
			log, _ = s.DatabaseOptRepo.GetLog(tenantId, condition.ID, uint(invokeId))

		}

		if log != nil {
			results = append(results, log)
		}
	}

	for _, condition := range postConditions {
		if condition.Disabled {
			continue
		}

		typ := condition.EntityType
		var log interface{}

		if typ == consts.ConditionTypeCheckpoint {
			log, _ = s.CheckpointRepo.GetLog(tenantId, condition.ID, uint(invokeId))

		}

		if log != nil {
			results = append(results, log)
		}
	}

	return
}

func (s *DebugInvokeService) GetAsInterface(tenantId consts.TenantId, id int) (debugData domain.DebugData,
	resultReq domain.DebugData, resultResp domain.DebugResponse, err error) {

	invocation, err := s.DebugInvokeRepo.Get(tenantId, uint(id))

	// deal with query params
	json.Unmarshal([]byte(invocation.ReqContent), &debugData)
	queryParams := []domain.Param{}

	if debugData.QueryParams != nil {
		for _, param := range *debugData.QueryParams {
			if param.ParamIn == consts.ParamInQuery { // ignore params from project settings
				queryParams = append(queryParams, param)
			}
		}
	}
	debugData.QueryParams = &queryParams

	// update request data
	debugPo := model.DebugInterface{}
	s.DebugInterfaceService.CopyValueFromRequest(tenantId, &debugPo, debugData)
	if resultReq.DebugInterfaceId > 0 {
		debugPo.ID = resultReq.DebugInterfaceId
	}
	err = s.ScenarioInterfaceRepo.SaveDebugData(tenantId, &debugPo)

	// save pre/post conditions
	preConditions := []domain.InterfaceExecCondition{}
	postConditions := []domain.InterfaceExecCondition{}
	json.Unmarshal([]byte(invocation.PreConditionsContent), &preConditions)
	json.Unmarshal([]byte(invocation.PostConditionsContent), &postConditions)
	s.ConditionRepo.ReplaceAll(tenantId, debugData.DebugInterfaceId, debugData.EndpointInterfaceId, preConditions, debugData.UsedBy, consts.ConditionSrcPre)
	s.ConditionRepo.ReplaceAll(tenantId, debugData.DebugInterfaceId, debugData.EndpointInterfaceId, postConditions, debugData.UsedBy, consts.ConditionSrcPre)

	// response data to show
	resultReq = debugData
	json.Unmarshal([]byte(invocation.RespContent), &resultResp)

	return
}

func (s *DebugInvokeService) Delete(tenantId consts.TenantId, id uint) (err error) {
	err = s.DebugRepo.Delete(tenantId, id)

	return
}
