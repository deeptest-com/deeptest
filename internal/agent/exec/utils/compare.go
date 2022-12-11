package utils

import (
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"strconv"
	"strings"
)

func Compare(operator consts.ComparisonOperator, actual, expect interface{}) (result consts.ResultStatus) {

	if operator == consts.Equal {
		if actual == expect {
			result = consts.Pass
		} else {
			result = consts.Fail
		}
	} else if operator == consts.NotEqual {
		if actual != expect {
			result = consts.Pass
		} else {
			result = consts.Fail
		}
	} else if operator == consts.Contain {
		if strings.Contains(interfaceToString(actual), interfaceToString(expect)) {
			result = consts.Pass
		} else {
			result = consts.Fail
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

func interfaceToString(interf interface{}) (ret string) {
	ret = fmt.Sprintf("%v", interf)

	return
}
