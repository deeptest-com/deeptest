package queryHelper

import (
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/model"
	"regexp"
	"strings"
)

func BoundaryQuery(content string, extractor *model.InterfaceExtractor) {
	regex := regexp.MustCompile(fmt.Sprintf("(?Ui)%s(.*)%s", extractor.BoundaryStart, extractor.BoundaryEnd))

	arrOfArr := regex.FindAllStringSubmatch(content, -1)

	results := make([]string, 0)
	for _, arr := range arrOfArr {
		result := ""
		if extractor.BoundaryIncluded {
			result = arr[0]
		} else {
			result = arr[1]
		}

		results = append(results, result)
	}

	if extractor.BoundaryIndex > 0 {
		if extractor.BoundaryIndex > len(results)-1 {
			extractor.BoundaryIndex = len(results) - 1
		}
		extractor.Result = results[extractor.BoundaryIndex]
	} else {
		extractor.Result = strings.Join(results, ", ")
	}
}
