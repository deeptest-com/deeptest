package agentExec

import (
	"errors"
	"github.com/Knetic/govaluate"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	commUtils "github.com/aaronchen2k/deeptest/internal/pkg/utils"
	valueGen "github.com/aaronchen2k/deeptest/internal/server/modules/helper/value"
	_stringUtils "github.com/aaronchen2k/deeptest/pkg/lib/string"
	"regexp"
	"strconv"
	"strings"
)

var (
	GovaluateFunctions = map[string]govaluate.ExpressionFunction{
		"length": func(args ...interface{}) (interface{}, error) {
			length := len(args[0].(string))
			return (float64)(length), nil
		},
		"match": func(args ...interface{}) (interface{}, error) {
			str := args[0].(string)
			regx := regexp.MustCompile("args[1].(string)")

			ret := regx.MatchString(str)
			return ret, nil
		},
		"uuid": func(args ...interface{}) (interface{}, error) {
			ret := _stringUtils.Uuid()
			return ret, nil
		},
	}
)

var (
	executableContainerProcessors = []string{
		consts.ProcessorLogic.ToString(),
		consts.ProcessorLoop.ToString(),
		consts.ProcessorData.ToString(),
	}

	noExecutableContainerProcessors = []string{
		consts.ProcessorRoot.ToString(),
		//consts.ProcessorThreadGroup.ToString(),
		consts.ProcessorGroup.ToString(),
	}

	actionProcessors = []string{
		consts.ProcessorTimer.ToString(),
		consts.ProcessorPrint.ToString(),
		consts.ProcessorVariable.ToString(),
		consts.ProcessorAssertion.ToString(),
		consts.ProcessorExtractor.ToString(),
		consts.ProcessorCookie.ToString(),
	}
)

func IsExecutableContainerProcessor(processor *Processor) bool {
	return _stringUtils.FindInArr(processor.EntityCategory.ToString(), executableContainerProcessors) &&
		processor.EntityType != consts.ProcessorLoopBreak
}

func IsNoExecutableContainerProcessor(processor *Processor) bool {
	return _stringUtils.FindInArr(processor.EntityCategory.ToString(), noExecutableContainerProcessors)
}

func IsActionProcessor(processor *Processor) bool {
	return _stringUtils.FindInArr(processor.EntityCategory.ToString(), actionProcessors) ||
		processor.EntityType == consts.ProcessorLoopBreak
}

func IsInterfaceProcessor(processor *Processor) bool {
	return processor.EntityCategory == consts.ProcessorInterface
}

func EvaluateGovaluateExpression(expression string, scopeId uint) (ret interface{}, err error) {
	expr := commUtils.RemoveLeftVariableSymbol(expression)

	valueExpression, err := govaluate.NewEvaluableExpressionWithFunctions(expr, GovaluateFunctions)
	if err != nil {
		ret = expression
		return
	}

	parameters, err := generateParams(expression, scopeId)
	if err != nil {
		return
	}

	ret, err = valueExpression.Evaluate(parameters)

	return
}

func generateParams(expression string, scopeId uint) (ret map[string]interface{}, err error) {
	ret = make(map[string]interface{}, 8)

	variables := GetVariablesInVariablePlaceholder(expression)

	for _, variableName := range variables {
		var vari domain.ExecVariable
		vari, err = GetVariable(scopeId, variableName)
		if err == nil {
			ret[variableName] = vari.Value
		}
	}

	return
}

func GenerateRangeItems(start, end, step interface{}, precision int, isRand bool, typ consts.DataType) (ret []interface{}, err error) {
	if typ == consts.Int {
		ret = valueGen.GenerateIntItems(start.(int64), end.(int64), int(step.(int64)), isRand, 1, "")

	} else if typ == consts.Float {
		ret = valueGen.GenerateFloatItems(start.(float64), end.(float64), step.(float64), isRand, precision, 1, "")

	} else if typ == consts.String {
		ret = valueGen.GenerateByteItems(start.(byte), end.(byte), step.(int), isRand, 1, "")

	}

	return
}

func GenerateListItems(listStr string) (ret []interface{}, typ consts.DataType, err error) {
	arr := strings.Split(listStr, ",")

	isInt := true
	isFloat := true
	for _, item := range arr {
		_, intErr := strconv.ParseInt(item, 10, 64)
		if intErr != nil {
			isInt = false

			_, floatErr := strconv.ParseFloat(item, 64)
			if floatErr != nil {
				isFloat = false
			}
		}
	}

	if isInt {
		typ = consts.Int
		for _, item := range arr {
			intVal, _ := strconv.ParseInt(item, 10, 64)
			ret = append(ret, intVal)
		}
	} else if isFloat {
		typ = consts.Float
		for _, item := range arr {
			floatVal, _ := strconv.ParseFloat(item, 64)
			ret = append(ret, floatVal)
		}
	} else { // string
		typ = consts.String
		for _, item := range arr {
			ret = append(ret, item)
		}
	}

	return
}

func GetRange(rangeStr, stepStr string) (start, end, step interface{}, precision int, typ consts.DataType, err error) {
	if stepStr == "" {
		stepStr = "1"
	}

	arr := strings.Split(rangeStr, "-")
	if len(arr) < 1 {
		err = errors.New("range string not right")
		return
	}

	startStr := arr[0]
	endStr := arr[1]

	startInt, intErr1 := strconv.ParseInt(startStr, 10, 64)
	endInt, intErr2 := strconv.ParseInt(endStr, 10, 64)
	stepInt, intErr3 := strconv.ParseInt(stepStr, 10, 64)
	if intErr1 == nil && intErr2 == nil && intErr3 == nil {
		start = startInt
		end = endInt
		step = stepInt
		typ = consts.Int

		return
	}

	startFloat, floatErr1 := strconv.ParseFloat(startStr, 64)
	endFloat, floatErr2 := strconv.ParseFloat(endStr, 64)
	stepFloat, floatErr3 := strconv.ParseInt(endStr, 10, 64)
	if floatErr1 == nil && floatErr2 == nil && floatErr3 == nil {

		precisionStart, step1 := valueGen.GetPrecision(startFloat, stepFloat)
		precisionEnd, step2 := valueGen.GetPrecision(endFloat, stepFloat)
		if precisionStart < precisionEnd {
			precision = precisionEnd
			step = step2
		} else {
			precision = precisionStart
			step = step1
		}

		if (startFloat > endFloat && stepFloat > 0) || (startFloat < endFloat && stepFloat < 0) {
			step = -1 * stepFloat
		}

		start = startFloat
		end = endFloat
		step = stepFloat
		typ = consts.Float

		return
	}

	start = startStr
	end = endStr
	step = stepStr
	typ = consts.String

	return
}

func GetVariablesInVariablePlaceholder(expression string) (ret []string) {
	re := regexp.MustCompile("(?siU)\\${(.*)}")
	matchResultArr := re.FindAllStringSubmatch(expression, -1)

	for _, childArr := range matchResultArr {
		variableName := childArr[1]
		ret = append(ret, variableName)
	}

	return
}
