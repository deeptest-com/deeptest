package repo

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
)

type EndpointFavoriteRepo struct {
	*BaseRepo `inject:""`
}

func (r *EndpointFavoriteRepo) Get(tenantId consts.TenantId, endpointId, userId uint) (ret model.EndpointFavorite) {
	r.GetDB(tenantId).Model(&model.EndpointFavorite{}).Where("endpoint_id  = ? and user_id = ?", endpointId, userId).First(&ret)
	return
}

func (r *EndpointFavoriteRepo) Delete(tenantId consts.TenantId, record model.EndpointFavorite) (err error) {
	err = r.GetDB(tenantId).Model(&model.EndpointFavorite{}).Delete(&record).Error
	return
}

func (r *EndpointFavoriteRepo) GetEndpointIds(tenantId consts.TenantId, userId uint) (ids []uint, err error) {
	var ret []model.EndpointFavorite
	err = r.GetDB(tenantId).Model(model.EndpointFavorite{}).Where(" user_id = ?", userId).Find(&ret).Error
	for _, item := range ret {
		ids = append(ids, item.EndpointId)
	}
	return
}

func (r *EndpointFavoriteRepo) Create(tenantId consts.TenantId, record model.EndpointFavorite) (err error) {
	err = r.GetDB(tenantId).Create(&record).Error
	return
}
