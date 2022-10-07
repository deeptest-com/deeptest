package agentExec

import (
	"encoding/json"
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/agent/exec/domain"
	"github.com/aaronchen2k/deeptest/internal/agent/exec/utils"
	queryHelper2 "github.com/aaronchen2k/deeptest/internal/agent/exec/utils/query"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/helper/websocket"
	stringUtils "github.com/aaronchen2k/deeptest/pkg/lib/string"

	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"github.com/jinzhu/copier"
	"strings"
)

type ProcessorInterface struct {
	ID uint `json:"id"`
	ProcessorEntity

	Request  domain.Request  `json:"request"`
	Response domain.Response `json:"response"`

	Extractors  []domain.ExecInterfaceExtractor
	Checkpoints []domain.ExecInterfaceCheckpoint

	Result Result `json:"id"`
}

func (p ProcessorInterface) Run(s *Session) (log Result, err error) {
	logUtils.Infof("interface entity")

	variableMap := GetVariableMap(p.ProcessorID)
	ReplaceAll(&p.Request, variableMap)

	p.Response, err = Invoke(p.Request)
	if err != nil {
		return
	}

	p.ExtractInterface(s)
	p.CheckInterface(s)

	reqContent, _ := json.Marshal(p.Request)
	respContent, _ := json.Marshal(p.Response)
	p.Result.ReqContent = string(reqContent)
	p.Result.RespContent = string(respContent)

	websocketHelper.SendExecMsg("exec interface", p.Result, s.WsMsg)

	return
}

func (p ProcessorInterface) ExtractInterface(s *Session) (err error) {
	for _, extractor := range p.Extractors {
		p.Extract(&extractor, p.Response)

		if err == nil { // gen report for processor
			interfaceExtractor := domain.ExecInterfaceExtractor{}
			copier.CopyWithOption(&interfaceExtractor, extractor, copier.Option{DeepCopy: true})

			p.Extractors = append(p.Extractors, interfaceExtractor)
		}
	}

	return
}

func (p *ProcessorInterface) CheckInterface(s *Session) (err error) {
	status := consts.Pass
	for _, checkpoint := range p.Checkpoints {
		p.Check(&checkpoint, p.Response)

		if checkpoint.ResultStatus == consts.Fail {
			status = consts.Fail
		}

		if err == nil {
			interfaceCheckpoint := domain.ExecInterfaceCheckpoint{}
			copier.CopyWithOption(&interfaceCheckpoint, checkpoint, copier.Option{DeepCopy: true})

			p.Checkpoints = append(p.Checkpoints, interfaceCheckpoint)
		}
	}

	p.Result.ResultStatus = status

	return
}

func (p ProcessorInterface) Extract(extractor *domain.ExecInterfaceExtractor, resp domain.Response) (err error) {
	p.ExtractValue(extractor, resp)
	SetVariable(p.ID, extractor.Variable, extractor.Result, extractor.Scope)

	return
}

func (p ProcessorInterface) ExtractValue(extractor *domain.ExecInterfaceExtractor, resp domain.Response) (err error) {
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
				extractor.Result = queryHelper2.JsonQuery(resp.Content, extractor.Expression)

			} else if utils.IsHtmlContent(resp.ContentType.String()) && extractor.Type == consts.HtmlQuery {
				extractor.Result = queryHelper2.HtmlQuery(resp.Content, extractor.Expression)

			} else if utils.IsXmlContent(resp.ContentType.String()) && extractor.Type == consts.XmlQuery {
				extractor.Result = queryHelper2.XmlQuery(resp.Content, extractor.Expression)

			} else if extractor.Type == consts.Boundary {
				extractor.Result = queryHelper2.BoundaryQuery(resp.Content, extractor.BoundaryStart, extractor.BoundaryEnd,
					extractor.BoundaryIndex, extractor.BoundaryIncluded)
			}
		}
	}

	extractor.Result = strings.TrimSpace(extractor.Result)

	return
}

func (p *ProcessorInterface) Check(checkpoint *domain.ExecInterfaceCheckpoint, resp domain.Response) (err error) {
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
		result, _ := EvaluateGovaluateExpressionByScope(checkpoint.Expression, p.ID)

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
		variable, _ := GetVariable(p.ID, checkpoint.ExtractorVariable)
		checkpoint.ActualResult = variable.Value.(string)

		checkpoint.ResultStatus = utils.Compare(checkpoint.Operator, checkpoint.ActualResult, checkpoint.Value)

		return
	}

	return
}
