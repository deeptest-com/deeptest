package agentExec

import (
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	httpHelper "github.com/aaronchen2k/deeptest/internal/pkg/helper/http"
	_httpUtils "github.com/aaronchen2k/deeptest/pkg/lib/http"
	"strings"
)

func Invoke(req *v1.BaseRequest) (resp v1.DebugResponse, err error) {
	GetRequestProps(req)

	req.Url, err = _httpUtils.AddDefaultUrlSchema(req.Url)
	if err != nil {
		return
	}

	if req.Method == consts.GET {
		resp, err = httpHelper.Get(*req)
	} else if req.Method == consts.POST {
		resp, err = httpHelper.Post(*req)
	} else if req.Method == consts.PUT {
		resp, err = httpHelper.Put(*req)
	} else if req.Method == consts.DELETE {
		resp, err = httpHelper.Delete(*req)
	} else if req.Method == consts.PATCH {
		resp, err = httpHelper.Patch(*req)
	} else if req.Method == consts.HEAD {
		resp, err = httpHelper.Head(*req)
	} else if req.Method == consts.CONNECT {
		resp, err = httpHelper.Connect(*req)
	} else if req.Method == consts.OPTIONS {
		resp, err = httpHelper.Options(*req)
	} else if req.Method == consts.TRACE {
		resp, err = httpHelper.Trace(*req)
	}

	GetContentProps(&resp)

	return
}

func GetRequestProps(req *v1.BaseRequest) {
	req.BodyLang = consts.LangTEXT

	arr := strings.Split(string(req.BodyType), "/")
	if len(arr) == 1 {
		return
	}

	typeName := arr[1]
	if typeName == "text" || typeName == "plain" {
		typeName = consts.LangTEXT.String()
	}

	req.BodyLang = consts.HttpRespLangType(typeName)
}

func GetContentProps(resp *v1.DebugResponse) {
	resp.ContentLang = consts.LangTEXT

	if resp.ContentLang == "" {
		return
	}

	arr := strings.Split(string(resp.ContentType), ";")

	arr1 := strings.Split(arr[0], "/")
	if len(arr1) == 1 {
		return
	}

	typeName := arr1[1]
	if typeName == "text" || typeName == "plain" {
		typeName = "plaintext"
	}
	resp.ContentLang = consts.HttpRespLangType(typeName)

	if len(arr) > 1 {
		arr2 := strings.Split(arr[1], "=")
		if len(arr2) > 1 {
			resp.ContentCharset = consts.HttpRespCharset(arr2[1])
		}
	}

	//ret.NodeContent = mockHelper.FormatXml(ret.NodeContent)

	return
}

func ReplaceRequestWithVars(req *v1.BaseRequest) {
	replaceUrl(req)
	replaceParams(req)
	replaceHeaders(req)
	replaceBody(req)
	replaceAuthor(req)
}

func replaceUrl(req *v1.BaseRequest) {
	req.Url = ReplaceVariableValue(req.Url)
}
func replaceParams(req *v1.BaseRequest) {
	for idx, param := range req.Params {
		req.Params[idx].Value = ReplaceVariableValue(param.Value)
	}
}
func replaceHeaders(req *v1.BaseRequest) {
	for idx, header := range req.Headers {
		req.Headers[idx].Value = ReplaceVariableValue(header.Value)
	}
}
func replaceBody(req *v1.BaseRequest) {
	req.Body = ReplaceVariableValue(req.Body)
}
func replaceAuthor(req *v1.BaseRequest) {
	if req.AuthorizationType == consts.BasicAuth {
		req.BasicAuth.Username = ReplaceVariableValue(req.BasicAuth.Username)
		req.BasicAuth.Password = ReplaceVariableValue(req.BasicAuth.Password)

	} else if req.AuthorizationType == consts.BearerToken {
		req.BearerToken.Token = ReplaceVariableValue(req.BearerToken.Token)

	} else if req.AuthorizationType == consts.OAuth2 {
		req.OAuth20.Name = ReplaceVariableValue(req.OAuth20.Name)
		req.OAuth20.CallbackUrl = ReplaceVariableValue(req.OAuth20.CallbackUrl)
		req.OAuth20.AuthURL = ReplaceVariableValue(req.OAuth20.AuthURL)
		req.OAuth20.AccessTokenURL = ReplaceVariableValue(req.OAuth20.AccessTokenURL)
		req.OAuth20.ClientID = ReplaceVariableValue(req.OAuth20.ClientID)
		req.OAuth20.Scope = ReplaceVariableValue(req.OAuth20.Scope)

	} else if req.AuthorizationType == consts.ApiKey {
		req.ApiKey.Key = ReplaceVariableValue(req.ApiKey.Key)
		req.ApiKey.Value = ReplaceVariableValue(req.ApiKey.Value)
		req.ApiKey.TransferMode = ReplaceVariableValue(req.ApiKey.TransferMode)
	}
}
