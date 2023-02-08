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

type ProcessorData struct {
	ID uint `json:"id" yaml:"id"`
	ProcessorEntityBase

	Type      consts.DataSource `json:"type,omitempty" yaml:"type,omitempty"`
	Url       string            `json:"url,omitempty" yaml:"url,omitempty"`
	Separator string            `json:"separator,omitempty" yaml:"separator,omitempty"`

	RepeatTimes int `json:"repeatTimes,omitempty" yaml:"repeatTimes,omitempty"`
	//StartIndex     int    `json:"startIndex,omitempty" yaml:"startIndex,omitempty"`
	//EndIndex       int    `json:"endIndex,omitempty" yaml:"endIndex,omitempty"`

	IsLoop int  `json:"isLoop,omitempty" yaml:"isLoop,omitempty"`
	IsRand bool `json:"isRand,omitempty" yaml:"isRand,omitempty"`
	IsOnce bool `json:"isOnce,omitempty" yaml:"isOnce,omitempty"`

	VariableName string `json:"variableName,omitempty" yaml:"variableName,omitempty"`
}

func (entity ProcessorData) Run(processor *Processor, session *Session) (err error) {
	logUtils.Infof("data entity")

	startTime := time.Now()
	processor.Result = &agentDomain.Result{
		ID:                int(entity.ProcessorID),
		Name:              entity.Name,
		ProcessorCategory: entity.ProcessorCategory,
		ProcessorType:     entity.ProcessorType,
		StartTime:         &startTime,
		ParentId:          int(entity.ParentID),
	}

	processor.Result.Iterator, processor.Result.Summary = entity.getIterator()

	processor.AddResultToParent()
	execUtils.SendExecMsg(*processor.Result, session.WsMsg)

	entity.runDataItems(session, processor, processor.Result.Iterator)

	endTime := time.Now()
	processor.Result.EndTime = &endTime

	//result = *processor.Result

	return
}

func (entity *ProcessorData) runDataItems(session *Session, processor *Processor, iterator agentDomain.ExecIterator) (err error) {
	for _, item := range iterator.Data {
		SetVariable(processor.ID, iterator.VariableName, item, consts.Local)

		for _, child := range processor.Children {
			child.Run(session)
		}
	}

	return
}

func (entity *ProcessorData) getIterator() (iterator agentDomain.ExecIterator, msg string) {
	if entity.ID == 0 {
		msg = "执行前请先配置处理器。"
		return
	}

	iterator, _ = entity.GenerateLoopList()
	msg = fmt.Sprintf("迭代数据\"%s\"。", entity.Url)

	iterator.VariableName = entity.VariableName

	return
}

func (entity *ProcessorData) GenerateLoopList() (ret agentDomain.ExecIterator, err error) {
	pth, _ := agentUtils.DownloadUploadedFile(entity.Url)
	if err != nil {
		logUtils.Infof("Download file %s failed", pth)
	}

	if entity.ProcessorType == consts.ProcessorDataText {
		ret.Data, err = agentUtils.ReadDataFromText(pth, entity.Separator)
	} else if entity.ProcessorType == consts.ProcessorDataExcel {
		ret.Data, err = agentUtils.ReadDataFromExcel(pth)
	}

	return
}
