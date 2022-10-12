package agentExec

import (
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/agent/exec/domain"
	"github.com/aaronchen2k/deeptest/internal/agent/exec/utils/exec"
	valueGen "github.com/aaronchen2k/deeptest/internal/agent/exec/utils/value"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
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

func (entity ProcessorCookie) Run(processor *Processor, session *Session) (log domain.Result, err error) {
	processor.Result = domain.Result{
		ID:                entity.ProcessorID,
		Name:              entity.Name,
		ProcessorCategory: entity.ProcessorCategory,
		ProcessorType:     entity.ProcessorType,
		ParentId:          entity.ParentID,
	}

	cookieName := entity.CookieName
	variableName := entity.VariableName
	defaultValue := entity.Default
	domain := entity.Domain
	expireTime := entity.ExpireTime
	expression := entity.RightValue
	typ := entity.ProcessorType

	if typ == consts.ProcessorCookieSet {
		var variableValue interface{}
		variableValue, err = EvaluateGovaluateExpressionByScope(expression, entity.ProcessorID)
		if err != nil {
			processor.Result.Summary = fmt.Sprintf("计算表达式\"%s\"错误 %s。", expression, err.Error())
			exec.SendExecMsg(processor.Result, session.WsMsg)
			return
		}

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
			exec.SendExecMsg(processor.Result, session.WsMsg)
			return
		}

		SetVariable(processor.ParentId, variableName, variableValue, consts.Local) // set in parent scope
		processor.Result.Summary = fmt.Sprintf("将%s%s值%v赋予变量%s。", cookieName, words, variableValue, variableName)

	} else if typ == consts.ProcessorCookieClear {
		ClearCookie(processor.ParentId, cookieName) // set in parent scope
		processor.Result.Summary = fmt.Sprintf("%s。", cookieName)
	}

	exec.SendExecMsg(processor.Result, session.WsMsg)

	return
}
