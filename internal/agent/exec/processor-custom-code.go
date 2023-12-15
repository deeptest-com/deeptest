package agentExec

import (
	"github.com/aaronchen2k/deeptest/internal/agent/exec/domain"
	"github.com/aaronchen2k/deeptest/internal/agent/exec/utils/exec"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	scriptHelper "github.com/aaronchen2k/deeptest/internal/pkg/helper/script"
	commonUtils "github.com/aaronchen2k/deeptest/pkg/lib/comm"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	uuid "github.com/satori/go.uuid"
	"time"
)

type ProcessorCustomCode struct {
	ID uint `json:"id" yaml:"id"`
	ProcessorEntityBase

	Content string `gorm:"type:longtext;" json:"content" yaml:"content"`
	Desc    string `json:"desc" yaml:"desc"`
}

func (entity ProcessorCustomCode) Run(processor *Processor, session *Session) (err error) {
	defer func() {
		if errX := recover(); errX != nil {
			processor.Error(session, errX)
		}
	}()
	logUtils.Infof("print entity")

	startTime := time.Now()
	processor.Result = &agentDomain.ScenarioExecResult{
		ID:                int(entity.ProcessorID),
		Name:              entity.Name,
		ProcessorCategory: entity.ProcessorCategory,
		ProcessorType:     entity.ProcessorType,
		StartTime:         &startTime,
		ParentId:          int(entity.ParentID),
		ScenarioId:        processor.ScenarioId,
		ProcessorId:       processor.ID,
		LogId:             uuid.NewV4(),
		ParentLogId:       processor.Parent.Result.LogId,
		Round:             processor.Round,
	}

	scriptBase := domain.ScriptBase{
		Content: entity.Content,
	}

	err = ExecScript(&scriptBase, processor.ProjectId, session.ExecUuid)
	scriptHelper.GenResultMsg(&scriptBase)
	//scriptBase.VariableSettings = VariableSettings

	for _, item := range GetGojaVariables(session.ExecUuid) {
		SetVariable(processor.ParentId, item.Name, item.Value, consts.ExtractorResultTypeObject, consts.Public, session.ExecUuid)
	}

	processor.Result.Summary = scriptBase.ResultStatus.String()

	processor.AddResultToParent()
	result := false
	if scriptBase.ResultStatus == consts.Pass {
		result = true
	}
	detail := map[string]interface{}{"name": entity.Name, "content": entity.Content, "result": result, "output": scriptBase.Output}
	processor.Result.Detail = commonUtils.JsonEncode(detail)
	execUtils.SendExecMsg(*processor.Result, consts.Processor, session.WsMsg)

	endTime := time.Now()
	processor.Result.EndTime = &endTime

	return
}
