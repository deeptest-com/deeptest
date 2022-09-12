package service

import (
	"encoding/json"
	"fmt"
	"github.com/Knetic/govaluate"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/business"
	serverDomain "github.com/aaronchen2k/deeptest/internal/server/modules/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/repo"
	"github.com/jinzhu/copier"
	"github.com/kataras/iris/v12/websocket"
	"regexp"
)

type ExecHelperService struct {
	ScenarioProcessorRepo *repo.ScenarioProcessorRepo `inject:""`
	ScenarioRepo          *repo.ScenarioRepo          `inject:""`
	TestResultRepo        *repo.ReportRepo            `inject:""`
	TestLogRepo           *repo.LogRepo               `inject:""`
	InterfaceRepo         *repo.InterfaceRepo         `inject:""`
	InterfaceService      *InterfaceService           `inject:""`
	ExecRequestService    *business.ExecRequest       `inject:""`
	ExecContext           *business.ExecContext       `inject:""`
	ExtractorService      *ExtractorService           `inject:""`
}

//func (s *ExecComm) ParseThreadGroup(processor *model.ProcessorThreadGroup, log *domain.ExecLog, msg *websocket.Message) (
//	result string, err error) {
//
//	return
//}

func (s *ExecHelperService) ParseLogic(logic *model.ProcessorLogic, parentLog *domain.ExecLog, msg *websocket.Message) (
	output domain.ExecOutput, err error) {
	if logic.ID == 0 {
		output.Msg = "执行前请先配置处理器。"
		return
	}

	typ := logic.ProcessorType

	if typ == consts.ProcessorLogicIf {
		var result interface{}
		result, err = s.ComputerExpress(logic.Expression, logic.ProcessorId)
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

func (s *ExecHelperService) ParseLoop(loop *model.ProcessorLoop, parentLog *domain.ExecLog, msg *websocket.Message) (
	output domain.ExecOutput, err error) {

	if loop.ID == 0 {
		output.Msg = "执行前请先配置处理器。"
		return
	}

	typ := loop.ProcessorType
	if typ == consts.ProcessorLoopTime {
		output.Times = loop.Times
		output.Msg = fmt.Sprintf("执行%d次。", output.Times)
		return
	} else if typ == consts.ProcessorLoopUntil {
		output.Expression = loop.UntilExpression
		output.Msg = fmt.Sprintf("直到%s。", output.Expression)
		return
	} else if typ == consts.ProcessorLoopIn {
		output.List = loop.List
		output.Msg = fmt.Sprintf("迭代列表%s。", output.List)
		return
	} else if typ == consts.ProcessorLoopRange {
		output.Range = loop.Range
		output.Msg = fmt.Sprintf("区间%s。", output.Range)
		return
	}

	return
}

func (s *ExecHelperService) ParseLoopBreak(loop *model.ProcessorLoop, parentLog *domain.ExecLog, msg *websocket.Message) (
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

func (s *ExecHelperService) ParseData(data *model.ProcessorData, parentLog *domain.ExecLog, msg *websocket.Message) (
	output domain.ExecOutput, err error) {
	output.Url = data.Url

	output.Msg = fmt.Sprintf("使用数据%s迭代变量%s。", output.Url, data.VariableName)

	return
}

func (s *ExecHelperService) ParseTimer(processor *model.ProcessorTimer, parentLog *domain.ExecLog, msg *websocket.Message) (
	output domain.ExecOutput, err error) {

	output.SleepTime = processor.SleepTime
	output.Msg = fmt.Sprintf("等待%d秒。", output.SleepTime)

	return
}

func (s *ExecHelperService) ParseVariable(processor *model.ProcessorVariable, parentLog *domain.ExecLog, msg *websocket.Message) (
	output domain.ExecOutput, err error) {

	output.VariableName = processor.VariableName
	expression := processor.RightValue
	typ := processor.ProcessorType

	if typ == consts.ProcessorVariableSet {
		output.VariableValue, err = s.ComputerExpress(expression, processor.ProcessorId)
		if err != nil {
			output.Msg = fmt.Sprintf("计算表达式子错误 %s。", err.Error())
			return
		}

		output.Msg = fmt.Sprintf("%s为%v。", output.VariableName, output.VariableValue)

	} else if typ == consts.ProcessorVariableClear {
		s.ExecContext.ClearVariable(parentLog.ProcessId, output.VariableName) // set in parent scope
		output.Msg = fmt.Sprintf("%s。", output.VariableName)
	}

	return
}

func (s *ExecHelperService) ParseAssertion(processor *model.ProcessorAssertion, parentLog *domain.ExecLog, msg *websocket.Message) (
	output domain.ExecOutput, err error) {

	expression := processor.Expression
	result, err := s.ComputerExpress(expression, processor.ProcessorId)
	output.Pass, _ = result.(bool)

	status := "失败"
	if output.Pass {
		status = "通过"
	}

	output.Msg = fmt.Sprintf("断言表达式'%s'结果为%s。", expression, status)

	return
}

func (s *ExecHelperService) ParseExtractor(extractor *model.ProcessorExtractor, parentLog *domain.ExecLog, msg *websocket.Message) (
	output domain.ExecOutput, err error) {

	brother, ok := getPreviousBrother(*parentLog)
	if !ok {
		output.Msg = fmt.Sprintf("前面节点不是接口，无法应用提取器。")
		return
	}

	resp := serverDomain.InvocationResponse{}
	json.Unmarshal([]byte(brother.RespContent), &resp)

	interfaceExtractor := model.InterfaceExtractor{}
	copier.CopyWithOption(&interfaceExtractor, extractor, copier.Option{DeepCopy: true})
	err = s.ExtractorService.ExtractValue(&interfaceExtractor, resp)
	if err != nil {
		output.Msg = fmt.Sprintf("%s提取器解析错误%s。", output.Type, err.Error())
		return
	}

	s.ExecContext.SetVariable(parentLog.ProcessId, extractor.Variable, extractor.Result) // set in parent scope
	output.Msg = fmt.Sprintf("将提取器%s的结果%v赋予变量%s。", output.Type, extractor.Result, extractor.Variable)

	return
}

func (s *ExecHelperService) ParseCookie(processor *model.ProcessorCookie, parentLog *domain.ExecLog, msg *websocket.Message) (
	output domain.ExecOutput, err error) {

	cookieName := processor.CookieName
	variableName := processor.VariableName
	domain := processor.Domain
	expireTime := processor.ExpireTime
	expression := processor.RightValue
	typ := processor.ProcessorType

	if typ == consts.ProcessorCookieSet {
		var variableValue interface{}
		variableValue, err = s.ComputerExpress(expression, processor.ProcessorId)
		if err != nil {
			output.Msg = fmt.Sprintf("计算表达式子错误 %s。", err.Error())
			return
		}

		s.ExecContext.SetCookie(parentLog.ProcessId, cookieName, variableValue, domain, expireTime) // set in parent scope
		output.Msg = fmt.Sprintf("设置Cookie %s的值为%v。", cookieName, variableValue)

	} else if typ == consts.ProcessorCookieGet {
		var variableValue interface{}
		variableValue = s.ExecContext.GetCookie(parentLog.ProcessId, cookieName, domain)
		if err != nil {
			output.Msg = fmt.Sprintf("获取Cookie %s的值错误 %s。", cookieName, err.Error())
			return
		}

		s.ExecContext.SetVariable(parentLog.ProcessId, variableName, variableValue) // set in parent scope
		output.Msg = fmt.Sprintf("获取Cookie %s的值%s，赋予变量%s。", cookieName, variableValue, variableName)

	} else if typ == consts.ProcessorCookieClear {
		s.ExecContext.ClearCookie(parentLog.ProcessId, cookieName) // set in parent scope
		output.Msg = fmt.Sprintf("清除Cookie%s。", cookieName)
	}

	return
}

func (s *ExecHelperService) GetVariableValueByName(processorId uint, name string) (ret interface{}, err error) {
	vari, err := s.ExecContext.GetVariable(processorId, name)
	ret = vari.Value

	return
}

func (s *ExecHelperService) ComputerExpress(expression string, scopeId uint) (ret interface{}, err error) {
	// remove variable symbol ${} for govaluate
	re := regexp.MustCompile("(?siU)\\${(.*)}")
	expr := re.ReplaceAllString(expression, "$1")

	valueExpression, err := govaluate.NewEvaluableExpression(expr)
	if err != nil {
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

	variables := getVariables(expression)

	for _, variableName := range variables {
		var vari domain.ExecVariable
		vari, err = s.ExecContext.GetVariable(scopeId, variableName)
		if err == nil {
			ret[variableName] = vari.Value
		}
	}

	return
}

func getVariables(expression string) (ret []string) {
	re := regexp.MustCompile("(?siU)\\${(.*)}")
	matchResultArr := re.FindAllStringSubmatch(expression, -1)

	for _, childArr := range matchResultArr {
		variableName := childArr[1]
		ret = append(ret, variableName)
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
