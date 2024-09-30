package repo

import (
	"github.com/deeptest-com/deeptest/internal/pkg/consts"
	"github.com/deeptest-com/deeptest/internal/server/modules/model"
	"gorm.io/gorm"
)

type ScenarioInterfaceRepo struct {
	*BaseRepo           `inject:""`
	DB                  *gorm.DB `inject:""`
	*DebugInterfaceRepo `inject:""`
}

func (r *ScenarioInterfaceRepo) Get(tenantId consts.TenantId, id uint) (po model.Processor, err error) {
	err = r.GetDB(tenantId).
		Where("id=?", id).
		Where("NOT deleted").
		First(&po).Error
	return
}

func (r *ScenarioInterfaceRepo) GetDetail(tenantId consts.TenantId, processorId uint) (debugInterface model.DebugInterface, err error) {
	if processorId <= 0 {
		return
	}

	processor, err := r.Get(tenantId, processorId)

	debugInterface, err = r.DebugInterfaceRepo.Get(tenantId, processor.EntityId)

	return
}

func (r *ScenarioInterfaceRepo) SaveDebugData(tenantId consts.TenantId, interf *model.DebugInterface) (err error) {
	r.GetDB(tenantId).Transaction(func(tx *gorm.DB) error {
		err = r.GetDB(tenantId).Save(interf).Error
		if err != nil {
			return err
		}

		err = r.UpdateParams(tenantId, interf.ID, interf.QueryParams, interf.PathParams)
		if err != nil {
			return err
		}

		err = r.UpdateBodyFormData(tenantId, interf.ID, interf.BodyFormData)
		if err != nil {
			return err
		}

		err = r.UpdateBodyFormUrlencoded(tenantId, interf.ID, interf.BodyFormUrlencoded)
		if err != nil {
			return err
		}

		err = r.UpdateHeaders(tenantId, interf.ID, interf.Headers)
		if err != nil {
			return err
		}

		err = r.UpdateBasicAuth(tenantId, interf.ID, interf.BasicAuth)
		if err != nil {
			return err
		}

		err = r.UpdateBearerToken(tenantId, interf.ID, interf.BearerToken)
		if err != nil {
			return err
		}

		err = r.UpdateOAuth20(tenantId, interf.ID, interf.OAuth20)
		if err != nil {
			return err
		}

		err = r.UpdateApiKey(tenantId, interf.ID, interf.ApiKey)
		if err != nil {
			return err
		}

		err = r.UpdateGlobalParams(tenantId, interf.ID, interf.GlobalParams)
		if err != nil {
			return err
		}
		return err
	})

	return
}
