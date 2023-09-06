package service

import (
	"errors"
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	repo "github.com/aaronchen2k/deeptest/internal/server/modules/repo"
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

func (s *EnvironmentService) List(projectId int) (envs []model.Environment, err error) {
	envs, err = s.EnvironmentRepo.List(projectId)

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

func (s *EnvironmentService) Get(id, projectId uint) (env model.Environment, err error) {
	if id > 0 {
		env, err = s.EnvironmentRepo.Get(id)
	} else {
		env, _ = s.EnvironmentRepo.GetByProject(projectId)
	}

	if env.ID > 0 {
		env.Vars, err = s.EnvironmentRepo.GetVars(env.ID)
	}

	return
}

func (s *EnvironmentService) Copy(envId int) (err error) {
	err = s.EnvironmentRepo.Copy(envId)

	return
}

func (s *EnvironmentService) Create(env *model.Environment, projectId uint) (err error) {
	env.Sort = s.EnvironmentRepo.GetMaxOrder(projectId)
	err = s.EnvironmentRepo.Save(env)
	err = s.ProjectRepo.UpdateDefaultEnvironment(projectId, env.ID)

	return
}

func (s *EnvironmentService) Update(env *model.Environment) (err error) {
	err = s.EnvironmentRepo.Save(env)

	return
}

func (s *EnvironmentService) Delete(reqId uint) (err error) {
	err = s.EnvironmentRepo.Delete(reqId)

	return
}

func (s *EnvironmentService) Change(id, projectId int) (err error) {
	err = s.ProjectRepo.UpdateDefaultEnvironment(uint(projectId), uint(id))

	return
}

func (s *EnvironmentService) GetVar(id uint) (env model.EnvironmentVar, err error) {
	env, err = s.EnvironmentRepo.GetVar(id)

	return
}

func (s *EnvironmentService) CreateVar(po *model.EnvironmentVar) (err error) {
	temp, _ := s.EnvironmentRepo.GetVarByName(po.Name, 0, po.EnvironmentId)

	if temp.ID > 0 {
		err = errors.New("")
		return
	}

	err = s.EnvironmentRepo.SaveVar(po)

	return
}

func (s *EnvironmentService) UpdateVar(po *model.EnvironmentVar) (err error) {
	temp, _ := s.EnvironmentRepo.GetVarByName(po.Name, po.ID, po.EnvironmentId)
	if temp.ID > 0 {
		err = errors.New("")
		return
	}

	err = s.EnvironmentRepo.SaveVar(po)

	return
}

func (s *EnvironmentService) DeleteVar(id uint) (err error) {
	err = s.EnvironmentRepo.DeleteVar(id)

	return
}

func (s *EnvironmentService) ClearAllVar(environmentId uint) (err error) {
	err = s.EnvironmentRepo.ClearAllVar(environmentId)

	return
}

func (s *EnvironmentService) DisableShareVar(id uint) (err error) {
	err = s.EnvironmentRepo.DisableShareVar(id)

	return
}

func (s *EnvironmentService) DisableAllShareVar(interfaceId uint) (err error) {
	//interf, _ := s.InterfaceRepo.Get(interfaceId)
	//
	//err = s.EnvironmentRepo.DisableAllShareVar(interf.ProjectId)

	return
}

func (s *EnvironmentService) Save(req v1.EnvironmentReq) (id uint, err error) {
	var environment model.Environment
	copier.CopyWithOption(&environment, req, copier.Option{DeepCopy: true})
	environment.Sort = s.EnvironmentRepo.GetMaxOrder(req.ProjectId)
	err = s.EnvironmentRepo.SaveEnvironment(&environment)
	id = environment.ID
	return
}

func (s *EnvironmentService) Clone(id uint) (environment *model.Environment, err error) {
	environment, err = s.EnvironmentRepo.GetEnvironmentById(id)
	if err != nil {
		return
	}
	err = s.EnvironmentRepo.GetEnvironmentDetail(environment)
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

	err = s.EnvironmentRepo.SaveEnvironment(environment)
	return
}

func (s *EnvironmentService) DeleteEnvironment(id uint) (err error) {
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
	err = s.EnvironmentRepo.DeleteEnvironment(id)
	return
}

func (s *EnvironmentService) ListAll(projectId uint) (res []model.Environment, err error) {
	res, err = s.EnvironmentRepo.GetListByProjectId(projectId)
	return
}

func (s *EnvironmentService) SaveGlobal(projectId uint, req []v1.EnvironmentVariable) (err error) {
	var vars []model.EnvironmentVar
	copier.CopyWithOption(&vars, req, copier.Option{DeepCopy: true})

	err = s.EnvironmentRepo.SaveVars(projectId, 0, vars)

	return
}

func (s *EnvironmentService) ListGlobal(projectId uint) (res []model.EnvironmentVar, err error) {
	res, err = s.EnvironmentRepo.ListGlobalVar(projectId)
	return
}

func (s *EnvironmentService) SaveParams(req v1.EnvironmentParamsReq) (err error) {
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
	err = s.EnvironmentRepo.SaveParams(req.ProjectId, params)
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

func (s *EnvironmentService) ListParams(projectId uint) (ret map[string]interface{}, err error) {
	return s.EnvironmentRepo.ListParams(projectId)
}

func (s *EnvironmentService) SaveOrder(req v1.EnvironmentIdsReq) (err error) {
	return s.EnvironmentRepo.SaveOrder(req)
}

func (s *EnvironmentService) GetVarsByServer(serverId uint) (ret []domain.GlobalVar, err error) {
	server, _ := s.ServeServerRepo.Get(serverId)

	pos, _ := s.EnvironmentRepo.GetVars(server.EnvironmentId)

	for _, po := range pos {
		ret = append(ret, domain.GlobalVar{
			Name:       po.Name,
			LocalValue: po.LocalValue,
		})
	}

	return
}
func (s *EnvironmentService) GetVarsByEnv(envId uint) (ret []domain.GlobalVar, err error) {
	pos, _ := s.EnvironmentRepo.GetVars(envId)

	for _, po := range pos {
		ret = append(ret, domain.GlobalVar{
			Name:       po.Name,
			LocalValue: po.LocalValue,
		})
	}

	return
}
func (s *EnvironmentService) GetGlobalVars(projectId uint) (ret []domain.GlobalVar, err error) {
	pos, _ := s.EnvironmentRepo.ListGlobalVar(projectId)

	for _, v := range pos {
		ret = append(ret, domain.GlobalVar{
			Name:        v.Name,
			LocalValue:  v.LocalValue,
			RemoteValue: v.RemoteValue,
		})
	}

	return
}
func (s *EnvironmentService) GetGlobalParams(projectId uint) (ret []domain.GlobalParam, err error) {
	pos, _ := s.EnvironmentRepo.ListParamModel(projectId)

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
