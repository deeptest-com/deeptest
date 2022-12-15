package agentExec

import (
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/agent/exec/domain"
	"github.com/aaronchen2k/deeptest/internal/agent/exec/utils/exec"
	valueGen "github.com/aaronchen2k/deeptest/internal/agent/exec/utils/value"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
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
	logUtils.Infof("cookie entity")

	startTime := time.Now()
	processor.Result = &domain.Result{
		ID:                int(entity.ProcessorID),
		Name:              entity.Name,
		ProcessorCategory: entity.ProcessorCategory,
		ProcessorType:     entity.ProcessorType,
		StartTime:         &startTime,
		ParentId:          int(entity.ParentID),
	}

	cookieName := entity.CookieName
	variableName := entity.VariableName
	defaultValue := entity.Default
	domain := entity.Domain
	expireTime := entity.ExpireTime
	rightValue := entity.RightValue
	typ := entity.ProcessorType

	if typ == consts.ProcessorCookieSet {

		variableMap := GetVariableMap(processor.ID)
		variableValue := ReplaceVariableValue(rightValue, variableMap)

		SetCookie(processor.ParentId, cookieName, variableValue, domain, expireTime) // set in parent scope

		processor.Result.Summary = fmt.Sprintf("%s为%v。", cookieName, variableValue)

	} else if typ == consts.ProcessorCookieGet {
		var variableValue interface{}
		cookie := GetCookie(processor.ParentId, cookieName, domain)
		variableValue = cookie.Value

		words := ""
		if variableValue == nil && defaultValue != "" {
			variableValue, _ = valueGen.ParseValue(defaultValue)
			words = "默认"
		}

		if err != nil {
			processor.Result.Summary = fmt.Sprintf("获取Cookie %s的值错误 %s。", cookieName, err.Error())
			processor.AddResultToParent()
			exec.SendExecMsg(*processor.Result, session.WsMsg)
			return
		}

		if variableValue == nil {
			variableValue = "空"
		}

		SetVariable(processor.ParentId, variableName, variableValue, consts.Local) // set in parent scope
		processor.Result.Summary = fmt.Sprintf("将%s%s值\"%v\"赋予变量%s。", cookieName, words, variableValue, variableName)

	} else if typ == consts.ProcessorCookieClear {
		ClearCookie(processor.ParentId, cookieName) // set in parent scope
		processor.Result.Summary = fmt.Sprintf("%s。", cookieName)
	}

	processor.AddResultToParent()
	exec.SendExecMsg(*processor.Result, session.WsMsg)

	endTime := time.Now()
	processor.Result.EndTime = &endTime

	return
}
