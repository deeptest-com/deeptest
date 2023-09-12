package agentExec

import (
	"encoding/json"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	checkpointHelper "github.com/aaronchen2k/deeptest/internal/pkg/helper/checkpoint"
	extractorHelper "github.com/aaronchen2k/deeptest/internal/pkg/helper/extractor"
	scriptHelper "github.com/aaronchen2k/deeptest/internal/pkg/helper/script"
)

func ExecPreConditions(execObj InterfaceExecObj) (err error) {
	for index, condition := range execObj.PreConditions {
		if condition.Type == consts.ConditionTypeScript {
			var scriptBase domain.ScriptBase
			json.Unmarshal(condition.Raw, &scriptBase)

			err = ExecScript(&scriptBase)
			scriptHelper.GenResultMsg(&scriptBase)
			scriptBase.VariableSettings = VariableSettings

			execObj.PreConditions[index].Raw, _ = json.Marshal(scriptBase)
		}
	}

	return
}

func ExecPostConditions(obj InterfaceExecObj, resp domain.DebugResponse) (err error) {
	for index, condition := range obj.PostConditions {
		if condition.Type == consts.ConditionTypeExtractor {
			var extractorBase domain.ExtractorBase
			json.Unmarshal(condition.Raw, &extractorBase)

			if extractorBase.Disabled || extractorBase.Variable == "" {
				continue
			}

			err = ExecExtract(&extractorBase, resp)
			extractorHelper.GenResultMsg(&extractorBase)

			if extractorBase.ResultStatus == consts.Pass {
				SetVariable(0, extractorBase.Variable, extractorBase.Result, consts.Public)
			}

			obj.PostConditions[index].Raw, _ = json.Marshal(extractorBase)

		} else if condition.Type == consts.ConditionTypeScript {
			var scriptBase domain.ScriptBase
			json.Unmarshal(condition.Raw, &scriptBase)
			if scriptBase.Disabled {
				continue
			}

			err = ExecScript(&scriptBase)
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
			checkpointHelper.GenResultMsg(&checkpointBase)

			obj.PostConditions[index].Raw, _ = json.Marshal(checkpointBase)
		} else if condition.Type == consts.ConditionTypeResponseDefine {
			var responseDefineBase domain.ResponseDefineBase
			json.Unmarshal(condition.Raw, &responseDefineBase)
			if responseDefineBase.Disabled {
				continue
			}

			err = ExecResponseDefine(&responseDefineBase, resp)
			obj.PostConditions[index].Raw, _ = json.Marshal(responseDefineBase)
			//openapi.GenResultMsg(&responseDefineBase)
		}
	}

	return
}
