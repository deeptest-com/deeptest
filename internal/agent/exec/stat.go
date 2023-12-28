package agentExec

import (
	"encoding/json"
	agentDomain "github.com/aaronchen2k/deeptest/internal/agent/exec/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
)

func ResetStat(execUuid string) {
	SetStat(execUuid, &agentDomain.InterfaceStat{})
}

func CountStat(execUuid string, result *agentDomain.ScenarioExecResult) agentDomain.InterfaceStat {
	stat := GetStat(execUuid)

	stat.InterfaceCount += 1
	stat.InterfaceDurationTotal += result.Cost
	stat.InterfaceDurationAverage = stat.InterfaceDurationTotal / stat.InterfaceCount

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
			stat.CheckpointPass += 1
		} else if checkpointBase.ResultStatus == consts.Fail {
			stat.CheckpointFail += 1
			result.ResultStatus = consts.Fail
		}

		if item.Type != consts.ConditionTypeResponseDefine {
			var responseDefineBase domain.ResponseDefineBase
			json.Unmarshal(item.Raw, &responseDefineBase)
			if responseDefineBase.Disabled {
				continue
			}
		}
	}

	if result.ResultStatus == consts.Pass {
		stat.InterfacePass += 1
	} else if result.ResultStatus == consts.Fail {
		stat.InterfaceFail += 1
	}

	SetStat(execUuid, stat)

	return *stat
}

func CountSkip(execUuid string, executedProcessorIds map[uint]bool, skippedChildren []*Processor) agentDomain.InterfaceStat {
	countedProcessorIds := map[uint]bool{}
	countSkipInterface(execUuid, executedProcessorIds, skippedChildren, &countedProcessorIds)

	return *GetStat(execUuid)
}

func countSkipInterface(execUuid string, executedProcessorIds map[uint]bool, skippedChildren []*Processor, countedProcessorIds *map[uint]bool) agentDomain.InterfaceStat {
	stat := GetStat(execUuid)

	for _, child := range skippedChildren {
		if child.Disable {
			continue
		}

		_, executed := executedProcessorIds[child.ID]
		_, counted := (*countedProcessorIds)[child.ID]
		if child.EntityType == consts.ProcessorInterfaceDefault && !executed && !counted {
			stat.InterfaceSkip += 1
			(*countedProcessorIds)[child.ID] = true
		}

		if len(child.Children) > 0 {
			countSkipInterface(execUuid, map[uint]bool{}, child.Children, countedProcessorIds)
		}
	}

	return *stat
}
