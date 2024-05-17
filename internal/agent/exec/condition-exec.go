package agentExec

import (
	"encoding/json"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	checkpointHelper "github.com/aaronchen2k/deeptest/internal/pkg/helper/checkpoint"
	databaseOptHelpper "github.com/aaronchen2k/deeptest/internal/pkg/helper/database-opt"
	extractorHelper "github.com/aaronchen2k/deeptest/internal/pkg/helper/extractor"
	scriptHelper "github.com/aaronchen2k/deeptest/internal/pkg/helper/script"
	_logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"regexp"
	"strings"
)

func ExecPreConditions(execObj *InterfaceExecObj, execUuid string) (err error) {
	preConditions := make([]domain.InterfaceExecCondition, 0) // will be changed and append items to it

	for _, condition := range execObj.PreConditions {
		if condition.Type == consts.ConditionTypeScript {
			DealwithScriptCondition(condition, nil, execObj.DebugData.ProjectId, &preConditions,
				execUuid, false, execObj.TenantId)

		} else if condition.Type == consts.ConditionTypeDatabase {
			DealwithDatabaseCondition(condition, &preConditions, execObj.TenantId, execObj.DebugData.ProjectId, execUuid)

		}
	}

	execObj.PreConditions = preConditions

	return
}

func ExecPostConditions(execObj *InterfaceExecObj, resp domain.DebugResponse, execUuid string) (resultStatus consts.ResultStatus, err error) {
	resultStatus = consts.Pass
	postConditions := make([]domain.InterfaceExecCondition, 0) // will be changed and append items to it

	for _, condition := range execObj.PostConditions {
		if condition.Type == consts.ConditionTypeScript {
			DealwithScriptCondition(condition, &resultStatus, execObj.DebugData.ProjectId, &postConditions,
				execUuid, true, execObj.TenantId)

		} else if condition.Type == consts.ConditionTypeDatabase {
			DealwithDatabaseCondition(condition, &postConditions, execObj.TenantId, execObj.DebugData.ProjectId, execUuid)

		} else if condition.Type == consts.ConditionTypeExtractor {
			DealwithExtractorCondition(condition, resp, &postConditions, execUuid)

		} else if condition.Type == consts.ConditionTypeResponseDefine {
			//DealwithResponseDefineCondition(condition, resp, &resultStatus, &postConditions)
		}
	}

	for _, condition := range execObj.PostConditions {
		if condition.Type == consts.ConditionTypeCheckpoint {
			DealwithDealwithCheckPointCondition(condition, resp, &resultStatus, &postConditions, execUuid)

		}
	}

	execObj.PostConditions = postConditions

	return
}

func DealwithScriptCondition(condition domain.InterfaceExecCondition, resultStatus *consts.ResultStatus,
	projectId uint, conditions *[]domain.InterfaceExecCondition, execUuid string, isPostCondition bool, tenantId consts.TenantId) {

	var scriptBase domain.ScriptBase
	json.Unmarshal(condition.Raw, &scriptBase)
	if scriptBase.Disabled {
		return
	}
	err := ExecScript(&scriptBase, tenantId, projectId, execUuid)
	if err != nil {
		_logUtils.Info("script exec failed")
	}

	scriptHelper.GenResultMsg(&scriptBase)
	scriptBase.VariableSettings = *GetGojaVariables(execUuid)

	condition.Raw, _ = json.Marshal(scriptBase)
	*conditions = append(*conditions, condition)

	for _, item := range *GetGojaLogs(execUuid) {
		if isPostCondition {
			createAssertFromScriptResult(item, conditions, resultStatus, scriptBase.ConditionId, scriptBase.ConditionEntityId)
		}
	}
}

func DealwithDatabaseCondition(condition domain.InterfaceExecCondition,
	postConditions *[]domain.InterfaceExecCondition, tenantId consts.TenantId, projectId uint, execUuid string) {

	status := consts.Pass

	var databaseOptBase domain.DatabaseOptBase
	json.Unmarshal(condition.Raw, &databaseOptBase)
	if databaseOptBase.Disabled {
		return
	}

	databaseOptBase.Sql = ReplaceVariableValue(databaseOptBase.Sql, tenantId, projectId, execUuid)

	err := ExecDbOpt(&databaseOptBase)
	if err != nil || databaseOptBase.ResultStatus == consts.Fail {
		status = consts.Fail
	}

	databaseOptHelpper.GenResultMsg(&databaseOptBase)

	if databaseOptBase.JsonPath != "" && databaseOptBase.Variable != "" && status == consts.Pass {
		SetVariable(0, databaseOptBase.Variable, databaseOptBase.Result, databaseOptBase.ResultType,
			consts.Public, execUuid)
	}

	condition.Raw, _ = json.Marshal(databaseOptBase)
	*postConditions = append(*postConditions, condition)
}

func DealwithExtractorCondition(condition domain.InterfaceExecCondition, resp domain.DebugResponse,
	postConditions *[]domain.InterfaceExecCondition, execUuid string) {

	var extractorBase domain.ExtractorBase
	json.Unmarshal(condition.Raw, &extractorBase)

	if extractorBase.Disabled || extractorBase.Variable == "" {
		return
	}

	err := ExecExtract(&extractorBase, resp)
	if err != nil {
		_logUtils.Info("extract failed")
	}

	extractorHelper.GenResultMsg(&extractorBase)

	if extractorBase.ResultStatus == consts.Pass {
		SetVariable(0, extractorBase.Variable, extractorBase.Result, extractorBase.ResultType, extractorBase.Scope, execUuid)
	}

	condition.Raw, _ = json.Marshal(extractorBase)
	*postConditions = append(*postConditions, condition)
}

func DealwithDealwithCheckPointCondition(condition domain.InterfaceExecCondition, resp domain.DebugResponse,
	status *consts.ResultStatus, postConditions *[]domain.InterfaceExecCondition, execUuid string) {

	var checkpointBase domain.CheckpointBase
	json.Unmarshal(condition.Raw, &checkpointBase)
	if checkpointBase.Disabled {
		return
	}

	err := ExecCheckPoint(&checkpointBase, resp, 0, execUuid)
	if err != nil || checkpointBase.ResultStatus == consts.Fail {
		*status = consts.Fail
	}

	checkpointHelper.GenResultMsg(&checkpointBase)

	condition.Raw, _ = json.Marshal(checkpointBase)
	*postConditions = append(*postConditions, condition)
}

func DealwithResponseDefineCondition(condition domain.InterfaceExecCondition, resp domain.DebugResponse,
	resultStatus *consts.ResultStatus, postConditions *[]domain.InterfaceExecCondition) {

	var responseDefineBase domain.ResponseDefineBase
	json.Unmarshal(condition.Raw, &responseDefineBase)
	if responseDefineBase.Disabled {
		return
	}

	err := ExecResponseDefine(&responseDefineBase, resp)
	if err != nil || responseDefineBase.ResultStatus == consts.Fail {
		*resultStatus = consts.Fail
	}

	condition.Raw, _ = json.Marshal(responseDefineBase)
	*postConditions = append(*postConditions, condition)
}

func createAssertFromScriptResult(output string, conditions *[]domain.InterfaceExecCondition,
	status *consts.ResultStatus, conditionId, conditionEntityId uint) {
	// Assertion Failed: [NAME] ERROR.
	// Assertion Pass: [NAME].

	regx := regexp.MustCompile(`Assertion (Failed|Pass) \[(.+)\](.*)\.`)
	arr := regx.FindAllStringSubmatch(output, -1)
	if len(arr) == 0 {
		return
	}

	statusStr, name, _ := ParseChaiAssertion(output)

	checkpoint := domain.CheckpointBase{
		Type:      consts.Script,
		ResultMsg: strings.Replace(strings.Trim(output, "\""), "AssertionError", "", -1),

		ConditionId:         conditionId,
		ConditionEntityId:   conditionEntityId,
		ConditionEntityType: consts.ConditionTypeCheckpoint,
	}

	if statusStr == "fail" {
		*status = consts.Fail
		checkpoint.ResultStatus = consts.Fail
	} else {
		checkpoint.ResultStatus = consts.Pass
	}

	newCheckPointCondition := domain.InterfaceExecCondition{
		Type: consts.ConditionTypeCheckpoint,
		Desc: name,
	}
	newCheckPointCondition.Raw, _ = json.Marshal(checkpoint)

	*conditions = append(*conditions, newCheckPointCondition)
}
