package service

import (
	"context"
	"fmt"
	"github.com/aaronchen2k/deeptest/cmd/agent/v1/domain"
	_httpUtils "github.com/aaronchen2k/deeptest/pkg/lib/http"
	"github.com/getkin/kin-openapi/openapi3"
	"regexp"
)

type SpecService struct {
}

func (s *SpecService) SubmitSpec(req domain.SubmitSpecReq) (doc3 *openapi3.T, info map[string]interface{}, err error) {
	ctx := context.Background()
	loader := &openapi3.Loader{Context: ctx, IsExternalRefsAllowed: true}

	if req.File != "" {
		doc3, err = loader.LoadFromFile(req.File)
	} else {
		var content []byte
		content, err = _httpUtils.Get(req.Url)
		if err != nil {
			return
		}

		doc3, err = loader.LoadFromData(content)
	}

	if err != nil {
		return
	}

	info = map[string]interface{}{}
	info["desc"] = s.parseDesc(doc3.Info.Description)

	return
}

func (s *SpecService) parseDesc(desc string) (ret [][]int) {
	compileRegex := regexp.MustCompile(`[\^\n]# (.+)`)
	arr := compileRegex.FindAllStringSubmatchIndex(desc, -1)

	str := desc[arr[0][2]:arr[0][3]]
	fmt.Sprintln(arr, str)

	for _, item := range arr {
		ret = append(ret, []int{item[2], item[3]})
	}

	if len(ret) == 0 {
		ret = append(ret, []int{0, 0})
	}

	return
}
