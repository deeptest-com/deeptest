package repo

import (
	serverDomain "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/server/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"gorm.io/gorm"
)

type DebugInterfaceRepo struct {
	*BaseRepo             `inject:""`
	DB                    *gorm.DB `inject:""`
	*ServeRepo            `inject:""`
	*ServeServerRepo      `inject:""`
	EndpointCaseRepo      *EndpointCaseRepo      `inject:""`
	DiagnoseInterfaceRepo *DiagnoseInterfaceRepo `inject:""`
}

func (r *DebugInterfaceRepo) Tested(id uint) (res bool, err error) {
	var count int64
	err = r.DB.Model(&model.DebugInterface{}).Where("id=?", id).Count(&count).Error
	if err != nil {
		return
	}
	res = count > 0
	return
}

func (r *DebugInterfaceRepo) UpdateOrder(pos serverConsts.DropPos, targetId uint) (parentId uint, ordr int) {
	if pos == serverConsts.Inner {
		parentId = targetId

		var preChild model.DebugInterface
		r.DB.Where("parent_id=?", parentId).
			Order("ordr DESC").Limit(1).
			First(&preChild)

		ordr = preChild.Ordr + 1

	} else if pos == serverConsts.Before {
		brother, _ := r.Get(targetId)
		parentId = brother.ParentId

		r.DB.Model(&model.DebugInterface{}).
			Where("NOT deleted AND parent_id=? AND ordr >= ?", parentId, brother.Ordr).
			Update("ordr", gorm.Expr("ordr + 1"))

		ordr = brother.Ordr

	} else if pos == serverConsts.After {
		brother, _ := r.Get(targetId)
		parentId = brother.ParentId

		r.DB.Model(&model.DebugInterface{}).
			Where("NOT deleted AND parent_id=? AND ordr > ?", parentId, brother.Ordr).
			Update("ordr", gorm.Expr("ordr + 1"))

		ordr = brother.Ordr + 1

	}

	return
}

func (r *DebugInterfaceRepo) ListByProject(projectId int) (pos []*model.DebugInterface, err error) {
	err = r.DB.
		Where("project_id=?", projectId).
		Where("NOT deleted").
		Order("parent_id ASC, ordr ASC").
		Find(&pos).Error
	return
}

func (r *DebugInterfaceRepo) Get(id uint) (po model.DebugInterface, err error) {
	err = r.DB.
		Where("id=?", id).
		Where("NOT deleted").
		First(&po).Error
	return
}

func (r *DebugInterfaceRepo) GetByCaseInterfaceId(caseInterfaceId uint) (po model.DebugInterface, err error) {
	err = r.DB.
		Where("case_interface_id=?", caseInterfaceId).
		Where("NOT deleted").
		First(&po).Error
	return
}

func (r *DebugInterfaceRepo) GetDetail(interfId uint) (interf model.DebugInterface, err error) {
	if interfId <= 0 {
		return
	}

	interf, err = r.Get(interfId)

	interf.QueryParams, interf.PathParams, _ = r.ListParams(interfId)
	interf.Headers, _ = r.ListHeaders(interfId)

	interf.BodyFormData, _ = r.ListBodyFormData(interfId)
	interf.BodyFormUrlencoded, _ = r.ListBodyFormUrlencoded(interfId)

	interf.BasicAuth, _ = r.GetBasicAuth(interfId)
	interf.BearerToken, _ = r.GetBearerToken(interfId)
	interf.OAuth20, _ = r.GetOAuth20(interfId)
	interf.ApiKey, _ = r.GetApiKey(interfId)

	return
}

func (r *DebugInterfaceRepo) Save(interf *model.DebugInterface) (err error) {
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
func (r *DebugInterfaceRepo) UpdateName(req serverDomain.EndpointCase) (err error) {
	err = r.DB.Model(&model.DebugInterface{}).
		Where("id = ?", req.Id).
		Update("name", req.Name).Error

	return
}

func (r *DebugInterfaceRepo) UpdateHeaders(id uint, headers []model.DebugInterfaceHeader) (err error) {
	err = r.RemoveHeaders(id)

	if len(headers) == 0 {
		return
	}

	var newHeaders []model.DebugInterfaceHeader
	for _, h := range headers {
		if h.Name == "" {
			continue
		}

		h.ID = 0
		h.InterfaceId = id

		newHeaders = append(newHeaders, h)
	}

	if len(newHeaders) == 0 {
		return
	}
	err = r.DB.Create(&newHeaders).Error

	return
}

func (r *DebugInterfaceRepo) RemoveHeaders(id uint) (err error) {
	err = r.DB.
		Where("interface_id = ?", id).
		Delete(&model.DebugInterfaceHeader{}, "").Error

	return
}

func (r *DebugInterfaceRepo) UpdateParams(id uint, queryParams, pathParams []model.DebugInterfaceParam) (err error) {
	err = r.RemoveParams(id)

	var params []model.DebugInterfaceParam

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
func (r *DebugInterfaceRepo) RemoveParams(id uint) (err error) {
	err = r.DB.
		Where("interface_id = ?", id).
		Delete(&model.DebugInterfaceParam{}, "").Error

	return
}

func (r *DebugInterfaceRepo) UpdateCookies(id uint, cookies []model.DebugInterfaceCookie) (err error) {
	err = r.RemoveCookie(id)

	if len(cookies) == 0 {
		return
	}

	var newCookies []model.DebugInterfaceCookie
	for _, c := range cookies {
		if c.Name == "" {
			continue
		}

		c.ID = 0
		c.InterfaceId = id

		newCookies = append(newCookies, c)
	}

	if len(newCookies) == 0 {
		return
	}
	err = r.DB.Create(&newCookies).Error

	return
}

func (r *DebugInterfaceRepo) RemoveCookie(id uint) (err error) {
	err = r.DB.
		Where("interface_id = ?", id).
		Delete(&model.DebugInterfaceCookie{}, "").Error

	return
}

func (r *DebugInterfaceRepo) UpdateProcessorId(id, processorId uint) (err error) {
	values := map[string]interface{}{
		"scenario_processor_id": processorId,
	}
	err = r.UpdateDebugInfo(id, values)

	return
}

func (r *DebugInterfaceRepo) UpdateBodyFormData(id uint, items []model.DebugInterfaceBodyFormDataItem) (err error) {
	err = r.RemoveBodyFormData(id)

	if len(items) == 0 {
		return
	}

	var list []model.DebugInterfaceBodyFormDataItem
	for _, item := range items {
		if item.Name == "" {
			continue
		}

		item.ID = 0
		item.InterfaceId = id

		list = append(list, item)
	}

	if len(list) == 0 {
		return
	}
	err = r.DB.Create(&list).Error

	return
}
func (r *DebugInterfaceRepo) RemoveBodyFormData(id uint) (err error) {
	err = r.DB.
		Where("interface_id = ?", id).
		Delete(&model.DebugInterfaceBodyFormDataItem{}, "").Error

	return
}

func (r *DebugInterfaceRepo) UpdateBodyFormUrlencoded(id uint, items []model.DebugInterfaceBodyFormUrlEncodedItem) (err error) {
	err = r.RemoveBodyFormUrlencoded(id)

	if len(items) == 0 {
		return
	}

	var list []model.DebugInterfaceBodyFormUrlEncodedItem
	for _, item := range items {
		if item.Name == "" {
			continue
		}

		item.ID = 0
		item.InterfaceId = id

		list = append(list, item)
	}

	if len(list) == 0 {
		return
	}
	err = r.DB.Create(&list).Error

	return
}
func (r *DebugInterfaceRepo) RemoveBodyFormUrlencoded(id uint) (err error) {
	err = r.DB.
		Where("interface_id = ?", id).
		Delete(&model.DebugInterfaceBodyFormUrlEncodedItem{}, "").Error

	return
}

func (r *DebugInterfaceRepo) UpdateBasicAuth(id uint, payload model.DebugInterfaceBasicAuth) (err error) {
	r.RemoveBasicAuth(id)

	payload.InterfaceId = id
	err = r.DB.Save(&payload).Error

	return
}
func (r *DebugInterfaceRepo) RemoveBasicAuth(id uint) (err error) {
	err = r.DB.
		Where("interface_id = ?", id).
		Delete(&model.DebugInterfaceBasicAuth{}, "").Error

	return
}

func (r *DebugInterfaceRepo) UpdateBearerToken(id uint, payload model.DebugInterfaceBearerToken) (err error) {
	r.RemoveBearerToken(id)

	payload.InterfaceId = id
	err = r.DB.Save(&payload).Error

	return
}
func (r *DebugInterfaceRepo) RemoveBearerToken(id uint) (err error) {
	err = r.DB.
		Where("interface_id = ?", id).
		Delete(&model.DebugInterfaceBearerToken{}, "").Error

	return
}

func (r *DebugInterfaceRepo) UpdateOAuth20(interfaceId uint, payload model.DebugInterfaceOAuth20) (err error) {
	r.RemoveOAuth20(interfaceId)

	payload.InterfaceId = interfaceId
	err = r.DB.Save(&payload).Error

	return
}
func (r *DebugInterfaceRepo) RemoveOAuth20(interfaceId uint) (err error) {
	err = r.DB.
		Where("interface_id = ?", interfaceId).
		Delete(&model.DebugInterfaceOAuth20{}).Error

	return
}

func (r *DebugInterfaceRepo) UpdateApiKey(id uint, payload model.DebugInterfaceApiKey) (err error) {
	r.RemoveApiKey(id)

	payload.InterfaceId = id
	err = r.DB.Save(&payload).Error

	return
}
func (r *DebugInterfaceRepo) RemoveApiKey(id uint) (err error) {
	err = r.DB.
		Where("interface_id = ?", id).
		Delete(&model.DebugInterfaceApiKey{}, "").Error

	return
}

func (r *DebugInterfaceRepo) Delete(id uint) (err error) {
	err = r.DB.Model(&model.DebugInterface{}).
		Where("id=?", id).
		Update("deleted", true).
		Error

	return
}

func (r *DebugInterfaceRepo) GetChildren(defId, fieldId uint) (children []*model.DebugInterface, err error) {
	err = r.DB.Where("defID=? AND parentID=?", defId, fieldId).Find(&children).Error
	return
}

func (r *DebugInterfaceRepo) SetIsRange(fieldId uint, b bool) (err error) {
	err = r.DB.Model(&model.DebugInterface{}).
		Where("id = ?", fieldId).Update("isRange", b).Error

	return
}

func (r *DebugInterfaceRepo) UpdateOrdAndParent(interf model.DebugInterface) (err error) {
	err = r.DB.Model(&interf).
		Updates(model.DebugInterface{InterfaceBase: model.InterfaceBase{Ordr: interf.Ordr, ParentId: interf.ParentId}}).
		Error

	return
}

func (r *DebugInterfaceRepo) SetOAuth2AccessToken(token string, interfaceId int) (err error) {
	err = r.DB.Model(&model.DebugInterfaceOAuth20{}).
		Where("interface_id = ?", interfaceId).
		Update("access_token", token).Error

	return
}

func (r *DebugInterfaceRepo) ListParams(interfaceId uint) (
	queryParams []model.DebugInterfaceParam, pathParams []model.DebugInterfaceParam, err error) {

	pos := []model.DebugInterfaceParam{}

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
func (r *DebugInterfaceRepo) ListHeaders(interfaceId uint) (pos []model.DebugInterfaceHeader, err error) {
	err = r.DB.
		Where("interface_id=?", interfaceId).
		Where("NOT deleted").
		Order("id ASC").
		Find(&pos).Error

	return
}
func (r *DebugInterfaceRepo) ListCookies(interfaceId uint) (pos []model.DebugInterfaceCookie, err error) {
	err = r.DB.
		Where("interface_id=?", interfaceId).
		Where("NOT deleted").
		Order("id ASC").
		Find(&pos).Error

	return
}
func (r *DebugInterfaceRepo) ListBodyFormData(interfaceId uint) (pos []model.DebugInterfaceBodyFormDataItem, err error) {
	err = r.DB.
		Where("interface_id=?", interfaceId).
		Where("NOT deleted").
		Order("id ASC").
		Find(&pos).Error

	return
}
func (r *DebugInterfaceRepo) ListBodyFormUrlencoded(interfaceId uint) (pos []model.DebugInterfaceBodyFormUrlEncodedItem, err error) {
	err = r.DB.
		Where("interface_id=?", interfaceId).
		Where("NOT deleted").
		Order("id ASC").
		Find(&pos).Error

	return
}

func (r *DebugInterfaceRepo) GetBasicAuth(id uint) (po model.DebugInterfaceBasicAuth, err error) {
	err = r.DB.
		Where("interface_id = ?", id).
		First(&po).Error

	return
}
func (r *DebugInterfaceRepo) GetBearerToken(id uint) (po model.DebugInterfaceBearerToken, err error) {
	err = r.DB.
		Where("interface_id = ?", id).
		First(&po).Error

	return
}
func (r *DebugInterfaceRepo) GetOAuth20(id uint) (po model.DebugInterfaceOAuth20, err error) {
	err = r.DB.
		Where("interface_id = ?", id).
		First(&po).Error

	return
}
func (r *DebugInterfaceRepo) GetApiKey(id uint) (po model.DebugInterfaceApiKey, err error) {
	err = r.DB.
		Where("interface_id = ?", id).
		First(&po).Error

	return
}

func (r *DebugInterfaceRepo) SaveInterfaces(interf *model.DebugInterface) (err error) {
	r.DB.Transaction(func(tx *gorm.DB) error {
		err = r.UpdateInterface(interf)
		if err != nil {
			return err
		}

		err = r.UpdateParams(interf.ID, interf.QueryParams, interf.PathParams)
		if err != nil {
			return err
		}

		err = r.UpdateHeaders(interf.ID, interf.Headers)
		if err != nil {
			return err
		}

		err = r.UpdateCookies(interf.ID, interf.Cookies)
		if err != nil {
			return err
		}

		return err
	})

	return
}

func (r *DebugInterfaceRepo) UpdateInterface(interf *model.DebugInterface) (err error) {
	err = r.BaseRepo.Save(interf.ID, interf)
	return
}

func (r *DebugInterfaceRepo) PopulateProps(po *model.DebugInterface) (err error) {
	po.QueryParams, po.PathParams, _ = r.ListParams(po.ID)
	po.Headers, _ = r.ListHeaders(po.ID)
	po.BodyFormData, _ = r.ListBodyFormData(po.ID)
	po.BodyFormUrlencoded, _ = r.ListBodyFormUrlencoded(po.ID)
	po.BasicAuth, _ = r.GetBasicAuth(po.ID)
	po.BearerToken, _ = r.GetBearerToken(po.ID)
	po.OAuth20, _ = r.GetOAuth20(po.ID)
	po.ApiKey, _ = r.GetApiKey(po.ID)

	return
}

func (r *DebugInterfaceRepo) UpdateDebugInfo(id uint, values map[string]interface{}) (err error) {
	err = r.DB.Model(&model.DebugInterface{}).
		Where("id=?", id).
		Updates(values).
		Error

	return
}

func (r *DebugInterfaceRepo) DeleteByProcessorIds(ids []uint) (err error) {
	if len(ids) > 0 {
		err = r.DB.Model(&model.DebugInterface{}).
			Where("scenario_processor_id IN (?)", ids).
			Update("deleted", true).
			Error
	}

	return
}

func (r *DebugInterfaceRepo) CreateDefault(src consts.ProcessorInterfaceSrc, projectId uint) (id uint, err error) {
	po := model.DebugInterface{
		ProcessorInterfaceSrc: src,

		InterfaceBase: model.InterfaceBase{
			InterfaceConfigBase: model.InterfaceConfigBase{
				Method: consts.GET,
			},
			ProjectId: projectId,
		},
	}

	serves, _ := r.ServeRepo.ListByProject(projectId)
	if len(serves) > 0 {
		po.ServeId = serves[0].ID

		server, _ := r.ServeServerRepo.GetDefaultByServe(po.ServeId)
		po.ServerId = server.ID
	}

	err = r.Save(&po)

	id = po.ID

	return
}

func (r *DebugInterfaceRepo) GetSourceNameById(id uint) (name string, err error) {
	debugInterface, err := r.Get(id)
	if err != nil {
		return
	}

	switch debugInterface.ProcessorInterfaceSrc {
	case consts.InterfaceSrcDefine:
		endpointInterface, err := r.EndpointInterfaceRepo.Get(debugInterface.EndpointInterfaceId)
		if err != nil {
			return "", err
		}
		name = endpointInterface.Name
	case consts.InterfaceSrcCase:
		endpointCase, err := r.EndpointCaseRepo.Get(debugInterface.CaseInterfaceId)
		if err != nil {
			return "", err
		}
		name = endpointCase.Name
	case consts.InterfaceSrcDiagnose:
		diagnoseInterface, err := r.DiagnoseInterfaceRepo.Get(debugInterface.DiagnoseInterfaceId)
		if err != nil {
			return "", err
		}
		name = diagnoseInterface.Title
	}

	return
}
