package execUtils

import (
	"encoding/json"
	agentDomain "github.com/aaronchen2k/deeptest/internal/agent/exec/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
)

var (
	stat = agentDomain.InterfaceStat{}
)

func InitStat() {
	stat = agentDomain.InterfaceStat{}
}

func CountStat(result *agentDomain.ScenarioExecResult) agentDomain.InterfaceStat {
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
	}

	if result.ResultStatus == consts.Pass {
		stat.InterfacePass += 1
	} else if result.ResultStatus == consts.Fail {
		stat.InterfaceFail += 1
	} else if result.ResultStatus == consts.Skip {
		stat.InterfaceSkip += 1
	}

	return stat
}
