package agentExec

import (
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/agent/exec/domain"
	"github.com/aaronchen2k/deeptest/internal/agent/exec/utils"
	"github.com/aaronchen2k/deeptest/internal/agent/exec/utils/exec"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"time"
)

type ProcessorLoop struct {
	ID uint `json:"id" yaml:"id"`
	ProcessorEntityBase

	Times        int    `json:"times" yaml:"times"` // time
	Range        string `json:"range" yaml:"range"` // range
	List         string `json:"list" yaml:"list"`   // in
	Step         string `json:"step" yaml:"step"`
	IsRand       bool   `json:"isRand" yaml:"isRand"`
	VariableName string `json:"variableName" yaml:"variableName"`

	UntilExpression   string `json:"untilExpression" yaml:"untilExpression"` // until
	BreakIfExpression string `json:"breakIfExpression" yaml:"breakIfExpression"`
}

func (entity *ProcessorLoop) Run(processor *Processor, session *Session) (err error) {
	logUtils.Infof("loop entity")

	startTime := time.Now()
	processor.Result = &domain.Result{
		ID:                int(entity.ProcessorID),
		Name:              entity.Name,
		ProcessorCategory: entity.ProcessorCategory,
		ProcessorType:     entity.ProcessorType,
		StartTime:         &startTime,
		ParentId:          int(entity.ParentID),
	}

	if entity.ProcessorType == consts.ProcessorLoopBreak {
		processor.Result.WillBreak, processor.Result.Summary = entity.getBeak()

		processor.AddResultToParent()
		exec.SendExecMsg(*processor.Result, session.WsMsg)

		return
	}

	processor.Result.Iterator, processor.Result.Summary = entity.getIterator()

	processor.AddResultToParent()
	exec.SendExecMsg(*processor.Result, session.WsMsg)

	if entity.ProcessorType == consts.ProcessorLoopUntil {
		entity.runLoopUntil(session, processor, processor.Result.Iterator)
	} else {
		entity.runLoopItems(session, processor, processor.Result.Iterator)
	}

	endTime := time.Now()
	processor.Result.EndTime = &endTime

	return
}

func (entity *ProcessorLoop) runLoopItems(s *Session, processor *Processor, iterator domain.ExecIterator) (err error) {
	for _, item := range iterator.Items {
		SetVariable(entity.ProcessorID, iterator.VariableName, item, consts.Local)

		for _, child := range processor.Children {
			child.Run(s)

			if child.Result.WillBreak {
				logUtils.Infof("break")
				goto LABEL
			}
		}
	}
LABEL:

	return
}

func (entity *ProcessorLoop) runLoopUntil(s *Session, processor *Processor, iterator domain.ExecIterator) (err error) {
	expression := iterator.UntilExpression

	for {
		result, err := EvaluateGovaluateExpressionByScope(expression, entity.ID)
		pass, ok := result.(bool)
		if err != nil || !ok || pass {
			break
		}

		for _, child := range processor.Children {
			(*child).Run(s)

			if child.Result.WillBreak {
				logUtils.Infof("break")
				goto LABEL
			}
		}
	}
LABEL:

	return
}

func (entity *ProcessorLoop) getBeak() (ret bool, msg string) {
	breakFrom := entity.ParentID
	breakIfExpress := entity.BreakIfExpression

	result, err := EvaluateGovaluateExpressionByScope(breakIfExpress, entity.ProcessorID)
	ret, ok := result.(bool)
	if err == nil && ok && ret {
		breakMap.Store(breakFrom, true)
		msg = "真"
	} else {
		msg = "假"
	}

	return
}

func (entity *ProcessorLoop) getIterator() (iterator domain.ExecIterator, msg string) {
	if entity.ID == 0 {
		msg = "执行前请先配置处理器。"
		return
	}

	if entity.ProcessorType == consts.ProcessorLoopTime {
		iterator, _ = entity.GenerateLoopTimes()
		msg = fmt.Sprintf("迭代\"%d\"次。", entity.Times)
	} else if entity.ProcessorType == consts.ProcessorLoopIn {
		iterator, _ = entity.GenerateLoopList()
		msg = fmt.Sprintf("迭代列表\"%s\"。", entity.List)
	} else if entity.ProcessorType == consts.ProcessorLoopRange {
		iterator, _ = entity.GenerateLoopRange()
		msg = fmt.Sprintf("迭代区间\"%s\"。", entity.Range)
	} else if entity.ProcessorType == consts.ProcessorLoopUntil {
		iterator.UntilExpression = entity.UntilExpression
		msg = fmt.Sprintf("迭代直到\"%s\"。", entity.UntilExpression)
	}

	iterator.VariableName = entity.VariableName

	return
}

func (entity *ProcessorLoop) GenerateLoopTimes() (ret domain.ExecIterator, err error) {
	if entity.Times > 0 {
		for i := 0; i < entity.Times; i++ {
			ret.Items = append(ret.Items, i+1)
		}
	}

	ret.DataType = consts.Int

	return
}
func (entity *ProcessorLoop) GenerateLoopRange() (ret domain.ExecIterator, err error) {
	start, end, step, precision, typ, err := utils.GetRange(entity.Range, entity.Step)
	if err == nil {
		ret.DataType = typ
		ret.Items, _ = utils.GenerateRangeItems(start, end, step, precision, entity.IsRand, typ)
	}

	return
}
func (entity *ProcessorLoop) GenerateLoopList() (ret domain.ExecIterator, err error) {
	ret.Items, ret.DataType, err = utils.GenerateListItems(entity.List)

	return
}
