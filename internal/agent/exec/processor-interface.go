package agentExec

import (
	"encoding/json"
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/agent/exec/domain"
	"github.com/aaronchen2k/deeptest/internal/agent/exec/utils"
	"github.com/aaronchen2k/deeptest/internal/agent/exec/utils/exec"
	queryHelper "github.com/aaronchen2k/deeptest/internal/agent/exec/utils/query"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	stringUtils "github.com/aaronchen2k/deeptest/pkg/lib/string"

	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"strings"
)

type ProcessorInterface struct {
	ID uint `json:"id"`
	ProcessorEntityBase

	domain.Request
	Response domain.Response `json:"response"`

	Extractors  []domain.Extractor
	Checkpoints []domain.Checkpoint
}

func (entity ProcessorInterface) Run(processor *Processor, session *Session) (log domain.Result, err error) {
	logUtils.Infof("interface entity")

	processor.Result = domain.Result{
		ID:                entity.ProcessorID,
		Name:              entity.Name,
		ProcessorCategory: entity.ProcessorCategory,
		ProcessorType:     entity.ProcessorType,
		ParentId:          entity.ParentID,
	}

	variableMap := GetVariableMap(entity.ProcessorID)
	ReplaceAll(&entity.Request, variableMap)

	GetRequestProps(&entity.Request)
	entity.Response, err = Invoke(entity.Request)
	GetContentProps(&entity.Response)

	reqContent, _ := json.Marshal(entity.Request)
	processor.Result.ReqContent = string(reqContent)

	respContent, _ := json.Marshal(entity.Response)
	processor.Result.RespContent = string(respContent)

	if err != nil {
		processor.Result.ResultStatus = consts.Fail
		processor.Result.Summary = err.Error()
		processor.Parent.Result.Children = append(processor.Parent.Result.Children, &processor.Result)
		exec.SendErrorMsg(processor.Result, session.WsMsg)
		return
	}

	entity.ExtractInterface(processor, session)
	entity.CheckInterface(processor, session)

	processor.Parent.Result.Children = append(processor.Parent.Result.Children, &processor.Result)
	exec.SendExecMsg(processor.Result, session.WsMsg)

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

func (entity ProcessorInterface) Extract(extractor *domain.Extractor, resp domain.Response) (err error) {
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
			if utils.IsJsonContent(resp.ContentType.String()) && extractor.Type == consts.JsonQuery {
				extractor.Result = queryHelper.JsonQuery(resp.Content, extractor.Expression)

			} else if utils.IsHtmlContent(resp.ContentType.String()) && extractor.Type == consts.HtmlQuery {
				extractor.Result = queryHelper.HtmlQuery(resp.Content, extractor.Expression)

			} else if utils.IsXmlContent(resp.ContentType.String()) && extractor.Type == consts.XmlQuery {
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

func (entity *ProcessorInterface) Check(checkpoint *domain.Checkpoint, resp domain.Response) (err error) {
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

		checkpoint.ResultStatus = utils.Compare(checkpoint.Operator, checkpoint.ActualResult, checkpoint.Value)

		return
	}

	return
}
