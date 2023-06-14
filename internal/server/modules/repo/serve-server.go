package repo

import (
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
)

type ServeServerRepo struct {
	*BaseRepo       `inject:""`
	CategoryRepo    *CategoryRepo    `inject:""`
	EnvironmentRepo *EnvironmentRepo `inject:""`
	EndpointRepo    *EndpointRepo    `inject:""`
}

func NewServerRepo() *ServeServerRepo {
	return &ServeServerRepo{}
}

func (r *ServeServerRepo) Get(id uint) (res model.ServeServer, err error) {
	err = r.DB.Where("NOT deleted").First(&res, id).Error
	return
}

func (r *ServeServerRepo) GetByEndpoint(endpointId uint) (res model.ServeServer, err error) {
	endpoint, _ := r.EndpointRepo.Get(endpointId)

	err = r.DB.Where("NOT deleted").First(&res, endpoint.ServerId).Error
	return
}

func (r *ServeServerRepo) GetDefaultByServe(serveId uint) (ret model.ServeServer, err error) {
	err = r.DB.Where("serve_id = ? AND NOT deleted", serveId).First(&ret).Error

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
