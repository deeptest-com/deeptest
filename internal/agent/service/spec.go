package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	v1 "github.com/aaronchen2k/deeptest/cmd/agent/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	httpHelper "github.com/aaronchen2k/deeptest/internal/pkg/helper/http"
	_domain "github.com/aaronchen2k/deeptest/pkg/domain"
	_httpUtils "github.com/aaronchen2k/deeptest/pkg/lib/http"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"github.com/getkin/kin-openapi/openapi3"
	"regexp"
)

func ParseSpec(req v1.ParseSpecReq) (err error) {
	ctx := context.Background()
	loader := &openapi3.Loader{Context: ctx, IsExternalRefsAllowed: true}

	var doc3 *openapi3.T
	if req.File != "" { // from (converted) file
		doc3, err = loader.LoadFromFile(req.File)
	} else { // from url
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

	err = postSpecToServer(doc3, req)
	if err != nil {
		return
	}

	//desc := s.parseDesc(doc3.Info.Description)

	return
}

func postSpecToServer(doc3 *openapi3.T, req v1.ParseSpecReq) (err error) {
	url := fmt.Sprintf("import/importSpec")

	body, err := json.Marshal(doc3)
	if err != nil {
		logUtils.Infof("marshal request data failed, error, %s", err.Error())
		return
	}

	httpReq := domain.BaseRequest{
		Url: _httpUtils.AddSepIfNeeded(req.ServerUrl) + url,
		QueryParams: &[]domain.Param{
			{Name: "targetId", Value: fmt.Sprintf("%d", req.TargetId)},
		},
		BodyType:          consts.ContentTypeJSON,
		Body:              string(body),
		AuthorizationType: consts.BearerToken,
		BearerToken: domain.BearerToken{
			Token: req.Token,
		},
	}

	resp, err := httpHelper.Post(httpReq)
	if err != nil {
		return
	}

	if resp.StatusCode != consts.OK.Int() {
		err = errors.New(fmt.Sprintf("get interface obj failed, response %v", resp))
		return
	}

	respContent := _domain.Response{}
	json.Unmarshal([]byte(resp.Content), &respContent)
	if respContent.Code != 0 {
		err = errors.New(fmt.Sprintf("get interface obj failed, response %v", resp))
		return
	}

	return
}

func parseDesc(desc string) (ret [][]int) {
	compileRegex := regexp.MustCompile(`[\^\n]# (.+)`)
	arr := compileRegex.FindAllStringSubmatchIndex(desc, -1)

	if len(arr) > 0 {
		str := desc[arr[0][2]:arr[0][3]]
		fmt.Sprintln(arr, str)

		for _, item := range arr {
			ret = append(ret, []int{item[2], item[3]})
		}
	}

	if len(ret) == 0 {
		ret = append(ret, []int{0, 0})
	}

	return
}
