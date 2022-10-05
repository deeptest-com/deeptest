package execHelper

import (
	"errors"
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	valueGen2 "github.com/aaronchen2k/deeptest/internal/server/modules/helper/value"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	fileUtils "github.com/aaronchen2k/deeptest/pkg/lib/file"
	"github.com/xuri/excelize/v2"
	"regexp"
	"strconv"
	"strings"
)

func Compare(operator consts.ComparisonOperator, actual, expect interface{}) (result consts.ResultStatus) {

	if operator == consts.Equal {
		if actual == expect {
			result = consts.Pass
		}
	} else if operator == consts.NotEqual {
		if actual != expect {
			result = consts.Pass
		}
	} else if operator == consts.Contain {
		if strings.Contains(interfaceToString(actual), interfaceToString(expect)) {
			result = consts.Pass
		}
	} else {
		result = CompareAsFloat(operator, actual, expect)
	}

	return
}

func CompareAsFloat(operator consts.ComparisonOperator, actual, expect interface{}) (
	result consts.ResultStatus) {

	result = consts.Fail

	actualFloat, err1 := strconv.ParseFloat(interfaceToString(actual), 64)
	expectFloat, err2 := strconv.ParseFloat(interfaceToString(expect), 64)

	if err1 != nil || err2 != nil { // not a number
		return
	}

	switch operator.String() {
	case consts.GreaterThan.String():
		result = GetResult(actualFloat > expectFloat)

	case consts.LessThan.String():
		result = GetResult(actualFloat < expectFloat)

	case consts.GreaterThanOrEqual.String():
		result = GetResult(actualFloat >= expectFloat)

	case consts.LessThanOrEqual.String():
		result = GetResult(actualFloat <= expectFloat)

	default:

	}

	return
}

func GetResult(b bool) (
	result consts.ResultStatus) {

	if b {
		result = consts.Pass
	} else {
		result = consts.Fail
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

		precisionStart, step1 := valueGen2.GetPrecision(startFloat, stepFloat)
		precisionEnd, step2 := valueGen2.GetPrecision(endFloat, stepFloat)
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

func GenerateRangeItems(start, end, step interface{}, precision int, isRand bool, typ consts.DataType) (ret []interface{}, err error) {
	if typ == consts.Int {
		ret = valueGen2.GenerateIntItems(start.(int64), end.(int64), int(step.(int64)), isRand, 1, "")

	} else if typ == consts.Float {
		ret = valueGen2.GenerateFloatItems(start.(float64), end.(float64), step.(float64), isRand, precision, 1, "")

	} else if typ == consts.String {
		ret = valueGen2.GenerateByteItems(start.(byte), end.(byte), step.(int), isRand, 1, "")

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

func GenerateDataItems(data model.ProcessorData) (ret []map[string]interface{}, err error) {
	url := data.Url

	if data.ProcessorType == consts.ProcessorDataText {
		ret, err = readDataFromText(url, data.Separator)
	} else if data.ProcessorType == consts.ProcessorDataExcel {
		ret, err = readDataFromExcel(url)
	}

	return
}
func readDataFromText(url, separator string) (ret []map[string]interface{}, err error) {
	content := fileUtils.ReadFile(url)
	arr := strings.Split(strings.ReplaceAll(content, "\r\n", "\n"), "\n")
	if len(arr) < 2 {
		return
	}

	colNameMap := map[int]string{}
	cols := strings.Split(arr[0], separator)
	for index, col := range cols {
		colNameMap[index] = col
	}

	for index, line := range arr {
		if index == 0 {
			continue
		}

		cols := strings.Split(line, separator)
		mp := map[string]interface{}{}
		for index, col := range cols {
			mp[colNameMap[index]] = col
		}
		ret = append(ret, mp)
	}

	return
}
func readDataFromExcel(url string) (ret []map[string]interface{}, err error) {
	excel, err := excelize.OpenFile(url)
	if err != nil {
		return
	}

	if len(excel.GetSheetList()) == 0 {
		return
	}

	firstSheet := excel.GetSheetList()[0]

	rows, err := excel.GetRows(firstSheet)
	if len(rows) < 2 {
		return
	}

	colNameMap := map[int]string{}
	for index, col := range rows[0] {
		col = strings.Replace(col, "'", "''", -1)
		colNameMap[index] = col
	}

	for rowIndex, row := range rows {
		if rowIndex == 0 {
			continue
		}

		mp := map[string]interface{}{}
		for index, col := range row {
			col = strings.Replace(col, "'", "''", -1)
			mp[colNameMap[index]] = col
		}
		ret = append(ret, mp)
	}

	return
}

func interfaceToString(interf interface{}) (ret string) {
	ret = fmt.Sprintf("%v", interf)

	return
}

func ParseValue(str string) (ret interface{}, typ consts.DataType) {
	isInt := true
	isFloat := true

	intVal := int64(0)
	floatVal := float64(0)

	intVal, intErr := strconv.ParseInt(str, 10, 64)
	if intErr != nil {
		isInt = false

		var floatErr error
		floatVal, floatErr = strconv.ParseFloat(str, 64)
		if floatErr != nil {
			isFloat = false
		}
	}

	if isInt {
		ret = intVal
	} else if isFloat {
		ret = floatVal
	} else {
		ret = str
	}

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
