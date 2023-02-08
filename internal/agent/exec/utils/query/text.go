package queryUtils

import (
	"regexp"
)

func RegxQuery(content string, expr string) (result string) {
	expr = "(?m)" + expr
	regx := regexp.MustCompile(expr)

	arr := regx.FindAllStringSubmatch(content, -1)

	if len(arr) == 0 || len(arr[0]) < 2 {
		return
	}

	result = arr[0][1]

	return
}
