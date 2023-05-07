package valueUtils

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
)

func GenerateIntItems(start int64, end int64, step interface{}, rand bool, repeat int, repeatTag string) (ret []interface{}) {
	ret = generateIntItemsByStep(start, end, step.(int), repeat, repeatTag)

	if rand {
		ret = randItems(ret)
	}

	return
}

func generateIntItemsByStep(start int64, end int64, step int, repeat int, repeatTag string) []interface{} {
	arr := make([]interface{}, 0)

	total := 0
	if repeatTag == "" {
		for i := 0; true; {
			val := start + int64(i*step)
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

			if total >= consts.MaxNum {
				break
			}
			i++
		}
	} else if repeatTag == "!" {
		for round := 0; round < repeat; round++ {
			for i := 0; true; {
				val := start + int64(i*step)
				if (val > end && step > 0) || (val < end && step < 0) {
					break
				}

				arr = append(arr, val)

				if total >= consts.MaxNum {
					break
				}
				i++
			}

			if total >= consts.MaxNum {
				break
			}
		}
	}

	return arr
}
