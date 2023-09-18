package test

import (
	"fmt"
	responseDefineHelper "github.com/aaronchen2k/deeptest/internal/pkg/helper/schema"
	_commUtils "github.com/aaronchen2k/deeptest/pkg/lib/comm"
	"testing"
)

func TestOpenapi2endpoint(t *testing.T) {
	json := `{"#/components/schemas/_domain.Response":{"type":"object","properties":{"code":{"type":"integer"},"data":{},"msg":{"type":"string"}}},"#/components/schemas/consts.FormDataType":{"type":"string"},"#/components/schemas/consts.HttpContentType":{"type":"string"},"#/components/schemas/consts.ResultStatus":{"type":"string"},"#/components/schemas/domain.BearerToken":{"type":"object","properties":{"token":{"type":"string"}}},"#/components/schemas/domain.InterfaceToEnvMap":{"type":"object"},"#/components/schemas/serverConsts.ProjectType":{"type":"string"},"#/components/schemas/serverDomain.DataReq":{"type":"object","properties":{"clearData":{"type":"boolean"},"sys":{"ref":"#/components/schemas/serverDomain.DataSys"}}},"#/components/schemas/serverDomain.DataSys":{"type":"object","properties":{"adminPassword":{"type":"string"}}},"#/components/schemas/serverDomain.ProjectReq":{"type":"object","properties":{"adminId":{"type":"integer"},"adminName":{"type":"string"},"createdAt":{"type":"string"},"desc":{"type":"string"},"id":{"type":"integer"},"includeExample":{"type":"boolean"},"logo":{"type":"string"},"name":{"type":"string"},"orgId":{"type":"integer"},"schemaId":{"type":"integer"},"shortName":{"type":"string"},"type":{"ref":"#/components/schemas/serverConsts.ProjectType"},"updatedAt":{"type":"string"}}}}`
	//responseDefineHelper.NewSchema2conv()
	fmt.Println(json, "+++")
	var schema map[string]*responseDefineHelper.SchemaRef
	_commUtils.JsonDecode(json, &schema)
	/*
		schemaRef := new(responseDefineHelper.SchemaRef)
		schemaRef.Sample = new(responseDefineHelper.Schema)
		schemaRef.Sample.Type = "string"
		schema = map[string]*responseDefineHelper.SchemaRef{
			"/components/schemas/_domain.Response": schemaRef,
		}
	*/
	x := _commUtils.JsonEncode(schema)
	fmt.Println(x, "+++")
	var schema1 map[string]responseDefineHelper.SchemaRef
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
