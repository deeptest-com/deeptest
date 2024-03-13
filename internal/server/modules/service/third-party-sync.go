package service

import (
	"encoding/json"
	"errors"
	"fmt"
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	integrationDomain "github.com/aaronchen2k/deeptest/integration/domain"
	leyan "github.com/aaronchen2k/deeptest/integration/leyan/service"
	"github.com/aaronchen2k/deeptest/internal/pkg/config"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/core/cron"
	"github.com/aaronchen2k/deeptest/internal/pkg/helper/openapi/thirdPart"
	serverConsts "github.com/aaronchen2k/deeptest/internal/server/consts"
	"github.com/aaronchen2k/deeptest/internal/server/core/cache"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	_commUtils "github.com/aaronchen2k/deeptest/pkg/lib/comm"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"github.com/getkin/kin-openapi/openapi3"
	"gorm.io/gorm"
	"time"
)

type ThirdPartySyncService struct {
	ThirdPartySyncRepo       *repo.ThirdPartySyncRepo    `inject:""`
	CategoryRepo             *repo.CategoryRepo          `inject:""`
	EndpointRepo             *repo.EndpointRepo          `inject:""`
	EndpointInterfaceRepo    *repo.EndpointInterfaceRepo `inject:""`
	UserRepo                 *repo.UserRepo              `inject:""`
	BaseRepo                 *repo.BaseRepo              `inject:""`
	RemoteService            *leyan.RemoteService        `inject:""`
	ServeService             *ServeService               `inject:""`
	EndpointService          *EndpointService            `inject:""`
	EndpointInterfaceService *EndpointInterfaceService   `inject:""`
	Cron                     *cron.ServerCron            `inject:""`
}

func (s *ThirdPartySyncService) GetToken(baseUrl string) (token string, err error) {
	loginByOauthReq := integrationDomain.LoginByOauthReq{
		LoginName: config.CONFIG.ThirdParty.Username,
		Password:  _commUtils.Sha256(config.CONFIG.ThirdParty.Password),
	}

	loginByOauthResData := s.RemoteService.LoginByOauth(loginByOauthReq, baseUrl)
	if loginByOauthResData.Code == "" {
		return "", errors.New("login fail")
	}

	getTokenFromCodeReq := integrationDomain.GetTokenFromCodeReq{
		Code: loginByOauthResData.Code,
	}

	getTokenFromCodeResData := s.RemoteService.GetTokenFromCode(getTokenFromCodeReq, baseUrl)
	token = getTokenFromCodeResData.Token

	return
}

func (s *ThirdPartySyncService) GetClasses(serviceCode, token string, baseUrl string) (classes []integrationDomain.FindClassByServiceCodeResData) {
	classes = s.RemoteService.LcQueryAgent(serviceCode, token, baseUrl)
	return
}

// GetFunctionsByClass 已废弃
func (s *ThirdPartySyncService) GetFunctionsByClass(serviceCode, classCode, token string, baseUrl string) (functions []string) {
	getFunctionsByClassReq := integrationDomain.GetFunctionsByClassReq{
		ServiceCode: serviceCode,
		ClassCode:   classCode,
	}
	getFunctionsByClassResData := s.RemoteService.GetFunctionsByClass(getFunctionsByClassReq, token, baseUrl)
	for _, v := range getFunctionsByClassResData {
		//不同步内部方法
		if v.MessageType == 1 {
			functions = append(functions, v.Code)
		}
	}

	return
}

func (s *ThirdPartySyncService) GetFunctionsByClassNew(serviceId, classCode, parentCodes, objId, token string, baseUrl string) (functions []string) {
	getFunctionsByClassReq := integrationDomain.QueryMsgReq{}
	getFunctionsByClassReq.ClassInfo.ParentCodes = parentCodes
	getFunctionsByClassReq.ClassInfo.ObjId = objId
	getFunctionsByClassReq.ClassInfo.Code = classCode
	getFunctionsByClassReq.ClassInfo.ServiceId = serviceId

	getFunctionsByClassResData := s.RemoteService.LcQueryMsg(getFunctionsByClassReq, token, baseUrl)
	for _, v := range getFunctionsByClassResData {
		//不同步继承方法和不允许被重写的内部方法
		//if v.IsExtend == consts.IntegrationFuncIsExtend || (v.MessageType == 0 && v.Overridable == consts.IntegrationFuncCanNotOverridable) {
		if v.IsExtend == consts.IntegrationFuncIsExtend {
			continue
		}

		functions = append(functions, v.Code)
	}

	return
}

func (s *ThirdPartySyncService) GetFunctionDetail(classCode, function, token string, baseUrl string) (data integrationDomain.MetaGetMethodDetailResData) {
	metaGetMethodDetailReq := integrationDomain.MetaGetMethodDetailReq{
		ClassName:   classCode,
		Method:      function,
		IncludeSelf: true,
	}
	data = s.RemoteService.MetaGetMethodDetail(metaGetMethodDetailReq, token, baseUrl)

	return
}

func (s *ThirdPartySyncService) SaveData(tenantId consts.TenantId) (err error) {
	thirdPartySyncStatus, _ := cache.GetCacheString("thirdPartySyncStatus")
	if thirdPartySyncStatus == "Start" {
		return
	}

	_ = cache.SetCache("thirdPartySyncStatus", "Start", 1*time.Hour)
	syncList, err := s.GetAllData(tenantId)
	if err != nil {
		return
	}

	for _, syncConfig := range syncList {
		if syncConfig.Switch == consts.SwitchOFF {
			continue
		}
		projectId, serveId, userId, baseUrl := syncConfig.ProjectId, syncConfig.ServeId, syncConfig.CreateUserId, syncConfig.Url

		token, err := s.GetToken(baseUrl)
		if err != nil {
			continue
		}

		classes := s.GetClasses(syncConfig.ServiceCode, token, baseUrl)
		for _, class := range classes {
			classCode := class.Code

			functionList := s.GetFunctionsByClassNew(class.ServiceId, classCode, class.ParentCodes, class.ObjId, token, baseUrl)
			if len(functionList) == 0 {
				continue
			}

			categoryId, err := s.SaveCategory(tenantId, class, projectId, syncConfig.ServeId)
			if err != nil {
				continue
			}

			for _, function := range functionList {
				path := "/" + syncConfig.ServiceCode + "/" + classCode + "/" + function
				functionDetail := s.GetFunctionDetail(classCode, function, token, baseUrl)
				if functionDetail.Code == "" {
					continue
				}

				fmt.Println(functionDetail)
				title := classCode + "-" + functionDetail.Code
				endpoint, err := s.EndpointRepo.GetByItem(tenantId, consts.ThirdPartySync, projectId, path, serveId, int64(categoryId))
				if err != nil && err != gorm.ErrRecordNotFound {
					continue
				}

				oldEndpointDetail, err := s.EndpointRepo.GetAll(tenantId, endpoint.ID, "v0.1.0")
				if err != nil && err != gorm.ErrRecordNotFound {
					continue
				}

				newEndpointDetail, err := s.GenerateEndpoint(functionDetail)
				if err != nil && err != gorm.ErrRecordNotFound {
					continue
				}

				oldEndpointDetail.ServeId = 0
				newEndpointDetail.ServeId = 0
				newSnapshot := _commUtils.JsonEncode(s.EndpointService.Yaml(tenantId, newEndpointDetail))
				if oldEndpointDetail.Snapshot == newSnapshot {
					continue
				}
				oldEndpointId := endpoint.ID

				oldEndpointDetailJson := _commUtils.JsonEncode(s.EndpointService.Yaml(tenantId, oldEndpointDetail))
				if endpoint.ID != 0 && oldEndpointDetail.Snapshot != oldEndpointDetailJson {
					s.EndpointRepo.UpdateSnapshot(tenantId, endpoint.ID, newSnapshot)
					continue
				}

				endpointId, err := s.SaveEndpoint(tenantId, title, projectId, serveId, userId, oldEndpointId, int64(categoryId), path, newSnapshot, consts.AutoAdd)
				if err != nil {
					continue
				}

				interfaceId, err := s.SaveEndpointInterface(tenantId, title, functionDetail, endpointId, projectId, path)
				if err != nil {
					continue
				}

				if err = s.SaveBody(tenantId, functionDetail, interfaceId); err != nil {
					continue
				}
			}
		}
	}

	cache.SetCache("thirdPartySyncStatus", "Stop", -1)
	return
}

func (s *ThirdPartySyncService) SaveCategory(tenantId consts.TenantId, class integrationDomain.FindClassByServiceCodeResData, projectId, serveId uint) (categoryId uint, err error) {
	rootNode, err := s.CategoryRepo.GetRootNode(tenantId, projectId, serverConsts.EndpointCategory)
	if err != nil {
		return
	}

	name := class.Code
	if class.Code != class.Name {
		name = class.Name + "(" + name + ")"
	}
	categoryReq := model.Category{
		Name:       name,
		ProjectId:  projectId,
		ServeId:    serveId,
		Type:       serverConsts.EndpointCategory,
		SourceType: consts.ThirdPartySync,
		ParentId:   int(rootNode.ID),
	}

	category, err := s.CategoryRepo.GetDetail(tenantId, categoryReq)
	if err != nil && err != gorm.ErrRecordNotFound {
		return
	}

	if category.ID != 0 {
		return category.ID, nil
	}

	err = s.CategoryRepo.Save(tenantId, &categoryReq)
	if err != nil {
		return 0, err
	}
	categoryId = categoryReq.ID

	return
}

func (s *ThirdPartySyncService) SaveEndpoint(tenantId consts.TenantId, title string, projectId, serveId, userId, oldEndpointId uint, categoryId int64, path, snapshot string, dataSyncType consts.DataSyncType) (endpointId uint, err error) {
	timeNow := time.Now()
	endpoint := model.Endpoint{
		Title:       title,
		ProjectId:   projectId,
		ServeId:     serveId,
		Path:        path,
		Status:      1,
		Snapshot:    snapshot,
		SourceType:  consts.ThirdPartySync,
		ChangedTime: &timeNow,
	}
	if dataSyncType == consts.FullCover {
		endpoint.ChangedStatus = consts.NoChanged
	}

	if oldEndpointId == 0 || dataSyncType == consts.Add {
		endpoint.CategoryId = categoryId
	}

	if userId != 0 {
		user, err := s.UserRepo.FindById(tenantId, userId)
		if err != nil {
			return 0, err
		}
		endpoint.CreateUser = user.Username
	}

	if oldEndpointId != 0 {
		endpoint.ID = oldEndpointId
	}

	err = s.EndpointRepo.SaveAll(tenantId, &endpoint)
	if err != nil {
		return 0, err
	}

	endpointId = endpoint.ID

	return
}

func (s *ThirdPartySyncService) SaveEndpointInterface(tenantId consts.TenantId, title string, functionDetail integrationDomain.MetaGetMethodDetailResData, endpointId, projectId uint, path string) (interfaceId uint, err error) {
	endpointInterface := model.EndpointInterface{
		InterfaceBase: model.InterfaceBase{
			Name:      title,
			ProjectId: projectId,
			InterfaceConfigBase: model.InterfaceConfigBase{
				Url:      path,
				Method:   "POST",
				BodyType: s.getBodyType(functionDetail.RequestType),
				Version:  "v0.1.0",
			},
		},
		EndpointId:    endpointId,
		SourceType:    consts.ThirdPartySync,
		ResponseCodes: "200",
	}
	err = s.EndpointInterfaceRepo.Save(tenantId, &endpointInterface)

	interfaceId = endpointInterface.ID
	return
}

func (s *ThirdPartySyncService) getBodyType(requestType string) (bodyType consts.HttpContentType) {
	if requestType == "JSON" {
		bodyType = consts.ContentTypeJSON
	} else if requestType == "FORM" {
		bodyType = consts.ContentTypeFormData
	}

	return
}

func (s *ThirdPartySyncService) getRequestBody(functionDetail integrationDomain.MetaGetMethodDetailResData) (requestBody string) {
	if functionDetail.RequestType == "JSON" {
		requestBody = functionDetail.RequestBody
	} else if functionDetail.RequestType == "FORM" {
		requestBody = functionDetail.RequestFormBody
	}

	return
}

func (s *ThirdPartySyncService) GetSchema(bodyString, requestType string) (schema *openapi3.SchemaRef) {
	if bodyString == "" {
		return
	}

	var schemas thirdPart.Schemas
	_ = json.Unmarshal([]byte(bodyString), &schemas)

	if requestType == "JSON" {
		schemas = schemas["root"].Properties
	}

	return thirdPart.NewThirdPart2conv().Convert(schemas)

}

func (s *ThirdPartySyncService) GenerateEndpoint(functionDetail integrationDomain.MetaGetMethodDetailResData) (res model.Endpoint, err error) {
	functionBody := functionDetail.RequestBody
	if functionDetail.RequestType == "FORM" {
		functionBody = functionDetail.RequestFormBody
	}

	requestBodySchema := s.GetSchema(functionBody, functionDetail.RequestType)
	requestSchemaString, _ := json.Marshal(requestBodySchema)

	requestBodyItem := model.EndpointInterfaceRequestBodyItem{}
	requestBodyItem.Content = string(requestSchemaString)

	requestBody := model.EndpointInterfaceRequestBody{}
	requestBody.MediaType = s.getBodyType(functionDetail.RequestType).String()
	requestBody.SchemaItem = requestBodyItem

	responseBodySchema := s.GetSchema(functionDetail.ResponseBody, functionDetail.RequestType)
	responseSchemaString, _ := json.Marshal(responseBodySchema)

	responseBodyItem := model.EndpointInterfaceResponseBodyItem{}
	responseBodyItem.Content = string(responseSchemaString)

	responseBody := model.EndpointInterfaceResponseBody{}
	responseBody.MediaType = s.getBodyType(functionDetail.RequestType).String()
	responseBody.SchemaItem = responseBodyItem
	responseBody.Code = "200"

	endpointInterface := model.EndpointInterface{}
	endpointInterface.BodyType = s.getBodyType(functionDetail.RequestType)
	endpointInterface.Method = consts.POST
	endpointInterface.RequestBody = requestBody
	endpointInterface.ResponseCodes = "200"
	endpointInterface.ResponseBodies = []model.EndpointInterfaceResponseBody{responseBody}

	res.Title = functionDetail.ClassCode + "-" + functionDetail.Code
	res.Path = "/" + functionDetail.ServiceCode + "/" + functionDetail.ClassCode + "/" + functionDetail.Code
	res.Interfaces = []model.EndpointInterface{endpointInterface}

	return
}

func (s *ThirdPartySyncService) SaveBody(tenantId consts.TenantId, functionDetail integrationDomain.MetaGetMethodDetailResData, interfaceId uint) (err error) {
	functionBody := functionDetail.RequestBody
	if functionDetail.RequestType == "FORM" {
		functionBody = functionDetail.RequestFormBody
	}

	requestBodySchema := s.GetSchema(functionBody, functionDetail.RequestType)
	responseBodySchema := s.GetSchema(functionDetail.ResponseBody, functionDetail.RequestType)
	fmt.Println(requestBodySchema, responseBodySchema)

	requestSchemaString, _ := json.Marshal(requestBodySchema)

	generateFromRequestReq := v1.GenerateFromRequestReq{
		ContentType: s.getBodyType(functionDetail.RequestType).String(),
		Data:        string(requestSchemaString),
		InterfaceId: interfaceId,
	}
	_, err = s.EndpointInterfaceService.GenerateFromRequest(tenantId, generateFromRequestReq)
	if err != nil {
		return
	}

	responseSchemaString, _ := json.Marshal(responseBodySchema)

	generateFromResponseReq := v1.GenerateFromResponseReq{
		Code:        "200",
		ContentType: s.getBodyType(functionDetail.ResponseType).String(),
		Data:        string(responseSchemaString),
		InterfaceId: interfaceId,
	}
	_, err = s.EndpointInterfaceService.GenerateFromResponse(tenantId, generateFromResponseReq)
	return
}

func (s *ThirdPartySyncService) UpdateExecTimeById(tenantId consts.TenantId, id uint) (err error) {
	return s.ThirdPartySyncRepo.UpdateExecTimeById(tenantId, id)
}

func (s *ThirdPartySyncService) AddThirdPartySyncCron(tenantId consts.TenantId) {
	name := fmt.Sprintf("ThirdPartySync_%v", tenantId)

	s.Cron.RemoveTask(name)

	s.Cron.AddCommonTask(name, "* * * * *", func() {
		err := s.SaveData(tenantId)
		if err != nil {
			logUtils.Error("third party 定时导入任务失败，错误原因：" + err.Error())
		}

		logUtils.Info("third party 定时任务结束")
	})
}

func (s *ThirdPartySyncService) GetAllData(tenantId consts.TenantId) (res []model.ThirdPartySync, err error) {
	return s.ThirdPartySyncRepo.AllData(tenantId)
}

func (s *ThirdPartySyncService) SyncFunctionBody(tenantId consts.TenantId, projectId, serveId, interfaceId uint, classCode, functionCode string) (err error) {
	syncConfig, err := s.ThirdPartySyncRepo.GetByProjectAndServe(tenantId, projectId, serveId)
	if err != nil {
		return
	}

	token, err := s.GetToken(syncConfig.Url)
	if err != nil {
		return
	}

	functionDetail := s.GetFunctionDetail(classCode, functionCode, token, syncConfig.Url)
	if functionDetail.Code == "" {
		return
	}

	err = s.SaveBody(tenantId, functionDetail, interfaceId)

	return
}

func (s *ThirdPartySyncService) ImportThirdPartyFunctions(tenantId consts.TenantId, req v1.ImportEndpointDataReq) (err error) {
	token, err := s.GetToken(req.FilePath)
	if err != nil {
		return
	}

	for _, function := range req.FunctionCodes {
		functionDetail := s.GetFunctionDetail(req.ClassCode, function, token, req.FilePath)
		if functionDetail.Code == "" {
			continue
		}

		//path := "/" + functionDetail.ServiceCode + "/" + req.ClassCode + "/" + function
		path := "/" + req.ClassCode + "/" + function
		if req.AddServicePrefix {
			path = "/" + functionDetail.ServiceCode + path
		}
		title := req.ClassCode + "-" + functionDetail.Code

		endpoint, err := s.EndpointRepo.GetByItem(tenantId, consts.ThirdPartySync, req.ProjectId, path, req.ServeId, req.CategoryId)
		if err != nil && err != gorm.ErrRecordNotFound {
			continue
		}

		newEndpointDetail, err := s.GenerateEndpoint(functionDetail)
		if err != nil && err != gorm.ErrRecordNotFound {
			continue
		}

		newSnapshot := _commUtils.JsonEncode(s.EndpointService.Yaml(tenantId, newEndpointDetail))

		if req.DataSyncType == consts.AutoAdd && endpoint.ID != 0 {
			oldEndpointDetail, err := s.EndpointRepo.GetAll(tenantId, endpoint.ID, "v0.1.0")
			if err != nil && err != gorm.ErrRecordNotFound {
				continue
			}
			oldEndpointDetail.ServeId = 0

			if oldEndpointDetail.Snapshot == newSnapshot {
				continue
			}

			oldEndpointDetailJson := _commUtils.JsonEncode(s.EndpointService.Yaml(tenantId, oldEndpointDetail))
			if oldEndpointDetail.Snapshot != oldEndpointDetailJson {
				err = s.EndpointRepo.UpdateSnapshot(tenantId, endpoint.ID, newSnapshot)
				continue
			}
		}

		var oldEndpointId uint
		if req.DataSyncType != consts.Add && endpoint.ID != 0 {
			oldEndpointId = endpoint.ID
		}

		endpointId, err := s.SaveEndpoint(tenantId, title, req.ProjectId, req.ServeId, req.UserId, oldEndpointId, req.CategoryId, path, newSnapshot, req.DataSyncType)
		if err != nil {
			continue
		}

		interfaceId, err := s.SaveEndpointInterface(tenantId, title, functionDetail, endpointId, req.ProjectId, path)
		if err != nil {
			continue
		}

		if err = s.SaveBody(tenantId, functionDetail, interfaceId); err != nil {
			continue
		}
	}

	return
}

func (s *ThirdPartySyncService) ListFunctionsByClass(baseUrl, classCode string) (res []integrationDomain.GetFunctionDetailsByClassResData, err error) {
	token, err := s.GetToken(baseUrl)
	if err != nil {
		err = errors.New("您输入的环境URL地址有误")
		return
	}

	functionList, err := s.RemoteService.GetFunctionDetailsByClass(classCode, token, baseUrl)
	if err != nil {
		err = errors.New("您输入的环境URL地址有误")
		return
	}

	if len(functionList) == 0 {
		err = errors.New("您输入的智能体未查找到消息列表，请检查智能体名是否输入有误")
		return
	}

	for _, function := range functionList {
		if function.MessageType == 1 {
			res = append(res, function)
		}
	}

	return
}
