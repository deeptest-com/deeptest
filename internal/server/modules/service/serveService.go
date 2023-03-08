package service

import (
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/helper/openapi"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	_domain "github.com/aaronchen2k/deeptest/pkg/domain"
	_commUtils "github.com/aaronchen2k/deeptest/pkg/lib/comm"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/jinzhu/copier"
	encoder "github.com/zwgblue/yaml-encoder"
)

type ServeService struct {
	ServeRepo *repo.ServeRepo `inject:""`
}

func NewServeService() *ServeService {
	return &ServeService{}
}

func (s *ServeService) ListByProject(projectId int) (ret []model.Serve, err error) {
	ret, err = s.ServeRepo.ListByProject(projectId)
	return
}

func (s *ServeService) Paginate(req v1.ServeReqPaginate) (ret _domain.PageData, err error) {
	ret, err = s.ServeRepo.Paginate(req)
	return
}

func (s *ServeService) Save(req v1.ServeReq) (res uint, err error) {
	var serve model.Serve
	copier.CopyWithOption(&serve, req, copier.Option{DeepCopy: true})
	err = s.ServeRepo.Save(serve.ID, &serve)
	return serve.ID, err
}

func (s *ServeService) GetById(id uint) (res model.Serve) {
	res, _ = s.ServeRepo.Get(id)
	return
}

func (s *ServeService) DeleteById(id uint) (err error) {
	err = s.ServeRepo.DeleteById(id)
	return
}

func (s *ServeService) DisableById(id uint) (err error) {
	err = s.ServeRepo.DisableById(id)
	return
}

func (s *ServeService) PaginateVersion(req v1.ServeVersionPaginate) (ret _domain.PageData, err error) {
	return s.ServeRepo.PaginateVersion(req)
}

func (s *ServeService) SaveVersion(req v1.ServeVersionReq) (res uint, err error) {
	var serveVersion model.ServeVersion
	copier.CopyWithOption(&serveVersion, req, copier.Option{DeepCopy: true})
	err, res = s.ServeRepo.Save(serveVersion.ID, &serveVersion), serveVersion.ID
	return
}

func (s *ServeService) DeleteVersionById(id uint) (err error) {
	err = s.ServeRepo.DeleteVersionById(id)
	return
}

func (s *ServeService) DisableVersionById(id uint) (err error) {
	err = s.ServeRepo.DisableVersionById(id)
	return
}

func (s *ServeService) ListServer(serveId uint) (res []model.ServeServer, err error) {
	res, err = s.ServeRepo.ListServer(serveId)
	return
}

func (s *ServeService) SaveServer(req v1.ServeServerReq) (res uint, err error) {
	var serve model.ServeServer
	copier.CopyWithOption(&serve, req, copier.Option{DeepCopy: true})
	err = s.ServeRepo.Save(serve.ID, &serve)
	return serve.ID, err
}

func (s *ServeService) Copy(id uint) (err error) {
	serve, _ := s.ServeRepo.Get(id)
	serve.ID = 0
	serve.CreatedAt = nil
	serve.UpdatedAt = nil
	return s.ServeRepo.Save(0, &serve)
}

func (s *ServeService) SaveSchema(req v1.ServeSchemaReq) (res uint, err error) {
	var serveSchema model.ComponentSchema
	copier.CopyWithOption(&serveSchema, req, copier.Option{DeepCopy: true})
	err = s.ServeRepo.Save(serveSchema.ID, &serveSchema)
	return serveSchema.ID, err
}

func (s *ServeService) PaginateSchema(req v1.ServeSchemaPaginate) (ret _domain.PageData, err error) {
	return s.ServeRepo.PaginateSchema(req)
}

func (s *ServeService) Example2Schema(data string) (schema openapi3.Schema) {
	schema2conv := openapi.NewSchema2conv()
	var obj interface{}
	schema = openapi3.Schema{}
	_commUtils.JsonDecode(data, &obj)
	_commUtils.JsonDecode("{\"id\":1,\"name\":\"user\"}", &obj)
	//_commUtils.JsonDecode("[\"0，2，3\"]", &obj)
	//_commUtils.JsonDecode("[]", &obj)
	//_commUtils.JsonDecode("[{\"id\":1,\"name\":\"user\"}]", &obj)
	//_commUtils.JsonDecode("{\"id\":[1,2,3],\"name\":\"user\"}", &obj)
	schema2conv.Example2Schema(obj, &schema)
	//fmt.Println(_commUtils.JsonEncode(schema), "++++++++++++")
	return
}

func (s *ServeService) DeleteSchemaById(id uint) (err error) {
	err = s.ServeRepo.DeleteSchemaById(id)
	return
}

func (s *ServeService) Schema2Example(data string) (obj interface{}) {
	schema2conv := openapi.NewSchema2conv()
	schema := openapi3.Schema{}
	_commUtils.JsonDecode(data, &schema)
	//_commUtils.JsonDecode("{\"type\":\"array\",\"items\":{\"type\":\"number\"}}", &schema)
	//_commUtils.JsonDecode("{\"properties\":{\"id\":{\"type\":\"number\"},\"name\":{\"type\":\"string\"}},\"type\":\"object\"}", &schema)
	//_commUtils.JsonDecode("{\"type\":\"array\",\"items\":{\"properties\":{\"id\":{\"type\":\"number\"},\"name\":{\"type\":\"string\"}},\"type\":\"object\"}}", &schema)
	obj = schema2conv.Schema2Example(schema)
	//fmt.Println(_commUtils.JsonEncode(obj), "++++++++++++")
	return
}

func (s *ServeService) Schema2Yaml(data string) (res string) {
	schema := openapi3.Schema{}
	_commUtils.JsonDecode(data, &schema)
	content, _ := encoder.NewEncoder(schema).Encode()
	return string(content)
}

func (s *ServeService) CopySchema(id uint) (schema model.ComponentSchema, err error) {
	schema, err = s.ServeRepo.GetSchema(id)
	if err != nil {
		return
	}
	schema.ID = 0
	schema.CreatedAt = nil
	schema.UpdatedAt = nil
	err = s.ServeRepo.Save(0, &schema)
	return
}
