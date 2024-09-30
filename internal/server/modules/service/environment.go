package service

import (
	"errors"
	v1 "github.com/deeptest-com/deeptest/cmd/server/v1/domain"
	"github.com/deeptest-com/deeptest/internal/pkg/consts"
	"github.com/deeptest-com/deeptest/internal/pkg/domain"
	"github.com/deeptest-com/deeptest/internal/server/modules/model"
	repo "github.com/deeptest-com/deeptest/internal/server/modules/repo"
	"github.com/jinzhu/copier"
)

type EnvironmentService struct {
	EnvironmentRepo *repo.EnvironmentRepo `inject:""`
	ScenarioRepo    *repo.ScenarioRepo    `inject:""`
	ProjectRepo     *repo.ProjectRepo     `inject:""`
	ServeRepo       *repo.ServeRepo       `inject:""`
	ServeServerRepo *repo.ServeServerRepo `inject:""`

	EndpointRepo          *repo.EndpointRepo          `inject:""`
	EndpointInterfaceRepo *repo.EndpointInterfaceRepo `inject:""`
	DebugInterfaceRepo    *repo.DebugInterfaceRepo    `inject:""`
}

func (s *EnvironmentService) List(tenantId consts.TenantId, projectId int) (envs []model.Environment, err error) {
	envs, err = s.EnvironmentRepo.List(tenantId, projectId)

	return
}

//func (s *EnvironmentService) ListVariableForExec(scenario modelRef.Scenario) (ret map[string]interface{}, err error) {
//	ret = map[string]interface{}{}
//
//	pos, err := s.EnvironmentRepo.ListVariableByProject(scenario.ProjectId)
//	if err != nil {
//		return
//	}
//
//	for _, po := range pos {
//		ret[po.Name] = po.RightValue
//	}
//
//	return
//}

func (s *EnvironmentService) Get(tenantId consts.TenantId, id, projectId uint) (env model.Environment, err error) {
	if id > 0 {
		env, err = s.EnvironmentRepo.Get(tenantId, id)
	} else {
		env, _ = s.EnvironmentRepo.GetByProject(tenantId, projectId)
	}

	if env.ID > 0 {
		env.Vars, err = s.EnvironmentRepo.GetVars(tenantId, env.ID)
	}

	return
}

func (s *EnvironmentService) Copy(tenantId consts.TenantId, envId int) (err error) {
	err = s.EnvironmentRepo.Copy(tenantId, envId)

	return
}

func (s *EnvironmentService) Create(tenantId consts.TenantId, env *model.Environment, projectId uint) (err error) {
	env.Sort = s.EnvironmentRepo.GetMaxOrder(tenantId, projectId)
	err = s.EnvironmentRepo.Save(tenantId, env)
	err = s.ProjectRepo.UpdateDefaultEnvironment(tenantId, projectId, env.ID)

	return
}

func (s *EnvironmentService) Update(tenantId consts.TenantId, env *model.Environment) (err error) {
	err = s.EnvironmentRepo.Save(tenantId, env)

	return
}

func (s *EnvironmentService) Delete(tenantId consts.TenantId, reqId uint) (err error) {
	err = s.EnvironmentRepo.Delete(tenantId, reqId)

	return
}

func (s *EnvironmentService) Change(tenantId consts.TenantId, id, projectId int) (err error) {
	err = s.ProjectRepo.UpdateDefaultEnvironment(tenantId, uint(projectId), uint(id))

	return
}

func (s *EnvironmentService) GetVar(tenantId consts.TenantId, id uint) (env model.EnvironmentVar, err error) {
	env, err = s.EnvironmentRepo.GetVar(tenantId, id)

	return
}

func (s *EnvironmentService) CreateVar(tenantId consts.TenantId, po *model.EnvironmentVar) (err error) {
	temp, _ := s.EnvironmentRepo.GetVarByName(tenantId, po.Name, 0, po.EnvironmentId)

	if temp.ID > 0 {
		err = errors.New("")
		return
	}

	err = s.EnvironmentRepo.SaveVar(tenantId, po)

	return
}

func (s *EnvironmentService) UpdateVar(tenantId consts.TenantId, po *model.EnvironmentVar) (err error) {
	temp, _ := s.EnvironmentRepo.GetVarByName(tenantId, po.Name, po.ID, po.EnvironmentId)
	if temp.ID > 0 {
		err = errors.New("")
		return
	}

	err = s.EnvironmentRepo.SaveVar(tenantId, po)

	return
}

func (s *EnvironmentService) DeleteVar(tenantId consts.TenantId, id uint) (err error) {
	err = s.EnvironmentRepo.DeleteVar(tenantId, id)

	return
}

func (s *EnvironmentService) ClearAllVar(tenantId consts.TenantId, environmentId uint) (err error) {
	err = s.EnvironmentRepo.ClearAllVar(tenantId, environmentId)

	return
}

func (s *EnvironmentService) DisableShareVar(tenantId consts.TenantId, id uint) (err error) {
	err = s.EnvironmentRepo.DisableShareVar(tenantId, id)

	return
}

func (s *EnvironmentService) DisableAllShareVar(interfaceId uint) (err error) {
	//interf, _ := s.InterfaceRepo.Get(interfaceId)
	//
	//err = s.EnvironmentRepo.DisableAllShareVar(interf.ProjectId)

	return
}

func (s *EnvironmentService) Save(tenantId consts.TenantId, req v1.EnvironmentReq) (id uint, err error) {
	var environment model.Environment
	copier.CopyWithOption(&environment, req, copier.Option{DeepCopy: true})
	if req.ID == 0 {
		environment.Sort = s.EnvironmentRepo.GetMaxOrder(tenantId, req.ProjectId)
	}
	err = s.EnvironmentRepo.SaveEnvironment(tenantId, &environment)
	id = environment.ID
	return
}

func (s *EnvironmentService) Clone(tenantId consts.TenantId, id uint) (environment *model.Environment, err error) {
	environment, err = s.EnvironmentRepo.GetEnvironmentById(tenantId, id)
	if err != nil {
		return
	}
	err = s.EnvironmentRepo.GetEnvironmentDetail(tenantId, environment)
	if err != nil {
		return
	}
	environment.ID = 0
	environment.Name = environment.Name + "_copy"

	for key, _ := range environment.ServeServers {
		environment.ServeServers[key].ID = 0
	}
	for key, _ := range environment.Vars {
		environment.Vars[key].ID = 0
	}

	err = s.EnvironmentRepo.SaveEnvironment(tenantId, environment)
	return
}

func (s *EnvironmentService) DeleteEnvironment(tenantId consts.TenantId, id uint) (err error) {
	/*
		var count int64
		count, err = s.ServeRepo.GetServerCountByEnvironmentId(id)
		if err != nil {
			return err
		}

		if count > 0 {
			err = fmt.Errorf("the environment has been associated with services and cannot be deleted")
			return err
		}
	*/
	err = s.EnvironmentRepo.DeleteEnvironment(tenantId, id)
	return
}

func (s *EnvironmentService) ListAll(tenantId consts.TenantId, projectId uint) (res []model.Environment, err error) {
	res, err = s.EnvironmentRepo.GetListByProjectId(tenantId, projectId)
	return
}

func (s *EnvironmentService) SaveGlobal(tenantId consts.TenantId, projectId uint, req []v1.EnvironmentVariable) (err error) {
	var vars []model.EnvironmentVar
	copier.CopyWithOption(&vars, req, copier.Option{DeepCopy: true})

	err = s.EnvironmentRepo.SaveVars(tenantId, projectId, 0, vars)

	return
}

func (s *EnvironmentService) ListGlobal(tenantId consts.TenantId, projectId uint) (res []model.EnvironmentVar, err error) {
	res, err = s.EnvironmentRepo.ListGlobalVar(tenantId, projectId)
	return
}

func (s *EnvironmentService) SaveParams(tenantId consts.TenantId, req v1.EnvironmentParamsReq) (err error) {
	var params []model.EnvironmentParam
	if req.Header != nil {
		params = append(params, s.getParams(req.ProjectId, "header", req.Header)...)
	}
	if req.Cookie != nil {
		params = append(params, s.getParams(req.ProjectId, "cookie", req.Cookie)...)
	}
	if req.Query != nil {
		params = append(params, s.getParams(req.ProjectId, "query", req.Query)...)
	}
	if req.Body != nil {
		params = append(params, s.getParams(req.ProjectId, "body", req.Body)...)
	}
	if req.Path != nil {
		params = append(params, s.getParams(req.ProjectId, "path", req.Path)...)
	}
	err = s.EnvironmentRepo.SaveParams(tenantId, req.ProjectId, params)
	return
}

func (s *EnvironmentService) getParams(projectId uint, in consts.ParamIn, ReqParams []v1.EnvironmentParam) (params []model.EnvironmentParam) {
	for _, item := range ReqParams {
		var param model.EnvironmentParam
		copier.CopyWithOption(&param, item, copier.Option{DeepCopy: true})

		param.ProjectId = projectId
		param.In = in

		params = append(params, param)
	}
	return
}

func (s *EnvironmentService) ListParams(tenantId consts.TenantId, projectId uint) (ret map[string]interface{}, err error) {
	return s.EnvironmentRepo.ListParams(tenantId, projectId)
}

func (s *EnvironmentService) SaveOrder(tenantId consts.TenantId, req v1.EnvironmentIdsReq) (err error) {
	return s.EnvironmentRepo.SaveOrder(tenantId, req)
}

func (s *EnvironmentService) GetVarsByServer(tenantId consts.TenantId, serverId uint) (ret []domain.GlobalVar, err error) {
	server, _ := s.ServeServerRepo.Get(tenantId, serverId)

	pos, _ := s.EnvironmentRepo.GetVars(tenantId, server.EnvironmentId)

	for _, po := range pos {
		ret = append(ret, domain.GlobalVar{
			Name:       po.Name,
			LocalValue: po.LocalValue,
		})
	}

	return
}
func (s *EnvironmentService) GetVarsByEnv(tenantId consts.TenantId, envId uint) (ret []domain.GlobalVar, err error) {
	pos, _ := s.EnvironmentRepo.GetVars(tenantId, envId)

	for _, po := range pos {
		ret = append(ret, domain.GlobalVar{
			Name:        po.Name,
			RemoteValue: po.RemoteValue,
		})
	}

	return
}
func (s *EnvironmentService) GetGlobalVars(tenantId consts.TenantId, projectId uint) (ret []domain.GlobalVar, err error) {
	pos, _ := s.EnvironmentRepo.ListGlobalVar(tenantId, projectId)

	for _, v := range pos {
		ret = append(ret, domain.GlobalVar{
			Name:        v.Name,
			LocalValue:  v.LocalValue,
			RemoteValue: v.RemoteValue,
		})
	}

	return
}
func (s *EnvironmentService) GetGlobalParams(tenantId consts.TenantId, projectId uint) (ret []domain.GlobalParam, err error) {
	pos, _ := s.EnvironmentRepo.ListParamModel(tenantId, projectId)

	for _, v := range pos {
		ret = append(ret, domain.GlobalParam{
			Name:         v.Name,
			Type:         v.Type,
			Required:     v.Required,
			DefaultValue: v.DefaultValue,
			In:           v.In,
		})
	}

	return
}

//func (s *EnvironmentService) GetDebugEnvByDebugInterfaceOrEndpointInterface(debugInterfaceId, endpointInterfaceId uint) (ret model.Environment, err error) {
//	var serveId uint
//
//	if debugInterfaceId > 0 {
//		debug, _ := s.DebugInterfaceRepo.Get(debugInterfaceId)
//		serveId = debug.ServerId
//	} else {
//		interf, _ := s.EndpointInterfaceRepo.Get(endpointInterfaceId)
//		endpoint, _ := s.EndpointRepo.Get(interf.EndpointId)
//		serveId = endpoint.ServerId
//	}
//
//	serveServer, _ := s.ServeServerRepo.Get(serveId)
//
//	ret, _ = s.EnvironmentRepo.Get(serveServer.EnvironmentId)
//
//	return
//}
