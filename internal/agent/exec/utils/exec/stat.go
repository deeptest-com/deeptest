package execUtils

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
	}

	if result.ResultStatus == consts.Pass {
		Stat.InterfacePass += 1
	} else if result.ResultStatus == consts.Fail {
		Stat.InterfaceFail += 1
	} else if result.ResultStatus == consts.Skip {
		Stat.InterfaceSkip += 1
	}

	return Stat
}
