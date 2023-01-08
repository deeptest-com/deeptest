package repo

import (
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"gorm.io/gorm"
)

type ScenarioInterfaceRepo struct {
	DB *gorm.DB `inject:""`

	ScenarioNodeRepo      *ScenarioNodeRepo      `inject:""`
	ExtractorRepo         *ExtractorRepo         `inject:""`
	CheckpointRepo        *CheckpointRepo        `inject:""`
	ScenarioInterfaceRepo *ScenarioInterfaceRepo `inject:""`
}

func (r *ScenarioInterfaceRepo) Get(interfaceId uint) (field model.ProcessorInterface, err error) {
	err = r.DB.
		Where("id=?", interfaceId).
		Where("NOT deleted").
		First(&field).Error
	return
}

func (r *ScenarioInterfaceRepo) GetDetail(interfId uint) (interf model.ProcessorInterface, err error) {
	if interfId > 0 {
		interf, err = r.Get(interfId)

		interf.Params, _ = r.ListParams(interfId)
		interf.Headers, _ = r.ListHeaders(interfId)
		interf.BodyFormData, _ = r.ListBodyFormData(interfId)
		interf.BodyFormUrlencoded, _ = r.ListBodyFormUrlencoded(interfId)

		interf.BasicAuth, _ = r.GetBasicAuth(interfId)
		interf.BearerToken, _ = r.GetBearerToken(interfId)
		interf.OAuth20, _ = r.GetOAuth20(interfId)
		interf.ApiKey, _ = r.GetApiKey(interfId)
	}

	return
}

func (r *ScenarioInterfaceRepo) SaveInterface(interfaceProcessor *model.ProcessorInterface) (err error) {
	err = r.DB.Save(interfaceProcessor).Error

	err = r.UpdateParams(interfaceProcessor.ID, interfaceProcessor.Params)
	if err != nil {
		return err
	}

	err = r.UpdateBodyFormData(interfaceProcessor.ID, interfaceProcessor.BodyFormData)
	if err != nil {
		return err
	}

	err = r.UpdateBodyFormUrlencoded(interfaceProcessor.ID, interfaceProcessor.BodyFormUrlencoded)
	if err != nil {
		return err
	}

	err = r.UpdateHeaders(interfaceProcessor.ID, interfaceProcessor.Headers)
	if err != nil {
		return err
	}

	err = r.UpdateBasicAuth(interfaceProcessor.ID, interfaceProcessor.BasicAuth)
	if err != nil {
		return err
	}

	err = r.UpdateBearerToken(interfaceProcessor.ID, interfaceProcessor.BearerToken)
	if err != nil {
		return err
	}

	err = r.UpdateOAuth20(interfaceProcessor.ID, interfaceProcessor.OAuth20)
	if err != nil {
		return err
	}

	err = r.UpdateApiKey(interfaceProcessor.ID, interfaceProcessor.ApiKey)
	if err != nil {
		return err
	}

	return
}

func (r *ScenarioInterfaceRepo) UpdateParams(id uint, params []model.ProcessorInterfaceParam) (err error) {
	err = r.RemoveParams(id)

	if len(params) == 0 {
		return
	}

	for idx, _ := range params {
		params[idx].ID = 0
		params[idx].InterfaceId = id
	}

	err = r.DB.Create(&params).Error

	return
}
func (r *ScenarioInterfaceRepo) RemoveParams(id uint) (err error) {
	err = r.DB.
		Where("interface_id = ?", id).
		Delete(&model.ProcessorInterfaceParam{}, "").Error

	return
}

func (r *ScenarioInterfaceRepo) UpdateBodyFormData(id uint, items []model.ProcessorInterfaceBodyFormDataItem) (err error) {
	err = r.RemoveBodyFormData(id)

	if len(items) == 0 {
		return
	}

	for idx, _ := range items {
		items[idx].ID = 0
		items[idx].InterfaceId = id
	}

	err = r.DB.Create(&items).Error

	return
}
func (r *ScenarioInterfaceRepo) RemoveBodyFormData(id uint) (err error) {
	err = r.DB.
		Where("interface_id = ?", id).
		Delete(&model.ProcessorInterfaceBodyFormDataItem{}, "").Error

	return
}

func (r *ScenarioInterfaceRepo) UpdateBodyFormUrlencoded(id uint, items []model.ProcessorInterfaceBodyFormUrlEncodedItem) (err error) {
	err = r.RemoveBodyFormUrlencoded(id)

	if len(items) == 0 {
		return
	}

	for idx, _ := range items {
		items[idx].ID = 0
		items[idx].InterfaceId = id
	}

	err = r.DB.Create(&items).Error

	return
}
func (r *ScenarioInterfaceRepo) RemoveBodyFormUrlencoded(id uint) (err error) {
	err = r.DB.
		Where("interface_id = ?", id).
		Delete(&model.ProcessorInterfaceBodyFormUrlEncodedItem{}, "").Error

	return
}

func (r *ScenarioInterfaceRepo) UpdateHeaders(id uint, headers []model.ProcessorInterfaceHeader) (err error) {
	err = r.RemoveHeaders(id)

	if len(headers) == 0 {
		return
	}

	for idx, _ := range headers {
		headers[idx].ID = 0
		headers[idx].InterfaceId = id
	}

	err = r.DB.Create(&headers).Error

	return
}
func (r *ScenarioInterfaceRepo) RemoveHeaders(id uint) (err error) {
	err = r.DB.
		Where("interface_id = ?", id).
		Delete(&model.ProcessorInterfaceHeader{}, "").Error

	return
}

func (r *ScenarioInterfaceRepo) UpdateBasicAuth(id uint, payload model.ProcessorInterfaceBasicAuth) (err error) {
	payload.InterfaceId = id
	err = r.DB.Save(&payload).Error

	return
}
func (r *ScenarioInterfaceRepo) RemoveBasicAuth(id uint) (err error) {
	err = r.DB.
		Where("interface_id = ?", id).
		Delete(&model.ProcessorInterfaceBasicAuth{}, "").Error

	return
}

func (r *ScenarioInterfaceRepo) UpdateBearerToken(id uint, payload model.ProcessorInterfaceBearerToken) (err error) {
	payload.InterfaceId = id
	err = r.DB.Save(&payload).Error

	return
}
func (r *ScenarioInterfaceRepo) RemoveBearerToken(id uint) (err error) {
	err = r.DB.
		Where("interface_id = ?", id).
		Delete(&model.ProcessorInterfaceBearerToken{}, "").Error

	return
}

func (r *ScenarioInterfaceRepo) UpdateOAuth20(interfaceId uint, payload model.ProcessorInterfaceOAuth20) (err error) {
	r.RemoveOAuth20(interfaceId)

	payload.InterfaceId = interfaceId
	err = r.DB.Save(&payload).Error

	return
}
func (r *ScenarioInterfaceRepo) RemoveOAuth20(interfaceId uint) (err error) {
	err = r.DB.
		Where("interface_id = ?", interfaceId).
		Delete(&model.ProcessorInterfaceOAuth20{}).Error

	return
}

func (r *ScenarioInterfaceRepo) UpdateApiKey(id uint, payload model.ProcessorInterfaceApiKey) (err error) {
	payload.InterfaceId = id
	err = r.DB.Save(&payload).Error

	return
}
func (r *ScenarioInterfaceRepo) RemoveApiKey(id uint) (err error) {
	err = r.DB.
		Where("interface_id = ?", id).
		Delete(&model.ProcessorInterfaceApiKey{}, "").Error

	return
}

func (r *ScenarioInterfaceRepo) ListParams(interfaceId uint) (pos []model.ProcessorInterfaceParam, err error) {
	err = r.DB.
		Where("interface_id=?", interfaceId).
		Where("NOT deleted").
		Order("id ASC").
		Find(&pos).Error
	return
}
func (r *ScenarioInterfaceRepo) ListHeaders(interfaceId uint) (pos []model.ProcessorInterfaceHeader, err error) {
	err = r.DB.
		Where("interface_id=?", interfaceId).
		Where("NOT deleted").
		Order("id ASC").
		Find(&pos).Error

	return
}
func (r *ScenarioInterfaceRepo) ListBodyFormData(interfaceId uint) (pos []model.ProcessorInterfaceBodyFormDataItem, err error) {
	err = r.DB.
		Where("interface_id=?", interfaceId).
		Where("NOT deleted").
		Order("id ASC").
		Find(&pos).Error

	return
}
func (r *ScenarioInterfaceRepo) ListBodyFormUrlencoded(interfaceId uint) (pos []model.ProcessorInterfaceBodyFormUrlEncodedItem, err error) {
	err = r.DB.
		Where("interface_id=?", interfaceId).
		Where("NOT deleted").
		Order("id ASC").
		Find(&pos).Error

	return
}

func (r *ScenarioInterfaceRepo) GetBasicAuth(id uint) (po model.ProcessorInterfaceBasicAuth, err error) {
	err = r.DB.
		Where("interface_id = ?", id).
		First(&po).Error

	return
}
func (r *ScenarioInterfaceRepo) GetBearerToken(id uint) (po model.ProcessorInterfaceBearerToken, err error) {
	err = r.DB.
		Where("interface_id = ?", id).
		First(&po).Error

	return
}
func (r *ScenarioInterfaceRepo) GetOAuth20(id uint) (po model.ProcessorInterfaceOAuth20, err error) {
	err = r.DB.
		Where("interface_id = ?", id).
		First(&po).Error

	return
}
func (r *ScenarioInterfaceRepo) GetApiKey(id uint) (po model.ProcessorInterfaceApiKey, err error) {
	err = r.DB.
		Where("interface_id = ?", id).
		First(&po).Error

	return
}
