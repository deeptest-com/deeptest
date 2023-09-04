package valueUtils

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"strconv"
)

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
