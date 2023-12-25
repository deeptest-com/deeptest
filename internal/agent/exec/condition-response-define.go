package agentExec

import (
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	responseDefineHelper "github.com/aaronchen2k/deeptest/internal/pkg/helper/schema"
	commonUtils "github.com/aaronchen2k/deeptest/pkg/lib/comm"
	"strings"

	"strconv"
)

func ExecResponseDefine(responseDefine *domain.ResponseDefineBase, res domain.DebugResponse) (err error) {
	if responseDefine.Code == "" {
		responseDefine.Code = "200"
	}
	code, err := strconv.Atoi(responseDefine.Code)
	if err != nil {
		panic(err)
	}

	responseDefine.ResultStatus = consts.Pass

	if code != int(res.StatusCode) {
		responseDefine.ResultStatus = consts.Fail
		responseDefine.ResultMsg = fmt.Sprintf("响应码应该为%v，实际为%v", code, res.StatusCode)
		return
	}

	mediaTypes := strings.Split(string(res.ContentType), ";")
	if responseDefine.MediaType != "" && !commonUtils.InArray(responseDefine.MediaType, mediaTypes) {
		responseDefine.ResultStatus = consts.Fail
		responseDefine.ResultMsg = fmt.Sprintf("响应体应该为%v，实际为%v", responseDefine.MediaType, res.ContentType)
		return
	}

	if responseDefine.Schema == "" {
		return
	}

	schema := new(responseDefineHelper.SchemaRef)
	commonUtils.JsonDecode(responseDefine.Schema, schema)
	var obj interface{}
	commonUtils.JsonDecode(res.Content, &obj)
	schema2conv := responseDefineHelper.NewSchema2conv()
	component := responseDefineHelper.NewComponents()
	commonUtils.JsonDecode(responseDefine.Component, component)
	schema2conv.Components = component
	ret := schema2conv.AssertDataForSchema(schema, obj)
	if !ret {
		responseDefine.ResultStatus = consts.Fail
		responseDefine.ResultMsg = "返回数据结构与接口定义不一致"
		return
	}

	responseDefine.ResultMsg = "返回数据结构校验通过"

	//responseDefine.Schema
	return
}
