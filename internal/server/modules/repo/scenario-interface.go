package repo

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"gorm.io/gorm"
)

type ScenarioInterfaceRepo struct {
	*BaseRepo `inject:""`
	DB        *gorm.DB `inject:""`
}

func (r *ScenarioInterfaceRepo) Get(id uint) (po model.ScenarioInterface, err error) {
	err = r.DB.
		Where("id=?", id).
		Where("NOT deleted").
		First(&po).Error
	return
}

func (r *ScenarioInterfaceRepo) ListParams(interfaceId uint) (
	queryParams []model.ScenarioInterfaceParam, pathParams []model.ScenarioInterfaceParam, err error) {

	pos := []model.ScenarioInterfaceParam{}

	err = r.DB.
		Where("interface_id=?", interfaceId).
		Where("NOT deleted").
		Order("id ASC").
		Find(&pos).Error

	for _, po := range pos {
		if po.ParamIn == consts.ParamInQuery {
			queryParams = append(queryParams, po)
		} else if po.ParamIn == consts.ParamInPath {
			pathParams = append(pathParams, po)
		}
	}

	return
}

func (r *ScenarioInterfaceRepo) ListHeaders(interfaceId uint) (pos []model.ScenarioInterfaceHeader, err error) {
	err = r.DB.
		Where("interface_id=?", interfaceId).
		Where("NOT deleted").
		Order("id ASC").
		Find(&pos).Error

	return
}

func (r *ScenarioInterfaceRepo) ListBodyFormData(interfaceId uint) (pos []model.ScenarioInterfaceBodyFormDataItem, err error) {
	err = r.DB.
		Where("interface_id=?", interfaceId).
		Where("NOT deleted").
		Order("id ASC").
		Find(&pos).Error

	return
}

func (r *ScenarioInterfaceRepo) ListBodyFormUrlencoded(interfaceId uint) (pos []model.ScenarioInterfaceBodyFormUrlEncodedItem, err error) {
	err = r.DB.
		Where("interface_id=?", interfaceId).
		Where("NOT deleted").
		Order("id ASC").
		Find(&pos).Error

	return
}

func (r *ScenarioInterfaceRepo) GetBasicAuth(id uint) (po model.ScenarioInterfaceBasicAuth, err error) {
	err = r.DB.
		Where("interface_id = ?", id).
		First(&po).Error

	return
}

func (r *ScenarioInterfaceRepo) GetBearerToken(id uint) (po model.ScenarioInterfaceBearerToken, err error) {
	err = r.DB.
		Where("interface_id = ?", id).
		First(&po).Error

	return
}

func (r *ScenarioInterfaceRepo) GetOAuth20(id uint) (po model.ScenarioInterfaceOAuth20, err error) {
	err = r.DB.
		Where("interface_id = ?", id).
		First(&po).Error

	return
}

func (r *ScenarioInterfaceRepo) GetApiKey(id uint) (po model.ScenarioInterfaceApiKey, err error) {
	err = r.DB.
		Where("interface_id = ?", id).
		First(&po).Error

	return
}

func (r *ScenarioInterfaceRepo) GetDetail(interfId uint) (interf model.ScenarioInterface, err error) {
	if interfId > 0 {
		interf, err = r.Get(interfId)

		interf.QueryParams, interf.PathParams, _ = r.ListParams(interfId)
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

func (r *ScenarioInterfaceRepo) HasScenarioInterfaceRecord(endpointInterfaceId uint) (id uint, err error) {
	var po model.ScenarioInterface

	err = r.DB.Model(&po).
		Where("endpoint_interface_id=?", endpointInterfaceId).
		First(&po).Error

	if err != nil {
		return
	}

	id = po.ID

	return
}

func (r *ScenarioInterfaceRepo) Save(interf *model.ScenarioInterface) (err error) {
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

		return err
	})

	return
}

func (r *ScenarioInterfaceRepo) UpdateParams(id uint, queryParams, pathParams []model.ScenarioInterfaceParam) (err error) {
	err = r.RemoveParams(id)

	var params []model.ScenarioInterfaceParam

	for _, p := range queryParams {

		if p.Name == "" {
			continue
		}

		p.ID = 0
		p.InterfaceId = id
		p.ParamIn = consts.ParamInQuery
		params = append(params, p)
	}

	for _, p := range pathParams {
		if p.Name == "" {
			continue
		}
		p.ID = 0
		p.InterfaceId = id
		p.ParamIn = consts.ParamInPath
		params = append(params, p)
	}

	if len(params) == 0 {
		return
	}

	err = r.DB.Create(&params).Error

	return
}

func (r *ScenarioInterfaceRepo) RemoveParams(id uint) (err error) {
	err = r.DB.
		Where("interface_id = ?", id).
		Delete(&model.ScenarioInterfaceParam{}, "").Error

	return
}

func (r *ScenarioInterfaceRepo) UpdateBodyFormData(id uint, items []model.ScenarioInterfaceBodyFormDataItem) (err error) {
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
		Delete(&model.ScenarioInterfaceBodyFormDataItem{}, "").Error

	return
}

func (r *ScenarioInterfaceRepo) UpdateBodyFormUrlencoded(id uint, items []model.ScenarioInterfaceBodyFormUrlEncodedItem) (err error) {
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
		Delete(&model.ScenarioInterfaceBodyFormUrlEncodedItem{}, "").Error

	return
}

func (r *ScenarioInterfaceRepo) UpdateHeaders(id uint, headers []model.ScenarioInterfaceHeader) (err error) {
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
		Delete(&model.ScenarioInterfaceHeader{}, "").Error

	return
}

func (r *ScenarioInterfaceRepo) UpdateBasicAuth(id uint, payload model.ScenarioInterfaceBasicAuth) (err error) {
	if err = r.RemoveBasicAuth(id); err != nil {
		return
	}

	payload.InterfaceId = id
	err = r.DB.Save(&payload).Error

	return
}

func (r *ScenarioInterfaceRepo) RemoveBasicAuth(id uint) (err error) {
	err = r.DB.
		Where("interface_id = ?", id).
		Delete(&model.ScenarioInterfaceBasicAuth{}, "").Error

	return
}

func (r *ScenarioInterfaceRepo) UpdateBearerToken(id uint, payload model.ScenarioInterfaceBearerToken) (err error) {
	if err = r.RemoveBearerToken(id); err != nil {
		return
	}

	payload.InterfaceId = id
	err = r.DB.Save(&payload).Error

	return
}

func (r *ScenarioInterfaceRepo) RemoveBearerToken(id uint) (err error) {
	err = r.DB.
		Where("interface_id = ?", id).
		Delete(&model.ScenarioInterfaceBearerToken{}, "").Error

	return
}

func (r *ScenarioInterfaceRepo) UpdateOAuth20(interfaceId uint, payload model.ScenarioInterfaceOAuth20) (err error) {
	if err = r.RemoveOAuth20(interfaceId); err != nil {
		return
	}

	payload.InterfaceId = interfaceId
	err = r.DB.Save(&payload).Error

	return
}

func (r *ScenarioInterfaceRepo) RemoveOAuth20(interfaceId uint) (err error) {
	err = r.DB.
		Where("interface_id = ?", interfaceId).
		Delete(&model.ScenarioInterfaceOAuth20{}).Error

	return
}

func (r *ScenarioInterfaceRepo) UpdateApiKey(id uint, payload model.ScenarioInterfaceApiKey) (err error) {
	if err = r.RemoveApiKey(id); err != nil {
		return
	}

	payload.InterfaceId = id
	err = r.DB.Save(&payload).Error

	return
}

func (r *ScenarioInterfaceRepo) RemoveApiKey(id uint) (err error) {
	err = r.DB.
		Where("interface_id = ?", id).
		Delete(&model.ScenarioInterfaceApiKey{}, "").Error

	return
}
