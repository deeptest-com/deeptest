package agentExec

import (
	"encoding/json"
	"github.com/aaronchen2k/deeptest/internal/agent/exec/domain"
	"github.com/aaronchen2k/deeptest/internal/agent/exec/utils/exec"
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

func (entity ProcessorInterface) Run(processor *Processor, session *Session) (err error) {
	defer func() {
		if errX := recover(); errX != nil {
			processor.Error(session, errX)
		}
	}()
	logUtils.Infof("interface entity")
	SetCurrDebugInterfaceId(session.ExecUuid, processor.EntityId)

	execStartTime := time.Now()
	processor.Result = &agentDomain.ScenarioExecResult{
		ID:                  int(entity.ProcessorID),
		Name:                processor.Name,
		ProcessorCategory:   entity.ProcessorCategory,
		ProcessorType:       entity.ProcessorType,
		StartTime:           &execStartTime,
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
	InitJsRuntime(processor.ProjectId, session.ExecUuid)
	SetReqValueToGoja(&baseRequest)

	// exec pre-condition
	entity.ExecPreConditions(processor, session)

	// dealwith variables
	ReplaceVariables(&baseRequest, session.ExecUuid)

	GetReqValueFromGoja(session.ExecUuid)

	// add cookies
	DealwithCookies(&baseRequest, entity.ProcessorID, session.ExecUuid)

	// gen request url
	GenRequestUrlWithBaseUrlAndPathParam(&baseRequest, processor.EntityId, entity.BaseUrl, session.ExecUuid)

	// send request
	requestStartTime := time.Now()
	entity.Response, err = Invoke(&baseRequest)
	requestEndTime := time.Now()

	// exec post-condition
	SetRespValueToGoja(&entity.Response)
	processor.Result.ResultStatus, _ = entity.ExecPostConditions(processor, &detail, session)
	GetRespValueFromGoja(session.ExecUuid)
	processor.Result.Detail = commonUtils.JsonEncode(detail)

	// get the response data updated by script post-condition
	if GetCurrResponse(session.ExecUuid).Data != nil {
		entity.Response = GetCurrResponse(session.ExecUuid)
	}

	// dealwith response
	ok := entity.GenResultFromResponse(processor, baseRequest, requestEndTime, requestStartTime, &detail, session, err)
	if !ok {
		return
	}

	for _, c := range entity.Response.Cookies {
		SetCookie(processor.ParentId, c.Name, c.Value, c.Domain, c.ExpireTime, session.ExecUuid)
	}

	execUtils.SendExecMsg(*processor.Result, consts.Processor, session.WsMsg)

	endTime := time.Now()
	processor.Result.EndTime = &endTime

	stat := CountStat(session.ExecUuid, processor.Result)
	execUtils.SendStatMsg(stat, session.WsMsg)
	processor.AddResultToParent()

	return
}

func (entity *ProcessorInterface) ExecPreConditions(processor *Processor, session *Session) (err error) {
	for _, condition := range entity.PreConditions {
		if condition.Type == consts.ConditionTypeScript {
			entity.DealwithScriptCondition(condition, nil, processor.ProjectId, &processor.Result.PreConditions,
				session.ExecUuid, false)

		} else if condition.Type == consts.ConditionTypeDatabase {
			entity.DealwithDatabaseOptCondition(condition, processor.ID, processor.ParentId, &processor.Result.PreConditions,
				session.ExecUuid)
		}
	}

	return
}

func (entity *ProcessorInterface) ExecPostConditions(processor *Processor, detail *map[string]interface{}, session *Session) (
	interfaceStatus consts.ResultStatus, err error) {
	interfaceStatus = processor.Result.ResultStatus
	for _, condition := range entity.PostConditions {
		if condition.Type == consts.ConditionTypeScript {
			entity.DealwithScriptCondition(condition, &interfaceStatus, processor.ProjectId, &processor.Result.PostConditions,
				session.ExecUuid, true)

		} else if condition.Type == consts.ConditionTypeDatabase {
			entity.DealwithDatabaseOptCondition(condition, processor.ID, processor.ParentId, &processor.Result.PostConditions, session.ExecUuid)

		} else if condition.Type == consts.ConditionTypeExtractor {
			entity.DealwithExtractorCondition(condition,
				processor.ID, processor.ParentId, &processor.Result.PostConditions, session.ExecUuid)

		} else if condition.Type == consts.ConditionTypeResponseDefine {
			entity.DealwithResponseDefineCondition(condition, &interfaceStatus, &processor.Result.PostConditions, detail, session.ExecUuid)

		}
	}

	for _, condition := range entity.PostConditions {
		if condition.Type == consts.ConditionTypeCheckpoint {
			entity.DealwithCheckpointCondition(condition, &interfaceStatus, &processor.Result.PostConditions,
				detail, session.ExecUuid)
		}
	}

	return
}

func (entity *ProcessorInterface) DealwithScriptCondition(condition domain.InterfaceExecCondition,
	interfaceStatus *consts.ResultStatus, projectId uint, conditions *[]domain.InterfaceExecCondition,
	execUuid string, isPostCondition bool) {

	var scriptBase domain.ScriptBase
	json.Unmarshal(condition.Raw, &scriptBase)
	if scriptBase.Disabled {
		return
	}

	err := ExecScript(&scriptBase, projectId, execUuid) // will set vari
	if err != nil {
	}

	scriptHelper.GenResultMsg(&scriptBase)
	scriptBase.VariableSettings = GetGojaVariables(execUuid)

	interfaceExecCondition := domain.InterfaceExecCondition{
		Type: condition.Type,
	}
	interfaceExecCondition.Raw, _ = json.Marshal(scriptBase)
	*conditions = append(*conditions, interfaceExecCondition)

	if isPostCondition {
		for _, item := range GetGojaLogs(execUuid) {
			createAssertFromScriptResult(item, conditions, interfaceStatus,
				scriptBase.ConditionId, scriptBase.ConditionEntityId)
		}
	}
}

func (entity *ProcessorInterface) DealwithDatabaseOptCondition(condition domain.InterfaceExecCondition,
	processorId, parentId uint, conditions *[]domain.InterfaceExecCondition, execUuid string) {

	var databaseOptBase domain.DatabaseOptBase
	json.Unmarshal(condition.Raw, &databaseOptBase)
	if databaseOptBase.Disabled {
		return
	}

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

		SetVariable(scopeId, databaseOptBase.Variable, databaseOptBase.Result, databaseOptBase.ResultType,
			consts.Public, execUuid)
	}

	condition.Raw, _ = json.Marshal(databaseOptBase)
	*conditions = append(*conditions, condition)
}

func (entity *ProcessorInterface) DealwithExtractorCondition(condition domain.InterfaceExecCondition,
	processorId, parentId uint, conditions *[]domain.InterfaceExecCondition, execUuid string) {

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

		SetVariable(scopeId, extractorBase.Variable, extractorBase.Result,
			extractorBase.ResultType, extractorBase.Scope, execUuid)
	}

	interfaceExecCondition := domain.InterfaceExecCondition{
		Type: condition.Type,
	}
	interfaceExecCondition.Raw, _ = json.Marshal(extractorBase)
	*conditions = append(*conditions, interfaceExecCondition)
}

func (entity *ProcessorInterface) DealwithCheckpointCondition(condition domain.InterfaceExecCondition,
	interfaceStatus *consts.ResultStatus, conditions *[]domain.InterfaceExecCondition,
	detail *map[string]interface{}, execUuid string) {

	var checkpointBase domain.CheckpointBase
	json.Unmarshal(condition.Raw, &checkpointBase)
	if checkpointBase.Disabled {
		return
	}

	resp := entity.Response
	err := ExecCheckPoint(&checkpointBase, resp, 0, execUuid)
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
	detail *map[string]interface{}, session *Session, err error) (ok bool) {

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
