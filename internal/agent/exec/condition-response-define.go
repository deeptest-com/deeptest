package agentExec

import (
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	responseDefineHelpper "github.com/aaronchen2k/deeptest/internal/pkg/helper/responeDefine"
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
	if !commonUtils.InArray(responseDefine.MediaType, mediaTypes) {
		responseDefine.ResultStatus = consts.Fail
		responseDefine.ResultMsg = fmt.Sprintf("响应体应该为%v，实际为%v", responseDefine.MediaType, res.ContentType)
		return
	}

	schema := new(responseDefineHelpper.SchemaRef)
	commonUtils.JsonDecode(responseDefine.Schema, schema)
	var obj interface{}
	commonUtils.JsonDecode(res.Content, &obj)
	newSchema2conv := responseDefineHelpper.NewSchema2conv()
	ret := newSchema2conv.AssertDataForSchema(schema, obj)
	if !ret {
		responseDefine.ResultStatus = consts.Fail
		responseDefine.ResultMsg = "返回数据结构与接口定义不一致"
		return
	}

	responseDefine.ResultMsg = "返回数据结构校验通过"

	//responseDefine.Schema
	return
}
