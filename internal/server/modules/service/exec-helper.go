package service

import (
	"encoding/json"
	"fmt"
	"github.com/Knetic/govaluate"
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	commUtils "github.com/aaronchen2k/deeptest/internal/pkg/utils"
	"github.com/aaronchen2k/deeptest/internal/server/modules/business"
	"github.com/aaronchen2k/deeptest/internal/server/modules/helper/exec"
	"github.com/aaronchen2k/deeptest/internal/server/modules/helper/expression"
	model "github.com/aaronchen2k/deeptest/internal/server/modules/model"
	repo2 "github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	"github.com/jinzhu/copier"
	"github.com/kataras/iris/v12/websocket"
	"regexp"
)

type ExecHelperService struct {
	ScenarioProcessorRepo *repo2.ScenarioProcessorRepo `inject:""`
	ScenarioRepo          *repo2.ScenarioRepo          `inject:""`
	TestResultRepo        *repo2.ReportRepo            `inject:""`
	TestLogRepo           *repo2.LogRepo               `inject:""`
	InterfaceRepo         *repo2.InterfaceRepo         `inject:""`
	InterfaceService      *InterfaceService            `inject:""`
	ExecContext           *business.ExecContext        `inject:""`
	ExtractorService      *ExtractorService            `inject:""`
}

//func (s *ExecComm) ParseThreadGroup(processor *model.ProcessorThreadGroup, log *domain.ExecLog, msg *websocket.Message) (
//	result string, err error) {
//
//	return
//}

func (s *ExecHelperService) EvaluateLogic(logic *model.ProcessorLogic, parentLog *domain.ExecLog, msg *websocket.Message) (
	output domain.ExecOutput, err error) {
	if logic.ID == 0 {
		output.Msg = "执行前请先配置处理器。"
		return
	}

	typ := logic.ProcessorType

	if typ == consts.ProcessorLogicIf {
		var result interface{}
		result, err = s.EvaluateGovaluateExpression(logic.Expression, logic.ProcessorID)
		if err != nil {
			output.Pass = false
			output.Msg = fmt.Sprintf("不通过")
			return
		}

		output.Pass = result.(bool)
		output.Msg = fmt.Sprintf("通过")
		return
	} else if typ == consts.ProcessorLogicElse {
		output.Pass = false

		brother, ok := getPreviousBrother(*parentLog)
		if ok && !brother.Output.Pass {
			output.Pass = true
			output.Msg = fmt.Sprintf("通过")
		}

		//output.Msg = fmt.Sprintf("不通过")
		return
	}

	return
}

func (s *ExecHelperService) EvaluateLoop(loop *model.ProcessorLoop, parentLog *domain.ExecLog, msg *websocket.Message) (
	output domain.ExecOutput, err error) {

	if loop.ID == 0 {
		output.Msg = "执行前请先配置处理器。"
		return
	}

	typ := loop.ProcessorType
	if typ == consts.ProcessorLoopTime {
		output.Times = loop.Times
		output.Msg = fmt.Sprintf("执行\"%d\"次。", output.Times)
		return
	} else if typ == consts.ProcessorLoopUntil {
		output.Expression = loop.UntilExpression
		output.Msg = fmt.Sprintf("直到\"%s\"。", output.Expression)
		return
	} else if typ == consts.ProcessorLoopIn {
		output.List = loop.List
		output.Msg = fmt.Sprintf("迭代列表\"%s\"。", output.List)
		return
	} else if typ == consts.ProcessorLoopRange {
		output.Range = loop.Range
		output.Msg = fmt.Sprintf("区间\"%s\"。", output.Range)
		return
	}

	return
}

func (s *ExecHelperService) EvaluateLoopBreak(loop *model.ProcessorLoop, parentLog *domain.ExecLog, msg *websocket.Message) (
	output domain.ExecOutput, err error) {

	if loop.ID == 0 {
		output.Msg = "执行前请先配置处理器。"
		return
	}

	output.Expression = loop.BreakIfExpression
	output.BreakFrom = parentLog.ProcessId
	output.Msg = fmt.Sprintf("-")

	return
}

func (s *ExecHelperService) EvaluateData(data *model.ProcessorData, parentLog *domain.ExecLog, msg *websocket.Message) (
	output domain.ExecOutput, err error) {
	output.Url = data.Url

	output.Msg = fmt.Sprintf("迭代\"%s\"为变量\"%s\"。", output.Url, data.VariableName)

	return
}

func (s *ExecHelperService) EvaluateTimer(processor *model.ProcessorTimer, parentLog *domain.ExecLog, msg *websocket.Message) (
	output domain.ExecOutput, err error) {

	output.SleepTime = processor.SleepTime
	output.Msg = fmt.Sprintf("等待\"%d\"秒。", output.SleepTime)

	return
}

func (s *ExecHelperService) EvaluatePrint(processor *model.ProcessorPrint, parentLog *domain.ExecLog, msg *websocket.Message) (
	output domain.ExecOutput, err error) {

	output.Expression = processor.Expression

	return
}

func (s *ExecHelperService) EvaluateVariable(processor *model.ProcessorVariable, parentLog *domain.ExecLog, msg *websocket.Message) (
	output domain.ExecOutput, err error) {

	output.VariableName = processor.VariableName
	expression := processor.RightValue
	typ := processor.ProcessorType

	if typ == consts.ProcessorVariableSet {
		output.VariableValue, err = s.EvaluateGovaluateExpression(expression, processor.ProcessorID)
		if err != nil {
			output.Msg = fmt.Sprintf("计算表达式\"%s\"错误 \"%s\"。", expression, err.Error())
			return
		}

		output.Msg = fmt.Sprintf("\"%s\"为\"%v\"。", output.VariableName, output.VariableValue)

	} else if typ == consts.ProcessorVariableClear {
		output.Msg = fmt.Sprintf("%s。", output.VariableName)
	}

	return
}

func (s *ExecHelperService) EvaluateAssertion(processor *model.ProcessorAssertion, parentLog *domain.ExecLog, msg *websocket.Message) (
	output domain.ExecOutput, err error) {

	output.Expression = processor.Expression

	return
}

func (s *ExecHelperService) EvaluateExtractor(extractor *model.ProcessorExtractor, parentLog *domain.ExecLog, msg *websocket.Message) (
	output domain.ExecOutput, err error) {

	brother, ok := getPreviousBrother(*parentLog)
	if !ok {
		output.Msg = fmt.Sprintf("先前节点不是接口，无法应用提取器。")
		return
	}

	resp := v1.InvocationResponse{}
	json.Unmarshal([]byte(brother.RespContent), &resp)

	interfaceExtractor := model.InterfaceExtractor{}
	copier.CopyWithOption(&interfaceExtractor, extractor, copier.Option{DeepCopy: true})

	interfaceExtractor.Src = consts.Body
	interfaceExtractor.Type = s.getExtractorTypeForProcessor(extractor.ProcessorType)

	err = s.ExtractorService.ExtractValue(&interfaceExtractor, resp)
	if err != nil {
		output.Msg = fmt.Sprintf("%s提取器解析错误 %s。", output.Type, err.Error())
		return
	}

	s.ExecContext.SetVariable(parentLog.ProcessId, extractor.Variable, interfaceExtractor.Result, false) // set in parent scope
	output.Msg = fmt.Sprintf("将结果\"%v\"赋予变量\"%s\"。", interfaceExtractor.Result, extractor.Variable)

	return
}

func (s *ExecHelperService) EvaluateCookie(processor *model.ProcessorCookie, parentLog *domain.ExecLog, msg *websocket.Message) (
	output domain.ExecOutput, err error) {

	cookieName := processor.CookieName
	variableName := processor.VariableName
	defaultValue := processor.Default
	domain := processor.Domain
	expireTime := processor.ExpireTime
	expression := processor.RightValue
	typ := processor.ProcessorType

	if typ == consts.ProcessorCookieSet {
		var variableValue interface{}
		variableValue, err = s.EvaluateGovaluateExpression(expression, processor.ProcessorID)
		if err != nil {
			output.Msg = fmt.Sprintf("计算表达式\"%s\"错误 %s。", expression, err.Error())
			return
		}

		s.ExecContext.SetCookie(parentLog.ProcessId, cookieName, variableValue, domain, expireTime) // set in parent scope
		output.Msg = fmt.Sprintf("%s为%v。", cookieName, variableValue)

	} else if typ == consts.ProcessorCookieGet {
		var variableValue interface{}
		cookie := s.ExecContext.GetCookie(parentLog.ProcessId, cookieName, domain)
		variableValue = cookie.Value

		words := ""
		if variableValue == nil && defaultValue != "" {
			variableValue, _ = execHelper.ParseValue(defaultValue)
			words = "默认"
		}

		if err != nil {
			output.Msg = fmt.Sprintf("获取Cookie %s的值错误 %s。", cookieName, err.Error())
			return
		}

		s.ExecContext.SetVariable(parentLog.ProcessId, variableName, variableValue, false) // set in parent scope
		output.Msg = fmt.Sprintf("将%s%s值%v赋予变量%s。", cookieName, words, variableValue, variableName)

	} else if typ == consts.ProcessorCookieClear {
		s.ExecContext.ClearCookie(parentLog.ProcessId, cookieName) // set in parent scope
		output.Msg = fmt.Sprintf("%s。", cookieName)
	}

	return
}

func (s *ExecHelperService) GetVariableValueByName(processorId uint, name string) (ret interface{}, err error) {
	vari, err := s.ExecContext.GetVariable(processorId, name)
	ret = vari.Value

	return
}

func (s *ExecHelperService) EvaluateGovaluateExpression(expression string, scopeId uint) (ret interface{}, err error) {
	expr := commUtils.RemoveLeftVariableSymbol(expression)

	valueExpression, err := govaluate.NewEvaluableExpressionWithFunctions(expr, expressionHelper.GovaluateFunctions)
	if err != nil {
		ret = expression
		return
	}

	parameters, err := s.generateParams(expression, scopeId)
	if err != nil {
		return
	}

	ret, err = valueExpression.Evaluate(parameters)

	return
}

func (s *ExecHelperService) generateParams(expression string, scopeId uint) (ret map[string]interface{}, err error) {
	ret = make(map[string]interface{}, 8)

	variables := execHelper.GetVariablesInVariablePlaceholder(expression)

	for _, variableName := range variables {
		var vari domain.ExecVariable
		vari, err = s.ExecContext.GetVariable(scopeId, variableName)
		if err == nil {
			ret[variableName] = vari.Value
		}
	}

	return
}

func (s *ExecHelperService) ReplaceVariablesWithVerbs(expression string) (ret string) {
	regx := regexp.MustCompile("(?siU)(\\${.*})")
	ret = regx.ReplaceAllString(expression, "%v")

	return
}

func (s *ExecHelperService) getExtractorTypeForProcessor(processorType consts.ProcessorType) (ret consts.ExtractorType) {
	if processorType == consts.ProcessorExtractorBoundary {
		ret = consts.Boundary
	} else if processorType == consts.ProcessorExtractorJsonQuery {
		ret = consts.JsonQuery
	} else if processorType == consts.ProcessorExtractorHtmlQuery {
		ret = consts.HtmlQuery
	} else if processorType == consts.ProcessorExtractorXmlQuery {
		ret = consts.XmlQuery
	}

	return
}

func getPreviousBrother(parent domain.ExecLog) (brother *domain.ExecLog, ok bool) {
	if len(*parent.Logs) > 0 {
		brother = (*parent.Logs)[len(*parent.Logs)-1]
		ok = true
		return
	}

	return
}
