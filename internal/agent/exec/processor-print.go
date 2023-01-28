package agentExec

import (
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/agent/exec/domain"
	"github.com/aaronchen2k/deeptest/internal/agent/exec/utils/exec"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"strings"
	"time"
)

type ProcessorPrint struct {
	ID uint `json:"id" yaml:"id"`
	ProcessorEntityBase

	Expression string `json:"expression" yaml:"expression"`
}

func (entity ProcessorPrint) Run(processor *Processor, session *Session) (err error) {
	logUtils.Infof("print entity")

	startTime := time.Now()
	processor.Result = &domain.Result{
		ID:                int(entity.ProcessorID),
		Name:              entity.Name,
		ProcessorCategory: entity.ProcessorCategory,
		ProcessorType:     entity.ProcessorType,
		StartTime:         &startTime,
		ParentId:          int(entity.ParentID),
	}

	variableMap := GetCachedVariableMapInContext(processor.ID)
	value := ReplaceVariableValue(entity.Expression, variableMap, Datapools)

	processor.Result.Summary = strings.ReplaceAll(fmt.Sprintf("%s为\"%v\"。",
		entity.Expression, value), "<nil>", "空")

	processor.AddResultToParent()
	exec.SendExecMsg(*processor.Result, session.WsMsg)

	endTime := time.Now()
	processor.Result.EndTime = &endTime

	return
}
