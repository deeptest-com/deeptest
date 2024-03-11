package agentExec

import (
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/helper/http"
	_httpUtils "github.com/aaronchen2k/deeptest/pkg/lib/http"
	_logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	_stringUtils "github.com/aaronchen2k/deeptest/pkg/lib/string"
	"strings"
)

func Invoke(req *domain.BaseRequest) (resp domain.DebugResponse, err error) {
	GetRequestProps(req)

	if DemoTestSite != "" {
		req.Url = _httpUtils.AddSepIfNeeded(DemoTestSite) + strings.ToLower(req.Method.String())

		notes := fmt.Sprintf("We change request url to %s on demo site.", req.Url)

		if req.QueryParams == nil {
			req.QueryParams = &[]domain.Param{}
		}
		*req.QueryParams = append(*req.QueryParams, domain.Param{
			Name:  "notes",
			Value: notes,
		})
		_logUtils.Infof(notes)
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

	GetContentProps(req, &resp)

	return
}

func GetRequestProps(req *domain.BaseRequest) {
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

func GetContentProps(req *domain.BaseRequest, resp *domain.DebugResponse) {
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
		typeName = "text"
	}

	if arr1[0] == "image" {
		typeName = strings.ToUpper(typeName)
	}
	resp.ContentLang = consts.HttpRespLangType(typeName)

	if len(arr) > 1 {
		arr2 := strings.Split(arr[1], "=")
		if len(arr2) > 1 {
			resp.ContentCharset = consts.HttpRespCharset(arr2[1])
		}
	}

	//resp.Content = mockHelper.FormatXml(resp.Content)

	fillCookieInHeader(req)

	return
}

func ReplaceVariables(req *domain.BaseRequest, execUuid string) {
	// 每个接口的局部参数覆盖全局参数
	mergeParams(req)

	replaceUrl(req, execUuid)

	replaceQueryParams(req, execUuid)
	replacePathParams(req, execUuid)
	replaceHeaders(req, execUuid)
	replaceCookies(req, execUuid)
	replaceFormBodies(req, execUuid)

	replaceBody(req, execUuid)
	replaceAuthor(req, execUuid)
}

func DealwithCookies(req *domain.BaseRequest, processorId uint, execUuid string) {
	if req.Cookies != nil {
		*req.Cookies = ListScopeCookie(processorId, execUuid)
	}
}

func replaceUrl(req *domain.BaseRequest, execUuid string) {
	// project's global params already be added
	req.Url = ReplaceVariableValue(req.Url, execUuid)
}
func replaceQueryParams(req *domain.BaseRequest, execUuid string) {
	if req.GlobalParams != nil {
		for _, p := range *req.GlobalParams {
			if !p.Disabled && p.In == consts.ParamInQuery {

				if req.QueryParams == nil {
					req.QueryParams = &[]domain.Param{}
				}
				*req.QueryParams = append(*req.QueryParams, domain.Param{
					Name:  p.Name,
					Value: p.DefaultValue,
				})
			}
		}
	}

	var queryParams []domain.Param

	if req.QueryParams != nil {
		for idx, param := range *req.QueryParams {
			if param.Disabled {
				continue
			}
			(*req.QueryParams)[idx].Value = ReplaceVariableValue(param.Value, execUuid)
			queryParams = append(queryParams, (*req.QueryParams)[idx])
		}
		req.QueryParams = &queryParams
	}

}

func replacePathParams(req *domain.BaseRequest, execUuid string) {
	var pathParams []domain.Param

	if req.PathParams != nil {
		for idx, param := range *req.PathParams {
			if param.Disabled || param.Name == "" {
				continue
			}
			(*req.PathParams)[idx].Value = ReplaceVariableValue(param.Value, execUuid)
			pathParams = append(pathParams, (*req.PathParams)[idx])
		}
		req.PathParams = &pathParams
	}

	return
}

func replaceHeaders(req *domain.BaseRequest, execUuid string) {
	if req.GlobalParams != nil {
		for _, p := range *req.GlobalParams {
			if p.In == consts.ParamInHeader && !p.Disabled {
				if req.Headers == nil {
					req.Headers = &[]domain.Header{}
				}
				*req.Headers = append(*req.Headers, domain.Header{
					Name:  p.Name,
					Value: p.DefaultValue,
				})
			}
		}
	}

	var headers []domain.Header
	if req.Headers != nil {
		for idx, header := range *req.Headers {
			if header.Disabled {
				continue
			}
			(*req.Headers)[idx].Value = ReplaceVariableValue(header.Value, execUuid)
			headers = append(headers, (*req.Headers)[idx])
		}
		req.Headers = &headers
	}

}
func replaceCookies(req *domain.BaseRequest, execUuid string) {
	if req.GlobalParams != nil {
		for _, p := range *req.GlobalParams {
			if p.In == consts.ParamInCookie && !p.Disabled {
				if req.Cookies == nil {
					req.Cookies = &[]domain.ExecCookie{}
				}

				*req.Cookies = append(*req.Cookies, domain.ExecCookie{
					Name:  p.Name,
					Value: p.DefaultValue,
				})
			}
		}
	}

	var cookies []domain.ExecCookie
	if req.Cookies != nil {
		for idx, cookie := range *req.Cookies {
			if cookie.Disabled {
				continue
			}
			(*req.Cookies)[idx].Value = ReplaceVariableValue(_stringUtils.InterfToStr(cookie.Value), execUuid)
			cookies = append(cookies, (*req.Cookies)[idx])
		}
		*req.Cookies = cookies
	}

}
func replaceFormBodies(req *domain.BaseRequest, execUuid string) {
	if req.GlobalParams != nil {
		for _, v := range *req.GlobalParams {
			if v.In == consts.ParamInBody && !v.Disabled {
				if req.BodyFormData == nil {
					req.BodyFormData = &[]domain.BodyFormDataItem{}
				}

				*req.BodyFormData = append(*req.BodyFormData, domain.BodyFormDataItem{
					Name:  v.Name,
					Value: v.DefaultValue,
				})

				if req.BodyFormUrlencoded == nil {
					req.BodyFormUrlencoded = &[]domain.BodyFormUrlEncodedItem{}
				}
				*req.BodyFormUrlencoded = append(*req.BodyFormUrlencoded, domain.BodyFormUrlEncodedItem{
					Name:  v.Name,
					Value: v.DefaultValue,
				})
			}
		}
	}
	if req.BodyFormData != nil {
		for idx, item := range *req.BodyFormData {
			(*req.BodyFormData)[idx].Value = ReplaceVariableValue(_stringUtils.InterfToStr(item.Value), execUuid)
		}
	}
	if req.BodyFormUrlencoded != nil {
		for idx, item := range *req.BodyFormUrlencoded {
			(*req.BodyFormUrlencoded)[idx].Value = ReplaceVariableValue(_stringUtils.InterfToStr(item.Value), execUuid)
		}
	}
}
func replaceBody(req *domain.BaseRequest, execUuid string) {
	req.Body = ReplaceVariableValueInBody(req.Body, execUuid)
}
func replaceAuthor(req *domain.BaseRequest, execUuid string) {
	if req.AuthorizationType == consts.BasicAuth {
		req.BasicAuth.Username = ReplaceVariableValue(req.BasicAuth.Username, execUuid)
		req.BasicAuth.Password = ReplaceVariableValue(req.BasicAuth.Password, execUuid)

	} else if req.AuthorizationType == consts.BearerToken {
		req.BearerToken.Token = ReplaceVariableValue(req.BearerToken.Token, execUuid)

	} else if req.AuthorizationType == consts.OAuth2 {
		req.OAuth20.Name = ReplaceVariableValue(req.OAuth20.Name, execUuid)
		req.OAuth20.CallbackUrl = ReplaceVariableValue(req.OAuth20.CallbackUrl, execUuid)
		req.OAuth20.AuthURL = ReplaceVariableValue(req.OAuth20.AuthURL, execUuid)
		req.OAuth20.AccessTokenURL = ReplaceVariableValue(req.OAuth20.AccessTokenURL, execUuid)
		req.OAuth20.ClientID = ReplaceVariableValue(req.OAuth20.ClientID, execUuid)
		req.OAuth20.Scope = ReplaceVariableValue(req.OAuth20.Scope, execUuid)

	} else if req.AuthorizationType == consts.ApiKey {
		req.ApiKey.Key = ReplaceVariableValue(req.ApiKey.Key, execUuid)
		req.ApiKey.Value = ReplaceVariableValue(req.ApiKey.Value, execUuid)
		req.ApiKey.TransferMode = ReplaceVariableValue(req.ApiKey.TransferMode, execUuid)
	}
}

func mergeParams(req *domain.BaseRequest) {
	if req.GlobalParams != nil {
		for key, globalParam := range *req.GlobalParams {
			if globalParam.In == consts.ParamInQuery {
				if req.QueryParams != nil {
					for _, item := range *req.QueryParams {
						if item.Name == globalParam.Name && !item.Disabled {
							(*req.GlobalParams)[key].Disabled = true
						}
					}
				}
			} else if globalParam.In == consts.ParamInHeader {
				if req.Headers != nil {
					for _, item := range *req.Headers {
						if item.Name == globalParam.Name && !item.Disabled {
							(*req.GlobalParams)[key].Disabled = true
						}
					}
				}
			} else if globalParam.In == consts.ParamInCookie {
				if req.Cookies != nil {
					for _, item := range *req.Cookies {
						if item.Name == globalParam.Name && !item.Disabled {
							(*req.GlobalParams)[key].Disabled = true
						}
					}
				}
			}
		}
	}

}

func MergeGlobalParams(globalParams, selfGlobalParam []domain.GlobalParam) (ret []domain.GlobalParam) {

	ret = globalParams
	for key, globalParam := range ret {
		for _, param := range selfGlobalParam {
			if param.Name == globalParam.Name && param.In == globalParam.In {
				ret[key].Disabled = param.Disabled
			}
		}
	}

	return
}

func fillCookieInHeader(req *domain.BaseRequest) {
	var cookies = ""

	if req.Cookies != nil {
		for _, cookie := range *req.Cookies {
			if cookie.Name == "" || cookie.Value == "" {
				continue
			}
			if cookies == "" {
				cookies += fmt.Sprintf("%s=%s", cookie.Name, cookie.Value)
			} else {
				cookies += fmt.Sprintf(";%s=%s", cookie.Name, cookie.Value)
			}
		}
	}

	if cookies != "" {
		if req.Headers == nil {
			req.Headers = &[]domain.Header{}
		}
		*req.Headers = append(*req.Headers, domain.Header{Name: "Cookie", Value: cookies})
	}

}
