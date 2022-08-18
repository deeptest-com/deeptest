package execHelper

import (
	"errors"
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	valueGen "github.com/aaronchen2k/deeptest/internal/server/modules/v1/helper/value"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/model"
	"github.com/kataras/iris/v12/websocket"
	"strconv"
	"strings"
)

//func ExecThreadGroup(processor *model.ProcessorThreadGroup, log *domain.Log, msg websocket.Message) (
//	result string, err error) {
//
//	return
//}

func ExecLogic(processor *model.ProcessorLogic, parentLog *domain.Log, msg websocket.Message) (
	result domain.Output, err error) {

	return
}

func ExecLoop(loop *model.ProcessorLoop, parentLog *domain.Log, msg websocket.Message) (
	output domain.Output, err error) {

	if loop.ID == 0 {
		output.Text = "执行前请先配置处理器。"
		return
	}

	typ := loop.ProcessorType
	if typ == consts.ProcessorLoopTime {
		output.Times = loop.Times
		output.Text = fmt.Sprintf("执行%d次。", output.Times)
		return
	} else if typ == consts.ProcessorLoopRange {
		output.Range = loop.Range
		output.Text = fmt.Sprintf("区间%s。", output.Range)
		return
	}

	return
}

func ExecData(processor *model.ProcessorData, parentLog *domain.Log, msg websocket.Message) (
	result domain.Output, err error) {

	return
}

func ExecTimer(processor *model.ProcessorTimer, parentLog *domain.Log, msg websocket.Message) (
	result domain.Output, err error) {

	return
}

func ExecVariable(processor *model.ProcessorVariable, parentLog *domain.Log, msg websocket.Message) (
	result domain.Output, err error) {

	return
}

func ExecAssertion(processor *model.ProcessorAssertion, parentLog *domain.Log, msg websocket.Message) (
	result domain.Output, err error) {

	return
}

func ExecExtractor(processor *model.ProcessorExtractor, parentLog *domain.Log, msg websocket.Message) (
	result domain.Output, err error) {

	return
}

func ExecCookie(processor *model.ProcessorCookie, parentLog *domain.Log, msg websocket.Message) (
	result domain.Output, err error) {

	return
}

func GetRange(output domain.Output, stepStr string) (start, end, step interface{}, precision int, typ consts.RangeType, err error) {
	rangeStr := output.Range
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
		typ = consts.RangeInt

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
		output.RangeEnd = endFloat
		step = stepFloat
		typ = consts.RangeFloat

		return
	}

	start = startStr
	end = endStr
	step = stepStr
	typ = consts.RangeString

	return
}

func GenerateRangeItems(start, end, step interface{}, precision int, isRand bool, typ consts.RangeType) (ret []interface{}, err error) {
	if typ == consts.RangeInt {
		ret = valueGen.GenerateIntItems(start.(int64), end.(int64), int(step.(int64)), isRand, 1, "")

	} else if typ == consts.RangeFloat {
		ret = valueGen.GenerateFloatItems(start.(float64), end.(float64), step.(float64), isRand, precision, 1, "")

	} else if typ == consts.RangeString {
		ret = valueGen.GenerateByteItems(start.(byte), end.(byte), step.(int), isRand, 1, "")

	}

	return
}
