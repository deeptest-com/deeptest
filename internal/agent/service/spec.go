package service

import (
	"context"
	"fmt"
	"github.com/aaronchen2k/deeptest/cmd/agent/v1/domain"
	"github.com/getkin/kin-openapi/openapi3"
	"regexp"
)

type SpecService struct {
}

func (s *SpecService) SubmitSpec(req domain.SubmitSpecReq) (doc3 *openapi3.T, mp map[string]string, err error) {
	ctx := context.Background()
	loader := &openapi3.Loader{Context: ctx, IsExternalRefsAllowed: true}

	doc3, err = loader.LoadFromFile(req.File)

	mp = map[string]string{}
	mp["desc"] = s.parseDesc(doc3.Info.Description)

	if err != nil {
		return
	}

	return
}

func (s *SpecService) parseDesc(desc string) (ret string) {
	compileRegex := regexp.MustCompile(`[\^\n]# (.+)`)
	arr := compileRegex.FindAllStringSubmatchIndex(desc, -1)

	str := desc[arr[0][2]:arr[0][3]]

	fmt.Sprintln(arr, str)

	return
}
