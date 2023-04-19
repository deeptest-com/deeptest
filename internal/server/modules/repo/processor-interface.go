package repo

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"gorm.io/gorm"
)

type ProcessorInterfaceRepo struct {
	DB *gorm.DB `inject:""`

	ScenarioNodeRepo       *ScenarioNodeRepo       `inject:""`
	ExtractorRepo          *ExtractorRepo          `inject:""`
	CheckpointRepo         *CheckpointRepo         `inject:""`
	ProcessorInterfaceRepo *ProcessorInterfaceRepo `inject:""`
}

func (r *ProcessorInterfaceRepo) Get(interfaceId uint) (po model.ProcessorInterface, err error) {
	err = r.DB.
		Where("id=?", interfaceId).
		Where("NOT deleted").
		First(&po).Error

	return
}

func (r *ProcessorInterfaceRepo) GetDetail(interfId uint) (interf model.ProcessorInterface, err error) {
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

func (r *ProcessorInterfaceRepo) SaveInterface(interfaceProcessor *model.ProcessorInterface) (err error) {
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

func (r *ProcessorInterfaceRepo) UpdateParams(id uint, params []model.ProcessorInterfaceParam) (err error) {
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
func (r *ProcessorInterfaceRepo) RemoveParams(id uint) (err error) {
	err = r.DB.
		Where("interface_id = ?", id).
		Delete(&model.ProcessorInterfaceParam{}, "").Error

	return
}

func (r *ProcessorInterfaceRepo) UpdateBodyFormData(id uint, items []model.ProcessorInterfaceBodyFormDataItem) (err error) {
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
func (r *ProcessorInterfaceRepo) RemoveBodyFormData(id uint) (err error) {
	err = r.DB.
		Where("interface_id = ?", id).
		Delete(&model.ProcessorInterfaceBodyFormDataItem{}, "").Error

	return
}

func (r *ProcessorInterfaceRepo) UpdateBodyFormUrlencoded(id uint, items []model.ProcessorInterfaceBodyFormUrlEncodedItem) (err error) {
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
func (r *ProcessorInterfaceRepo) RemoveBodyFormUrlencoded(id uint) (err error) {
	err = r.DB.
		Where("interface_id = ?", id).
		Delete(&model.ProcessorInterfaceBodyFormUrlEncodedItem{}, "").Error

	return
}

func (r *ProcessorInterfaceRepo) UpdateHeaders(id uint, headers []model.ProcessorInterfaceHeader) (err error) {
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
func (r *ProcessorInterfaceRepo) RemoveHeaders(id uint) (err error) {
	err = r.DB.
		Where("interface_id = ?", id).
		Delete(&model.ProcessorInterfaceHeader{}, "").Error

	return
}

func (r *ProcessorInterfaceRepo) UpdateBasicAuth(id uint, payload model.ProcessorInterfaceBasicAuth) (err error) {
	payload.InterfaceId = id
	err = r.DB.Save(&payload).Error

	return
}
func (r *ProcessorInterfaceRepo) RemoveBasicAuth(id uint) (err error) {
	err = r.DB.
		Where("interface_id = ?", id).
		Delete(&model.ProcessorInterfaceBasicAuth{}, "").Error

	return
}

func (r *ProcessorInterfaceRepo) UpdateBearerToken(id uint, payload model.ProcessorInterfaceBearerToken) (err error) {
	payload.InterfaceId = id
	err = r.DB.Save(&payload).Error

	return
}
func (r *ProcessorInterfaceRepo) RemoveBearerToken(id uint) (err error) {
	err = r.DB.
		Where("interface_id = ?", id).
		Delete(&model.ProcessorInterfaceBearerToken{}, "").Error

	return
}

func (r *ProcessorInterfaceRepo) UpdateOAuth20(interfaceId uint, payload model.ProcessorInterfaceOAuth20) (err error) {
	r.RemoveOAuth20(interfaceId)

	payload.InterfaceId = interfaceId
	err = r.DB.Save(&payload).Error

	return
}
func (r *ProcessorInterfaceRepo) RemoveOAuth20(interfaceId uint) (err error) {
	err = r.DB.
		Where("interface_id = ?", interfaceId).
		Delete(&model.ProcessorInterfaceOAuth20{}).Error

	return
}

func (r *ProcessorInterfaceRepo) UpdateApiKey(id uint, payload model.ProcessorInterfaceApiKey) (err error) {
	payload.InterfaceId = id
	err = r.DB.Save(&payload).Error

	return
}
func (r *ProcessorInterfaceRepo) RemoveApiKey(id uint) (err error) {
	err = r.DB.
		Where("interface_id = ?", id).
		Delete(&model.ProcessorInterfaceApiKey{}, "").Error

	return
}

func (r *ProcessorInterfaceRepo) ListParams(interfaceId uint) (pos []model.ProcessorInterfaceParam, err error) {
	err = r.DB.
		Where("interface_id=?", interfaceId).
		Where("NOT deleted").
		Order("id ASC").
		Find(&pos).Error
	return
}
func (r *ProcessorInterfaceRepo) ListHeaders(interfaceId uint) (pos []model.ProcessorInterfaceHeader, err error) {
	err = r.DB.
		Where("interface_id=?", interfaceId).
		Where("NOT deleted").
		Order("id ASC").
		Find(&pos).Error

	return
}
func (r *ProcessorInterfaceRepo) ListBodyFormData(interfaceId uint) (pos []model.ProcessorInterfaceBodyFormDataItem, err error) {
	err = r.DB.
		Where("interface_id=?", interfaceId).
		Where("NOT deleted").
		Order("id ASC").
		Find(&pos).Error

	return
}
func (r *ProcessorInterfaceRepo) ListBodyFormUrlencoded(interfaceId uint) (pos []model.ProcessorInterfaceBodyFormUrlEncodedItem, err error) {
	err = r.DB.
		Where("interface_id=?", interfaceId).
		Where("NOT deleted").
		Order("id ASC").
		Find(&pos).Error

	return
}

func (r *ProcessorInterfaceRepo) GetBasicAuth(id uint) (po model.ProcessorInterfaceBasicAuth, err error) {
	err = r.DB.
		Where("interface_id = ?", id).
		First(&po).Error

	return
}
func (r *ProcessorInterfaceRepo) GetBearerToken(id uint) (po model.ProcessorInterfaceBearerToken, err error) {
	err = r.DB.
		Where("interface_id = ?", id).
		First(&po).Error

	return
}
func (r *ProcessorInterfaceRepo) GetOAuth20(id uint) (po model.ProcessorInterfaceOAuth20, err error) {
	err = r.DB.
		Where("interface_id = ?", id).
		First(&po).Error

	return
}
func (r *ProcessorInterfaceRepo) GetApiKey(id uint) (po model.ProcessorInterfaceApiKey, err error) {
	err = r.DB.
		Where("interface_id = ?", id).
		First(&po).Error

	return
}

func (r *ProcessorInterfaceRepo) ListInvocation(processorInterfaceId uint) (pos []model.ProcessorInvocation, err error) {
	err = r.DB.
		Select("id", "name").
		Where("processor_interface_id=?", processorInterfaceId).
		Where("NOT deleted").
		Order("created_at DESC").
		Find(&pos).Error
	return
}

func (r *ProcessorInterfaceRepo) Update(interf model.ProcessorInterface) (err error) {
	r.DB.Transaction(func(tx *gorm.DB) error {
		err = r.DB.Updates(interf).Error
		if err != nil {
			return err
		}

		err = r.UpdateParams(interf.ID, interf.Params)
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

		return err
	})

	return
}

func (r *ProcessorInterfaceRepo) GetProcessor(scenarioId, endpointId uint, method consts.HttpMethod) (processor model.ProcessorInterface, err error) {
	err = r.DB.Where("scenario_id = ? AND endpoint_id = ? AND method = ?", scenarioId, endpointId, method).First(&processor).Error
	return
}

/*
func (r *ProcessorInterfaceRepo) SaveProcessor(interf model.Interface) (err error) {
	var processor model.interfaces2debug
	processor, err = r.GetProcessor(0, interf.EndpointId, interf.Method)
	if err != nil {
		return err
	}
	processor := openapi.interfaces2debug(interf)
	processorInterface := interfaces2debug.Convert()
	if processor.ID == 0 {
		err = r.SaveInterface(processorInterface)
	} else {
		processorInterface.ID = processor.ID
		err = r.Update(*processorInterface)
	}
	if err != nil {
		return err
	}
	return
}

*/

func (r *ProcessorInterfaceRepo) GetList(projectId, scenarioId uint) (processors []model.ProcessorInterface, err error) {
	err = r.DB.Where("project_id=? and scenario_id=?", projectId, scenarioId).Find(&processors).Error
	if err != nil {
		return
	}
	for key, processor := range processors {
		processors[key], err = r.GetDetail(processor.ID)
		if err != nil {
			return
		}
	}

	return
}

func (r *ProcessorInterfaceRepo) GetCountByEndpointId(endpointId uint) (count int64, err error) {
	err = r.DB.Model(&model.ProcessorInterface{}).Where("NOT deleted and endpoint_id=?", endpointId).Count(&count).Error
	return
}
