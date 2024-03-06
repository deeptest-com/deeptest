package repo

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
)

type ServeServerRepo struct {
	*BaseRepo       `inject:""`
	CategoryRepo    *CategoryRepo    `inject:""`
	EnvironmentRepo *EnvironmentRepo `inject:""`

	DebugInterfaceRepo    *DebugInterfaceRepo    `inject:""`
	EndpointInterfaceRepo *EndpointInterfaceRepo `inject:""`
	EndpointRepo          *EndpointRepo          `inject:""`
}

func (r *ServeServerRepo) Get(tenantId consts.TenantId, id uint) (res model.ServeServer, err error) {
	err = r.GetDB(tenantId).Where("NOT deleted").First(&res, id).Error
	return
}

func (r *ServeServerRepo) GetByDebugInfo(tenantId consts.TenantId, debugInterfaceId, endpointInterfaceId uint) (ret model.ServeServer, err error) {
	serverId := uint(0)

	if debugInterfaceId > 0 {
		debugInterface, _ := r.DebugInterfaceRepo.Get(tenantId, debugInterfaceId)
		serverId = debugInterface.ServerId

	} else if endpointInterfaceId > 0 {
		endpointInterface, _ := r.EndpointInterfaceRepo.Get(tenantId, endpointInterfaceId)
		endpoint, _ := r.EndpointRepo.Get(tenantId, endpointInterface.EndpointId)
		serverId = endpoint.ServerId
	}

	ret, _ = r.Get(tenantId, serverId)

	return
}

func (r *ServeServerRepo) GetByEndpoint(tenantId consts.TenantId, endpointId uint) (res model.ServeServer, err error) {
	endpoint, _ := r.EndpointRepo.Get(tenantId, endpointId)

	err = r.GetDB(tenantId).Where("NOT deleted").First(&res, endpoint.ServerId).Error
	return
}

func (r *ServeServerRepo) GetDefaultByServe(tenantId consts.TenantId, serveId uint) (ret model.ServeServer, err error) {
	servers := []model.ServeServer{}
	err = r.GetDB(tenantId).Where("serve_id = ? AND NOT deleted", serveId).
		Find(&servers).Error

	minEnvironmentSort := -1

	for _, server := range servers {
		var environment model.Environment
		environment, err = r.EnvironmentRepo.Get(tenantId, server.EnvironmentId)
		if err != nil {
			return
		}

		if minEnvironmentSort < 0 {
			minEnvironmentSort = int(environment.Sort)
			ret = server
			continue
		}

		if minEnvironmentSort > int(environment.Sort) {
			minEnvironmentSort = int(environment.Sort)
			ret = server
		}
	}

	return
}

func (r *ServeServerRepo) FindByServeAndExecEnv(tenantId consts.TenantId, serveId, environmentId uint) (ret model.ServeServer, err error) {
	err = r.GetDB(tenantId).
		Where("serve_id = ? AND environment_id =? AND NOT deleted", serveId, environmentId).
		First(&ret).Error

	return
}

func (r *ServeServerRepo) SetUrl(tenantId consts.TenantId, serveId uint, url string) (err error) {
	err = r.GetDB(tenantId).Model(model.ServeServer{}).Where("serve_id=? and  url=?", serveId, "http://localhost").Update("url", url).Error
	return
}

func (r *ServeServerRepo) BatchCreate(tenantId consts.TenantId, req []model.ServeServer) error {
	return r.GetDB(tenantId).Create(req).Error
}

func (r *ServeServerRepo) UpdateUrlByServeAndServer(tenantId consts.TenantId, serveId, serverId uint, url string) error {
	err := r.GetDB(tenantId).Model(&model.ServeServer{}).
		Where("serve_id = ?", serveId).
		Where("environment_id = ?", serverId).
		Update("url", url).Error

	return err
}

func (r *ServeServerRepo) GetByServeAndUrl(tenantId consts.TenantId, serveId uint, url string) (ret model.ServeServer, err error) {
	err = r.GetDB(tenantId).
		Where("serve_id = ? AND url =? AND NOT deleted", serveId, url).
		First(&ret).Error

	return
}
