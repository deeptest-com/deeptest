package agentExec

import (
	"encoding/json"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	checkpointHelper "github.com/aaronchen2k/deeptest/internal/pkg/helper/checkpoint"
	databaseOptHelpper "github.com/aaronchen2k/deeptest/internal/pkg/helper/database-opt"
	extractorHelper "github.com/aaronchen2k/deeptest/internal/pkg/helper/extractor"
	scriptHelper "github.com/aaronchen2k/deeptest/internal/pkg/helper/script"
	"log"
	"regexp"
	"strings"
)

func ExecPreConditions(execObj *InterfaceExecObj, execUuid string) (status consts.ResultStatus, err error) {
	status = consts.Pass

	preConditions := make([]domain.InterfaceExecCondition, 0) // will be changed and append items to it

	for _, condition := range execObj.PreConditions {
		if condition.Type == consts.ConditionTypeScript {
			DealwithScriptCondition(condition, &status, execObj.DebugData.ProjectId, &preConditions,
				execUuid, true)

		} else if condition.Type == consts.ConditionTypeDatabase {
			DealwithDatabaseCondition(condition, &status, &preConditions, execUuid)

		}
	}

	execObj.PreConditions = preConditions

	return
}

func ExecPostConditions(execObj *InterfaceExecObj, resp domain.DebugResponse, execUuid string) (status consts.ResultStatus, err error) {
	status = consts.Pass
	postConditions := make([]domain.InterfaceExecCondition, 0) // will be changed and append items to it

	for _, condition := range execObj.PostConditions {
		if condition.Type == consts.ConditionTypeScript {
			DealwithScriptCondition(condition, &status, execObj.DebugData.ProjectId, &postConditions,
				execUuid, true)

		} else if condition.Type == consts.ConditionTypeDatabase {
			DealwithDatabaseCondition(condition, &status, &postConditions, execUuid)

		} else if condition.Type == consts.ConditionTypeExtractor {
			DealwithExtractorCondition(condition, resp, &status, &postConditions, execUuid)

		} else if condition.Type == consts.ConditionTypeResponseDefine {
			DealwithResponseDefineCondition(condition, resp, &status, &postConditions)
		}
	}

	for _, condition := range execObj.PostConditions {
		if condition.Type == consts.ConditionTypeCheckpoint {
			DealwithDealwithCheckPointCondition(condition, resp, &status, &postConditions, execUuid)

		}
	}

	execObj.PostConditions = postConditions

	return
}

func DealwithScriptCondition(condition domain.InterfaceExecCondition, status *consts.ResultStatus,
	projectId uint, conditions *[]domain.InterfaceExecCondition, execUuid string, isPostCondition bool) {

	var scriptBase domain.ScriptBase
	json.Unmarshal(condition.Raw, &scriptBase)
	if scriptBase.Disabled {
		return
	}

	err := ExecScript(&scriptBase, projectId, execUuid)
	if err != nil {
		*status = consts.Fail
	}

	scriptHelper.GenResultMsg(&scriptBase)
	scriptBase.VariableSettings = GetGojaVariables(execUuid)

	condition.Raw, _ = json.Marshal(scriptBase)
	*conditions = append(*conditions, condition)

	if isPostCondition {
		for _, item := range GetGojaLogs(execUuid) {
			// Assertion Failed: [NAME] ERROR.
			// Assertion Pass: [NAME].

			regx := regexp.MustCompile(`Assertion (Failed|Pass) \[(.+)\](.*)\.`)
			arr := regx.FindAllStringSubmatch(item, -1)
			log.Println(arr)

			if len(arr) == 0 {
				continue
			}

			statusStr := strings.ToLower(arr[0][1])
			name := arr[0][2]
			//err := arr[0][3]

			checkpoint := domain.CheckpointBase{
				Type:      consts.Script,
				ResultMsg: strings.Replace(strings.Trim(item, "\""), "AssertionError", "", -1),

				ConditionId:         scriptBase.ConditionId,
				ConditionEntityId:   scriptBase.ConditionEntityId,
				ConditionEntityType: consts.ConditionTypeCheckpoint,
			}

			if statusStr == "failed" {
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
	}
}

func DealwithDatabaseCondition(condition domain.InterfaceExecCondition, status *consts.ResultStatus,
	postConditions *[]domain.InterfaceExecCondition, execUuid string) {

	var databaseOptBase domain.DatabaseOptBase
	json.Unmarshal(condition.Raw, &databaseOptBase)
	if databaseOptBase.Disabled {
		return
	}

	err := ExecDbOpt(&databaseOptBase)
	if err != nil || databaseOptBase.ResultStatus == consts.Fail {
		*status = consts.Fail
	}

	databaseOptHelpper.GenResultMsg(&databaseOptBase)

	if databaseOptBase.JsonPath != "" && databaseOptBase.Variable != "" && *status == consts.Pass {
		SetVariable(0, databaseOptBase.Variable, databaseOptBase.Result, databaseOptBase.ResultType,
			consts.Public, execUuid)
	} else {
		*status = consts.Fail
	}

	condition.Raw, _ = json.Marshal(databaseOptBase)
	*postConditions = append(*postConditions, condition)
}

func DealwithExtractorCondition(condition domain.InterfaceExecCondition, resp domain.DebugResponse,
	status *consts.ResultStatus, postConditions *[]domain.InterfaceExecCondition, execUuid string) {

	var extractorBase domain.ExtractorBase
	json.Unmarshal(condition.Raw, &extractorBase)

	if extractorBase.Disabled || extractorBase.Variable == "" {
		return
	}

	err := ExecExtract(&extractorBase, resp)
	if err != nil {
		*status = consts.Fail
	}

	extractorHelper.GenResultMsg(&extractorBase)

	if extractorBase.ResultStatus == consts.Pass {
		SetVariable(0, extractorBase.Variable, extractorBase.Result, extractorBase.ResultType, extractorBase.Scope, execUuid)
	} else {
		*status = consts.Fail
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
	status *consts.ResultStatus, postConditions *[]domain.InterfaceExecCondition) {

	var responseDefineBase domain.ResponseDefineBase
	json.Unmarshal(condition.Raw, &responseDefineBase)
	if responseDefineBase.Disabled {
		return
	}

	err := ExecResponseDefine(&responseDefineBase, resp)
	if err != nil || responseDefineBase.ResultStatus == consts.Fail {
		*status = consts.Fail
	}

	condition.Raw, _ = json.Marshal(responseDefineBase)
	//openapi.GenResultMsg(&responseDefineBase)
	*postConditions = append(*postConditions, condition)
}
