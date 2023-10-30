package repo

import (
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"gorm.io/gorm"
)

type ScenarioInterfaceRepo struct {
	*BaseRepo `inject:""`
	DB        *gorm.DB `inject:""`

	*DebugInterfaceRepo `inject:""`
}

func (r *ScenarioInterfaceRepo) Get(id uint) (po model.Processor, err error) {
	err = r.DB.
		Where("id=?", id).
		Where("NOT deleted").
		First(&po).Error
	return
}

func (r *ScenarioInterfaceRepo) GetDetail(processorId uint) (debugInterface model.DebugInterface, err error) {
	if processorId <= 0 {
		return
	}

	processor, err := r.Get(processorId)

	debugInterface, err = r.DebugInterfaceRepo.Get(processor.EntityId)

	return
}

func (r *ScenarioInterfaceRepo) SaveDebugData(interf *model.DebugInterface) (err error) {
	r.DB.Transaction(func(tx *gorm.DB) error {
		err = r.DB.Save(interf).Error
		if err != nil {
			return err
		}

		err = r.UpdateParams(interf.ID, interf.QueryParams, interf.PathParams)
		if err != nil {
			return err
		}

		err = r.UpdateBodyFormData(interf.ID, interf.BodyFormData)
		if err != nil {
			return err
		}

		err = r.UpdateBodyFormUrlencoded(interf.ID, interf.BodyFormUrlencoded)
		if err != nil {
			return err
		}

		err = r.UpdateHeaders(interf.ID, interf.Headers)
		if err != nil {
			return err
		}

		err = r.UpdateBasicAuth(interf.ID, interf.BasicAuth)
		if err != nil {
			return err
		}

		err = r.UpdateBearerToken(interf.ID, interf.BearerToken)
		if err != nil {
			return err
		}

		err = r.UpdateOAuth20(interf.ID, interf.OAuth20)
		if err != nil {
			return err
		}

		err = r.UpdateApiKey(interf.ID, interf.ApiKey)
		if err != nil {
			return err
		}

		err = r.UpdateGlobalParams(interf.ID, interf.GlobalParams)
		if err != nil {
			return err
		}
		return err
	})

	return
}
