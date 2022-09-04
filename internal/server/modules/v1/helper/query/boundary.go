package queryHelper

import (
	"fmt"
	"regexp"
	"strings"
)

func BoundaryQuery(content string, boundaryStart, boundaryEnd string, boundaryIndex int, boundaryIncluded bool) (result string) {
	regex := regexp.MustCompile(fmt.Sprintf("(?Ui)%s(.*)%s", boundaryStart, boundaryEnd))

	arrOfArr := regex.FindAllStringSubmatch(content, -1)

	results := make([]string, 0)
	for _, arr := range arrOfArr {
		result := ""
		if boundaryIncluded {
			result = arr[0]
		} else {
			result = arr[1]
		}

		results = append(results, result)
	}

	if boundaryIndex > 0 {
		if boundaryIndex > len(results)-1 {
			boundaryIndex = len(results) - 1
		}
		result = results[boundaryIndex]
	} else {
		result = strings.Join(results, ", ")
	}

	return
}
