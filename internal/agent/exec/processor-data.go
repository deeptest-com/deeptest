package agentExec

import (
	"errors"
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/agent/exec/domain"
	"github.com/aaronchen2k/deeptest/internal/agent/exec/utils/exec"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	commUtils "github.com/aaronchen2k/deeptest/internal/pkg/utils"
	commonUtils "github.com/aaronchen2k/deeptest/pkg/lib/comm"
	_intUtils "github.com/aaronchen2k/deeptest/pkg/lib/int"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	uuid "github.com/satori/go.uuid"
	"time"
)

type ProcessorData struct {
	ID uint `json:"id" yaml:"id"`
	ProcessorEntityBase

	Src          consts.DataItSrc  `json:"src" yaml:"src"`
	Type         consts.DataItType `json:"type,omitempty" yaml:"type,omitempty"`
	Url          string            `json:"url,omitempty" yaml:"url,omitempty"`
	DatapoolId   uint              `json:"datapoolId,omitempty" yaml:"datapoolId,omitempty"`
	DatapoolName string            `json:"datapoolName,omitempty" yaml:"datapoolName,omitempty"`
	Separator    string            `json:"separator,omitempty" yaml:"separator,omitempty"`

	RepeatTimes int `json:"repeatTimes,omitempty" yaml:"repeatTimes,omitempty"`
	//StartIndex     int    `json:"startIndex,omitempty" yaml:"startIndex,omitempty"`
	//EndIndex       int    `json:"endIndex,omitempty" yaml:"endIndex,omitempty"`

	IsLoop int  `json:"isLoop,omitempty" yaml:"isLoop,omitempty"`
	IsRand bool `json:"isRand,omitempty" yaml:"isRand,omitempty"`
	IsOnce bool `json:"isOnce,omitempty" yaml:"isOnce,omitempty"`

	VariableName string `json:"variableName,omitempty" yaml:"variableName,omitempty"`
}

func (entity ProcessorData) Run(processor *Processor, session *ExecSession) (err error) {
	logUtils.Infof("data entity")

	startTime := time.Now()
	processor.Result = &agentExecDomain.ScenarioExecResult{
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

	processor.Result.Iterator, processor.Result.Summary = entity.getIterator(session)

	detail := map[string]interface{}{"variableName": entity.VariableName, "url": entity.Url, "separator": entity.Separator, "repeatTimes": entity.RepeatTimes}
	processor.Result.Detail = commonUtils.JsonEncode(detail)

	processor.AddResultToParent()
	execUtils.SendExecMsg(*processor.Result, consts.Processor, session.WsMsg)

	entity.runDataItems(session, processor, processor.Result.Iterator)

	endTime := time.Now()
	processor.Result.EndTime = &endTime

	return
}

func (entity *ProcessorData) runDataItems(session *ExecSession, processor *Processor, iterator agentExecDomain.ExecIterator) (err error) {
	defer func() {
		if errX := recover(); errX != nil {
			processor.Error(session, errX)
		}
	}()

	ctx := session.Ctx
	executedProcessorIds := map[uint]bool{}

	for i := 0; i < entity.RepeatTimes; i++ {
		if entity.IsRand {
			iterator.Data = randArr(iterator.Data)
		}
		for index, item := range iterator.Data {
			if DemoTestSite != "" && index > 2 {
				break
			}

			SetVariable(session, processor.ID, iterator.VariableName, item, consts.ExtractorResultTypeString, consts.Public)

			round := ""

			for _, child := range processor.Children {
				select {
				case <-ctx.Done():
					break

				default:
				}
				if child.Disable {
					continue
				}

				executedProcessorIds[child.ID] = true

				if round == "" {
					round = fmt.Sprintf("第 %v 轮，%v = %v", i+1, iterator.VariableName, commonUtils.JsonEncode(item))
					child.Round = round
				}

				child.Run(session)
			}
		}
	}

	stat := CountSkip(session.ExecUuid, executedProcessorIds, processor.Children)
	execUtils.SendStatMsg(stat, session.WsMsg)

	return
}

func (entity *ProcessorData) getIterator(session *ExecSession) (iterator agentExecDomain.ExecIterator, msg string) {
	if entity.ID == 0 {
		msg = "执行前请先配置处理器。"
		return
	}

	iterator, _ = entity.GenerateLoopList(session)
	msg = fmt.Sprintf("迭代数据\"%s\"。", entity.Url)

	iterator.VariableName = entity.VariableName

	return
}

func (entity *ProcessorData) GenerateLoopList(session *ExecSession) (ret agentExecDomain.ExecIterator, err error) {
	if entity.Src == consts.SrcDatapool {
		for name, dp := range session.ExecScene.Datapools {
			if name == entity.DatapoolName {
				ret.Data = dp
				break
			}
		}

	} else {
		pth, _ := DownloadUploadedFile(session, entity.Url)
		if err != nil {
			logUtils.Infof("Download file %s failed", pth)
		}

		if entity.ProcessorType != consts.ProcessorDataDefault {
			err = errors.New("wrong data processor")
			return
		}

		format := commUtils.GetDataFileFormat(pth)
		if format == consts.FormatText || format == consts.FormatCsv {
			ret.Data, err = ReadDataFromText(pth, entity.Separator)
		} else if format == consts.FormatExcel {
			ret.Data, err = ReadDataFromExcel(pth)
		}
		/*
			if entity.IsRand {
				ret.Data = randArr(ret.Data)
			}
		*/
	}

	return
}

func randArr(arr []domain.VarKeyValuePair) (ret []domain.VarKeyValuePair) {
	indexArr := _intUtils.GenUniqueRandNum(0, len(arr), len(arr))

	for _, item := range indexArr {
		ret = append(ret, arr[item])
	}

	return
}
