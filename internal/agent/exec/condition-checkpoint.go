package agentExec

import (
	"encoding/json"
	"fmt"
	agentUtils "github.com/aaronchen2k/deeptest/internal/agent/exec/utils"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	extractorHelper "github.com/aaronchen2k/deeptest/internal/pkg/helper/extractor"
	_stringUtils "github.com/aaronchen2k/deeptest/pkg/lib/string"
	"strconv"
	"strings"
)

func ExecCheckPoint(checkpoint *domain.CheckpointBase, resp domain.DebugResponse, processorId uint, session *ExecSession) (err error) {
	checkpoint.ResultStatus = consts.Pass

	// Judgement 表达式
	if checkpoint.Type == consts.Judgement {
		boolResult, variablesArr, err1 := computerExpr(checkpoint.Expression, session)
		checkpoint.Variables = getVariableArrDesc(variablesArr)
		checkpoint.ActualResult = fmt.Sprintf("%v", boolResult)

		if err1 != nil {
			checkpoint.ResultStatus = consts.Fail
			err = err1
			return
		}

		ret, ok := boolResult.(bool)
		if ok && ret {
			checkpoint.ResultStatus = consts.Pass
		} else {
			checkpoint.ResultStatus = consts.Fail
		}

		return
	}

	// 计算表达式
	checkpointValue, variablesArr, err := computerExpr(checkpoint.Value, session)
	checkpoint.Variables = getVariableArrDesc(variablesArr)
	checkpointValue = _stringUtils.InterfToStr(checkpointValue)

	// Response ResultStatus
	if checkpoint.Type == consts.ResponseStatus {
		checkpoint.ActualResult = fmt.Sprintf("%d", resp.StatusCode)
		if err != nil {
			checkpoint.ExpectResult = fmt.Sprintf("%v", checkpointValue)
			checkpoint.ResultStatus = consts.Fail
			return
		}

		checkpointValueStr := fmt.Sprintf("%v", checkpointValue)
		expectCodeNum, err1 := strconv.Atoi(checkpointValueStr)

		expectValue := fmt.Sprintf("%d", expectCodeNum)
		if err1 != nil {
			expectValue = checkpointValueStr
			checkpoint.ResultStatus = consts.Fail

			return
		}

		checkpoint.ExpectResult = fmt.Sprintf("%v", expectValue)

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
		checkpoint.ExpectResult = _stringUtils.InterfToStr(checkpointValue)
		if err != nil {
			checkpoint.ExpectResult = fmt.Sprintf("%v", checkpointValue)
			checkpoint.ResultStatus = consts.Fail
			return
		}

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

	// Response Body
	if checkpoint.Type == consts.ResponseBody {
		var jsonData interface{}
		json.Unmarshal([]byte(resp.Content), &jsonData)

		checkpoint.ActualResult = "<RESPONSE_BODY>"
		checkpoint.ExpectResult = _stringUtils.InterfToStr(checkpointValue)
		if err != nil {
			checkpoint.ResultStatus = consts.Fail
			return
		}

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

	// Extractor Vari
	if checkpoint.Type == consts.ExtractorVari {
		variable, _ := GetVariable(checkpoint.ExtractorVariable, session.GetCurrScenarioProcessorId(), session)

		checkpoint.ActualResult = fmt.Sprintf("%v", variable.Value)
		checkpoint.ExpectResult = _stringUtils.InterfToStr(checkpointValue)
		if err != nil {
			checkpoint.ResultStatus = consts.Fail
			return
		}

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
		checkpoint.ExpectResult = _stringUtils.InterfToStr(checkpointValue)

		checkpoint.ResultStatus = agentUtils.Compare(checkpoint.Operator, checkpoint.ActualResult, checkpointValue)

		return
	}

	return
}

func computerExpr(expression string, session *ExecSession) (
	expectResult interface{}, params domain.VarKeyValuePair, err error) {

	expectResult, params, err = NewGojaSimple().ExecJsFuncSimple(expression, session, true)

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
