package service

import (
	"fmt"
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/helper/openapi"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	_domain "github.com/aaronchen2k/deeptest/pkg/domain"
	"reflect"
)

type EndpointService struct {
	EndpointRepo *repo.EndpointRepo `inject:""`
	ServeRepo    *repo.ServeRepo    `inject:""`
}

func NewEndpointService() *EndpointService {
	return &EndpointService{}
}

func (s *EndpointService) Paginate(req v1.EndpointReqPaginate) (ret _domain.PageData, err error) {
	ret, err = s.EndpointRepo.Paginate(req)
	return
}

func (s *EndpointService) Save(endpoint model.Endpoint) (res uint, err error) {
	//fmt.Println(_commUtils.JsonEncode(endpoint), "++++++", _commUtils.JsonEncode(req))
	err = s.EndpointRepo.SaveAll(&endpoint)
	return endpoint.ID, err
}

func (s *EndpointService) GetById(id uint) (res model.Endpoint) {
	res, _ = s.EndpointRepo.GetAll(id)
	return
}

func (s *EndpointService) DeleteById(id uint) (err error) {
	err = s.EndpointRepo.DeleteById(id)
	return
}

func (s *EndpointService) DisableById(id uint) (err error) {
	err = s.EndpointRepo.DisableById(id)
	return
}

func (s *EndpointService) Copy(id uint) (res uint, err error) {
	endpoint, _ := s.EndpointRepo.GetAll(id)
	//s.removeIds(&endpoint)
	//fmt.Println(endpoint)
	err = s.EndpointRepo.SaveAll(&endpoint)
	//fmt.Println(endpoint.PathParams[0].ID)
	return endpoint.ID, err
}

func (s *EndpointService) removeIds(object interface{}) interface{} {
	//fmt.Println(reflect.ValueOf(object).Kind(), "++++++")

	//fmt.Println("canset++++++", c.Kind(), reflect.Struct)
	T := reflect.TypeOf(object)
	V := reflect.ValueOf(object)
	if T.Kind() == reflect.Ptr {
		T = T.Elem()
		V = V.Elem()
	}
	fmt.Println(T.Kind(), "+++")
	if T.Kind() == reflect.Struct {
		/*
			if V.FieldByName("ID").CanSet() {
				V.FieldByName("ID").SetUint(0)
				fmt.Println("-----")
				//fmt.Println(object.ID, "+++++")
			}
		*/
		for i := 0; i < T.NumField(); i++ {
			fmt.Println(T.Field(i).Type.Kind(), T.Field(i).Name)
			if T.Field(i).Type.Kind() == reflect.Struct {
				s.removeIds(V.FieldByName(T.Field(i).Name))
			} else if T.Field(i).Type.Kind() == reflect.Slice {
				//fmt.Println(T.Field(i).Name, T.Field(i).Type.Kind(), V.FieldByName(T.Field(i).Name), reflect.TypeOf(V.FieldByName(T.Field(i).Name).Interface()))
				s.removeIds(V.FieldByName(T.Field(i).Name).Interface())
			} else if T.Field(i).Type.Kind() == reflect.Uint && T.Field(i).Name == "ID" {

				V.FieldByName("ID").SetUint(0)
			}
		}
	} else if T.Kind() == reflect.Slice {
		for i := 0; i < V.Len(); i++ {
			fmt.Println(i, V.Index(i).Interface())
			s.removeIds(V.Index(i).Addr())
		}
	}
	return object
}

func (s *EndpointService) Yaml(endpoint model.Endpoint) (res interface{}) {
	serve, err := s.ServeRepo.Get(endpoint.ServeId)
	if err != nil {
		return
	}
	serve2conv := openapi.NewServe2conv(serve, []model.Endpoint{endpoint})
	res = serve2conv.ToV3()
	fmt.Println(res, "+++++++++++++++++++")
	return
}
