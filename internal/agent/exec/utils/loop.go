package agentUtils

import (
	"errors"
	valueUtils "github.com/aaronchen2k/deeptest/internal/agent/exec/utils/value"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"strconv"
	"strings"
)

func GenerateRangeItems(start, end, step interface{}, precision int, isRand bool, typ consts.DataType) (ret []interface{}, err error) {
	if typ == consts.Int {
		ret = valueUtils.GenerateIntItems(start.(int64), end.(int64), int(step.(int64)), isRand, 1, "")

	} else if typ == consts.Float {
		ret = valueUtils.GenerateFloatItems(start.(float64), end.(float64), step.(float64), isRand, precision, 1, "")

	} else if typ == consts.String {
		ret = valueUtils.GenerateByteItems(start.(byte), end.(byte), step.(int), isRand, 1, "")

	}

	return
}

func GenerateListItems(listStr string, isRand bool) (ret []interface{}, typ consts.DataType, err error) {
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

	if isRand {
		ret = valueUtils.RandItems(ret)
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

		precisionStart, step1 := valueUtils.GetPrecision(startFloat, stepFloat)
		precisionEnd, step2 := valueUtils.GetPrecision(endFloat, stepFloat)
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
