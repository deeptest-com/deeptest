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

			err = ExecExtract(&extractorBase, resp)
			extractorHelper.GenResultMsg(&extractorBase)

			obj.PostConditions[index].Raw, _ = json.Marshal(extractorBase)

		} else if condition.Type == consts.ConditionTypeCheckpoint {
			var checkpointBase domain.CheckpointBase
			json.Unmarshal(condition.Raw, &checkpointBase)

			err = ExecCheckPoint(&checkpointBase, resp, 0)
			checkpointHelper.GenResultMsg(&checkpointBase)

			obj.PostConditions[index].Raw, _ = json.Marshal(checkpointBase)

		} else if condition.Type == consts.ConditionTypeScript {
			var scriptBase domain.ScriptBase
			json.Unmarshal(condition.Raw, &scriptBase)

			err = ExecScript(&scriptBase)
			scriptHelper.GenResultMsg(&scriptBase)

			obj.PostConditions[index].Raw, _ = json.Marshal(scriptBase)
		}
	}

	return
}
