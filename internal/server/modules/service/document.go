package service

import (
	"encoding/base64"
	"fmt"
	domain "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	commUtils "github.com/aaronchen2k/deeptest/pkg/lib/comm"
	"github.com/jinzhu/copier"
	"strconv"
	"strings"
)

type DocumentService struct {
	EndpointRepo          *repo.EndpointRepo          `inject:""`
	ProjectRepo           *repo.ProjectRepo           `inject:""`
	ServeRepo             *repo.ServeRepo             `inject:""`
	EnvironmentRepo       *repo.EnvironmentRepo       `inject:""`
	EndpointDocumentRepo  *repo.EndpointDocumentRepo  `inject:""`
	EndpointSnapshotRepo  *repo.EndpointSnapshotRepo  `inject:""`
	EndpointInterfaceRepo *repo.EndpointInterfaceRepo `inject:""`
	EndpointService       *EndpointService            `inject:""`
	ServeService          *ServeService               `inject:""`
}

const (
	EncryptKey = "docencryptkey123"
)

func (s *DocumentService) Content(tenantId consts.TenantId, req domain.DocumentReq) (res domain.DocumentRep, err error) {
	var projectId, documentId uint
	var endpointIds, serveIds []uint
	var needDetail bool

	projectId, serveIds, endpointIds, documentId, needDetail = req.ProjectId, req.ServeIds, req.EndpointIds, req.DocumentId, req.NeedDetail

	var endpoints map[uint][]domain.EndpointReq
	endpoints, err = s.GetEndpoints(tenantId, &projectId, &serveIds, &endpointIds, documentId, needDetail)
	if err != nil {
		return
	}

	res = s.GetProject(tenantId, projectId)

	res.Serves = s.GetServes(tenantId, serveIds, endpoints)

	return
}

func (s *DocumentService) GetEndpoints(tenantId consts.TenantId, projectId *uint, serveIds, endpointIds *[]uint, documentId uint, needDetail bool) (res map[uint][]domain.EndpointReq, err error) {
	var endpoints []*model.Endpoint

	if documentId != 0 {
		endpoints, err = s.EndpointSnapshotRepo.GetByDocumentId(tenantId, documentId)
	} else if *projectId != 0 {
		endpoints, err = s.EndpointRepo.GetByProjectId(tenantId, *projectId, needDetail)
	} else if len(*serveIds) != 0 {
		endpoints, err = s.EndpointRepo.GetByServeIds(tenantId, *serveIds, needDetail)
	} else if len(*endpointIds) != 0 {
		endpoints, err = s.EndpointRepo.GetByEndpointIds(tenantId, *endpointIds, needDetail)
	}

	if err != nil {
		return
	}

	for key, endpoint := range endpoints {
		for k, _ := range endpoint.Interfaces {
			s.MergeGlobalParams(&endpoints[key].Interfaces[k])
		}
	}

	s.FillRefId(tenantId, endpoints)

	res = s.GetEndpointsInfo(tenantId, projectId, serveIds, endpoints)

	return
}

func (s *DocumentService) GetEndpointsInfo(tenantId consts.TenantId, projectId *uint, serveIds *[]uint, endpoints []*model.Endpoint) (res map[uint][]domain.EndpointReq) {
	res = make(map[uint][]domain.EndpointReq)

	serves := make(map[uint]uint)
	for _, item := range endpoints {
		var endpoint domain.EndpointReq
		//ret, _ := s.EndpointRepo.GetAll(item.ID, "v0.1.0")
		copier.CopyWithOption(&endpoint, &item, copier.Option{IgnoreEmpty: true, DeepCopy: true})
		res[endpoint.ServeId] = append(res[endpoint.ServeId], endpoint)
		if _, ok := serves[endpoint.ServeId]; !ok {
			*serveIds = append(*serveIds, endpoint.ServeId)
			serves[endpoint.ServeId] = endpoint.ServeId
		}

		*projectId = endpoint.ProjectId
	}
	return
}

func (s *DocumentService) GetProject(tenantId consts.TenantId, projectId uint) (doc domain.DocumentRep) {
	project, err := s.ProjectRepo.Get(tenantId, projectId)
	if err != nil {
		return
	}
	copier.CopyWithOption(&doc, &project, copier.Option{IgnoreEmpty: true, DeepCopy: true})

	doc.GlobalParams, _ = s.EnvironmentRepo.ListParams(tenantId, projectId)
	doc.GlobalVars = s.GetGlobalVars(tenantId, projectId)
	doc.Components = s.GetSchemas(tenantId, projectId)
	return
}

func (s *DocumentService) GetServes(tenantId consts.TenantId, serveIds []uint, endpoints map[uint][]domain.EndpointReq) (serves []domain.DocumentServe) {
	res, _ := s.ServeRepo.GetServesByIds(tenantId, serveIds)
	//schemas := s.GetSchemas(serveIds)
	securities := s.GetSecurities(tenantId, serveIds)
	servers := s.GetServers(tenantId, serveIds)
	for _, item := range res {
		var serve domain.DocumentServe
		copier.CopyWithOption(&serve, &item, copier.Option{IgnoreEmpty: true, DeepCopy: true})
		serve.Endpoints = endpoints[uint(serve.ID)]
		//serve.Component = schemas[uint(serve.ID)]
		serve.Securities = securities[uint(serve.ID)]
		serve.Servers = servers[uint(serve.ID)]
		s.mocks(tenantId, serve.Endpoints, serve.Servers)
		serves = append(serves, serve)
	}
	return
}

func (s *DocumentService) mocks(tenantId consts.TenantId, endpoints []domain.EndpointReq, servers []domain.ServeServer) {
	var serve []model.ServeServer
	copier.CopyWithOption(&serve, &servers, copier.Option{IgnoreEmpty: true, DeepCopy: true})
	for key1, endpoint := range endpoints {
		for key2, interf := range endpoint.Interfaces {
			var interfaceDetail model.EndpointInterface
			copier.CopyWithOption(&interfaceDetail, &interf, copier.Option{IgnoreEmpty: true, DeepCopy: true})
			endpoints[key1].Interfaces[key2].Mock = s.mock(serve, interfaceDetail)
		}
	}

}

func (s *DocumentService) GetSchemas(tenantId consts.TenantId, projectId uint) (schemas []domain.ServeSchemaReq) {
	res, _ := s.ServeService.GetComponents(tenantId, projectId)
	for _, item := range res {
		var schema domain.ServeSchemaReq
		copier.CopyWithOption(&schema, &item, copier.Option{IgnoreEmpty: true, DeepCopy: true})
		schemas = append(schemas, schema)
	}
	return
}

func (s *DocumentService) GetServers(tenantId consts.TenantId, serveIds []uint) (servers map[uint][]domain.ServeServer) {
	servers = make(map[uint][]domain.ServeServer)
	res, _ := s.ServeRepo.GetServers(tenantId, serveIds)
	for _, item := range res {
		var server domain.ServeServer
		copier.CopyWithOption(&server, &item, copier.Option{IgnoreEmpty: true, DeepCopy: true})
		servers[server.ServeId] = append(servers[server.ServeId], server)
	}
	return
}

func (s *DocumentService) GetSecurities(tenantId consts.TenantId, serveIds []uint) (securities map[uint][]domain.ServeSecurityReq) {
	securities = make(map[uint][]domain.ServeSecurityReq)
	res, _ := s.ServeRepo.GetSecurities(tenantId, serveIds)
	for _, item := range res {
		var security domain.ServeSecurityReq
		copier.CopyWithOption(&security, &item, copier.Option{IgnoreEmpty: true, DeepCopy: true})
		securities[uint(security.ServeId)] = append(securities[uint(security.ServeId)], security)
	}
	return
}

func (s *DocumentService) GetGlobalVars(tenantId consts.TenantId, projectId uint) (globalVars []domain.EnvironmentParam) {
	res, _ := s.EnvironmentRepo.ListGlobalVar(tenantId, projectId)
	copier.CopyWithOption(&globalVars, &res, copier.Option{IgnoreEmpty: true, DeepCopy: true})
	return
}

func (s *DocumentService) GetDocumentVersionList(tenantId consts.TenantId, projectId uint, needLatest bool) (documents []model.EndpointDocument, err error) {
	if needLatest {
		latestDocument := model.EndpointDocument{
			Name:    "实时版本",
			Version: "latest",
		}
		documents = append(documents, latestDocument)
	}

	documentsTmp, err := s.EndpointDocumentRepo.ListByProject(tenantId, projectId)
	if err != nil {
		return
	}

	documents = append(documents, documentsTmp...)
	return
}

func (s *DocumentService) Publish(tenantId consts.TenantId, req domain.DocumentVersionReq, projectId uint) (documentId uint, err error) {
	documentId, err = s.EndpointSnapshotRepo.BatchCreateSnapshot(tenantId, req, projectId)
	return
}

func (s *DocumentService) RemoveSnapshot(tenantId consts.TenantId, snapshotId uint) (err error) {
	err = s.EndpointSnapshotRepo.DeleteById(tenantId, snapshotId)
	return
}

func (s *DocumentService) UpdateSnapshotContent(tenantId consts.TenantId, id uint, endpoint model.Endpoint) (err error) {
	err = s.EndpointSnapshotRepo.UpdateContent(tenantId, id, endpoint)
	return
}

func (s *DocumentService) UpdateDocument(tenantId consts.TenantId, req domain.UpdateDocumentVersionReq) (err error) {
	err = s.EndpointDocumentRepo.Update(tenantId, req)
	return
}

func (s *DocumentService) GenerateShareLink(tenantId consts.TenantId, req domain.DocumentShareReq) (link string, err error) {
	encryptValue := strconv.Itoa(int(req.ProjectId)) + "-" + strconv.Itoa(int(req.DocumentId)) + "-" + strconv.Itoa(int(req.EndpointId))
	res, err := commUtils.AesCBCEncrypt([]byte(encryptValue), []byte(EncryptKey))
	link = base64.RawURLEncoding.EncodeToString(res)
	return
}

func (s *DocumentService) DecryptShareLink(tenantId consts.TenantId, link string) (req domain.DocumentShareReq, err error) {
	linkByte, err := base64.RawURLEncoding.DecodeString(link)
	if err != nil {
		return
	}

	decryptValue, err := commUtils.AesCBCDecrypt(linkByte, []byte(EncryptKey))
	if err != nil {
		return
	}

	DocumentShareArr := strings.Split(string(decryptValue), "-")

	projectId, _ := strconv.Atoi(DocumentShareArr[0])
	documentId, _ := strconv.Atoi(DocumentShareArr[1])
	endpointId, _ := strconv.Atoi(DocumentShareArr[2])
	req.ProjectId = uint(projectId)
	req.DocumentId = uint(documentId)
	req.EndpointId = uint(endpointId)

	return
}

func (s *DocumentService) GetEndpointsByShare(tenantId consts.TenantId, projectId, endpointId *uint, serveIds *[]uint, documentId uint) (res map[uint][]domain.EndpointReq, err error) {
	var endpoints []*model.Endpoint
	if documentId != 0 {
		if *endpointId != 0 {
			endpoints, err = s.EndpointSnapshotRepo.GetByDocumentIdAndEndpointId(tenantId, documentId, *endpointId)
		} else {
			endpoints, err = s.EndpointSnapshotRepo.GetByDocumentId(tenantId, documentId)
		}
	} else if *projectId != 0 {
		if *endpointId != 0 {
			endpoints, err = s.EndpointRepo.GetByEndpointIds(tenantId, []uint{*endpointId}, false)
		} else {
			endpoints, err = s.EndpointRepo.GetByProjectId(tenantId, *projectId, false)
		}
	}
	if err != nil {
		return
	}

	if err != nil {
		return
	}

	res = s.GetEndpointsInfo(tenantId, projectId, serveIds, endpoints)

	return
}

func (s *DocumentService) ContentByShare(tenantId consts.TenantId, link string) (res domain.DocumentRep, err error) {
	var projectId, documentId, endpointId uint
	var serveIds []uint

	req, err := s.DecryptShareLink(tenantId, link)
	if err != nil {
		return
	}

	projectId, endpointId, documentId = req.ProjectId, req.EndpointId, req.DocumentId

	endpoints, err := s.GetEndpointsByShare(tenantId, &projectId, &endpointId, &serveIds, documentId)
	if err != nil {
		return
	}

	var version string
	if documentId == 0 {
		version = "latest"
	} else {
		document, err := s.EndpointDocumentRepo.GetById(tenantId, documentId)
		if err != nil {
			return res, err
		}
		version = document.Version
		documentId = document.ID
	}

	res = s.GetProject(tenantId, projectId)

	res.Serves = s.GetServes(tenantId, serveIds, endpoints)
	res.Version = version
	res.DocumentId = documentId

	return
}

func (s *DocumentService) GetDocumentDetail(tenantId consts.TenantId, documentId, endpointId, interfaceId uint) (res map[string]interface{}, err error) {
	var interfaceDetail model.EndpointInterface

	if documentId == 0 {
		interfaceDetail, err = s.EndpointInterfaceRepo.GetDetail(tenantId, interfaceId)
	} else {
		interfaceDetail, err = s.EndpointSnapshotRepo.GetInterfaceDetail(tenantId, documentId, endpointId, interfaceId)
	}

	if err != nil {
		return
	}

	endpoint, err := s.EndpointRepo.Get(tenantId, interfaceDetail.EndpointId)
	if err != nil {
		return
	}

	serveId := endpoint.ServeId
	serves, err := s.ServeRepo.GetServers(tenantId, []uint{serveId})
	if err != nil {
		return
	}

	s.EndpointService.SchemaConv(tenantId, &interfaceDetail, interfaceDetail.ProjectId)
	s.MergeGlobalParams(&interfaceDetail)

	res = make(map[string]interface{})
	res["interface"] = interfaceDetail
	res["servers"] = serves
	res["mock"] = s.mock(serves, interfaceDetail)

	return
}

func (s *DocumentService) mock(serves []model.ServeServer, interfaceDetail model.EndpointInterface) (ret []interface{}) {
	url := s.getMockEnvironment(serves)
	if url == "" {
		return
	}
	responseBodies := interfaceDetail.ResponseBodies
	path := interfaceDetail.Url
	for _, item := range responseBodies {
		ret = append(ret, map[string]interface{}{"name": item.Code, "url": fmt.Sprintf("%s%s?id=%d&code=%s", url, path, interfaceDetail.ID, item.Code)})
	}

	return
}

func (s *DocumentService) getMockEnvironment(serves []model.ServeServer) string {
	for _, item := range serves {
		if item.EnvironmentName == "Mock环境" {
			return item.Url
		}
	}

	return ""
}

func (s *DocumentService) MergeGlobalParams(endpointInterface *model.EndpointInterface) {
	for _, globalParam := range endpointInterface.GlobalParams {
		if globalParam.Disabled {
			continue
		}
		if globalParam.In == consts.ParamInQuery {
			endpointInterface.Params = append(endpointInterface.Params, model.EndpointInterfaceParam{SchemaParam: model.SchemaParam{Name: globalParam.Name, Type: string(globalParam.Type), Example: globalParam.DefaultValue, Default: globalParam.DefaultValue, Value: globalParam.DefaultValue, IsGlobal: true}})
		} else if globalParam.In == consts.ParamInCookie {
			endpointInterface.Cookies = append(endpointInterface.Cookies, model.EndpointInterfaceCookie{SchemaParam: model.SchemaParam{Name: globalParam.Name, Type: string(globalParam.Type), Example: globalParam.DefaultValue, Default: globalParam.DefaultValue, Value: globalParam.DefaultValue, IsGlobal: true}})
		} else if globalParam.In == consts.ParamInHeader {
			endpointInterface.Headers = append(endpointInterface.Headers, model.EndpointInterfaceHeader{SchemaParam: model.SchemaParam{Name: globalParam.Name, Type: string(globalParam.Type), Example: globalParam.DefaultValue, Default: globalParam.DefaultValue, Value: globalParam.DefaultValue, IsGlobal: true}})
		}

	}

}

func (s *DocumentService) FillRefId(tenantId consts.TenantId, endpoints []*model.Endpoint) {
	if len(endpoints) == 0 {
		return
	}
	projectId := endpoints[0].ProjectId
	components := s.ServeService.Components(tenantId, projectId)
	for _, endpoint := range endpoints {
		s.EndpointService.SchemasConv(tenantId, endpoint, components)
	}

}
