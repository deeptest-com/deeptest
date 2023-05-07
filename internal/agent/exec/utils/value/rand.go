package valueUtils

func RandItems(items []interface{}) (ret []interface{}) {
	length := len(items)

	for i := 0; i < len(items); i++ {
		ret = append(ret, items[RandNum(10000)%length])
	}

	return
}
