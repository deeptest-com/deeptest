package agentExec

import (
	"encoding/json"
	"fmt"
	agentUtils "github.com/aaronchen2k/deeptest/internal/agent/exec/utils"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	extractorHelper "github.com/aaronchen2k/deeptest/internal/pkg/helper/extractor"
	_stringUtils "github.com/aaronchen2k/deeptest/pkg/lib/string"
	"strings"
)

func ExecCheckPoint(checkpoint *domain.CheckpointBase, resp domain.DebugResponse, processorId uint, session *ExecSession) (err error) {
	checkpoint.ResultStatus = consts.Pass

	// Judgement 表达式
	if checkpoint.Type == consts.Judgement {
		result, variablesArr := computerExpr(checkpoint.Expression, session)

		checkpoint.Variables = getVariableArrDesc(variablesArr)
		checkpoint.ActualResult = fmt.Sprintf("%v", result)

		ret, ok := result.(bool)
		if ok && ret {
			checkpoint.ResultStatus = consts.Pass
		} else {
			checkpoint.ResultStatus = consts.Fail
		}

		return
	}

	// 计算表达式
	checkpointValue, variablesArr := computerExpr(checkpoint.Value, session)
	checkpointValue = _stringUtils.InterfToStr(checkpointValue)
	checkpoint.Variables = getVariableArrDesc(variablesArr)

	// Response ResultStatus
	if checkpoint.Type == consts.ResponseStatus {
		expectCodeNum := _stringUtils.ParseInt(fmt.Sprintf("%v", checkpointValue))

		checkpoint.ActualResult = fmt.Sprintf("%d", resp.StatusCode)

		if checkpoint.Operator == consts.Equal && resp.StatusCode == expectCodeNum {
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

		if checkpoint.Operator == consts.Equal && headerValue == checkpointValue {
			checkpoint.ResultStatus = consts.Pass
		} else if checkpoint.Operator == consts.NotEqual && headerValue != checkpointValue {
			checkpoint.ResultStatus = consts.Pass
		} else if checkpoint.Operator == consts.Contain && strings.Contains(headerValue, _stringUtils.InterfToStr(checkpointValue)) {
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
		if checkpoint.Operator == consts.Equal && resp.Content == checkpointValue {
			checkpoint.ResultStatus = consts.Pass
		} else if checkpoint.Operator == consts.NotEqual && resp.Content != checkpointValue {
			checkpoint.ResultStatus = consts.Pass
		} else if checkpoint.Operator == consts.Contain && strings.Contains(resp.Content, _stringUtils.InterfToStr(checkpointValue)) {
			checkpoint.ResultStatus = consts.Pass
		} else {
			checkpoint.ResultStatus = consts.Fail
		}

		return
	}

	// ExtractorVari
	if checkpoint.Type == consts.ExtractorVari {
		variable, _ := GetVariable(checkpoint.ExtractorVariable, session.GetCurrScenarioProcessorId(), session)

		checkpoint.ActualResult = fmt.Sprintf("%v", variable.Value)

		checkpoint.ResultStatus = agentUtils.Compare(checkpoint.Operator, checkpoint.ActualResult, checkpointValue)

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

		checkpoint.ResultStatus = agentUtils.Compare(checkpoint.Operator, checkpoint.ActualResult, checkpointValue)

		return
	}

	return
}

func computerExpr(expression string, session *ExecSession) (result interface{}, params domain.VarKeyValuePair) {
	result, params = NewGojaSimple().ExecJsFuncSimple(expression, session, true)

	return
}

func getVariableArrDesc(variablesArr domain.VarKeyValuePair) (ret string) {
	variablesBytes, _ := json.Marshal(combineSameVars(variablesArr))
	ret = string(variablesBytes)

	return
}

func combineSameVars(variables domain.VarKeyValuePair) (ret domain.VarKeyValuePair) {
	ret = domain.VarKeyValuePair{}

	for key, val := range variables {
		name := strings.TrimLeft(key, "+")
		ret[name] = val
	}

	return
}
