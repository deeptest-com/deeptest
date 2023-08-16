package test

import (
	"fmt"
	responseDefineHelper "github.com/aaronchen2k/deeptest/internal/pkg/helper/responeDefine"
	_commUtils "github.com/aaronchen2k/deeptest/pkg/lib/comm"
	"testing"
)

func TestOpenapi2endpoint(t *testing.T) {
	json := `{"properties":{"chat_tool":{"allOf":[{"$ref":"#/components/schemas/enums.PlatType"}],"description":"第三方平台类型"},"data":{"description":"请求参数，json enums.ApproveParam","type":"string"},"server_type":{"allOf":[{"$ref":"#/components/schemas/enums.BusinessType"}],"description":"请求服务类型"},"sign":{"description":"参数签名","type":"string"}},"type":"object"}`
	responseDefineHelper.NewSchema2conv()
	var schema responseDefineHelper.Schema
	_commUtils.JsonDecode(json, &schema)
	x := _commUtils.JsonEncode(schema)
	fmt.Println(x, "+++")
	var schema1 responseDefineHelper.Schema
	_commUtils.JsonDecode(x, &schema1)
	y := _commUtils.JsonEncode(schema1)
	fmt.Println(y, "---")

	/*
		doc := new(openapi3.T)
		endpoint, _, _ := openapi.NewOpenapi2endpoint(doc).Convert()
		fmt.Println(endpoint)
		return
	*/

}
