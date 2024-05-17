package agentExec

import (
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/agent/exec/domain"
	"github.com/aaronchen2k/deeptest/internal/agent/exec/utils/exec"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	commonUtils "github.com/aaronchen2k/deeptest/pkg/lib/comm"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	uuid "github.com/satori/go.uuid"
	"time"
)

type ProcessorCookie struct {
	ID uint `json:"id" yaml:"id"`
	ProcessorEntityBase

	CookieName   string     `json:"cookieName" yaml:"cookieName"`
	VariableName string     `json:"variableName" yaml:"variableName"`
	RightValue   string     `json:"rightValue" yaml:"rightValue"`
	Domain       string     `json:"domain" yaml:"domain"`
	ExpireTime   *time.Time `json:"expireTime" yaml:"expireTime"`

	Children []interface{} `json:"children" yaml:"children" gorm:"-"`
}

func (entity ProcessorCookie) Run(processor *Processor, session *Session) (err error) {
	defer func() {
		if errX := recover(); errX != nil {
			processor.Error(session, errX)
		}
	}()
	logUtils.Infof("cookie entity")

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

	cookieName := entity.CookieName
	domain := entity.Domain
	expireTime := entity.ExpireTime
	rightValue := entity.RightValue
	typ := entity.ProcessorType

	detail := map[string]interface{}{"name": entity.Name, "cookieName": cookieName}
	if typ == consts.ProcessorCookieSet {
		variableValue := ReplaceVariableValue(rightValue, session.TenantId, session.ProjectId, session.ExecUuid)

		SetCookie(processor.ParentId, cookieName, variableValue, domain, expireTime, session.ExecUuid) // set in parent scope

		processor.Result.Summary = fmt.Sprintf("%s为%v。", cookieName, variableValue)
		detail["variableValue"] = variableValue
		processor.Result.Detail = commonUtils.JsonEncode(detail)

	} else if typ == consts.ProcessorCookieClear {
		ClearCookie(processor.ParentId, cookieName, session.ExecUuid) // set in parent scope
		processor.Result.Summary = fmt.Sprintf("%s。", cookieName)
		processor.Result.Detail = commonUtils.JsonEncode(detail)
	}

	processor.AddResultToParent()
	execUtils.SendExecMsg(*processor.Result, consts.Processor, session.WsMsg)

	endTime := time.Now()
	processor.Result.EndTime = &endTime

	return
}
