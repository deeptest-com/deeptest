package service

import (
	agentExec "github.com/aaronchen2k/deeptest/internal/agent/exec"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
)

type DebugSceneService struct {
	EndpointInterfaceRepo *repo.EndpointInterfaceRepo `inject:""`
	EndpointRepo          *repo.EndpointRepo          `inject:""`
	ServeServerRepo       *repo.ServeServerRepo       `inject:""`
	ScenarioProcessorRepo *repo.ScenarioProcessorRepo `inject:""`
	EnvironmentRepo       *repo.EnvironmentRepo       `inject:""`
	DiagnoseInterfaceRepo *repo.DiagnoseInterfaceRepo `inject:""`
	ProfileRepo           *repo.ProfileRepo           `inject:""`

	ShareVarService *ShareVarService `inject:""`

	EnvironmentService *EnvironmentService `inject:""`
}

func (s *DebugSceneService) LoadScene(debugData *domain.DebugData, userId uint) (
	baseUrl string, shareVars []domain.GlobalVar, envVars []domain.GlobalVar,
	globalVars []domain.GlobalVar, globalParams []domain.GlobalParam) {

	debugServeId := debugData.ServeId
	debugServerId := debugData.ServerId

	if debugData.EndpointInterfaceId > 0 && (debugServeId <= 0 || debugServerId <= 0) {
		interf, _ := s.EndpointInterfaceRepo.Get(debugData.EndpointInterfaceId)
		endpoint, _ := s.EndpointRepo.Get(interf.EndpointId)

		if debugServeId <= 0 {
			debugServeId = endpoint.ServeId
		}
		if debugServerId <= 0 {
			debugServerId = endpoint.ServerId
		}
	}

	serveServer, _ := s.ServeServerRepo.Get(debugServerId)

	if debugData.DiagnoseInterfaceId > 0 {
		baseUrl = debugData.BaseUrl
	} else {
		baseUrl = serveServer.Url
	}

	envId := serveServer.EnvironmentId
	if userId != 0 {
		projectUserServer, _ := s.EnvironmentRepo.GetProjectUserServer(debugData.ProjectId, userId)
		if projectUserServer.ServerId != 0 {
			envId = projectUserServer.ServerId
		}
	}

	environment, _ := s.EnvironmentRepo.Get(envId)

	if debugData.ProjectId == 0 {
		debugData.ProjectId = environment.ProjectId
	}

	shareVars, _ = s.ShareVarService.ListForDebug(debugServeId, debugData.ScenarioProcessorId, debugData.UsedBy)
	envVars, _ = s.EnvironmentService.GetVarsByEnv(envId)
	globalVars, _ = s.EnvironmentService.GetGlobalVars(environment.ProjectId)
	globalParams, _ = s.EnvironmentService.GetGlobalParams(environment.ProjectId)

	//合并全局参数
	globalParams = agentExec.MergeGlobalParams(globalParams, debugData.GlobalParams)
	endpointInterfaceGlobalParams, _ := s.EndpointInterfaceRepo.GetGlobalParams(debugData.EndpointInterfaceId, debugData.ProjectId)
	globalParams = s.MergeGlobalParams(endpointInterfaceGlobalParams, globalParams)

	return
}

func (s *DebugSceneService) MergeGlobalParams(endpointInterfaceGlobalParams []model.EndpointInterfaceGlobalParam, globalParams []domain.GlobalParam) (ret []domain.GlobalParam) {

	for _, item := range globalParams {
		b := true
		for _, param := range endpointInterfaceGlobalParams {
			if param.Name == item.Name && param.In == item.In && param.Disabled {
				b = false
				break
			}
		}

		if b {
			ret = append(ret, item)
		}

	}

	return
}
