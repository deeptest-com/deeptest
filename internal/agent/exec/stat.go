package agentExec

import (
	"encoding/json"
	agentDomain "github.com/aaronchen2k/deeptest/internal/agent/exec/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
)

var (
	Stat = agentDomain.InterfaceStat{}
)

func ResetStat() {
	Stat = agentDomain.InterfaceStat{}
}

func CountStat(result *agentDomain.ScenarioExecResult) agentDomain.InterfaceStat {
	Stat.InterfaceCount += 1
	Stat.InterfaceDurationTotal += result.Cost
	Stat.InterfaceDurationAverage = Stat.InterfaceDurationTotal / Stat.InterfaceCount

	result.ResultStatus = consts.Pass

	for _, item := range result.PostConditions {
		if item.Type != consts.ConditionTypeCheckpoint {
			continue
		}

		var checkpointBase domain.CheckpointBase
		json.Unmarshal(item.Raw, &checkpointBase)
		if checkpointBase.Disabled {
			continue
		}

		if checkpointBase.ResultStatus == consts.Pass {
			Stat.CheckpointPass += 1
		} else if checkpointBase.ResultStatus == consts.Fail {
			Stat.CheckpointFail += 1
			result.ResultStatus = consts.Fail
		}

		if item.Type != consts.ConditionTypeResponseDefine {
			var responseDefineBase domain.ResponseDefineBase
			json.Unmarshal(item.Raw, &responseDefineBase)
			if responseDefineBase.Disabled {
				continue
			}

			// not to count
			//if responseDefineBase.ResultStatus == consts.Pass {
			//	Stat.CheckpointPass += 1
			//} else if responseDefineBase.ResultStatus == consts.Fail {
			//	Stat.CheckpointFail += 1
			//	result.ResultStatus = consts.Fail
			//}
		}
	}

	if result.ResultStatus == consts.Pass {
		Stat.InterfacePass += 1
	} else if result.ResultStatus == consts.Fail {
		Stat.InterfaceFail += 1
	}

	return Stat
}

func CountSkip(executedProcessorIds map[uint]bool, skippedChildren []*Processor) agentDomain.InterfaceStat {
	countedProcessorIds := map[uint]bool{}
	countSkipInterface(executedProcessorIds, skippedChildren, &countedProcessorIds)

	return Stat
}

func countSkipInterface(executedProcessorIds map[uint]bool, skippedChildren []*Processor, countedProcessorIds *map[uint]bool) agentDomain.InterfaceStat {
	for _, child := range skippedChildren {
		if child.Disable {
			continue
		}

		_, executed := executedProcessorIds[child.ID]
		_, counted := (*countedProcessorIds)[child.ID]
		if child.EntityType == consts.ProcessorInterfaceDefault && !executed && !counted {
			Stat.InterfaceSkip += 1
			(*countedProcessorIds)[child.ID] = true
		}

		if len(child.Children) > 0 {
			countSkipInterface(map[uint]bool{}, child.Children, countedProcessorIds)
		}
	}

	return Stat
}
