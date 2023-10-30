package agentExec

import (
	"encoding/json"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	checkpointHelper "github.com/aaronchen2k/deeptest/internal/pkg/helper/checkpoint"
	databaseOptHelpper "github.com/aaronchen2k/deeptest/internal/pkg/helper/database-opt"
	extractorHelper "github.com/aaronchen2k/deeptest/internal/pkg/helper/extractor"
	scriptHelper "github.com/aaronchen2k/deeptest/internal/pkg/helper/script"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
)

func ExecPreConditions(execObj InterfaceExecObj) (status consts.ResultStatus, err error) {
	status = consts.Pass

	for index, condition := range execObj.PreConditions {
		if condition.Type == consts.ConditionTypeScript {
			var scriptBase domain.ScriptBase
			json.Unmarshal(condition.Raw, &scriptBase)

			err = ExecScript(&scriptBase, execObj.DebugData.ProjectId)
			if err != nil {
				logUtils.Info(err.Error())
				status = consts.Fail
				return
			}
			scriptHelper.GenResultMsg(&scriptBase)
			scriptBase.VariableSettings = VariableSettings

			execObj.PreConditions[index].Raw, _ = json.Marshal(scriptBase)
		}
	}

	return
}

func ExecPostConditions(obj InterfaceExecObj, resp domain.DebugResponse) (status consts.ResultStatus, err error) {
	status = consts.Pass

	for index, condition := range obj.PostConditions {
		if condition.Type == consts.ConditionTypeExtractor {
			var extractorBase domain.ExtractorBase
			json.Unmarshal(condition.Raw, &extractorBase)

			if extractorBase.Disabled || extractorBase.Variable == "" {
				continue
			}

			err = ExecExtract(&extractorBase, resp)
			if err != nil {
				status = consts.Fail
			}

			extractorHelper.GenResultMsg(&extractorBase)

			if extractorBase.ResultStatus == consts.Pass {
				SetVariable(0, extractorBase.Variable, extractorBase.Result, consts.Public)
			} else {
				status = consts.Fail
			}

			obj.PostConditions[index].Raw, _ = json.Marshal(extractorBase)

		} else if condition.Type == consts.ConditionTypeScript {
			var scriptBase domain.ScriptBase
			json.Unmarshal(condition.Raw, &scriptBase)
			if scriptBase.Disabled {
				continue
			}

			err = ExecScript(&scriptBase, obj.DebugData.ProjectId)
			if err != nil {
				status = consts.Fail
			}

			scriptHelper.GenResultMsg(&scriptBase)
			scriptBase.VariableSettings = VariableSettings

			obj.PostConditions[index].Raw, _ = json.Marshal(scriptBase)

		} else if condition.Type == consts.ConditionTypeCheckpoint {
			var checkpointBase domain.CheckpointBase
			json.Unmarshal(condition.Raw, &checkpointBase)
			if checkpointBase.Disabled {
				continue
			}

			err = ExecCheckPoint(&checkpointBase, resp, 0)
			if err != nil {
				status = consts.Fail
			}

			checkpointHelper.GenResultMsg(&checkpointBase)

			obj.PostConditions[index].Raw, _ = json.Marshal(checkpointBase)

		} else if condition.Type == consts.ConditionTypeDatabase {
			var databaseOptBase domain.DatabaseOptBase
			json.Unmarshal(condition.Raw, &databaseOptBase)
			if databaseOptBase.Disabled {
				continue
			}

			err = ExecDbOpt(&databaseOptBase)
			if err != nil {
				status = consts.Fail
			}

			databaseOptHelpper.GenResultMsg(&databaseOptBase)

			obj.PostConditions[index].Raw, _ = json.Marshal(databaseOptBase)

		} else if condition.Type == consts.ConditionTypeResponseDefine {
			var responseDefineBase domain.ResponseDefineBase
			json.Unmarshal(condition.Raw, &responseDefineBase)
			if responseDefineBase.Disabled {
				continue
			}

			err = ExecResponseDefine(&responseDefineBase, resp)
			if err != nil {
				status = consts.Fail
			}

			obj.PostConditions[index].Raw, _ = json.Marshal(responseDefineBase)
			//openapi.GenResultMsg(&responseDefineBase)
		}
	}

	return
}
