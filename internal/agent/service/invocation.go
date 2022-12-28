package service

import (
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	httpHelper "github.com/aaronchen2k/deeptest/internal/pkg/helper/http"
	"github.com/aaronchen2k/deeptest/pkg/lib/http"
	"strings"
)

type InvocationService struct {
}

func (s *InvocationService) Test(req v1.InvocationRequest) (ret v1.InvocationResponse, err error) {
	req.Url, err = _httpUtils.AddDefaultUrlSchema(req.Url)
	if err != nil {
		return
	}

	if req.Method == consts.GET {
		ret, err = httpHelper.Get(req.BaseRequest)
	} else if req.Method == consts.POST {
		ret, err = httpHelper.Post(req.BaseRequest)
	} else if req.Method == consts.PUT {
		ret, err = httpHelper.Put(req.BaseRequest)
	} else if req.Method == consts.DELETE {
		ret, err = httpHelper.Delete(req.BaseRequest)
	} else if req.Method == consts.PATCH {
		ret, err = httpHelper.Patch(req.BaseRequest)
	} else if req.Method == consts.HEAD {
		ret, err = httpHelper.Head(req.BaseRequest)
	} else if req.Method == consts.CONNECT {
		ret, err = httpHelper.Connect(req.BaseRequest)
	} else if req.Method == consts.OPTIONS {
		ret, err = httpHelper.Options(req.BaseRequest)
	} else if req.Method == consts.TRACE {
		ret, err = httpHelper.Trace(req.BaseRequest)
	}

	s.GetContentProps(&ret)

	ret.Id = req.Id

	return
}

func (s *InvocationService) GetContentProps(ret *v1.InvocationResponse) {
	ret.ContentLang = consts.LangTEXT

	if ret.ContentLang == "" {
		return
	}

	arr := strings.Split(string(ret.ContentType), ";")
	arr1 := strings.Split(arr[0], "/")
	if len(arr1) == 1 {
		return
	}

	typeName := arr1[1]
	if typeName == "text" || typeName == "plain" {
		typeName = consts.LangTEXT.String()
	}
	ret.ContentLang = consts.HttpRespLangType(typeName)

	if len(arr) > 1 {
		arr2 := strings.Split(arr[1], "=")
		if len(arr2) > 1 {
			ret.ContentCharset = consts.HttpRespCharset(arr2[1])
		}
	}

	//ret.NodeContent = mockHelper.FormatXml(ret.NodeContent)

	return
}
