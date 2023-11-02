package service

import (
	"encoding/json"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	model "github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
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

	PreConditionRepo   *repo.PreConditionRepo   `inject:""`
	PostConditionRepo  *repo.PostConditionRepo  `inject:""`
	ExtractorRepo      *repo.ExtractorRepo      `inject:""`
	CheckpointRepo     *repo.CheckpointRepo     `inject:""`
	ScriptRepo         *repo.ScriptRepo         `inject:""`
	ResponseDefineRepo *repo.ResponseDefineRepo `inject:""`

	ScenarioInterfaceRepo *repo.ScenarioInterfaceRepo `inject:""`
}

func (s *DebugInvokeService) SubmitResult(req domain.SubmitDebugResultRequest) (invoke model.DebugInvoke, err error) {
	usedBy := req.Request.UsedBy
	var endpointId, serveId, processorId, scenarioId, projectId uint

	if usedBy == consts.InterfaceDebug {
		endpointId, serveId = s.DebugInterfaceService.GetEndpointAndServeIdForEndpointInterface(req.Request.EndpointInterfaceId)

		endpoint, _ := s.EndpointRepo.Get(endpointId)
		serveId = endpoint.ServeId
		projectId = endpoint.ProjectId

	} else if usedBy == consts.ScenarioDebug {
		processorId = req.Request.ScenarioProcessorId
		scenarioId = s.DebugInterfaceService.GetScenarioIdForDebugInterface(req.Request.ScenarioProcessorId)

		scenario, _ := s.ScenarioRepo.Get(scenarioId)
		projectId = scenario.ProjectId

	} else if usedBy == consts.DiagnoseDebug {
		diagnoseInterface, _ := s.DiagnoseInterfaceRepo.Get(req.Request.DiagnoseInterfaceId)

		serveId = diagnoseInterface.ServeId
		projectId = diagnoseInterface.ProjectId

	} else if usedBy == consts.CaseDebug {
		caseInterface, _ := s.EndpointCaseRepo.Get(req.Request.CaseInterfaceId)

		serveId = caseInterface.ServeId
		projectId = caseInterface.ProjectId
	}

	invoke, err = s.Create(req, serveId, processorId, scenarioId, projectId)

	s.ExecConditionService.SavePreConditionResult(invoke.ID,
		req.Request.DebugInterfaceId, req.Request.CaseInterfaceId, req.Request.EndpointInterfaceId,
		serveId, processorId, scenarioId, usedBy,
		req.PreConditions)

	s.ExecConditionService.SavePostConditionResult(invoke.ID,
		req.Request.DebugInterfaceId, req.Request.CaseInterfaceId, req.Request.EndpointInterfaceId,
		serveId, processorId, scenarioId, usedBy,
		req.PostConditions)

	if err != nil {
		return
	}

	return
}

func (s *DebugInvokeService) Create(req domain.SubmitDebugResultRequest,
	serveId, scenarioProcessorId, scenarioId, projectId uint) (po model.DebugInvoke, err error) {

	debugInterface, _ := s.DebugInterfaceRepo.Get(req.Request.DebugInterfaceId)

	po = model.DebugInvoke{
		ServeId: serveId,

		ScenarioProcessorId: scenarioProcessorId,
		ScenarioId:          scenarioId,

		InvocationBase: model.InvocationBase{
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

	err = s.DebugInvokeRepo.Save(&po)

	return
}

func (s *DebugInvokeService) ListByInterface(debugInterfaceId, endpointInterfaceId uint) (invocations []model.DebugInvoke, err error) {
	invocations, err = s.DebugRepo.List(debugInterfaceId, endpointInterfaceId)

	return
}

func (s *DebugInvokeService) GetLastResp(debugInterfaceId, endpointInterfaceId uint) (ret iris.Map, err error) {
	po, _ := s.DebugRepo.GetLast(debugInterfaceId, endpointInterfaceId)

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

func (s *DebugInvokeService) GetResult(invokeId int) (results []interface{}, err error) {
	invocation, err := s.DebugInvokeRepo.Get(uint(invokeId))

	conditions, err := s.PostConditionRepo.List(invocation.DebugInterfaceId, invocation.EndpointInterfaceId, consts.ConditionCategoryResult, "")

	for _, condition := range conditions {
		typ := condition.EntityType
		var log interface{}

		if typ == consts.ConditionTypeCheckpoint {
			log, _ = s.CheckpointRepo.GetLog(condition.ID, uint(invokeId))
			results = append(results, log)

		} else if typ == consts.ConditionTypeResponseDefine {
			log, _ = s.ResponseDefineRepo.GetLog(condition.ID, uint(invokeId))
			results = append(results, log)

		} else if typ == consts.ConditionTypeScript {
			logs, _ := s.CheckpointRepo.GetLogFromScriptAssert(condition.ID, uint(invokeId))
			for _, item := range logs {
				results = append(results, item)
			}
		}
	}

	return
}

func (s *DebugInvokeService) GetLog(invokeId int) (results []interface{}, err error) {
	invocation, err := s.DebugInvokeRepo.Get(uint(invokeId))

	preConditions, err := s.PreConditionRepo.List(invocation.DebugInterfaceId, invocation.EndpointInterfaceId, "")
	for _, condition := range preConditions {
		typ := condition.EntityType
		var log interface{}

		if typ == consts.ConditionTypeScript {
			log, _ = s.ScriptRepo.GetLog(condition.ID, uint(invokeId))
		}

		results = append(results, log)
	}

	postConditions, err := s.PostConditionRepo.List(invocation.DebugInterfaceId, invocation.EndpointInterfaceId, consts.ConditionCategoryConsole, "")
	for _, condition := range postConditions {
		typ := condition.EntityType
		var log interface{}

		if typ == consts.ConditionTypeExtractor {
			log, _ = s.ExtractorRepo.GetLog(condition.ID, uint(invokeId))

		} else if typ == consts.ConditionTypeCheckpoint {
			log, _ = s.CheckpointRepo.GetLog(condition.ID, uint(invokeId))

		} else if typ == consts.ConditionTypeScript {
			log, _ = s.ScriptRepo.GetLog(condition.ID, uint(invokeId))

		}

		results = append(results, log)
	}

	return
}

func (s *DebugInvokeService) GetAsInterface(id int) (debugData domain.DebugData,
	resultReq domain.DebugData, resultResp domain.DebugResponse, err error) {

	invocation, err := s.DebugInvokeRepo.Get(uint(id))

	// deal with query params
	json.Unmarshal([]byte(invocation.ReqContent), &debugData)
	queryParams := []domain.Param{}
	for _, param := range debugData.QueryParams {
		if param.ParamIn == consts.ParamInQuery { // ignore params from project settings
			queryParams = append(queryParams, param)
		}
	}
	debugData.QueryParams = queryParams

	// update request data
	debugPo := model.DebugInterface{}
	s.DebugInterfaceService.CopyValueFromRequest(&debugPo, debugData)
	if resultReq.DebugInterfaceId > 0 {
		debugPo.ID = resultReq.DebugInterfaceId
	}
	err = s.ScenarioInterfaceRepo.SaveDebugData(&debugPo)

	// save pre/post conditions
	preConditions := []domain.InterfaceExecCondition{}
	postConditions := []domain.InterfaceExecCondition{}
	json.Unmarshal([]byte(invocation.PreConditionsContent), &preConditions)
	json.Unmarshal([]byte(invocation.PostConditionsContent), &postConditions)
	s.PreConditionRepo.ReplaceAll(debugData.DebugInterfaceId, debugData.EndpointInterfaceId, preConditions, debugData.UsedBy)
	s.PostConditionRepo.ReplaceAll(debugData.DebugInterfaceId, debugData.EndpointInterfaceId, postConditions, debugData.UsedBy)

	// response data to show
	resultReq = debugData
	json.Unmarshal([]byte(invocation.RespContent), &resultResp)

	return
}

func (s *DebugInvokeService) Delete(id uint) (err error) {
	err = s.DebugRepo.Delete(id)

	return
}
