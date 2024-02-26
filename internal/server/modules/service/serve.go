package service

import (
	"fmt"
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/core/cron"
	schemaHelper "github.com/aaronchen2k/deeptest/internal/pkg/helper/schema"
	serverConsts "github.com/aaronchen2k/deeptest/internal/server/consts"
	"github.com/aaronchen2k/deeptest/internal/server/core/cache"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	_domain "github.com/aaronchen2k/deeptest/pkg/domain"
	_commUtils "github.com/aaronchen2k/deeptest/pkg/lib/comm"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/jinzhu/copier"
	encoder "github.com/zwgblue/yaml-encoder"
	"gorm.io/gorm"
	"strconv"
	"strings"
)

type ServeService struct {
	ServeRepo                *repo.ServeRepo             `inject:""`
	ServeServerRepo          *repo.ServeServerRepo       `inject:""`
	EndpointRepo             *repo.EndpointRepo          `inject:""`
	EndpointInterfaceRepo    *repo.EndpointInterfaceRepo `inject:""`
	ProjectRepo              *repo.ProjectRepo           `inject:""`
	EnvironmentRepo          *repo.EnvironmentRepo       `inject:""`
	CategoryRepo             *repo.CategoryRepo          `inject:""`
	Cron                     *cron.ServerCron            `inject:""`
	EndpointInterfaceService *EndpointInterfaceService   `inject:""`
	CategoryService          *CategoryService            `inject:""`
	ComponentService         *ComponentService           `inject:""`
}

func (s *ServeService) ListByProject(tenantId consts.TenantId, projectId int, userId uint) (ret []model.Serve, currServe model.Serve, err error) {
	ret, err = s.ServeRepo.ListByProject(tenantId, uint(projectId))

	currServe, err = s.ServeRepo.GetCurrServeByUser(tenantId, userId)

	if currServe.ProjectId != uint(projectId) { //重新更新默认服务
		if len(ret) > 0 {
			currServe, err = s.ServeRepo.ChangeServe(tenantId, ret[0].ID, userId)
		}
	}

	return
}

func (s *ServeService) Paginate(tenantId consts.TenantId, req v1.ServeReqPaginate) (ret _domain.PageData, err error) {
	ret, err = s.ServeRepo.Paginate(tenantId, req)
	return
}

func (s *ServeService) Save(tenantId consts.TenantId, req v1.ServeReq) (res uint, err error) {
	var serve model.Serve
	if s.ServeRepo.ServeExist(tenantId, uint(req.ID), req.ProjectId, req.Name) {
		err = fmt.Errorf("serve name already exist")
		return
	}
	copier.CopyWithOption(&serve, req, copier.Option{DeepCopy: true})
	err = s.ServeRepo.SaveServe(tenantId, &serve)
	return serve.ID, err
}

func (s *ServeService) GetById(tenantId consts.TenantId, id uint) (res model.Serve) {
	res, _ = s.ServeRepo.Get(tenantId, id)
	return
}

func (s *ServeService) DeleteById(tenantId consts.TenantId, id uint) (err error) {
	/*
		err = s.canDelete(id)
		if err != nil {
			return err
		}
	*/
	err = s.ServeRepo.DeleteById(tenantId, id)
	return
}

func (s *ServeService) canDelete(tenantId consts.TenantId, id uint) (err error) {
	var count int64
	count, err = s.EndpointRepo.GetCountByServeId(tenantId, id)
	if err != nil {
		return
	}
	if count != 0 {
		err = fmt.Errorf("interfaces are created under the service,not allowed to delete")
	}

	return
}

func (s *ServeService) DisableById(tenantId consts.TenantId, id uint) (err error) {
	err = s.ServeRepo.DisableById(tenantId, id)
	return
}

func (s *ServeService) PaginateVersion(tenantId consts.TenantId, req v1.ServeVersionPaginate) (ret _domain.PageData, err error) {
	return s.ServeRepo.PaginateVersion(tenantId, req)
}

func (s *ServeService) SaveVersion(tenantId consts.TenantId, req v1.ServeVersionReq) (res uint, err error) {
	var serveVersion model.ServeVersion
	if s.ServeRepo.VersionExist(tenantId, req.ID, uint(req.ServeId), req.Value) {
		err = fmt.Errorf("serve version already exist")
		return
	}
	copier.CopyWithOption(&serveVersion, req, copier.Option{DeepCopy: true})
	err, res = s.ServeRepo.SaveVersion(tenantId, serveVersion.ID, &serveVersion), serveVersion.ID
	return
}

func (s *ServeService) DeleteVersionById(tenantId consts.TenantId, id uint) (err error) {
	err = s.ServeRepo.DeleteVersionById(tenantId, id)
	return
}

func (s *ServeService) DisableVersionById(tenantId consts.TenantId, id uint) (err error) {
	err = s.ServeRepo.DisableVersionById(tenantId, id)
	return
}

func (s *ServeService) ListServer(tenantId consts.TenantId, req v1.ServeServer, projectId, userId uint) (res []model.ServeServer, currServer model.ServeServer, err error) {
	if req.ServeId == 0 {
		server, _ := s.ServeServerRepo.Get(tenantId, req.ServerId)
		req.ServeId = server.ServeId
	}

	res, err = s.ServeRepo.ListServer(tenantId, req.ServeId)
	if err != nil {
		return
	}

	currServer, err = s.ServeRepo.GetCurrServerByUser(tenantId, projectId, req.ServeId, userId)
	if currServer.ServeId != req.ServeId {
		if len(res) != 0 {
			currServer, err = s.ChangeServer(tenantId, projectId, userId, req.ServeId, res[0].EnvironmentId)
		}
	}

	return
}

func (s *ServeService) ChangeServer(tenantId consts.TenantId, projectId, userId, serveId, serverId uint) (currServer model.ServeServer, err error) {
	if err = s.EnvironmentRepo.SetProjectUserServer(tenantId, projectId, userId, serverId); err != nil {
		return
	}

	currServer, err = s.ServeRepo.GetCurrServerByUser(tenantId, projectId, serveId, userId)
	return
}

func (s *ServeService) SaveServer(tenantId consts.TenantId, req v1.ServeServer) (res uint, err error) {
	var serve model.ServeServer
	copier.CopyWithOption(&serve, req, copier.Option{DeepCopy: true})
	err = s.ServeRepo.Save(tenantId, serve.ID, &serve)
	return serve.ID, err
}

func (s *ServeService) Copy(tenantId consts.TenantId, id uint) (err error) {
	serve, _ := s.ServeRepo.Get(tenantId, id)
	serve.ID = 0
	serve.Name += "_copy"
	serve.CreatedAt = nil
	serve.UpdatedAt = nil
	return s.ServeRepo.Save(tenantId, 0, &serve)
}

func (s *ServeService) SaveSchema(tenantId consts.TenantId, req v1.ServeSchemaReq) (res uint, err error) {
	var serveSchema model.ComponentSchema

	if req.ID != 0 {
		err = s.CategoryRepo.UpdateNameByEntityId(tenantId, req.ID, req.Name, serverConsts.SchemaCategory)
	}

	copier.CopyWithOption(&serveSchema, req, copier.Option{DeepCopy: true})
	serveSchema.Ref, err = s.ServeRepo.GetSchemaRef(tenantId, serveSchema.ID)
	if err != nil {
		return
	}

	err = s.ServeRepo.Save(tenantId, serveSchema.ID, &serveSchema)
	if err != nil {
		return

	}
	res = serveSchema.ID
	return

}

func (s *ServeService) SaveSecurity(tenantId consts.TenantId, req v1.ServeSecurityReq) (res uint, err error) {
	var serveSecurity model.ComponentSchemaSecurity
	if s.ServeRepo.SecurityExist(tenantId, uint(req.ID), uint(req.ServeId), req.Name) {
		err = fmt.Errorf("security name already exist")
		return
	}
	copier.CopyWithOption(&serveSecurity, req, copier.Option{DeepCopy: true})
	err = s.ServeRepo.Save(tenantId, serveSecurity.ID, &serveSecurity)
	return serveSecurity.ID, err
}

func (s *ServeService) PaginateSchema(tenantId consts.TenantId, req v1.ServeSchemaPaginate) (ret _domain.PageData, err error) {
	return s.ServeRepo.PaginateSchema(tenantId, req)
}

func (s *ServeService) GetSchema(tenantId consts.TenantId, id uint) (schema model.ComponentSchema, err error) {
	schema, err = s.ServeRepo.GetSchema(tenantId, id)
	if err == nil {
		schema.Content = s.FillSchemaRefId(tenantId, schema.ProjectId, schema.Content, nil)
	}
	return

}

func (s *ServeService) PaginateSecurity(tenantId consts.TenantId, req v1.ServeSecurityPaginate) (ret _domain.PageData, err error) {
	return s.ServeRepo.PaginateSecurity(tenantId, req)
}

func (s *ServeService) Example2Schema(data string) (schema schemaHelper.Schema) {
	schema2conv := schemaHelper.NewSchema2conv()
	var obj interface{}
	schema = schemaHelper.Schema{}
	_commUtils.JsonDecode(data, &obj)
	//_commUtils.JsonDecode("{\"id\":1,\"name\":\"user\"}", &obj)
	//_commUtils.JsonDecode("[\"0，2，3\"]", &obj)
	//_commUtils.JsonDecode("[]", &obj)
	//_commUtils.JsonDecode("[{\"id\":1,\"name\":\"user\"}]", &obj)
	//_commUtils.JsonDecode("{\"id\":[1,2,3],\"name\":\"user\"}", &obj)
	schema2conv.Example2Schema(obj, &schema)
	//fmt.Println(_commUtils.JsonEncode(schema), "++++++++++++")
	return
}

func (s *ServeService) DeleteSchemaById(tenantId consts.TenantId, id uint) (err error) {
	/*
		var schema model.ComponentSchema
		schema, err = s.ServeRepo.GetSchema(id)
		if err != nil {
			return err
		}


		var count int64
		count, err = s.EndpointInterfaceRepo.GetCountByRef(schema.Ref)
		if err != nil {
			return
		}

		if count > 0 {
			err = fmt.Errorf("the schema has been referenced and cannot be deleted")
			return
		}
	*/

	err = s.ServeRepo.DeleteSchemaById(tenantId, id)
	if err != nil {
		return
	}

	err = s.CategoryRepo.DeleteByEntityId(tenantId, id)

	return
}

func (s *ServeService) DeleteSecurityId(tenantId consts.TenantId, id uint) (err error) {
	err = s.ServeRepo.DeleteSecurityId(tenantId, id)
	return
}

func (s *ServeService) Schema2Example(tenantId consts.TenantId, projectId uint, data string) (obj interface{}) {
	schema2conv := schemaHelper.NewSchema2conv()
	schema2conv.Components = s.Components(tenantId, projectId)
	//schema1 := openapi3.Schema{}
	//_commUtils.JsonDecode(data, &schema),
	//_commUtils.JsonDecode("{\"type\":\"array\",\"items\":{\"type\":\"number\"}}", &schema)
	//_commUtils.JsonDecode("{\"properties\":{\"id\":{\"type\":\"number\"},\"name\":{\"type\":\"string\"}},\"type\":\"object\"}", &schema)
	//_commUtils.JsonDecode("{\"type\":\"array\",\"items\":{\"properties\":{\"id\":{\"type\":\"number\"},\"name\":{\"type\":\"string\"}},\"type\":\"object\"}}", &schema)
	schema := schemaHelper.SchemaRef{}
	//data = "{\"type\":\"object\",\"properties\":{\"name1\":{\"type\":\"object\",\"ref\":\"#/components/schemas/user1\",\"name\":\"user1\"},\"name2\":{\"type\":\"string\"},\"name3\":{\"type\":\"string\"}}}"
	_commUtils.JsonDecode(data, &schema)
	//_commUtils.JsonDecode("{\"type\":\"array\",\"items\":{\"type\":\"number\"}}", &schema1)
	//copier.CopyWithOption(&schema, a, copier.Option{DeepCopy: true})
	//fmt.Println(schema, "+++++++++++++")
	obj = schema2conv.Schema2Example(schema)
	//fmt.Println(schema.Items, "+++++", schema1.Items, _commUtils.JsonEncode(obj), "++++++++++++")
	return
}

func (s *ServeService) Components(tenantId consts.TenantId, projectId uint) (components *schemaHelper.Components) {
	components = schemaHelper.NewComponents()

	result, err := s.ServeRepo.GetSchemasByProjectId(tenantId, projectId)

	if err != nil {
		return
	}

	for _, item := range result {
		var schema schemaHelper.SchemaRef
		_commUtils.JsonDecode(item.Content, &schema)
		components.Add(item.ID, item.Ref, &schema)
	}

	return

}

func (s *ServeService) Schema2Yaml(data string) (res string) {
	schema := openapi3.Schema{}
	_commUtils.JsonDecode(data, &schema)
	jsonStr := _commUtils.JsonEncode(schema)
	var ret interface{}
	_commUtils.JsonDecode(jsonStr, &ret)
	content, _ := encoder.NewEncoder(ret).Encode()
	return string(content)
}

func (s *ServeService) CopySchema(tenantId consts.TenantId, id uint) (category model.Category, err error) {
	schema, err := s.ServeRepo.GetSchema(tenantId, id)

	if err != nil {
		return
	}

	category, err = s.CategoryRepo.GetByEntityId(tenantId, schema.ID, serverConsts.SchemaCategory)
	if err != nil {
		return
	}

	schema.ID = 0
	schema.CreatedAt = nil
	schema.UpdatedAt = nil

	schema.Name = "CopyOf" + schema.Name
	err = s.ServeRepo.Save(tenantId, 0, &schema)
	if err != nil {
		return
	}

	category.ID = 0
	category.CreatedAt = nil
	category.UpdatedAt = nil
	category.Name = schema.Name
	category.EntityId = schema.ID

	err = s.CategoryRepo.Save(tenantId, &category)
	return
}

func (s *ServeService) CopySchemaOther(tenantId consts.TenantId, id uint) (entityId uint, err error) {
	schema, err := s.ServeRepo.GetSchema(tenantId, id)
	if err != nil {
		return
	}

	schema.ID = 0
	schema.CreatedAt = nil
	schema.UpdatedAt = nil
	schema.Name = "CopyOf" + schema.Name
	err = s.ServeRepo.Save(tenantId, 0, &schema)
	if err != nil {
		return
	}

	entityId = schema.ID

	return
}

func (s *ServeService) BindEndpoint(tenantId consts.TenantId, req v1.ServeVersionBindEndpointReq) (err error) {
	var serveEndpointVersion []model.ServeEndpointVersion
	for _, endpointVersion := range req.EndpointVersions {
		serveEndpointVersion = append(serveEndpointVersion, model.ServeEndpointVersion{EndpointId: endpointVersion.EndpointId, EndpointVersion: endpointVersion.Version, ServeId: req.ServeId, ServeVersion: req.ServeVersion})
	}
	err = s.ServeRepo.BindEndpoint(tenantId, req.ServeId, req.ServeVersion, serveEndpointVersion)
	return
}

func (s *ServeService) ChangeServe(tenantId consts.TenantId, serveId, userId uint) (serve model.Serve, err error) {
	serve, err = s.ServeRepo.ChangeServe(tenantId, serveId, userId)

	return
}

func (s *ServeService) AddServerForHistory(tenantId consts.TenantId, req v1.HistoryServeAddServesReq) (err error) {
	projects, err := s.ProjectRepo.ListAll(tenantId)
	if len(projects) == 0 {
		return
	}

	for _, v := range projects {
		v := v
		go func() {
			server, err := s.EnvironmentRepo.GetByProjectAndName(tenantId, v.ID, req.ServerName)
			if err != nil && err != gorm.ErrRecordNotFound {
				return
			}

			serveList, err := s.ServeRepo.ListByProject(tenantId, v.ID)
			if err != nil {
				return
			}
			host, _ := cache.GetCacheString("host")

			//新增Mock环境
			if server.ID == 0 {
				server.Name = req.ServerName
				server.ProjectId = v.ID
				server.Sort = s.EnvironmentRepo.GetMaxOrder(tenantId, v.ID)
				err = s.EnvironmentRepo.Save(tenantId, &server)
				if err != nil {
					return
				}

				serveServer := make([]model.ServeServer, 0)
				for _, serve := range serveList {
					url := host + "/mocks/" + strconv.Itoa(int(serve.ID))
					if req.Url != "" {
						url = req.Url
					}
					if url == "" {
						return
					}

					serveServerTmp := model.ServeServer{ServeId: serve.ID, EnvironmentId: server.ID, Url: url, Description: server.Name}
					serveServer = append(serveServer, serveServerTmp)
				}
				if err = s.ServeServerRepo.BatchCreate(tenantId, serveServer); err != nil {
					return
				}
			} else {
				//更新url
				for _, serve := range serveList {
					url := host + "mocks/" + strconv.Itoa(int(serve.ID))
					if req.Url != "" {
						url = req.Url
					}
					if url == "" {
						return
					}
					_ = s.ServeServerRepo.UpdateUrlByServeAndServer(tenantId, serve.ID, server.ID, url)

				}
			}
		}()
	}
	return
}

/*
func (s *ServeService) SaveSwaggerSync(req v1.SwaggerSyncReq) (data model.SwaggerSync, err error) {
	var swaggerSync model.SwaggerSync
	copier.CopyWithOption(&swaggerSync, req, copier.Option{DeepCopy: true})
	serve, _ := s.ServeRepo.GetDefault(req.ProjectId)
	swaggerSync.ServeId = int(serve.ID)
	err = s.ServeRepo.SaveSwaggerSync(&swaggerSync)
	data = swaggerSync
	s.AddSwaggerCron(data)
	return
}

func (s *ServeService) SwaggerSyncDetail(projectId uint) (data model.SwaggerSync, err error) {
	return s.ServeRepo.GetSwaggerSync(projectId)
}

func (s *ServeService) SwaggerSyncList() (data []model.SwaggerSync, err error) {
	return s.ServeRepo.GetSwaggerSyncList()
}

func (s *ServeService) GetSwaggerSyncById(id uint) (data model.SwaggerSync, err error) {
	data, err = s.ServeRepo.GetSwaggerSyncById(id)
	return
}

func (s *ServeService) UpdateSwaggerSyncExecTimeById(id uint) (err error) {
	return s.ServeRepo.UpdateSwaggerSyncExecTimeById(id)
}

func (s *ServeService) AddSwaggerCron(item model.SwaggerSync) {

	name := "swaggerSync" + "_" + strconv.Itoa(int(item.ID))
	s.Cron.RemoveTask(name)

	if item.Switch == consts.SwitchOFF {
		return
	}

	taskId := item.ID
	s.Cron.AddCommonTask(name, item.Cron, func() {
		task, err := s.GetSwaggerSyncById(taskId)
		logUtils.Info("swagger定时任务开启：" + _commUtils.JsonEncode(item))
		if err != nil {
			logUtils.Errorf("swagger定时导入任务失败,任务ID：%v,错误原因：%v", name, err.Error())
			return
		}
		if task.Switch == consts.SwitchOFF {
			logUtils.Infof("swagger定时导入关闭,任务ID:%v", name)
			return
		}
		req := v1.ImportEndpointDataReq{ProjectId: uint(task.ProjectId), ServeId: uint(task.ServeId), CategoryId: int64(task.CategoryId), OpenUrlImport: true, DriverType: convert.SWAGGER, FilePath: task.Url, DataSyncType: consts.FullCover}
		err = s.EndpointInterfaceService.ImportEndpointData(req)
		if err != nil {
			logUtils.Error("swagger定时导入任务失败，错误原因：" + err.Error())
		}

		//更新实现执行时间
		s.UpdateSwaggerSyncExecTimeById(taskId)
		logUtils.Info("swagger定时任务结束：" + _commUtils.JsonEncode(item))
	})

}

*/

func (s *ServeService) FillSchemaRefId(tenantId consts.TenantId, projectId uint, schemaStr string, components *schemaHelper.Components) string {
	schema2conv := schemaHelper.NewSchema2conv()
	if components == nil {
		schema2conv.Components = s.Components(tenantId, projectId)
	} else {
		schema2conv.Components = components
	}
	schema := new(schemaHelper.SchemaRef)
	_commUtils.JsonDecode(schemaStr, schema)
	schema2conv.FillRefId(schema)
	return _commUtils.JsonEncode(schema)
}

func (s *ServeService) dependComponents(schemaStr string, components, dependComponents *schemaHelper.Components) {
	schema := new(schemaHelper.SchemaRef)
	schemaStr = strings.ReplaceAll(schemaStr, "\\u0026", "&")
	schemaStr = strings.ReplaceAll(schemaStr, "\n", "")
	schemaStr = strings.ReplaceAll(schemaStr, "\"ref\":", "\"$ref\":")
	_commUtils.JsonDecode(schemaStr, schema)
	schema2conv := schemaHelper.NewSchema2conv()
	schema2conv.Components = components
	schema2conv.SchemaComponents(schema, dependComponents)
}

func (s *ServeService) GetComponents(tenantId consts.TenantId, projectId uint) (result []model.ComponentSchema, err error) {

	result, err = s.ServeRepo.GetSchemasByProjectId(tenantId, projectId)
	if err != nil {
		return
	}
	components := s.Components(tenantId, projectId)
	for key, item := range result {
		result[key].Content = s.FillSchemaRefId(tenantId, item.ProjectId, item.Content, components)
	}

	return

}
