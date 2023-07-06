package service

import (
	"fmt"
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	httpHelper "github.com/aaronchen2k/deeptest/internal/pkg/helper/http"
	"github.com/aaronchen2k/deeptest/internal/pkg/helper/openapi"
	"github.com/aaronchen2k/deeptest/internal/pkg/helper/openapi/convert"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	_domain "github.com/aaronchen2k/deeptest/pkg/domain"
	commonUtils "github.com/aaronchen2k/deeptest/pkg/lib/comm"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"go.uber.org/zap"
	"io/ioutil"
	"sync"
)

var l sync.RWMutex

type EndpointInterfaceService struct {
	EndpointInterfaceRepo *repo.EndpointInterfaceRepo `inject:""`
	EndpointService       *EndpointService            `inject:""`
}

func (s *EndpointInterfaceService) Paginate(req v1.EndpointInterfaceReqPaginate) (ret _domain.PageData, err error) {
	ret, err = s.EndpointInterfaceRepo.Paginate(req)
	return
}

func (s *EndpointInterfaceService) ImportEndpointData(req v1.ImportEndpointDataReq) (err error) {
	var data []byte
	if req.OpenUrlImport {
		request := domain.BaseRequest{Url: req.FilePath}
		var response domain.DebugResponse
		response, err = httpHelper.Get(request)
		data = []byte(response.Content)
	} else {
		data, err = ioutil.ReadFile(req.FilePath)
	}

	if err != nil {
		logUtils.Errorf("load end point data err ", zap.String("错误:", err.Error()))
		return err
	}

	req.DriverType, err = s.resetDriverType(req.DriverType, data)
	if err != nil {
		return
	}

	handler := convert.NewHandler(req.DriverType, data, req.FilePath)
	doc, err := handler.ToOpenapi()
	if err != nil {
		return err
	}

	openapi2endpoint := openapi.NewOpenapi2endpoint(doc, req.CategoryId)
	endpoints, dirs, components := openapi2endpoint.Convert()
	go s.EndpointService.SaveEndpoints(endpoints, dirs, components, req)

	return

}

func (s *EndpointInterfaceService) resetDriverType(driverType convert.DriverType, data []byte) (newDriverType convert.DriverType, err error) {
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
