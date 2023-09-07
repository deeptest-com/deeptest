package repo

import (
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

func (r *ServeServerRepo) Get(id uint) (res model.ServeServer, err error) {
	err = r.DB.Where("NOT deleted").First(&res, id).Error
	return
}

func (r *ServeServerRepo) GetByDebugInfo(debugInterfaceId, endpointInterfaceId uint) (ret model.ServeServer, err error) {
	serverId := uint(0)

	if debugInterfaceId > 0 {
		debugInterface, _ := r.DebugInterfaceRepo.Get(debugInterfaceId)
		serverId = debugInterface.ServerId

	} else if endpointInterfaceId > 0 {
		endpointInterface, _ := r.EndpointInterfaceRepo.Get(endpointInterfaceId)
		endpoint, _ := r.EndpointRepo.Get(endpointInterface.EndpointId)
		serverId = endpoint.ServerId
	}

	ret, _ = r.Get(serverId)

	return
}

func (r *ServeServerRepo) GetByEndpoint(endpointId uint) (res model.ServeServer, err error) {
	endpoint, _ := r.EndpointRepo.Get(endpointId)

	err = r.DB.Where("NOT deleted").First(&res, endpoint.ServerId).Error
	return
}

func (r *ServeServerRepo) GetDefaultByServe(serveId uint) (ret model.ServeServer, err error) {
	servers := []model.ServeServer{}
	err = r.DB.Where("serve_id = ? AND NOT deleted", serveId).
		Find(&servers).Error

	minEnvironmentSort := -1

	for _, server := range servers {
		var environment model.Environment
		environment, err = r.EnvironmentRepo.Get(server.EnvironmentId)
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

func (r *ServeServerRepo) FindByServeAndExecEnv(serveId, environmentId uint) (ret model.ServeServer, err error) {
	err = r.DB.
		Where("serve_id = ? AND environment_id =? AND NOT deleted", serveId, environmentId).
		First(&ret).Error

	return
}

func (r *ServeServerRepo) SetUrl(serveId uint, url string) (err error) {
	err = r.DB.Model(model.ServeServer{}).Where("serve_id=?", serveId).Update("url", url).Error
	return
}

func (r *ServeServerRepo) BatchCreate(req []model.ServeServer) error {
	return r.DB.Create(req).Error
}

func (r *ServeServerRepo) UpdateUrlByServerAndServer(serveId, serverId uint, url string) error {
	err := r.DB.Model(&model.ServeServer{}).
		Where("serve_id = ?", serveId).
		Where("environment_id = ?", serverId).
		Update("url", url).Error

	return err
}
