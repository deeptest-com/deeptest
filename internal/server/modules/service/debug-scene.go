package service

import (
	agentExec "github.com/aaronchen2k/deeptest/internal/agent/exec"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
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
	ServeRepo             *repo.ServeRepo             `inject:""`

	ShareVarService *ShareVarService `inject:""`

	EnvironmentService *EnvironmentService `inject:""`
}

func (s *DebugSceneService) LoadScene(tenantId consts.TenantId, debugData *domain.DebugData, userIdForDisplay, environmentIdForExec uint) (
	baseUrl string, shareVars []domain.GlobalVar, envVars []domain.GlobalVar,
	globalVars []domain.GlobalVar, globalParams []domain.GlobalParam) {

	debugServeId := debugData.ServeId
	debugServerId := debugData.ServerId

	if debugData.EndpointInterfaceId > 0 && (debugServeId <= 0 || debugServerId <= 0) {
		interf, _ := s.EndpointInterfaceRepo.Get(tenantId, debugData.EndpointInterfaceId)
		endpoint, _ := s.EndpointRepo.Get(tenantId, interf.EndpointId)

		if debugServeId <= 0 {
			debugServeId = endpoint.ServeId
		}
		if debugServerId <= 0 {
			debugServerId = endpoint.ServerId
		}
	}

	if environmentIdForExec > 0 {
		serveServer, _ := s.ServeRepo.GetCurrServerByUser(tenantId, debugData.ProjectId, debugServeId, userIdForDisplay)
		debugServerId = serveServer.ID
	}

	serveServer, _ := s.ServeServerRepo.Get(tenantId, debugServerId)

	if debugData.DiagnoseInterfaceId > 0 {
		baseUrl = debugData.BaseUrl
	} else {
		baseUrl = serveServer.Url
	}

	// get environment
	envId := serveServer.EnvironmentId
	if environmentIdForExec > 0 { // exec loading
		envId = environmentIdForExec

	} else if userIdForDisplay != 0 { // display loading
		projectUserServer, _ := s.EnvironmentRepo.GetProjectUserServer(tenantId, debugData.ProjectId, userIdForDisplay)
		if projectUserServer.ServerId != 0 {
			envId = projectUserServer.ServerId
		}
	}

	environment, _ := s.EnvironmentRepo.Get(tenantId, envId)

	if debugData.ProjectId == 0 {
		debugData.ProjectId = environment.ProjectId
	}

	if userIdForDisplay > 0 {
		shareVars, _ = s.ShareVarService.ListForDebug(tenantId, debugServeId, debugData.ScenarioProcessorId, debugData.UsedBy)
		envVars, _ = s.EnvironmentService.GetVarsByEnv(tenantId, environmentIdForExec)
		globalVars, _ = s.EnvironmentService.GetGlobalVars(tenantId, environment.ProjectId)
	}

	// dealwith global params
	globalParams, _ = s.EnvironmentService.GetGlobalParams(tenantId, environment.ProjectId)

	//	if git s > 0 { // merge global params
	globalParams = agentExec.MergeGlobalParams(globalParams, *debugData.GlobalParams)
	endpointInterfaceGlobalParams, _ := s.EndpointInterfaceRepo.GetGlobalParams(tenantId, debugData.EndpointInterfaceId, debugData.ProjectId)
	globalParams = s.MergeGlobalParams(endpointInterfaceGlobalParams, globalParams)
	//	}

	if globalParams == nil {
		globalParams = []domain.GlobalParam{}
	}

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
