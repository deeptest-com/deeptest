package service

import (
	domain "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	"github.com/jinzhu/copier"
)

type DocumentService struct {
	EndpointRepo    *repo.EndpointRepo    `inject:""`
	ProjectRepo     *repo.ProjectRepo     `inject:""`
	ServeRepo       *repo.ServeRepo       `inject:""`
	EnvironmentRepo *repo.EnvironmentRepo `inject:""`
}

func (s *DocumentService) Content(req domain.DocumentReq) (res domain.DocumentRep, err error) {
	var projectId uint
	var endpointIds, serveIds []uint

	projectId, serveIds, endpointIds = req.ProjectId, req.ServeIds, req.EndpointIds

	var endpoints map[uint][]domain.EndpointReq
	endpoints, err = s.GetEndpoints(&projectId, &serveIds, &endpointIds)
	if err != nil {
		return
	}

	res = s.GetProject(projectId)

	res.Serves = s.GetServes(serveIds, endpoints)

	return
}

func (s *DocumentService) GetEndpoints(projectId *uint, serveIds, endpointIds *[]uint) (res map[uint][]domain.EndpointReq, err error) {
	var endpoints []*model.Endpoint

	if *projectId != 0 {
		endpoints, err = s.EndpointRepo.GetByProjectId(*projectId)
	} else if len(*serveIds) != 0 {
		endpoints, err = s.EndpointRepo.GetByServeIds(*serveIds)
	} else if len(*endpointIds) != 0 {
		endpoints, err = s.EndpointRepo.GetByEndpointIds(*endpointIds)
	}

	if err != nil {
		return
	}

	res = s.GetEndpointsInfo(projectId, serveIds, endpoints)

	return
}

func (s *DocumentService) GetEndpointsInfo(projectId *uint, serveIds *[]uint, endpoints []*model.Endpoint) (res map[uint][]domain.EndpointReq) {
	res = make(map[uint][]domain.EndpointReq)

	serves := make(map[uint]uint)
	for _, item := range endpoints {
		var endpoint domain.EndpointReq
		ret, _ := s.EndpointRepo.GetAll(item.ID, "v0.1.0")
		copier.CopyWithOption(&endpoint, &ret, copier.Option{IgnoreEmpty: true, DeepCopy: true})
		res[endpoint.ServeId] = append(res[endpoint.ServeId], endpoint)
		if _, ok := serves[endpoint.ServeId]; !ok {
			*serveIds = append(*serveIds, endpoint.ServeId)
			serves[endpoint.ServeId] = endpoint.ServeId
		}

		*projectId = endpoint.ProjectId
	}
	return
}

func (s *DocumentService) GetProject(projectId uint) (doc domain.DocumentRep) {
	project, err := s.ProjectRepo.Get(projectId)
	if err != nil {
		return
	}
	copier.CopyWithOption(&doc, &project, copier.Option{IgnoreEmpty: true, DeepCopy: true})

	doc.GlobalParams, _ = s.EnvironmentRepo.ListParams(projectId)
	doc.GlobalVars = s.GetGlobalVars(projectId)
	return
}

func (s *DocumentService) GetServes(serveIds []uint, endpoints map[uint][]domain.EndpointReq) (serves []domain.DocumentServe) {
	res, _ := s.ServeRepo.GetServesByIds(serveIds)
	schemas := s.GetSchemas(serveIds)
	securities := s.GetSecurities(serveIds)
	servers := s.GetServers(serveIds)
	for _, item := range res {
		var serve domain.DocumentServe
		copier.CopyWithOption(&serve, &item, copier.Option{IgnoreEmpty: true, DeepCopy: true})
		serve.Endpoints = endpoints[uint(serve.ID)]
		serve.Component = schemas[uint(serve.ID)]
		serve.Securities = securities[uint(serve.ID)]
		serve.Servers = servers[uint(serve.ID)]
		serves = append(serves, serve)
	}
	return
}

func (s *DocumentService) GetSchemas(serveIds []uint) (schemas map[uint][]domain.ServeSchemaReq) {
	schemas = make(map[uint][]domain.ServeSchemaReq)
	res, _ := s.ServeRepo.GetSchemas(serveIds)
	for _, item := range res {
		var schema domain.ServeSchemaReq
		copier.CopyWithOption(&schema, &item, copier.Option{IgnoreEmpty: true, DeepCopy: true})
		schemas[uint(schema.ServeId)] = append(schemas[uint(schema.ServeId)], schema)
	}
	return
}

func (s *DocumentService) GetServers(serveIds []uint) (servers map[uint][]domain.ServeServer) {
	servers = make(map[uint][]domain.ServeServer)
	res, _ := s.ServeRepo.GetServers(serveIds)
	for _, item := range res {
		var server domain.ServeServer
		copier.CopyWithOption(&server, &item, copier.Option{IgnoreEmpty: true, DeepCopy: true})
		servers[server.ServeId] = append(servers[server.ServeId], server)
	}
	return
}

func (s *DocumentService) GetSecurities(serveIds []uint) (securities map[uint][]domain.ServeSecurityReq) {
	securities = make(map[uint][]domain.ServeSecurityReq)
	res, _ := s.ServeRepo.GetSecurities(serveIds)
	for _, item := range res {
		var security domain.ServeSecurityReq
		copier.CopyWithOption(&security, &item, copier.Option{IgnoreEmpty: true, DeepCopy: true})
		securities[uint(security.ServeId)] = append(securities[uint(security.ServeId)], security)
	}
	return
}

func (s *DocumentService) GetGlobalVars(projectId uint) (globalVars []domain.EnvironmentParam) {
	res, _ := s.EnvironmentRepo.ListGlobalVar(projectId)
	copier.CopyWithOption(&globalVars, &res, copier.Option{IgnoreEmpty: true, DeepCopy: true})
	return
}
