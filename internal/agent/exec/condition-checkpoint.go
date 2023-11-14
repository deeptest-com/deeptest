package agentExec

import (
	"encoding/json"
	"fmt"
	agentUtils "github.com/aaronchen2k/deeptest/internal/agent/exec/utils"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	extractorHelper "github.com/aaronchen2k/deeptest/internal/pkg/helper/extractor"
	stringUtils "github.com/aaronchen2k/deeptest/pkg/lib/string"
	"strings"
)

func ExecCheckPoint(checkpoint *domain.CheckpointBase, resp domain.DebugResponse, processorId uint, execUuid string) (err error) {
	checkpoint.ResultStatus = consts.Pass

	// Response ResultStatus
	if checkpoint.Type == consts.ResponseStatus {
		expectCode := stringUtils.ParseInt(checkpoint.Value)

		checkpoint.ActualResult = fmt.Sprintf("%d", resp.StatusCode)

		if checkpoint.Operator == consts.Equal && resp.StatusCode == consts.HttpRespCode(expectCode) {
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
		var result interface{}

		expr := ReplaceDatapoolVariInGovaluateExpress(checkpoint.Expression, execUuid)

		if processorId > 0 { // exec interface processor in scenario
			result, _ = EvaluateGovaluateExpressionByProcessorScope(expr, processorId, execUuid)
		} else { // exec by interface invocation
			result, _ = EvaluateGovaluateExpressionWithDebugVariables(expr, execUuid)
		}

		checkpoint.ActualResult = fmt.Sprintf("%v", result)

		ret, ok := result.(bool)
		if ok && ret {
			checkpoint.ResultStatus = consts.Pass
		} else {
			checkpoint.ResultStatus = consts.Fail
		}

		return
	}

	// ExtractorVari
	if checkpoint.Type == consts.ExtractorVari {
		variable, _ := GetVariable(GetCurrScenarioProcessorId(execUuid), checkpoint.ExtractorVariable, execUuid)

		checkpoint.ActualResult = fmt.Sprintf("%v", variable.Value)

		checkpoint.ResultStatus = agentUtils.Compare(checkpoint.Operator, checkpoint.ActualResult, checkpoint.Value)

		return
	}

	// Extractor
	if checkpoint.Type == consts.Extractor {
		extractor := domain.ExtractorBase{
			Type:       checkpoint.ExtractorType,
			Expression: checkpoint.ExtractorExpression,
		}
		extractorHelper.Extract(&extractor, resp)

		checkpoint.ActualResult = fmt.Sprintf("%v", extractor.Result)

		checkpoint.ResultStatus = agentUtils.Compare(checkpoint.Operator, checkpoint.ActualResult, checkpoint.Value)

		return
	}

	return
}
