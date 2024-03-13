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

func ExecPreConditions(session *ExecSession, execObj *InterfaceExecObj) (err error) {
	preConditions := make([]domain.InterfaceExecCondition, 0) // will be changed and append items to it

	for _, condition := range execObj.PreConditions {
		if condition.Type == consts.ConditionTypeScript {
			DealwithScriptCondition(session, condition, nil, execObj.DebugData.ProjectId, &preConditions,
				false)

		} else if condition.Type == consts.ConditionTypeDatabase {
			DealwithDatabaseCondition(session, condition, &preConditions)

		}
	}

	execObj.PreConditions = preConditions

	return
}

func ExecPostConditions(session *ExecSession, execObj *InterfaceExecObj, resp domain.DebugResponse) (resultStatus consts.ResultStatus, err error) {
	resultStatus = consts.Pass
	postConditions := make([]domain.InterfaceExecCondition, 0) // will be changed and append items to it

	for _, condition := range execObj.PostConditions {
		if condition.Type == consts.ConditionTypeScript {
			DealwithScriptCondition(session, condition, &resultStatus, execObj.DebugData.ProjectId, &postConditions,
				true)

		} else if condition.Type == consts.ConditionTypeDatabase {
			DealwithDatabaseCondition(session, condition, &postConditions)

		} else if condition.Type == consts.ConditionTypeExtractor {
			DealwithExtractorCondition(session, condition, resp, &postConditions)

		} else if condition.Type == consts.ConditionTypeResponseDefine {
			DealwithResponseDefineCondition(condition, resp, &resultStatus, &postConditions)
		}
	}

	for _, condition := range execObj.PostConditions {
		if condition.Type == consts.ConditionTypeCheckpoint {
			DealwithDealwithCheckPointCondition(session, condition, resp, &resultStatus, &postConditions)

		}
	}

	execObj.PostConditions = postConditions

	return
}

func DealwithScriptCondition(session *ExecSession, condition domain.InterfaceExecCondition, resultStatus *consts.ResultStatus,
	projectId uint, conditions *[]domain.InterfaceExecCondition, isPostCondition bool) {

	var scriptBase domain.ScriptBase
	json.Unmarshal(condition.Raw, &scriptBase)
	if scriptBase.Disabled {
		return
	}

	err := ExecScript(session, &scriptBase, projectId)
	if err != nil {
		_logUtils.Info("script exec failed")
	}

	scriptHelper.GenResultMsg(&scriptBase)
	scriptBase.VariableSettings = *session.GojaVariables

	condition.Raw, _ = json.Marshal(scriptBase)
	*conditions = append(*conditions, condition)

	for _, item := range *session.GojaLogs {
		if isPostCondition {
			createAssertFromScriptResult(item, conditions, resultStatus, scriptBase.ConditionId, scriptBase.ConditionEntityId)
		}
	}
}

func DealwithDatabaseCondition(session *ExecSession, condition domain.InterfaceExecCondition,
	postConditions *[]domain.InterfaceExecCondition) {

	status := consts.Pass

	var databaseOptBase domain.DatabaseOptBase
	json.Unmarshal(condition.Raw, &databaseOptBase)
	if databaseOptBase.Disabled {
		return
	}

	databaseOptBase.Sql = ReplaceVariableValue(session, databaseOptBase.Sql)

	err := ExecDbOpt(&databaseOptBase)
	if err != nil || databaseOptBase.ResultStatus == consts.Fail {
		status = consts.Fail
	}

	databaseOptHelpper.GenResultMsg(&databaseOptBase)

	if databaseOptBase.JsonPath != "" && databaseOptBase.Variable != "" && status == consts.Pass {
		SetVariable(session, 0, databaseOptBase.Variable, databaseOptBase.Result, databaseOptBase.ResultType,
			consts.Public)
	}

	condition.Raw, _ = json.Marshal(databaseOptBase)
	*postConditions = append(*postConditions, condition)
}

func DealwithExtractorCondition(session *ExecSession, condition domain.InterfaceExecCondition, resp domain.DebugResponse,
	postConditions *[]domain.InterfaceExecCondition) {

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
		SetVariable(session, 0, extractorBase.Variable, extractorBase.Result, extractorBase.ResultType, extractorBase.Scope)
	}

	condition.Raw, _ = json.Marshal(extractorBase)
	*postConditions = append(*postConditions, condition)
}

func DealwithDealwithCheckPointCondition(session *ExecSession, condition domain.InterfaceExecCondition, resp domain.DebugResponse,
	status *consts.ResultStatus, postConditions *[]domain.InterfaceExecCondition) {

	var checkpointBase domain.CheckpointBase
	json.Unmarshal(condition.Raw, &checkpointBase)
	if checkpointBase.Disabled {
		return
	}

	err := ExecCheckPoint(session, &checkpointBase, resp, 0)
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
