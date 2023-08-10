package agentExec

import (
	"encoding/json"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	checkpointHelper "github.com/aaronchen2k/deeptest/internal/pkg/helper/checkpoint"
	extractorHelper "github.com/aaronchen2k/deeptest/internal/pkg/helper/extractor"
	scriptHelper "github.com/aaronchen2k/deeptest/internal/pkg/helper/script"
)

func ExecPreConditions(obj *InterfaceExecObj) (err error) {
	for index, condition := range obj.PreConditions {
		if condition.Type == consts.ConditionTypeScript {
			var scriptBase domain.ScriptBase
			json.Unmarshal(condition.Raw, &scriptBase)

			err = ExecScript(&scriptBase)
			scriptHelper.GenResultMsg(&scriptBase)
			scriptBase.VariableSettings = VariableSettings

			obj.PreConditions[index].Raw, _ = json.Marshal(scriptBase)
		}
	}

	return
}

func ExecPostConditions(obj *InterfaceExecObj, resp domain.DebugResponse) (err error) {
	for index, condition := range obj.PostConditions {
		if condition.Type == consts.ConditionTypeExtractor {
			var extractorBase domain.ExtractorBase
			json.Unmarshal(condition.Raw, &extractorBase)

			if extractorBase.Disabled || extractorBase.Variable == "" {
				continue
			}

			err = ExecExtract(&extractorBase, resp)
			extractorHelper.GenResultMsg(&extractorBase)

			SetVariable(0, extractorBase.Variable, extractorBase.Result, consts.Public)

			obj.PostConditions[index].Raw, _ = json.Marshal(extractorBase)

		} else if condition.Type == consts.ConditionTypeCheckpoint {
			var checkpointBase domain.CheckpointBase
			json.Unmarshal(condition.Raw, &checkpointBase)
			if checkpointBase.Disabled {
				continue
			}

			err = ExecCheckPoint(&checkpointBase, resp, 0)
			checkpointHelper.GenResultMsg(&checkpointBase)

			obj.PostConditions[index].Raw, _ = json.Marshal(checkpointBase)

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
		}
	}

	return
}
