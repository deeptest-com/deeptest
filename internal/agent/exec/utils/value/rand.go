package valueUtils

import _intUtils "github.com/aaronchen2k/deeptest/pkg/lib/int"

func RandItems(items []interface{}) (ret []interface{}) {
	indexArr := _intUtils.GenUniqueRandNum(0, len(items), len(items))

	for _, item := range indexArr {
		ret = append(ret, items[item])
	}

	return
}
