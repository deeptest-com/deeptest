package service

import (
	"errors"
	"fmt"
	"github.com/474420502/requests"
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	builtin "github.com/aaronchen2k/deeptest/internal/pkg/buildin"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	curlHelper "github.com/aaronchen2k/deeptest/internal/pkg/helper/curl"
	"github.com/aaronchen2k/deeptest/internal/pkg/helper/openapi"
	serverConsts "github.com/aaronchen2k/deeptest/internal/server/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	_domain "github.com/aaronchen2k/deeptest/pkg/domain"
	_commUtils "github.com/aaronchen2k/deeptest/pkg/lib/comm"
	"net/http"
	"net/url"
	"strings"
	"sync"
)

type EndpointService struct {
	EndpointRepo             *repo.EndpointRepo          `inject:""`
	ServeRepo                *repo.ServeRepo             `inject:""`
	EndpointInterfaceRepo    *repo.EndpointInterfaceRepo `inject:""`
	ServeServerRepo          *repo.ServeServerRepo       `inject:""`
	UserRepo                 *repo.UserRepo              `inject:""`
	CategoryRepo             *repo.CategoryRepo          `inject:""`
	DiagnoseInterfaceService *DiagnoseInterfaceService   `inject:""`
}

func (s *EndpointService) Paginate(req v1.EndpointReqPaginate) (ret _domain.PageData, err error) {
	ret, err = s.EndpointRepo.Paginate(req)
	return
}

func (s *EndpointService) Save(endpoint model.Endpoint) (res uint, err error) {

	if endpoint.ServerId == 0 {
		server, _ := s.ServeServerRepo.GetDefaultByServe(endpoint.ServeId)
		endpoint.ServerId = server.ID
	}

	if endpoint.Curl != "" {
		err = s.curlToEndpoint(&endpoint)
		if err != nil {
			return
		}
	}

	err = s.EndpointRepo.SaveAll(&endpoint)

	return endpoint.ID, err
}

func (s *EndpointService) GetById(id uint, version string) (res model.Endpoint) {
	res, _ = s.EndpointRepo.GetAll(id, version)
	return
}

func (s *EndpointService) DeleteById(id uint) (err error) {
	var count int64
	count, err = s.EndpointRepo.GetUsedCountByEndpointId(id)
	if err != nil {
		return err
	}

	if count > 0 {
		err = fmt.Errorf("this interface has already been used by scenarios, not allowed to delete")
		return err
	}

	err = s.EndpointRepo.DeleteById(id)
	err = s.EndpointInterfaceRepo.DeleteByEndpoint(id)

	return
}

func (s *EndpointService) DisableById(id uint) (err error) {
	err = s.EndpointRepo.UpdateStatus(id, serverConsts.Abandoned)
	return
}

func (s *EndpointService) Publish(id uint) (err error) {
	err = s.EndpointRepo.UpdateStatus(id, serverConsts.Published)
	return
}

func (s *EndpointService) Develop(id uint) (err error) {
	err = s.EndpointRepo.UpdateStatus(id, serverConsts.Developing)
	return
}

func (s *EndpointService) Copy(id uint, version string) (res uint, err error) {
	endpoint, _ := s.EndpointRepo.GetAll(id, version)
	s.removeIds(&endpoint)
	endpoint.Title += "_copy"
	err = s.EndpointRepo.SaveAll(&endpoint)
	return endpoint.ID, err
}

func (s *EndpointService) removeIds(endpoint *model.Endpoint) {
	endpoint.ID = 0
	endpoint.CreatedAt = nil
	endpoint.UpdatedAt = nil
	for key, _ := range endpoint.PathParams {
		endpoint.PathParams[key].ID = 0
	}
	for key, _ := range endpoint.Interfaces {
		endpoint.Interfaces[key].ID = 0
		endpoint.Interfaces[key].RequestBody.ID = 0
		endpoint.Interfaces[key].RequestBody.SchemaItem.ID = 0
		for key1, _ := range endpoint.Interfaces[key].Headers {
			endpoint.Interfaces[key].Headers[key1].ID = 0
		}
		for key1, _ := range endpoint.Interfaces[key].Cookies {
			endpoint.Interfaces[key].Cookies[key1].ID = 0
		}
		for key1, _ := range endpoint.Interfaces[key].Params {
			endpoint.Interfaces[key].Params[key1].ID = 0
		}
		for key1, _ := range endpoint.Interfaces[key].ResponseBodies {
			endpoint.Interfaces[key].ResponseBodies[key1].ID = 0
			endpoint.Interfaces[key].ResponseBodies[key1].SchemaItem.ID = 0
			for key2, _ := range endpoint.Interfaces[key].ResponseBodies[key1].Headers {
				endpoint.Interfaces[key].ResponseBodies[key1].Headers[key2].ID = 0
			}
		}
	}

}

func (s *EndpointService) Yaml(endpoint model.Endpoint) (res interface{}) {
	serve, err := s.ServeRepo.Get(endpoint.ServeId)
	if err != nil {
		return
	}

	serveComponent, err := s.ServeRepo.GetSchemasByServeId(serve.ID)
	if err != nil {
		return
	}
	serve.Components = serveComponent
	serveServer, err := s.ServeRepo.ListServer(serve.ID)
	if err != nil {
		return
	}
	serve.Servers = serveServer

	Securities, err := s.ServeRepo.ListSecurity(serve.ID)
	if err != nil {
		return
	}
	serve.Securities = Securities

	serve2conv := openapi.NewServe2conv(serve, []model.Endpoint{endpoint})
	res = serve2conv.ToV3()
	return
}

func (s *EndpointService) UpdateStatus(id uint, status int64) (err error) {
	err = s.EndpointRepo.UpdateStatus(id, status)
	return
}

func (s *EndpointService) BatchDelete(ids []uint) (err error) {
	err = s.EndpointRepo.DeleteByIds(ids)
	return
}

func (s *EndpointService) GetVersionsByEndpointId(endpointId uint) (res []model.EndpointVersion, err error) {
	res, err = s.EndpointRepo.GetVersionsByEndpointId(endpointId)
	return
}

func (s *EndpointService) GetLatestVersion(endpointId uint) (version string) {
	version = "v0.1.0"
	if res, err := s.EndpointRepo.GetLatestVersion(endpointId); err != nil {
		version = res.Version
	}
	return
}

func (s *EndpointService) AddVersion(version *model.EndpointVersion) (err error) {
	err = s.EndpointRepo.FindVersion(version)
	if err != nil {
		err = s.EndpointRepo.Save(0, version)
	} else {
		err = fmt.Errorf("version already exists")
	}
	return
}

func (s *EndpointService) SaveEndpoints(endpoints []*model.Endpoint, dirs *openapi.Dirs, components map[string]*model.ComponentSchema, req v1.ImportEndpointDataReq) (err error) {

	if dirs.Id == 0 || dirs.Id == -1 {
		root, _ := s.CategoryRepo.ListByProject(serverConsts.EndpointCategory, req.ProjectId, 0)
		dirs.Id = int64(root[0].ID)
	}
	wg := new(sync.WaitGroup)
	wg.Add(2)
	s.createDirs(dirs, req)
	go s.createComponents(wg, components, req)
	go s.createEndpoints(wg, endpoints, dirs, req)

	//err = s.EndpointRepo.CreateEndpoints(endpoints)
	wg.Wait()
	return
}

func (s *EndpointService) createEndpoints(wg *sync.WaitGroup, endpoints []*model.Endpoint, dirs *openapi.Dirs, req v1.ImportEndpointDataReq) (err error) {
	defer func() {
		wg.Done()
	}()
	user, _ := s.UserRepo.FindById(req.UserId)
	for _, endpoint := range endpoints {
		endpoint.ProjectId, endpoint.ServeId, endpoint.CategoryId, endpoint.CreateUser = req.ProjectId, req.ServeId, req.CategoryId, user.Name
		endpoint.Status = 1
		endpoint.CategoryId = s.getCategoryId(endpoint.Tags, dirs)
		_, err = s.Save(*endpoint)
		if err != nil {
			return
		}
	}

	return
}

func (s *EndpointService) createComponents(wg *sync.WaitGroup, components map[string]*model.ComponentSchema, req v1.ImportEndpointDataReq) {
	defer func() {
		wg.Done()
	}()
	var NewComponents []*model.ComponentSchema
	for _, component := range components {
		component.ServeId = int64(req.ServeId)
		NewComponents = append(NewComponents, component)
	}
	s.ServeRepo.CreateSchemas(NewComponents)
}

func (s *EndpointService) createDirs(data *openapi.Dirs, req v1.ImportEndpointDataReq) (err error) {
	for name, dirs := range data.Dirs {
		category := model.Category{Name: name, ParentId: int(data.Id), ProjectId: req.ProjectId, UseID: req.UserId, Type: serverConsts.EndpointCategory}
		err = s.CategoryRepo.Save(&category)
		if err != nil {
			return
		}

		dirs.Id = int64(category.ID)
		err = s.createDirs(dirs, req)
		if err != nil {
			return err
		}
	}
	return
}

func (s *EndpointService) getCategoryId(tags []string, dirs *openapi.Dirs) int64 {
	rootId := dirs.Id
	for _, tag := range tags {
		dirs = dirs.Dirs[tag]
	}
	if dirs.Id == rootId {
		return -1
	}
	return dirs.Id
}

func (s *EndpointService) BatchUpdateByField(req v1.BatchUpdateReq) (err error) {
	valueType := builtin.InterfaceType(req.Value)
	if _commUtils.InSlice(req.FieldName, []string{"status", "categoryId"}) {
		if !_commUtils.InSlice(valueType, []string{"int", "float64"}) {
			err = errors.New("数据类型错误")
		}

		var value int64
		switch valueType {
		case "int":
			value = int64(req.Value.(int))
		case "float64":
			value = int64(req.Value.(float64))
		}

		if req.FieldName == "status" {
			err = s.EndpointRepo.BatchUpdateStatus(req.EndpointIds, value)
		} else if req.FieldName == "categoryId" {
			err = s.EndpointRepo.BatchUpdateCategory(req.EndpointIds, value)
		}
	} else {
		err = errors.New("字段错误")
	}
	return
}

func (s *EndpointService) curlToEndpoint(endpoint *model.Endpoint) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("curl格式错误")
		}
	}()
	curlObj := curlHelper.Parse(endpoint.Curl)
	wf := curlObj.CreateTemporary(curlObj.CreateSession())

	endpoint.Path = curlObj.ParsedURL.Path

	endpoint.Interfaces = s.getInterfaces(curlObj, wf)

	return
}

func (s *EndpointService) getInterfaces(cURL *curlHelper.CURL, wf *requests.Temporary) (interfaces []model.EndpointInterface) {
	interf := model.EndpointInterface{}
	interf.Params = s.getQueryParams(wf.GetQuery())
	interf.Headers = s.getHeaders(wf.Header)
	interf.Cookies = s.getCookies(wf.Cookies)
	bodyType := ""
	contentType := strings.Split(cURL.ContentType, ";")
	if len(contentType) > 1 {
		bodyType = contentType[0]
	}
	interf.BodyType = consts.HttpContentType(bodyType)
	interf.RequestBody = s.getRequestBody(wf.Body.String())
	interf.RequestBody.MediaType = string(interf.BodyType)
	interf.Method = s.getMethod(bodyType, cURL.Method)
	interfaces = append(interfaces, interf)

	return
}

func (s *EndpointService) getMethod(contentType, method string) (ret consts.HttpMethod) {
	ret = consts.HttpMethod(method)

	if contentType == "application/json" {
		ret = "POST"
	}

	return
}

func (s *EndpointService) getQueryParams(params url.Values) (ret []model.EndpointInterfaceParam) {
	m := map[string]bool{}
	for key, arr := range params {
		for _, item := range arr {
			if _, ok := m[key]; ok {
				continue
			}
			ret = append(ret, model.EndpointInterfaceParam{
				SchemaParam: model.SchemaParam{Name: key, Type: "string", Value: item, Default: item, Example: item},
			})
			m[key] = true
		}
	}

	return
}

func (s *EndpointService) getHeaders(header http.Header) (ret []model.EndpointInterfaceHeader) {
	for key, arr := range header {
		for _, item := range arr {
			ret = append(ret, model.EndpointInterfaceHeader{
				SchemaParam: model.SchemaParam{Name: key, Type: "string", Value: item, Default: item, Example: item},
			})
		}
	}

	return
}

func (s *EndpointService) getCookies(cookies map[string]*http.Cookie) (ret []model.EndpointInterfaceCookie) {
	for _, item := range cookies {
		ret = append(ret, model.EndpointInterfaceCookie{
			SchemaParam: model.SchemaParam{Name: item.Name, Type: "string", Value: item.Value, Default: item.Value, Example: item.Value},
		})
	}

	return
}

func (s *EndpointService) getRequestBody(body string) (requestBody model.EndpointInterfaceRequestBody) {
	requestBody = model.EndpointInterfaceRequestBody{}

	if body != "" {
		var examples []map[string]string
		examples = append(examples, map[string]string{"content": body, "name": "defaultExample"})
		requestBody.Examples = _commUtils.JsonEncode(examples)
	}

	requestBody.SchemaItem = s.getRequestBodyItem(body)
	return
}

func (s *EndpointService) getRequestBodyItem(body string) (requestBodyItem model.EndpointInterfaceRequestBodyItem) {
	requestBodyItem = model.EndpointInterfaceRequestBodyItem{}
	requestBodyItem.Type = "object"
	schema2conv := openapi.NewSchema2conv()
	var obj interface{}
	schema := openapi.Schema{}
	_commUtils.JsonDecode(body, &obj)
	schema2conv.Example2Schema(obj, &schema)
	requestBodyItem.Content = _commUtils.JsonEncode(schema)
	return
}
