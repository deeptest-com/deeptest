package agentExec

import (
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
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

func (p ProcessorLoop) Run(s *Session) (log Log, err error) {
	logUtils.Infof("loop entity")

	log.Name = p.Name

	if p.ProcessorType == consts.ProcessorLoopBreak {
		log.WillBreak, log.Output = getBeak(p)
	} else {
		log.Iterator, log.Output = getIterator(p)
	}

	return
}

func getBeak(loop ProcessorLoop) (ret bool, output string) {
	breakFrom := loop.ParentID
	breakIfExpress := loop.BreakIfExpression

	result, err := EvaluateGovaluateExpression(breakIfExpress, loop.ProcessorID)
	ret, ok := result.(bool)
	if err == nil && ok && ret {
		breakMap.Store(breakFrom, true)
		output = "真"
	} else {
		output = "假"
	}

	return
}

func getIterator(loop ProcessorLoop) (iterator domain.ExecIterator, msg string) {
	if loop.ID == 0 {
		msg = "执行前请先配置处理器。"
		return
	}

	if loop.ProcessorType == consts.ProcessorLoopTime {
		iterator, _ = GenerateLoopTimes(loop)
		msg = fmt.Sprintf("迭代\"%d\"次。", loop.Times)
	} else if loop.ProcessorType == consts.ProcessorLoopIn {
		iterator, _ = GenerateLoopList(loop)
		msg = fmt.Sprintf("迭代列表\"%s\"。", loop.List)
	} else if loop.ProcessorType == consts.ProcessorLoopRange {
		iterator, _ = GenerateLoopRange(loop)
		msg = fmt.Sprintf("迭代区间\"%s\"。", loop.Range)
	} else if loop.ProcessorType == consts.ProcessorLoopUntil {
		iterator.UntilExpression = loop.UntilExpression
		msg = fmt.Sprintf("迭代直到\"%s\"。", loop.UntilExpression)
	}

	iterator.VariableName = loop.VariableName

	return
}

func GenerateLoopTimes(loop ProcessorLoop) (ret domain.ExecIterator, err error) {
	if loop.Times > 0 {
		for i := 0; i < loop.Times; i++ {
			ret.Items = append(ret.Items, i+1)
		}
	}

	ret.DataType = consts.Int

	return
}
func GenerateLoopRange(loop ProcessorLoop) (ret domain.ExecIterator, err error) {
	start, end, step, precision, typ, err := GetRange(loop.Range, loop.Step)
	if err == nil {
		ret.DataType = typ
		ret.Items, _ = GenerateRangeItems(start, end, step, precision, loop.IsRand, typ)
	}

	return
}
func GenerateLoopList(loop ProcessorLoop) (ret domain.ExecIterator, err error) {
	ret.Items, ret.DataType, err = GenerateListItems(loop.List)

	return
}

//else if typ == consts.ProcessorLoopBreak {
//output.Expression = loop.BreakIfExpression
//output.BreakFrom = parentLog.ProcessId
//
//output, _ = EvaluateLoopBreak(&loop, parentLog, msg)
//
//breakFrom := output.BreakFrom
//breakIfExpress := output.Expression
//
//result, err := EvaluateGovaluateExpression(breakIfExpress, p.ID)
//pass, ok := result.(bool)
//if err == nil && ok && pass {
//breakMap.Store(breakFrom, true)
//ret = "真"
//} else {
//ret = "假"
//}
//
//return
//}

func IsLoopTimesPass(loop ProcessorLoop, output domain.ExecOutput) bool {
	return loop.ProcessorType == consts.ProcessorLoopTime && output.Times > 0
}
func IsLoopUntilPass(loop ProcessorLoop, output domain.ExecOutput) bool {
	return loop.ProcessorType == consts.ProcessorLoopUntil && output.Expression != ""
}
func IsLoopInPass(loop ProcessorLoop, output domain.ExecOutput) bool {
	return loop.ProcessorType == consts.ProcessorLoopIn && output.List != ""
}
func IsLoopRangePass(loop ProcessorLoop, output domain.ExecOutput) bool {
	return loop.ProcessorType == consts.ProcessorLoopRange && output.Range != ""
}
func IsLoopLoopBreak(loop ProcessorLoop, output domain.ExecOutput) bool {
	return loop.ProcessorType == consts.ProcessorLoopBreak
}
