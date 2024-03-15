package agentExec

import (
	"encoding/json"
	"github.com/aaronchen2k/deeptest/internal/agent/exec/domain"
	"github.com/aaronchen2k/deeptest/internal/agent/exec/utils/exec"
	performanceUtils "github.com/aaronchen2k/deeptest/internal/agent/exec/utils/performance"
	ptproto "github.com/aaronchen2k/deeptest/internal/performance/proto"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	checkpointHelper "github.com/aaronchen2k/deeptest/internal/pkg/helper/checkpoint"
	databaseOptHelpper "github.com/aaronchen2k/deeptest/internal/pkg/helper/database-opt"
	extractorHelper "github.com/aaronchen2k/deeptest/internal/pkg/helper/extractor"
	scriptHelper "github.com/aaronchen2k/deeptest/internal/pkg/helper/script"
	commonUtils "github.com/aaronchen2k/deeptest/pkg/lib/comm"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"github.com/jinzhu/copier"
	uuid "github.com/satori/go.uuid"
	"time"
)

type ProcessorInterface struct {
	ID uint `json:"id"`
	ProcessorEntityBase

	domain.BaseRequest
	Response domain.DebugResponse `json:"response"`

	BaseUrl string `json:"baseUrl"`

	PreConditions  []domain.InterfaceExecCondition `json:"preConditions"`
	PostConditions []domain.InterfaceExecCondition `json:"postConditions"`
}

func (entity ProcessorInterface) Run(processor *Processor, session *ExecSession) (err error) {
	defer func() {
		if errX := recover(); errX != nil {
			processor.Error(session, errX)
		}
	}()
	logUtils.Infof("interface entity")
	session.CurrScenarioProcessorId = processor.ID
	session.CurrDebugInterfaceId = processor.EntityId

	startTime := time.Now()
	processor.Result = &agentExecDomain.ScenarioExecResult{
		ID:                  int(entity.ProcessorID),
		Name:                processor.Name,
		ProcessorCategory:   entity.ProcessorCategory,
		ProcessorType:       entity.ProcessorType,
		StartTime:           &startTime,
		ParentId:            int(entity.ParentID),
		EndpointInterfaceId: processor.EndpointInterfaceId,
		DebugInterfaceId:    processor.EntityId,
		ProcessorId:         processor.ID,
		ScenarioId:          processor.ScenarioId,
		LogId:               uuid.NewV4(),
		ParentLogId:         processor.Parent.Result.LogId,
		Round:               processor.Round,
		ResultStatus:        consts.Pass,
	}

	detail := map[string]interface{}{}
	//在循环过程中，processor 被执行多次，变量替换会受到影响，第一次跌替换之后，就不能根据实际情况替换了
	var baseRequest domain.BaseRequest
	copier.CopyWithOption(&baseRequest, &entity.BaseRequest, copier.Option{IgnoreEmpty: true, DeepCopy: true})

	// init context
	//InitJsRuntime(processor.ProjectId, session.ExecUuid)
	SetReqValueToGoja(session, &baseRequest)

	// exec pre-condition
	entity.ExecPreConditions(session, processor)

	// dealwith variables
	ReplaceVariables(session, &baseRequest)

	GetReqValueFromGoja(session)

	// add cookies
	DealwithCookies(session, entity.ProcessorID, &baseRequest)

	// gen request url
	GenRequestUrlWithBaseUrlAndPathParam(session, &baseRequest, processor.EntityId, entity.BaseUrl)

	// send request
	requestStartTime := time.Now()
	entity.Response, err = Invoke(&baseRequest)
	requestEndTime := time.Now()

	// exec post-condition
	SetRespValueToGoja(session, &entity.Response)
	processor.Result.ResultStatus, _ = entity.ExecPostConditions(session, processor, &detail)
	GetRespValueFromGoja(session)
	processor.Result.Detail = commonUtils.JsonEncode(detail)

	// get the response data updated by script post-condition
	if session.CurrResponse.Data != nil {
		entity.Response = session.CurrResponse
	}

	// dealwith response
	ok := entity.GenResultFromResponse(processor, baseRequest, requestEndTime, requestStartTime, &detail, session, err)
	if !ok {
		return
	}

	for _, c := range entity.Response.Cookies {
		SetCookie(session, processor.ParentId, c.Name, c.Value, c.Domain, c.ExpireTime)
	}

	execUtils.SendExecMsg(*processor.Result, consts.Processor, session.WsMsg)

	endTime := time.Now()
	processor.Result.EndTime = &endTime

	stat := CountInterfaceStat(session.ExecUuid, processor.Result)
	execUtils.SendStatMsg(stat, session.WsMsg)
	processor.AddResultToParent()

	if session.InfluxdbSender == nil { // not performance testing
		return
	}

	// send performance metrics
	execParams := performanceUtils.GetExecParamsInCtx(session.Ctx)

	result := ptproto.PerformanceExecResp{
		Timestamp: time.Now().UnixMilli(),
		RunnerId:  execParams.RunnerId,
		Room:      execParams.Room,

		Requests: []*ptproto.PerformanceExecRecord{
			{
				RecordId:   int32(processor.ID),
				RecordName: processor.Name,

				StartTime: startTime.UnixMilli(),
				EndTime:   endTime.UnixMilli(),
				Duration:  int32(endTime.UnixMilli() - startTime.UnixMilli()), // 毫秒
				Status:    processor.Result.ResultStatus.String(),

				VuId: int32(session.VuNo),
			},
		},
	}
	session.InfluxdbSender.Send(result)

	return
}

func (entity *ProcessorInterface) ExecPreConditions(session *ExecSession, processor *Processor) (err error) {
	for _, condition := range entity.PreConditions {
		if condition.Type == consts.ConditionTypeScript {
			entity.DealwithScriptCondition(session, condition, nil, processor.ProjectId, &processor.Result.PreConditions,
				false)

		} else if condition.Type == consts.ConditionTypeDatabase {
			entity.DealwithDatabaseOptCondition(session, condition, processor.ID, processor.ParentId,
				&processor.Result.PreConditions)
		}
	}

	return
}

func (entity *ProcessorInterface) ExecPostConditions(session *ExecSession, processor *Processor, detail *map[string]interface{}) (
	interfaceStatus consts.ResultStatus, err error) {
	interfaceStatus = processor.Result.ResultStatus
	for _, condition := range entity.PostConditions {
		if condition.Type == consts.ConditionTypeScript {
			entity.DealwithScriptCondition(session, condition, &interfaceStatus, processor.ProjectId,
				&processor.Result.PostConditions, true)

		} else if condition.Type == consts.ConditionTypeDatabase {
			entity.DealwithDatabaseOptCondition(session, condition, processor.ID, processor.ParentId, &processor.Result.PostConditions)

		} else if condition.Type == consts.ConditionTypeExtractor {
			entity.DealwithExtractorCondition(session, condition,
				processor.ID, processor.ParentId, &processor.Result.PostConditions)

		} else if condition.Type == consts.ConditionTypeResponseDefine {
			entity.DealwithResponseDefineCondition(condition, &interfaceStatus, &processor.Result.PostConditions, detail, session.ExecUuid)

		}
	}

	for _, condition := range entity.PostConditions {
		if condition.Type == consts.ConditionTypeCheckpoint {
			entity.DealwithCheckpointCondition(session, condition, &interfaceStatus, &processor.Result.PostConditions,
				detail)
		}
	}

	return
}

func (entity *ProcessorInterface) DealwithScriptCondition(session *ExecSession, condition domain.InterfaceExecCondition,
	interfaceStatus *consts.ResultStatus, projectId uint, conditions *[]domain.InterfaceExecCondition,
	isPostCondition bool) {

	var scriptBase domain.ScriptBase
	json.Unmarshal(condition.Raw, &scriptBase)
	if scriptBase.Disabled {
		return
	}

	err := ExecScript(session, &scriptBase, projectId) // will set vari
	if err != nil {
	}

	scriptHelper.GenResultMsg(&scriptBase)
	scriptBase.VariableSettings = *session.GojaVariables

	interfaceExecCondition := domain.InterfaceExecCondition{
		Type: condition.Type,
	}
	interfaceExecCondition.Raw, _ = json.Marshal(scriptBase)
	*conditions = append(*conditions, interfaceExecCondition)

	if isPostCondition {
		for _, item := range *session.GojaLogs {
			createAssertFromScriptResult(item, conditions, interfaceStatus,
				scriptBase.ConditionId, scriptBase.ConditionEntityId)
		}
	}
}

func (entity *ProcessorInterface) DealwithDatabaseOptCondition(session *ExecSession, condition domain.InterfaceExecCondition,
	processorId, parentId uint, conditions *[]domain.InterfaceExecCondition) {

	var databaseOptBase domain.DatabaseOptBase
	json.Unmarshal(condition.Raw, &databaseOptBase)
	if databaseOptBase.Disabled {
		return
	}

	databaseOptBase.Sql = ReplaceVariableValue(session, databaseOptBase.Sql)

	conditionStatus := true
	err := ExecDbOpt(&databaseOptBase)
	if err != nil || databaseOptBase.ResultStatus == consts.Fail {
		conditionStatus = false
	}

	databaseOptHelpper.GenResultMsg(&databaseOptBase)

	if databaseOptBase.JsonPath != "" && databaseOptBase.Variable != "" && conditionStatus { // will set vari
		scopeId := parentId
		if databaseOptBase.Scope == consts.Private { // put vari in its own scope if Private
			scopeId = processorId
		}

		SetVariable(session, scopeId, databaseOptBase.Variable, databaseOptBase.Result, databaseOptBase.ResultType,
			consts.Public)
	}

	condition.Raw, _ = json.Marshal(databaseOptBase)
	*conditions = append(*conditions, condition)
}

func (entity *ProcessorInterface) DealwithExtractorCondition(session *ExecSession, condition domain.InterfaceExecCondition,
	processorId, parentId uint, conditions *[]domain.InterfaceExecCondition) {

	var extractorBase domain.ExtractorBase
	json.Unmarshal(condition.Raw, &extractorBase)

	if extractorBase.Disabled || extractorBase.Variable == "" {
		return
	}

	resp := entity.Response

	err := ExecExtract(&extractorBase, resp)
	if err != nil || extractorBase.ResultStatus == consts.Fail {
		logUtils.Infof("extract failed")
	}

	extractorHelper.GenResultMsg(&extractorBase)

	if extractorBase.ResultStatus == consts.Pass {
		scopeId := parentId
		if extractorBase.Scope == consts.Private { // put vari in its own scope if Private
			scopeId = processorId
		}

		SetVariable(session, scopeId, extractorBase.Variable, extractorBase.Result,
			extractorBase.ResultType, extractorBase.Scope)
	}

	interfaceExecCondition := domain.InterfaceExecCondition{
		Type: condition.Type,
	}
	interfaceExecCondition.Raw, _ = json.Marshal(extractorBase)
	*conditions = append(*conditions, interfaceExecCondition)
}

func (entity *ProcessorInterface) DealwithCheckpointCondition(session *ExecSession, condition domain.InterfaceExecCondition,
	interfaceStatus *consts.ResultStatus, conditions *[]domain.InterfaceExecCondition,
	detail *map[string]interface{}) {

	var checkpointBase domain.CheckpointBase
	json.Unmarshal(condition.Raw, &checkpointBase)
	if checkpointBase.Disabled {
		return
	}

	resp := entity.Response
	err := ExecCheckPoint(session, &checkpointBase, resp, 0)
	if err != nil || checkpointBase.ResultStatus == consts.Fail {
		*interfaceStatus = consts.Fail
	}

	checkpointHelper.GenResultMsg(&checkpointBase)
	if checkpointBase.ResultStatus == consts.Fail {
		*interfaceStatus = consts.Fail
	}

	interfaceExecCondition := domain.InterfaceExecCondition{
		Type: condition.Type,
	}

	interfaceExecCondition.Raw, _ = json.Marshal(checkpointBase)
	*conditions = append(*conditions, interfaceExecCondition)

	if _, ok := (*detail)["checkpoint"]; !ok {
		(*detail)["checkpoint"] = []map[string]interface{}{}
	}
	(*detail)["checkpoint"] = append((*detail)["checkpoint"].([]map[string]interface{}), map[string]interface{}{
		"resultStatus": checkpointBase.ResultStatus, "resultMsg": checkpointBase.ResultMsg,
	})
}

func (entity *ProcessorInterface) DealwithResponseDefineCondition(condition domain.InterfaceExecCondition,
	interfaceStatus *consts.ResultStatus, conditions *[]domain.InterfaceExecCondition,
	detail *map[string]interface{}, execUuid string) {

	var responseDefineBase domain.ResponseDefineBase
	json.Unmarshal(condition.Raw, &responseDefineBase)
	if responseDefineBase.Disabled {
		return
	}

	resp := entity.Response

	err := ExecResponseDefine(&responseDefineBase, resp)
	if err != nil || responseDefineBase.ResultStatus == consts.Fail {
		*interfaceStatus = consts.Fail
	}

	interfaceExecCondition := domain.InterfaceExecCondition{
		Type: condition.Type,
	}

	interfaceExecCondition.Raw, _ = json.Marshal(responseDefineBase)
	*conditions = append(*conditions, interfaceExecCondition)

	(*detail)["responseDefine"] = map[string]interface{}{"resultStatus": responseDefineBase.ResultStatus, "resultMsg": responseDefineBase.ResultMsg}

}

func (entity *ProcessorInterface) GenResultFromResponse(
	processor *Processor, baseRequest domain.BaseRequest, requestEndTime, requestStartTime time.Time,
	detail *map[string]interface{}, session *ExecSession, err error) (ok bool) {

	processor.Result.Cost = requestEndTime.UnixMilli() - requestStartTime.UnixMilli()
	reqContent, _ := json.Marshal(baseRequest)
	processor.Result.ReqContent = string(reqContent)
	respContent, _ := json.Marshal(entity.Response)
	processor.Result.RespContent = string(respContent)

	if err != nil {
		processor.Result.ResultStatus = consts.Fail
		processor.Result.Summary = err.Error()

		(*detail)["result"] = entity.Response.Content
		processor.Result.Detail = commonUtils.JsonEncode(*detail)
		execUtils.SendErrorMsg(*processor.Result, consts.Processor, session.WsMsg)
		processor.AddResultToParent()

		return
	}

	ok = true
	return
}
