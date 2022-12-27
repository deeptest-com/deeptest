package queryHelper

import (
	"regexp"
)

func RegxQuery(content string, expr string) (result string) {
	expr = "(?m)" + expr
	regx := regexp.MustCompile(expr)

	arr := regx.FindAllStringSubmatch(content, -1)

	if len(arr) == 0 {
		return
	}

	result = arr[0][1]

	return
}
