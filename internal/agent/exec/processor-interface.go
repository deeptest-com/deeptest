package agentExec

import (
	"encoding/json"
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/agent/exec/domain"
	agentUtils "github.com/aaronchen2k/deeptest/internal/agent/exec/utils"
	"github.com/aaronchen2k/deeptest/internal/agent/exec/utils/exec"
	queryHelper "github.com/aaronchen2k/deeptest/internal/agent/exec/utils/query"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	httpHelper "github.com/aaronchen2k/deeptest/internal/pkg/helper/http"
	stringUtils "github.com/aaronchen2k/deeptest/pkg/lib/string"
	"time"

	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"strings"
)

type ProcessorInterface struct {
	ID uint `json:"id"`
	ProcessorEntityBase

	v1.BaseRequest
	Response v1.DebugResponse `json:"response"`

	Extractors  []agentDomain.Extractor
	Checkpoints []agentDomain.Checkpoint
}

func (entity ProcessorInterface) Run(processor *Processor, session *Session) (err error) {
	logUtils.Infof("interface entity")

	startTime := time.Now()
	processor.Result = &agentDomain.ScenarioExecResult{
		ID:                int(entity.ProcessorID),
		Name:              entity.Name,
		ProcessorCategory: entity.ProcessorCategory,
		ProcessorType:     entity.ProcessorType,
		StartTime:         &startTime,
		ParentId:          int(entity.ParentID),
	}

	// exec pre-request script
	ExecJs(entity.PreRequestScript)

	// replace variables
	ReplaceRequestWithVars(&entity.BaseRequest)

	// send request
	entity.Response, err = Invoke(&entity.BaseRequest)

	reqContent, _ := json.Marshal(entity.BaseRequest)
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

	entity.ExtractInterface(processor, session)
	entity.CheckInterface(processor, session)

	processor.AddResultToParent()
	execUtils.SendExecMsg(*processor.Result, session.WsMsg)

	endTime := time.Now()
	processor.Result.EndTime = &endTime

	return
}

func (entity *ProcessorInterface) ExtractInterface(processor *Processor, session *Session) (err error) {
	for _, extractor := range entity.Extractors {
		err = entity.Extract(&extractor, entity.Response)
		SetVariable(entity.ParentID, extractor.Variable, extractor.Result, extractor.Scope)

		if err == nil { // gen report for processor
			processor.Result.ExtractorsResult = append(processor.Result.ExtractorsResult, extractor)
		}
	}

	return
}

func (entity *ProcessorInterface) CheckInterface(processor *Processor, session *Session) (err error) {
	status := consts.Pass

	for _, checkpoint := range entity.Checkpoints {
		entity.Check(&checkpoint, entity.Response)

		if checkpoint.ResultStatus == consts.Fail {
			status = consts.Fail
		}

		if err == nil {
			processor.Result.CheckpointsResult = append(processor.Result.CheckpointsResult, checkpoint)
		}
	}

	processor.Result.ResultStatus = status

	return
}

func (entity *ProcessorInterface) Extract(extractor *agentDomain.Extractor, resp v1.DebugResponse) (err error) {
	extractor.Result = ""

	if extractor.Disabled {
		extractor.Result = ""
	} else {
		if extractor.Src == consts.Header {
			for _, h := range resp.Headers {
				if h.Name == extractor.Key {
					extractor.Result = h.Value
					break
				}
			}
		} else {
			if httpHelper.IsJsonContent(resp.ContentType.String()) && extractor.Type == consts.JsonQuery {
				extractor.Result = queryHelper.JsonQuery(resp.Content, extractor.Expression)

			} else if httpHelper.IsHtmlContent(resp.ContentType.String()) && extractor.Type == consts.HtmlQuery {
				extractor.Result = queryHelper.HtmlQuery(resp.Content, extractor.Expression)

			} else if httpHelper.IsXmlContent(resp.ContentType.String()) && extractor.Type == consts.XmlQuery {
				extractor.Result = queryHelper.XmlQuery(resp.Content, extractor.Expression)

			} else if extractor.Type == consts.Boundary {
				extractor.Result = queryHelper.BoundaryQuery(resp.Content, extractor.BoundaryStart, extractor.BoundaryEnd,
					extractor.BoundaryIndex, extractor.BoundaryIncluded)
			}
		}
	}

	extractor.Result = strings.TrimSpace(extractor.Result)

	return
}

func (entity *ProcessorInterface) Check(checkpoint *agentDomain.Checkpoint, resp v1.DebugResponse) (err error) {
	if checkpoint.Disabled {
		checkpoint.ResultStatus = ""
		return
	}

	checkpoint.ResultStatus = consts.Pass

	// Response ResultStatus
	if checkpoint.Type == consts.ResponseStatus {
		expectCode := stringUtils.ParseInt(checkpoint.Value)

		checkpoint.ActualResult = fmt.Sprintf("%d", resp.StatusCode.Int())

		if checkpoint.Operator == consts.Equal && resp.StatusCode.Int() == expectCode {
			checkpoint.ResultStatus = consts.Pass
		} else {
			checkpoint.ResultStatus = consts.Fail
		}

		return
	}

	// Response Header
	if checkpoint.Type == consts.ResponseHeader {
		headerValue := ""
		for _, h := range resp.Headers {
			if h.Name == checkpoint.Expression {
				headerValue = h.Value
				break
			}
		}

		checkpoint.ActualResult = headerValue

		if checkpoint.Operator == consts.Equal && headerValue == checkpoint.Value {
			checkpoint.ResultStatus = consts.Pass
		} else if checkpoint.Operator == consts.NotEqual && headerValue != checkpoint.Value {
			checkpoint.ResultStatus = consts.Pass
		} else if checkpoint.Operator == consts.Contain && strings.Contains(headerValue, checkpoint.Value) {
			checkpoint.ResultStatus = consts.Pass
		} else {
			checkpoint.ResultStatus = consts.Fail
		}

		return
	}

	var jsonData interface{}
	json.Unmarshal([]byte(resp.Content), &jsonData)

	checkpoint.ActualResult = "<RESPONSE_BODY>"

	// Response Body
	if checkpoint.Type == consts.ResponseBody {
		if checkpoint.Operator == consts.Equal && resp.Content == checkpoint.Value {
			checkpoint.ResultStatus = consts.Pass
		} else if checkpoint.Operator == consts.NotEqual && resp.Content != checkpoint.Value {
			checkpoint.ResultStatus = consts.Pass
		} else if checkpoint.Operator == consts.Contain && strings.Contains(resp.Content, checkpoint.Value) {
			checkpoint.ResultStatus = consts.Pass
		} else {
			checkpoint.ResultStatus = consts.Fail
		}

		return
	}

	// Judgement
	if checkpoint.Type == consts.Judgement {
		result, _ := EvaluateGovaluateExpressionByScope(checkpoint.Expression, entity.ProcessorID)

		ret, ok := result.(bool)
		if ok && ret {
			checkpoint.ResultStatus = consts.Pass
		} else {
			checkpoint.ResultStatus = consts.Fail
		}
		checkpoint.ActualResult = fmt.Sprintf("%v", ret)

		return
	}

	// Extractor
	if checkpoint.Type == consts.Extractor {
		// get extractor variable value saved by previous extract opt
		variable, _ := GetVariable(entity.ProcessorID, checkpoint.ExtractorVariable)
		checkpoint.ActualResult = variable.Value.(string)

		checkpoint.ResultStatus = agentUtils.Compare(checkpoint.Operator, checkpoint.ActualResult, checkpoint.Value)

		return
	}

	return
}
