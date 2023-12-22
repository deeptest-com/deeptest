package service

import (
	"fmt"
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
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
}

func (s *ServeService) ListByProject(projectId int, userId uint) (ret []model.Serve, currServe model.Serve, err error) {
	ret, err = s.ServeRepo.ListByProject(uint(projectId))

	currServe, err = s.ServeRepo.GetCurrServeByUser(userId)

	if currServe.ProjectId != uint(projectId) { //重新更新默认服务
		if len(ret) > 0 {
			currServe, err = s.ServeRepo.ChangeServe(ret[0].ID, userId)
		}
	}

	return
}

func (s *ServeService) Paginate(req v1.ServeReqPaginate) (ret _domain.PageData, err error) {
	ret, err = s.ServeRepo.Paginate(req)
	return
}

func (s *ServeService) Save(req v1.ServeReq) (res uint, err error) {
	var serve model.Serve
	if s.ServeRepo.ServeExist(uint(req.ID), req.ProjectId, req.Name) {
		err = fmt.Errorf("serve name already exist")
		return
	}
	copier.CopyWithOption(&serve, req, copier.Option{DeepCopy: true})
	err = s.ServeRepo.SaveServe(&serve)
	return serve.ID, err
}

func (s *ServeService) GetById(id uint) (res model.Serve) {
	res, _ = s.ServeRepo.Get(id)
	return
}

func (s *ServeService) DeleteById(id uint) (err error) {
	/*
		err = s.canDelete(id)
		if err != nil {
			return err
		}
	*/
	err = s.ServeRepo.DeleteById(id)
	return
}

func (s *ServeService) canDelete(id uint) (err error) {
	var count int64
	count, err = s.EndpointRepo.GetCountByServeId(id)
	if err != nil {
		return
	}
	if count != 0 {
		err = fmt.Errorf("interfaces are created under the service,not allowed to delete")
	}

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
	if s.ServeRepo.VersionExist(uint(req.ID), uint(req.ServeId), req.Value) {
		err = fmt.Errorf("serve version already exist")
		return
	}
	copier.CopyWithOption(&serveVersion, req, copier.Option{DeepCopy: true})
	err, res = s.ServeRepo.SaveVersion(serveVersion.ID, &serveVersion), serveVersion.ID
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

func (s *ServeService) ListServer(req v1.ServeServer, projectId, userId uint) (res []model.ServeServer, currServer model.ServeServer, err error) {
	if req.ServeId == 0 {
		server, _ := s.ServeServerRepo.Get(req.ServerId)
		req.ServeId = server.ServeId
	}

	res, err = s.ServeRepo.ListServer(req.ServeId)
	if err != nil {
		return
	}

	currServer, err = s.ServeRepo.GetCurrServerByUser(projectId, req.ServeId, userId)
	if currServer.ServeId != req.ServeId {
		if len(res) != 0 {
			currServer, err = s.ChangeServer(projectId, userId, req.ServeId, res[0].EnvironmentId)
		}
	}

	return
}

func (s *ServeService) ChangeServer(projectId, userId, serveId, serverId uint) (currServer model.ServeServer, err error) {
	if err = s.EnvironmentRepo.SetProjectUserServer(projectId, userId, serverId); err != nil {
		return
	}

	currServer, err = s.ServeRepo.GetCurrServerByUser(projectId, serveId, userId)
	return
}

func (s *ServeService) SaveServer(req v1.ServeServer) (res uint, err error) {
	var serve model.ServeServer
	copier.CopyWithOption(&serve, req, copier.Option{DeepCopy: true})
	err = s.ServeRepo.Save(serve.ID, &serve)
	return serve.ID, err
}

func (s *ServeService) Copy(id uint) (err error) {
	serve, _ := s.ServeRepo.Get(id)
	serve.ID = 0
	serve.Name += "_copy"
	serve.CreatedAt = nil
	serve.UpdatedAt = nil
	return s.ServeRepo.Save(0, &serve)
}

func (s *ServeService) SaveSchema(req v1.ServeSchemaReq) (res uint, err error) {
	var serveSchema model.ComponentSchema
	//if req.ID == 0 && s.ServeRepo.SchemaExist(uint(req.ID), uint(req.ServeId), req.Name) {
	//	err = fmt.Errorf("schema name already exist")
	//	return
	//}
	copier.CopyWithOption(&serveSchema, req, copier.Option{DeepCopy: true})
	joinedPath, err := s.CategoryService.GetJoinedPath(serverConsts.SchemaCategory, req.ProjectId, uint(req.TargetId))
	if err != nil {
		return
	}

	serveSchema.Ref = "#/components/schemas" + joinedPath + "/" + serveSchema.Name
	err = s.ServeRepo.Save(serveSchema.ID, &serveSchema)

	if req.ID == 0 {
		createCategoryReq := v1.CategoryCreateReq{Name: req.Name, TargetId: req.TargetId, ProjectId: req.ProjectId, Type: serverConsts.SchemaCategory, Mode: "child", EntityId: serveSchema.ID}
		_, _ = s.CategoryService.Create(createCategoryReq)
	} else {
		category, err := s.CategoryRepo.GetByEntityId(serveSchema.ID)
		if err != nil {
			return res, err
		}

		category.Name = serveSchema.Name
		err = s.CategoryRepo.Save(&category)
	}

	return serveSchema.ID, err
}

func (s *ServeService) SaveSecurity(req v1.ServeSecurityReq) (res uint, err error) {
	var serveSecurity model.ComponentSchemaSecurity
	if s.ServeRepo.SecurityExist(uint(req.ID), uint(req.ServeId), req.Name) {
		err = fmt.Errorf("security name already exist")
		return
	}
	copier.CopyWithOption(&serveSecurity, req, copier.Option{DeepCopy: true})
	err = s.ServeRepo.Save(serveSecurity.ID, &serveSecurity)
	return serveSecurity.ID, err
}

func (s *ServeService) PaginateSchema(req v1.ServeSchemaPaginate) (ret _domain.PageData, err error) {
	return s.ServeRepo.PaginateSchema(req)
}

func (s *ServeService) GetSchema(id uint) (schema model.ComponentSchema, err error) {
	schema, err = s.ServeRepo.GetSchema(id)
	if err != nil {
		schema.Content = s.FillSchemaRefId(schema.ProjectId, schema.Content)
	}
	return
}

func (s *ServeService) PaginateSecurity(req v1.ServeSecurityPaginate) (ret _domain.PageData, err error) {
	return s.ServeRepo.PaginateSecurity(req)
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

func (s *ServeService) DeleteSchemaById(id uint) (err error) {
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

	err = s.ServeRepo.DeleteSchemaById(id)
	if err != nil {
		return
	}

	err = s.CategoryRepo.DeleteByEntityId(id)
	return
}

func (s *ServeService) DeleteSecurityId(id uint) (err error) {
	err = s.ServeRepo.DeleteSecurityId(id)
	return
}

func (s *ServeService) Schema2Example(projectId uint, data string) (obj interface{}) {
	schema2conv := schemaHelper.NewSchema2conv()
	schema2conv.Components = s.Components(projectId, data)
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

func (s *ServeService) Components(projectId uint, options ...string) (components schemaHelper.Components) {
	components = schemaHelper.Components{}

	var refIds []interface{}
	if len(options) > 0 {
		schema2conv := schemaHelper.NewSchema2conv()
		var schemaRef schemaHelper.SchemaRef
		_commUtils.JsonDecode(options[0], &schemaRef)
		refIds = schema2conv.GetRefIds(&schemaRef)
	}

	result, err := s.ServeRepo.GetSchemasByProjectId(projectId, refIds)
	if err != nil {
		return
	}

	for _, item := range result {
		var schema schemaHelper.SchemaRef
		_commUtils.JsonDecode(item.Content, &schema)
		schema.RefId = item.ID
		components[item.Ref] = &schema
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

func (s *ServeService) CopySchema(id uint) (schema model.ComponentSchema, err error) {
	schema, err = s.ServeRepo.GetSchema(id)
	if err != nil {
		return
	}

	category, err := s.CategoryRepo.GetByEntityId(schema.ID)
	if err != nil {
		return
	}

	schema.ID = 0
	schema.CreatedAt = nil
	schema.UpdatedAt = nil
	schema.Name = "CopyOf" + schema.Name
	err = s.ServeRepo.Save(0, &schema)
	if err != nil {
		return
	}

	category.ID = 0
	category.CreatedAt = nil
	category.UpdatedAt = nil
	category.Name = schema.Name
	category.EntityId = schema.ID

	err = s.CategoryRepo.Save(&category)
	return
}

func (s *ServeService) BindEndpoint(req v1.ServeVersionBindEndpointReq) (err error) {
	var serveEndpointVersion []model.ServeEndpointVersion
	for _, endpointVersion := range req.EndpointVersions {
		serveEndpointVersion = append(serveEndpointVersion, model.ServeEndpointVersion{EndpointId: endpointVersion.EndpointId, EndpointVersion: endpointVersion.Version, ServeId: req.ServeId, ServeVersion: req.ServeVersion})
	}
	err = s.ServeRepo.BindEndpoint(req.ServeId, req.ServeVersion, serveEndpointVersion)
	return
}

func (s *ServeService) ChangeServe(serveId, userId uint) (serve model.Serve, err error) {
	serve, err = s.ServeRepo.ChangeServe(serveId, userId)

	return
}

func (s *ServeService) AddServerForHistory(req v1.HistoryServeAddServesReq) (err error) {
	projects, err := s.ProjectRepo.ListAll()
	if len(projects) == 0 {
		return
	}

	for _, v := range projects {
		v := v
		go func() {
			server, err := s.EnvironmentRepo.GetByProjectAndName(v.ID, req.ServerName)
			if err != nil && err != gorm.ErrRecordNotFound {
				return
			}

			serveList, err := s.ServeRepo.ListByProject(v.ID)
			if err != nil {
				return
			}
			host, _ := cache.GetCacheString("host")

			//新增Mock环境
			if server.ID == 0 {
				server.Name = req.ServerName
				server.ProjectId = v.ID
				server.Sort = s.EnvironmentRepo.GetMaxOrder(v.ID)
				err = s.EnvironmentRepo.Save(&server)
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
				if err = s.ServeServerRepo.BatchCreate(serveServer); err != nil {
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
					_ = s.ServeServerRepo.UpdateUrlByServerAndServer(serve.ID, server.ID, url)

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

func (s *ServeService) FillSchemaRefId(projectId uint, schemaStr string) string {
	schema2conv := schemaHelper.NewSchema2conv()
	schema2conv.Components = s.Components(projectId, schemaStr)
	schema := new(schemaHelper.SchemaRef)
	_commUtils.JsonDecode(schemaStr, schema)
	schema2conv.FillRefId(schema)
	return _commUtils.JsonEncode(schema)
}
