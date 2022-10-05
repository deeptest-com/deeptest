package agentExec

import (
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
)

type ProcessorLoop struct {
	ID uint `json:"id" yaml:"id"`
	ProcessorEntity

	Times        int    `json:"times" yaml:"times"` // time
	Range        string `json:"range" yaml:"range"` // range
	List         string `json:"list" yaml:"list"`   // in
	Step         string `json:"step" yaml:"step"`
	IsRand       bool   `json:"isRand" yaml:"isRand"`
	VariableName string `json:"variableName" yaml:"variableName"`

	UntilExpression   string `json:"untilExpression" yaml:"untilExpression"` // until
	BreakIfExpression string `json:"breakIfExpression" yaml:"breakIfExpression"`
}

func (p ProcessorLoop) Run(s *Session) (log Log, variableName string, variableValues []interface{}, err error) {
	logUtils.Infof("loop")

	log = Log{
		Name:   p.Name,
		Output: p.getMsg(),
	}

	return
}

func (p ProcessorLoop) getMsg() (msg string) {
	if p.ID == 0 {
		msg = "执行前请先配置处理器。"
		return
	}

	typ := p.ProcessorType
	if typ == consts.ProcessorLoopTime {
		msg = fmt.Sprintf("执行\"%d\"次。", p.Times)
		return
	} else if typ == consts.ProcessorLoopUntil {
		msg = fmt.Sprintf("直到\"%s\"。", p.UntilExpression)
		return
	} else if typ == consts.ProcessorLoopIn {
		msg = fmt.Sprintf("迭代列表\"%s\"。", p.List)
		return
	} else if typ == consts.ProcessorLoopRange {
		msg = fmt.Sprintf("区间\"%s\"。", p.Range)
		return
	}

	return
}
