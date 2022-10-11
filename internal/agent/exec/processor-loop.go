package agentExec

import (
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/agent/exec/domain"
	"github.com/aaronchen2k/deeptest/internal/agent/exec/utils"
	"github.com/aaronchen2k/deeptest/internal/agent/exec/utils/exec"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
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

func (entity ProcessorLoop) Run(processor *Processor, session *Session) (log domain.Result, err error) {
	logUtils.Infof("loop entity")

	log = domain.Result{
		ID:                entity.ProcessorID,
		Name:              entity.Name,
		ProcessorCategory: entity.ProcessorCategory,
		ProcessorType:     entity.ProcessorType,
		ParentId:          entity.ParentID,
	}

	if entity.ProcessorType == consts.ProcessorLoopBreak {
		log.WillBreak, log.Output = entity.getBeak()
		processor.Result = log

		exec.SendExecMsg(processor.Result, session.WsMsg)
		return
	}

	log.Iterator, log.Output = entity.getIterator()

	processor.Result = log
	exec.SendExecMsg(processor.Result, session.WsMsg)

	if entity.ProcessorType == consts.ProcessorLoopUntil {
		entity.runLoopUntil(session, processor, log.Iterator)
	} else {
		entity.runLoopItems(session, processor, log.Iterator)
	}

	return
}

func (entity ProcessorLoop) getBeak() (ret bool, output string) {
	breakFrom := entity.ParentID
	breakIfExpress := entity.BreakIfExpression

	result, err := EvaluateGovaluateExpressionByScope(breakIfExpress, entity.ProcessorID)
	ret, ok := result.(bool)
	if err == nil && ok && ret {
		breakMap.Store(breakFrom, true)
		output = "真"
	} else {
		output = "假"
	}

	return
}

func (entity ProcessorLoop) getIterator() (iterator domain.ExecIterator, msg string) {
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

func (entity *ProcessorLoop) runLoopUntil(s *Session, processor *Processor, iterator domain.ExecIterator) (err error) {
	expression := iterator.UntilExpression

	for {
		result, err := EvaluateGovaluateExpressionByScope(expression, entity.ID)
		pass, ok := result.(bool)
		if err != nil || !ok || pass {
			break
		}

		for _, child := range processor.Children {
			childLog, _ := (*child).Run(s)

			if childLog.WillBreak {
				logUtils.Infof("break")
				goto LABEL
			}
		}
	}
LABEL:

	return
}

func (entity *ProcessorLoop) runLoopItems(s *Session, processor *Processor, iterator domain.ExecIterator) (err error) {
	for _, item := range iterator.Items {
		SetVariable(entity.ID, iterator.VariableName, item, consts.Local)

		for _, child := range processor.Children {
			childLog, _ := child.Run(s)

			if childLog.WillBreak {
				logUtils.Infof("break")
				goto LABEL
			}
		}
	}
LABEL:

	return
}

func (entity ProcessorLoop) GenerateLoopTimes() (ret domain.ExecIterator, err error) {
	if entity.Times > 0 {
		for i := 0; i < entity.Times; i++ {
			ret.Items = append(ret.Items, i+1)
		}
	}

	ret.DataType = consts.Int

	return
}
func (entity ProcessorLoop) GenerateLoopRange() (ret domain.ExecIterator, err error) {
	start, end, step, precision, typ, err := utils.GetRange(entity.Range, entity.Step)
	if err == nil {
		ret.DataType = typ
		ret.Items, _ = utils.GenerateRangeItems(start, end, step, precision, entity.IsRand, typ)
	}

	return
}
func (entity ProcessorLoop) GenerateLoopList() (ret domain.ExecIterator, err error) {
	ret.Items, ret.DataType, err = utils.GenerateListItems(entity.List)

	return
}
