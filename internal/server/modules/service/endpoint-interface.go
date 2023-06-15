package service

import (
	"encoding/json"
	"fmt"
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/helper/openapi"
	"github.com/aaronchen2k/deeptest/internal/pkg/helper/openapi/convert"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	_domain "github.com/aaronchen2k/deeptest/pkg/domain"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"go.uber.org/zap"
	"io/ioutil"
)

type EndpointInterfaceService struct {
	EndpointInterfaceRepo *repo.EndpointInterfaceRepo `inject:""`
	EndpointService       *EndpointService            `inject:""`
}

func NewEndpointInterfaceService() *EndpointInterfaceService {
	return &EndpointInterfaceService{}
}

func (s *EndpointInterfaceService) Paginate(req v1.EndpointInterfaceReqPaginate) (ret _domain.PageData, err error) {
	ret, err = s.EndpointInterfaceRepo.Paginate(req)
	return
}

func (s *EndpointInterfaceService) ImportEndpointData(req v1.ImportEndpointDataReq) (err error) {
	data, err := ioutil.ReadFile(req.FilePath)
	if err != nil {
		logUtils.Errorf("load end point data err ", zap.String("错误:", err.Error()))
		return err
	}

	handler := convert.NewHandler(req.DriverType, data, req.FilePath)
	doc, err := handler.ToOpenapi()
	fmt.Println(doc, "xxx")
	if err != nil {
		return err
	}
	//var x interface{}
	x, _ := json.Marshal(doc)
	fmt.Println(string(x))
	fmt.Println(json.Marshal(doc))
	openapi2endpoint := openapi.NewOpenapi2endpoint(doc)
	endpoints := openapi2endpoint.Convert()
	x, _ = json.Marshal(endpoints)
	fmt.Println(string(x))
	fmt.Println(endpoints)
	err = s.EndpointService.SaveEndpoints(endpoints, req)

	return

}
