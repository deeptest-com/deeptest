package service

import (
	"fmt"
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	httpHelper "github.com/aaronchen2k/deeptest/internal/pkg/helper/http"
	"github.com/aaronchen2k/deeptest/internal/pkg/helper/openapi"
	"github.com/aaronchen2k/deeptest/internal/pkg/helper/openapi/convert"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	_domain "github.com/aaronchen2k/deeptest/pkg/domain"
	commonUtils "github.com/aaronchen2k/deeptest/pkg/lib/comm"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"go.uber.org/zap"
	"io/ioutil"
)

type EndpointInterfaceService struct {
	EndpointInterfaceRepo *repo.EndpointInterfaceRepo `inject:""`
	EndpointService       *EndpointService            `inject:""`
}

func (s *EndpointInterfaceService) Paginate(tenantId consts.TenantId, req v1.EndpointInterfaceReqPaginate) (ret _domain.PageData, err error) {
	ret, err = s.EndpointInterfaceRepo.Paginate(tenantId, req)
	return
}

func (s *EndpointInterfaceService) ImportEndpointData(tenantId consts.TenantId, req v1.ImportEndpointDataReq) (err error) {
	var data []byte
	if req.OpenUrlImport {
		request := domain.BaseRequest{Url: req.FilePath}
		var response domain.DebugResponse
		response, err = httpHelper.Get(request, nil)
		data = []byte(response.Content)
	} else {
		data, err = ioutil.ReadFile(req.FilePath)
	}

	if err != nil {
		logUtils.Errorf("load end point data err ", zap.String("错误:", err.Error()))
		return err
	}

	req.DriverType, err = s.resetDriverType(tenantId, req.DriverType, data)
	if err != nil {
		return
	}

	handler := convert.NewHandler(req.DriverType, data, req.FilePath)
	doc, err := handler.ToOpenapi()
	if err != nil {
		logUtils.Errorf("load end point data err ", zap.String("错误:", err.Error()))
		return fmt.Errorf("文件格式错误")
	}

	openapi2endpoint := openapi.NewOpenapi2endpoint(doc, req.CategoryId)
	endpoints, dirs, components := openapi2endpoint.Convert()
	go s.EndpointService.SaveEndpoints(tenantId, endpoints, dirs, components, req)

	return

}

func (s *EndpointInterfaceService) resetDriverType(tenantId consts.TenantId, driverType convert.DriverType, data []byte) (newDriverType convert.DriverType, err error) {
	if driverType == convert.SWAGGER {
		res := make(map[string]interface{})
		err = commonUtils.JsonDecode(string(data), &res)
		if err != nil {
			return
		}

		if _, ok := res["openapi"].(string); ok {
			newDriverType = convert.SWAGGER3
			return
		}

		if _, ok := res["swagger"]; !ok {
			err = fmt.Errorf("file type error")
			return
		}

		if version, ok := res["swagger"].(string); ok && (version == "3.0" || version == "2.0") {
			newDriverType = convert.SWAGGER2
			if version == "3.0" {
				newDriverType = convert.SWAGGER3
			}
			return
		}

		err = fmt.Errorf("file type error")
		return
	}
	newDriverType = driverType
	return
}

func (s *EndpointInterfaceService) GenerateFromResponse(tenantId consts.TenantId, req v1.GenerateFromResponseReq) (responseBody model.EndpointInterfaceResponseBody, err error) {
	responseBody = model.EndpointInterfaceResponseBody{}
	responseBodyItem := model.EndpointInterfaceResponseBodyItem{}
	responseBody, err = s.EndpointInterfaceRepo.GetResponseBody(tenantId, req.InterfaceId, req.Code)
	if err == nil {
		responseBodyItem, err = s.EndpointInterfaceRepo.GetResponseBodyItem(tenantId, responseBody.ID)
		if err != nil {
			return
		}
	}
	responseBody.Code = req.Code
	responseBody.InterfaceId = req.InterfaceId
	responseBody.Description = req.Description
	responseBody.MediaType = req.ContentType
	responseBodyItem.Content = req.Data
	responseBody.SchemaItem = responseBodyItem

	err = s.EndpointInterfaceRepo.SaveResponseBody(tenantId, &responseBody)

	return

}

func (s *EndpointInterfaceService) GenerateFromRequest(tenantId consts.TenantId, req v1.GenerateFromRequestReq) (requestBody model.EndpointInterfaceRequestBody, err error) {
	requestBody = model.EndpointInterfaceRequestBody{}
	requestBodyItem := model.EndpointInterfaceRequestBodyItem{}
	requestBody, err = s.EndpointInterfaceRepo.GetRequestBody(tenantId, req.InterfaceId)
	if err == nil {
		requestBodyItem, err = s.EndpointInterfaceRepo.GetRequestBodyItem(tenantId, requestBody.ID)
		if err != nil {
			return
		}
	}
	requestBody.InterfaceId = req.InterfaceId
	requestBody.Description = req.Description
	requestBody.MediaType = req.ContentType
	requestBodyItem.Content = req.Data
	requestBody.SchemaItem = requestBodyItem

	err = s.EndpointInterfaceRepo.UpdateRequestBody(tenantId, &requestBody)

	return

}
