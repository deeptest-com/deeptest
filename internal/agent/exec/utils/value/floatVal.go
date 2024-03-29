package valueUtils

import (
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"strconv"
	"strings"
)

func GenerateFloatItems(start float64, end float64, step float64, rand bool, precision, repeat int, repeatTag string) (
	ret []interface{}) {

	if start > end && step > 0 || start < end && step < 0 {
		step = step * -1
	}

	ret = generateFloatItemsByStep(start, end, step, precision, repeat, repeatTag)

	if rand {
		ret = RandItems(ret)
	}

	return
}

func generateFloatItemsByStep(start float64, end float64, step float64, precision, repeat int, repeatTag string) []interface{} {
	arr := make([]interface{}, 0)

	total := 0

	if repeatTag == "" {
		for i := 0; true; {
			val := start + float64(i)*step
			val = ChangePrecision(val, precision)

			if (val > end && step > 0) || (val < end && step < 0) {
				break
			}

			for round := 0; round < repeat; round++ {
				arr = append(arr, val)

				total++
				if total > consts.MaxNum {
					break
				}
			}
			if total > consts.MaxNum {
				break
			}
			i++
		}
	} else if repeatTag == "!" {
		for round := 0; round < repeat; round++ {
			for i := 0; true; {
				val := start + float64(i)*step
				val = ChangePrecision(val, precision)

				if (val > end && step > 0) || (val < end && step < 0) {
					break
				}

				arr = append(arr, val)

				if total > consts.MaxNum {
					break
				}
				i++
			}

			total++
			if total > consts.MaxNum {
				break
			}
		}
	}

	return arr
}

func GetPrecision(base float64, step interface{}) (precision int, newStep float64) {
	var flt float64 = 1
	if step == nil {
		step = flt
	}
	str1 := strconv.FormatFloat(base, 'f', -1, 64)
	var stepFloat float64 = 1
	switch step.(type) {
	case float64:
		stepFloat = step.(float64)
	case int:
		stepFloat = float64(step.(int))
	}
	str2 := strconv.FormatFloat(stepFloat, 'f', -1, 64)

	index1 := strings.LastIndex(str1, ".")
	index2 := strings.LastIndex(str2, ".")

	if index1 < index2 {
		precision = len(str1) - index1 - 1
	} else {
		precision = len(str2) - index2 - 1
	}

	if step == nil || step == 0 {
		newStep = float64(1)
		for i := 0; i < precision; i++ {
			newStep = newStep / 10
		}
	} else {
		switch step.(type) {
		case float64:
			newStep = step.(float64)
		case int:
			newStep = float64(step.(int))
		}
	}

	return
}

func InterfaceToStr(val interface{}) (ret string) {
	ret = fmt.Sprintf("%v", val)
	return
}

func ChangePrecision(flt float64, precision int) float64 {
	format := fmt.Sprintf("%%.%df", precision)
	ret, _ := strconv.ParseFloat(fmt.Sprintf(format, flt), 64)
	return ret
}
