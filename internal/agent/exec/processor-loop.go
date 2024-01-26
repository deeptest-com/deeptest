package agentExec

import (
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/agent/exec/domain"
	agentUtils "github.com/aaronchen2k/deeptest/internal/agent/exec/utils"
	"github.com/aaronchen2k/deeptest/internal/agent/exec/utils/exec"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	commUtils "github.com/aaronchen2k/deeptest/internal/pkg/utils"
	commonUtils "github.com/aaronchen2k/deeptest/pkg/lib/comm"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	uuid "github.com/satori/go.uuid"
	"reflect"
	"strings"
	"time"
)

type ProcessorLoop struct {
	ID uint `json:"id" yaml:"id"`
	ProcessorEntityBase

	Times int    `json:"times" yaml:"times"` // time
	Range string `json:"range" yaml:"range"` // range

	InType   string `json:"inType" yaml:"inType"`     // in
	Variable string `json:"variable" yaml:"variable"` // array
	List     string `json:"list" yaml:"list"`         // list

	Step         string `json:"step" yaml:"step"`
	IsRand       bool   `json:"isRand" yaml:"isRand"`
	VariableName string `json:"variableName" yaml:"variableName"`

	UntilExpression   string `json:"untilExpression" yaml:"untilExpression"` // until
	BreakIfExpression string `json:"breakIfExpression" yaml:"breakIfExpression"`
}

func (entity ProcessorLoop) Run(processor *Processor, session *Session) (err error) {
	defer func() {
		if errX := recover(); errX != nil {
			processor.Error(session, errX)
		}
	}()
	logUtils.Infof("loop entity")

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

	processor.Result.Detail = commonUtils.JsonEncode(entity)
	execUtils.SendExecMsg(*processor.Result, consts.Processor, session.WsMsg)

	processor.Result.Iterator, processor.Result.Summary = entity.getIterator(session)

	if entity.ProcessorType == consts.ProcessorLoopUntil {
		entity.runLoopUntil(session, processor, processor.Result.Iterator)
	} else {
		entity.runLoopItems(session, processor, processor.Result.Iterator)
	}

	processor.AddResultToParent()

	endTime := time.Now()
	processor.Result.EndTime = &endTime

	return
}

func (entity *ProcessorLoop) runLoopItems(session *Session, processor *Processor, iterator agentDomain.ExecIterator) (err error) {
	executedProcessorIds := map[uint]bool{}

	for index, item := range iterator.Items {
		if GetForceStopExec(session.ExecUuid) {
			break
		}
		if DemoTestSite != "" && index > 2 {
			break
		}
		/*
			msg := agentDomain.ScenarioExecResult{
				ParentId:          int(processor.ID),
				Summary:           fmt.Sprintf("%d. %s为%v", index+1, iterator.VariableName, item),
				Name:              "循环变量",
				ProcessorCategory: consts.ProcessorPrint,
			}
			execUtils.SendExecMsg(msg, session.WsMsg)
		*/

		SetVariable(entity.ProcessorID, iterator.VariableName, item, consts.ExtractorResultTypeString, consts.Public, session.ExecUuid)

		round := ""
		for _, child := range processor.Children {
			if GetForceStopExec(session.ExecUuid) {
				break
			}
			if child.Disable {
				continue
			}

			executedProcessorIds[child.ID] = true

			//执行轮次
			if round == "" {
				if entity.ProcessorType == consts.ProcessorLoopTime {
					round = fmt.Sprintf("第 %v 轮", index+1)
				} else {
					desc, _ := commUtils.ConvertValueForPersistence(item)
					round = fmt.Sprintf("第 %v 轮，%v = %v", index+1, iterator.VariableName, desc)
				}

				child.Round = round
			}

			(*child).Run(session)
		}

		// check break
		result := agentDomain.ScenarioExecResult{}
		result.WillBreak, result.Summary, result.Detail = entity.getBeak(session.ExecUuid)
		if result.WillBreak {
			execUtils.SendExecMsg(result, consts.Processor, session.WsMsg)
			break
		}
	}

	stat := CountSkip(session.ExecUuid, executedProcessorIds, processor.Children)
	execUtils.SendStatMsg(stat, session.WsMsg)

	return
}

func (entity *ProcessorLoop) runLoopUntil(session *Session, processor *Processor, iterator agentDomain.ExecIterator) (err error) {
	expression := iterator.UntilExpression

	executedProcessorIds := map[uint]bool{}
	index := 0
	for {
		if GetForceStopExec(session.ExecUuid) {
			break
		}
		if DemoTestSite != "" && index > 2 {
			break
		}
		index += 1

		result, _, err := EvaluateGovaluateExpressionByProcessorScope(expression, entity.ProcessorID, session.ExecUuid)
		pass, ok := result.(bool)
		if err != nil || !ok || pass {
			result := agentDomain.ScenarioExecResult{
				WillBreak: true,
				Summary:   fmt.Sprintf("条件%s满足，跳出循环。", expression),
			}
			execUtils.SendExecMsg(result, consts.Processor, session.WsMsg)

			break
		}

		round := ""
		for _, child := range processor.Children {
			if GetForceStopExec(session.ExecUuid) {
				break
			}
			if child.Disable {
				continue
			}

			executedProcessorIds[child.ID] = true

			if round == "" {
				round = fmt.Sprintf("第 %v 轮", index)
				child.Round = round
			}

			(*child).Run(session)

			if child.Result.WillBreak {
				logUtils.Infof("break")
				goto LABEL
			}
		}

		if index >= consts.MaxLoopTimeForInterfaceTest {
			logUtils.Infof("break for reach MaxLoopTimeForInterfaceTest")
			panic(fmt.Sprintf("循环执行次数达到上限%d次", consts.MaxLoopTimeForInterfaceTest))
			//goto LABEL
		}
	}

LABEL:
	stat := CountSkip(session.ExecUuid, executedProcessorIds, processor.Children)
	execUtils.SendStatMsg(stat, session.WsMsg)

	return
}

func (entity *ProcessorLoop) getBeak(execUuid string) (ret bool, msg string, detailStr string) {
	breakIfExpress := strings.TrimSpace(entity.BreakIfExpression)

	if breakIfExpress == "" {
		return
	}

	expr := ReplaceDatapoolVariInGovaluateExpress(breakIfExpress, execUuid)
	result, _, _ := EvaluateGovaluateExpressionByProcessorScope(expr, entity.ProcessorID, execUuid)

	ret, ok := result.(bool)
	pass := false
	if ok && ret {
		msg = "真"
		pass = true
	} else {
		msg = "假"
	}

	detail := map[string]interface{}{"expression": breakIfExpress, "result": pass}
	detailStr = commonUtils.JsonEncode(detail)

	return
}

func (entity *ProcessorLoop) getIterator(session *Session) (iterator agentDomain.ExecIterator, msg string) {
	if entity.ID == 0 {
		msg = "执行前请先配置处理器。"
		return
	}

	if entity.ProcessorType == consts.ProcessorLoopTime {
		iterator, _ = entity.GenerateLoopTimes()
		msg = fmt.Sprintf("迭代\"%d\"次。", entity.Times)

	} else if entity.ProcessorType == consts.ProcessorLoopIn {
		if entity.InType == "variable" {
			iterator, _ = entity.GenerateLoopVariable(session.ExecUuid)
			msg = fmt.Sprintf("\"%s\"。", entity.Variable)

		} else if entity.InType == "list" {
			iterator, _ = entity.GenerateLoopList()
			msg = fmt.Sprintf("\"%s\"。", entity.List)
		}

	} else if entity.ProcessorType == consts.ProcessorLoopRange {
		iterator, _ = entity.GenerateLoopRange()
		msg = fmt.Sprintf("\"%s\"。", entity.Range)

	} else if entity.ProcessorType == consts.ProcessorLoopUntil {
		iterator.UntilExpression = entity.UntilExpression
		msg = fmt.Sprintf("\"%s\"。", entity.UntilExpression)
	}

	iterator.VariableName = entity.VariableName

	return
}

func (entity *ProcessorLoop) GenerateLoopTimes() (ret agentDomain.ExecIterator, err error) {
	if entity.Times > 0 {
		for i := 0; i < entity.Times; i++ {
			ret.Items = append(ret.Items, i+1)
		}
	}

	ret.DataType = consts.Int

	return
}
func (entity *ProcessorLoop) GenerateLoopRange() (ret agentDomain.ExecIterator, err error) {
	start, end, step, precision, typ, err := agentUtils.GetRange(entity.Range, entity.Step)
	if err == nil {
		ret.DataType = typ
		ret.Items, _ = agentUtils.GenerateRangeItems(start, end, step, precision, entity.IsRand, typ)
	}

	return
}
func (entity *ProcessorLoop) GenerateLoopVariable(execUuid string) (ret agentDomain.ExecIterator, err error) {
	variableObj, err := GetVariable(entity.ProcessorID, entity.Variable, execUuid)
	if err != nil {
		return
	}

	if variableObj.ValueType == consts.ExtractorResultTypeObject {
		val, err1 := commUtils.ConvertValueForUse(variableObj.Value, variableObj.ValueType)
		if err1 != nil || val == nil {
			return
		}

		typ := reflect.TypeOf(val)
		if typ.Kind() == reflect.Array || typ.Kind() == reflect.Slice {
			for _, item := range val.([]interface{}) {
				ret.Items = append(ret.Items, item)
			}
		}
	}

	return
}
func (entity *ProcessorLoop) GenerateLoopList() (ret agentDomain.ExecIterator, err error) {
	ret.Items, ret.DataType, err = agentUtils.GenerateListItems(entity.List, entity.IsRand)

	return
}
