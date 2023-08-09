package agentExec

import (
	"encoding/json"
	"github.com/aaronchen2k/deeptest/internal/agent/exec/domain"
	"github.com/aaronchen2k/deeptest/internal/agent/exec/utils/exec"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"github.com/jinzhu/copier"
	uuid "github.com/satori/go.uuid"
	"strings"
	"time"
)

type ProcessorInterface struct {
	ID uint `json:"id"`
	ProcessorEntityBase

	domain.BaseRequest
	Response domain.DebugResponse `json:"response"`

	BaseUrl string `json:"baseUrl"`

	PreConditions  []domain.InterfaceExecCondition `json:"preConditions"`
	PostConditions []domain.InterfaceExecCondition `json:"postConditions"`
}

func (entity ProcessorInterface) Run(processor *Processor, session *Session) (err error) {
	logUtils.Infof("interface entity")
	CurrDebugInterfaceId = processor.EntityId

	startTime := time.Now()
	processor.Result = &agentDomain.ScenarioExecResult{
		ID:                  int(entity.ProcessorID),
		Name:                entity.Name,
		ProcessorCategory:   entity.ProcessorCategory,
		ProcessorType:       entity.ProcessorType,
		StartTime:           &startTime,
		ParentId:            int(entity.ParentID),
		EndpointInterfaceId: processor.EndpointInterfaceId,
		DebugInterfaceId:    processor.EntityId,
		ProcessorId:         processor.ID,
		ScenarioId:          processor.ScenarioId,
		LogId:               uuid.NewV4(),
		ParentLogId:         processor.Parent.Result.LogId,
	}

	//在循环过程中，processor 被执行多次，变量替换会受到影响，第一次跌替换之后，就不能根据实际情况替换了
	var baseRequest domain.BaseRequest
	copier.CopyWithOption(&baseRequest, &entity.BaseRequest, copier.Option{IgnoreEmpty: true, DeepCopy: true})

	// exec pre-condition
	entity.ExecPreConditions(processor, session)

	// dealwith variables
	ReplaceVariables(&baseRequest, consts.ScenarioDebug)

	// add cookies
	DealwithCookies(&baseRequest, entity.ProcessorID)

	// gen request url
	GenRequestUrlWithBaseUrlAndPathParam(&baseRequest, processor.EntityId, entity.BaseUrl)

	// send request
	//startTime := time.UnixNano()
	entity.Response, err = Invoke(&baseRequest)

	processor.Result.Cost = time.Now().UnixMilli() - startTime.UnixMilli()
	reqContent, _ := json.Marshal(baseRequest)
	processor.Result.ReqContent = string(reqContent)
	respContent, _ := json.Marshal(entity.Response)
	processor.Result.RespContent = string(respContent)

	if err != nil {
		processor.Result.ResultStatus = consts.Fail
		processor.Result.Summary = err.Error()
		processor.AddResultToParent()
		execUtils.SendErrorMsg(*processor.Result, session.WsMsg)
		return
	}

	// exec post-condition
	entity.ExecPostConditions(processor, session)

	for _, c := range entity.Response.Cookies {
		SetCookie(processor.ParentId, c.Name, c.Value, c.Domain, c.ExpireTime)
	}

	processor.AddResultToParent()
	execUtils.SendExecMsg(*processor.Result, session.WsMsg)

	endTime := time.Now()
	processor.Result.EndTime = &endTime

	return
}

func (entity *ProcessorInterface) ExecPreConditions(processor *Processor, session *Session) (err error) {
	return
}
func (entity *ProcessorInterface) ExecPostConditions(processor *Processor, session *Session) (err error) {
	return
}

//func (entity *ProcessorInterface) ExtractInterface(processor *Processor, session *Session) (err error) {
//	for _, extractor := range entity.Extractors {
//		err = entity.extract(&extractor, entity.Response)
//		SetVariable(entity.ParentID, extractor.Variable, extractor.Result, extractor.Scope)
//
//		if err == nil { // gen report for processor
//			processor.Result.ExtractorsResult = append(processor.Result.ExtractorsResult, extractor)
//		}
//	}
//
//	return
//}
//
//func (entity *ProcessorInterface) CheckInterface(processor *Processor, session *Session) (err error) {
//	status := consts.Pass
//
//	for _, checkpoint := range entity.Checkpoints {
//		ExecCheck(&checkpoint, entity.Response, entity.ProcessorID)
//
//		if checkpoint.ResultStatus == consts.Fail {
//			status = consts.Fail
//		}
//
//		if err == nil {
//			processor.Result.CheckpointsResult = append(processor.Result.CheckpointsResult, checkpoint)
//		}
//	}
//
//	processor.Result.ResultStatus = status
//
//	return
//}

func (entity *ProcessorInterface) extract(extractor *domain.ExtractorBase, resp domain.DebugResponse) (err error) {
	err = ExecExtract(extractor, resp)

	extractor.Result = strings.TrimSpace(extractor.Result)

	return
}
