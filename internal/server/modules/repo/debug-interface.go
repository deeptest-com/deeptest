package repo

import (
	serverDomain "github.com/deeptest-com/deeptest/cmd/server/v1/domain"
	"github.com/deeptest-com/deeptest/internal/pkg/consts"
	"github.com/deeptest-com/deeptest/internal/server/consts"
	"github.com/deeptest-com/deeptest/internal/server/modules/model"
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

func (r *DebugInterfaceRepo) Tested(tenantId consts.TenantId, id uint) (res bool, err error) {
	var count int64
	err = r.GetDB(tenantId).Model(&model.DebugInterface{}).Where("id=?", id).Count(&count).Error
	if err != nil {
		return
	}
	res = count > 0
	return
}

func (r *DebugInterfaceRepo) UpdateOrder(tenantId consts.TenantId, pos serverConsts.DropPos, targetId uint) (parentId uint, ordr int) {
	if pos == serverConsts.Inner {
		parentId = targetId

		var preChild model.DebugInterface
		r.GetDB(tenantId).Where("parent_id=?", parentId).
			Order("ordr DESC").Limit(1).
			First(&preChild)

		ordr = preChild.Ordr + 1

	} else if pos == serverConsts.Before {
		brother, _ := r.Get(tenantId, targetId)
		parentId = brother.ParentId

		r.GetDB(tenantId).Model(&model.DebugInterface{}).
			Where("NOT deleted AND parent_id=? AND ordr >= ?", parentId, brother.Ordr).
			Update("ordr", gorm.Expr("ordr + 1"))

		ordr = brother.Ordr

	} else if pos == serverConsts.After {
		brother, _ := r.Get(tenantId, targetId)
		parentId = brother.ParentId

		r.GetDB(tenantId).Model(&model.DebugInterface{}).
			Where("NOT deleted AND parent_id=? AND ordr > ?", parentId, brother.Ordr).
			Update("ordr", gorm.Expr("ordr + 1"))

		ordr = brother.Ordr + 1

	}

	return
}

func (r *DebugInterfaceRepo) ListByProject(tenantId consts.TenantId, projectId int) (pos []*model.DebugInterface, err error) {
	err = r.GetDB(tenantId).
		Where("project_id=?", projectId).
		Where("NOT deleted").
		Order("parent_id ASC, ordr ASC").
		Find(&pos).Error
	return
}

func (r *DebugInterfaceRepo) Get(tenantId consts.TenantId, id uint) (po model.DebugInterface, err error) {
	err = r.GetDB(tenantId).
		Where("id=?", id).
		Where("NOT deleted").
		First(&po).Error
	return
}

func (r *DebugInterfaceRepo) GetByCaseInterfaceId(tenantId consts.TenantId, caseInterfaceId uint) (po model.DebugInterface, err error) {
	err = r.GetDB(tenantId).
		Where("case_interface_id=?", caseInterfaceId).
		Where("NOT deleted").
		First(&po).Error
	return
}

func (r *DebugInterfaceRepo) GetDetail(tenantId consts.TenantId, interfId uint) (interf model.DebugInterface, err error) {
	if interfId <= 0 {
		return
	}

	interf, err = r.Get(tenantId, interfId)

	interf.QueryParams, interf.PathParams, _ = r.ListParams(tenantId, interfId)
	interf.Headers, _ = r.ListHeaders(tenantId, interfId)

	interf.BodyFormData, _ = r.ListBodyFormData(tenantId, interfId)
	interf.BodyFormUrlencoded, _ = r.ListBodyFormUrlencoded(tenantId, interfId)

	interf.BasicAuth, _ = r.GetBasicAuth(tenantId, interfId)
	interf.BearerToken, _ = r.GetBearerToken(tenantId, interfId)
	interf.OAuth20, _ = r.GetOAuth20(tenantId, interfId)
	interf.ApiKey, _ = r.GetApiKey(tenantId, interfId)

	interf.GlobalParams, _ = r.GetGlobalParams(tenantId, interfId)
	interf.Cookies, _ = r.ListCookies(tenantId, interfId)

	return
}

func (r *DebugInterfaceRepo) Save(tenantId consts.TenantId, interf *model.DebugInterface) (err error) {
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

		err = r.UpdateCookies(tenantId, interf.ID, interf.Cookies)
		if err != nil {
			return err
		}

		return err
	})

	return
}
func (r *DebugInterfaceRepo) UpdateName(tenantId consts.TenantId, req serverDomain.EndpointCase) (err error) {
	err = r.GetDB(tenantId).Model(&model.DebugInterface{}).
		Where("id = ?", req.Id).
		Update("name", req.Name).Error

	return
}

func (r *DebugInterfaceRepo) UpdateHeaders(tenantId consts.TenantId, id uint, headers []model.DebugInterfaceHeader) (err error) {
	err = r.RemoveHeaders(tenantId, id)

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
	err = r.GetDB(tenantId).Create(&newHeaders).Error

	return
}

func (r *DebugInterfaceRepo) RemoveHeaders(tenantId consts.TenantId, id uint) (err error) {
	err = r.GetDB(tenantId).
		Where("interface_id = ?", id).
		Delete(&model.DebugInterfaceHeader{}, "").Error

	return
}

func (r *DebugInterfaceRepo) UpdateParams(tenantId consts.TenantId, id uint, queryParams, pathParams []model.DebugInterfaceParam) (err error) {
	err = r.RemoveParams(tenantId, id)

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

	err = r.GetDB(tenantId).Create(&params).Error

	return
}
func (r *DebugInterfaceRepo) RemoveParams(tenantId consts.TenantId, id uint) (err error) {
	err = r.GetDB(tenantId).
		Where("interface_id = ?", id).
		Delete(&model.DebugInterfaceParam{}, "").Error

	return
}

func (r *DebugInterfaceRepo) UpdateCookies(tenantId consts.TenantId, id uint, cookies []model.DebugInterfaceCookie) (err error) {
	err = r.RemoveCookie(tenantId, id)

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
	err = r.GetDB(tenantId).Create(&newCookies).Error

	return
}

func (r *DebugInterfaceRepo) RemoveCookie(tenantId consts.TenantId, id uint) (err error) {
	err = r.GetDB(tenantId).
		Where("interface_id = ?", id).
		Delete(&model.DebugInterfaceCookie{}, "").Error

	return
}

func (r *DebugInterfaceRepo) UpdateProcessorId(tenantId consts.TenantId, id, processorId uint) (err error) {
	values := map[string]interface{}{
		"scenario_processor_id": processorId,
	}
	err = r.UpdateDebugInfo(tenantId, id, values)

	return
}

func (r *DebugInterfaceRepo) UpdateBodyFormData(tenantId consts.TenantId, id uint, items []model.DebugInterfaceBodyFormDataItem) (err error) {
	err = r.RemoveBodyFormData(tenantId, id)

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
	err = r.GetDB(tenantId).Create(&list).Error

	return
}
func (r *DebugInterfaceRepo) RemoveBodyFormData(tenantId consts.TenantId, id uint) (err error) {
	err = r.GetDB(tenantId).
		Where("interface_id = ?", id).
		Delete(&model.DebugInterfaceBodyFormDataItem{}, "").Error

	return
}

func (r *DebugInterfaceRepo) UpdateBodyFormUrlencoded(tenantId consts.TenantId, id uint, items []model.DebugInterfaceBodyFormUrlEncodedItem) (err error) {
	err = r.RemoveBodyFormUrlencoded(tenantId, id)

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
	err = r.GetDB(tenantId).Create(&list).Error

	return
}
func (r *DebugInterfaceRepo) RemoveBodyFormUrlencoded(tenantId consts.TenantId, id uint) (err error) {
	err = r.GetDB(tenantId).
		Where("interface_id = ?", id).
		Delete(&model.DebugInterfaceBodyFormUrlEncodedItem{}, "").Error

	return
}

func (r *DebugInterfaceRepo) UpdateBasicAuth(tenantId consts.TenantId, id uint, payload model.DebugInterfaceBasicAuth) (err error) {
	r.RemoveBasicAuth(tenantId, id)

	payload.InterfaceId = id
	err = r.GetDB(tenantId).Save(&payload).Error

	return
}
func (r *DebugInterfaceRepo) RemoveBasicAuth(tenantId consts.TenantId, id uint) (err error) {
	err = r.GetDB(tenantId).
		Where("interface_id = ?", id).
		Delete(&model.DebugInterfaceBasicAuth{}, "").Error

	return
}

func (r *DebugInterfaceRepo) UpdateBearerToken(tenantId consts.TenantId, id uint, payload model.DebugInterfaceBearerToken) (err error) {
	r.RemoveBearerToken(tenantId, id)

	payload.InterfaceId = id
	err = r.GetDB(tenantId).Save(&payload).Error

	return
}
func (r *DebugInterfaceRepo) RemoveBearerToken(tenantId consts.TenantId, id uint) (err error) {
	err = r.GetDB(tenantId).
		Where("interface_id = ?", id).
		Delete(&model.DebugInterfaceBearerToken{}, "").Error

	return
}

func (r *DebugInterfaceRepo) UpdateOAuth20(tenantId consts.TenantId, interfaceId uint, payload model.DebugInterfaceOAuth20) (err error) {
	r.RemoveOAuth20(tenantId, interfaceId)

	payload.InterfaceId = interfaceId
	err = r.GetDB(tenantId).Save(&payload).Error

	return
}
func (r *DebugInterfaceRepo) RemoveOAuth20(tenantId consts.TenantId, interfaceId uint) (err error) {
	err = r.GetDB(tenantId).
		Where("interface_id = ?", interfaceId).
		Delete(&model.DebugInterfaceOAuth20{}).Error

	return
}

func (r *DebugInterfaceRepo) UpdateApiKey(tenantId consts.TenantId, id uint, payload model.DebugInterfaceApiKey) (err error) {
	r.RemoveApiKey(tenantId, id)

	payload.InterfaceId = id
	err = r.GetDB(tenantId).Save(&payload).Error

	return
}
func (r *DebugInterfaceRepo) RemoveApiKey(tenantId consts.TenantId, id uint) (err error) {
	err = r.GetDB(tenantId).
		Where("interface_id = ?", id).
		Delete(&model.DebugInterfaceApiKey{}, "").Error

	return
}

func (r *DebugInterfaceRepo) Delete(tenantId consts.TenantId, id uint) (err error) {
	err = r.GetDB(tenantId).Model(&model.DebugInterface{}).
		Where("id=?", id).
		Update("deleted", true).
		Error

	return
}

func (r *DebugInterfaceRepo) GetChildren(tenantId consts.TenantId, defId, fieldId uint) (children []*model.DebugInterface, err error) {
	err = r.GetDB(tenantId).Where("defID=? AND parentID=?", defId, fieldId).Find(&children).Error
	return
}

func (r *DebugInterfaceRepo) SetIsRange(tenantId consts.TenantId, fieldId uint, b bool) (err error) {
	err = r.GetDB(tenantId).Model(&model.DebugInterface{}).
		Where("id = ?", fieldId).Update("isRange", b).Error

	return
}

func (r *DebugInterfaceRepo) UpdateOrdAndParent(tenantId consts.TenantId, interf model.DebugInterface) (err error) {
	err = r.GetDB(tenantId).Model(&interf).
		Updates(model.DebugInterface{InterfaceBase: model.InterfaceBase{Ordr: interf.Ordr, ParentId: interf.ParentId}}).
		Error

	return
}

func (r *DebugInterfaceRepo) SetOAuth2AccessToken(tenantId consts.TenantId, token string, interfaceId int) (err error) {
	err = r.GetDB(tenantId).Model(&model.DebugInterfaceOAuth20{}).
		Where("interface_id = ?", interfaceId).
		Update("access_token", token).Error

	return
}

func (r *DebugInterfaceRepo) ListParams(tenantId consts.TenantId, interfaceId uint) (
	queryParams []model.DebugInterfaceParam, pathParams []model.DebugInterfaceParam, err error) {

	pos := []model.DebugInterfaceParam{}

	err = r.GetDB(tenantId).
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
func (r *DebugInterfaceRepo) ListHeaders(tenantId consts.TenantId, interfaceId uint) (pos []model.DebugInterfaceHeader, err error) {
	err = r.GetDB(tenantId).
		Where("interface_id=?", interfaceId).
		Where("NOT deleted").
		Order("id ASC").
		Find(&pos).Error

	return
}
func (r *DebugInterfaceRepo) ListCookies(tenantId consts.TenantId, interfaceId uint) (pos []model.DebugInterfaceCookie, err error) {
	err = r.GetDB(tenantId).
		Where("interface_id=?", interfaceId).
		Where("NOT deleted").
		Order("id ASC").
		Find(&pos).Error

	return
}
func (r *DebugInterfaceRepo) ListBodyFormData(tenantId consts.TenantId, interfaceId uint) (pos []model.DebugInterfaceBodyFormDataItem, err error) {
	err = r.GetDB(tenantId).
		Where("interface_id=?", interfaceId).
		Where("NOT deleted").
		Order("id ASC").
		Find(&pos).Error

	return
}
func (r *DebugInterfaceRepo) ListBodyFormUrlencoded(tenantId consts.TenantId, interfaceId uint) (pos []model.DebugInterfaceBodyFormUrlEncodedItem, err error) {
	err = r.GetDB(tenantId).
		Where("interface_id=?", interfaceId).
		Where("NOT deleted").
		Order("id ASC").
		Find(&pos).Error

	return
}

func (r *DebugInterfaceRepo) GetBasicAuth(tenantId consts.TenantId, id uint) (po model.DebugInterfaceBasicAuth, err error) {
	err = r.GetDB(tenantId).
		Where("interface_id = ?", id).
		First(&po).Error

	return
}
func (r *DebugInterfaceRepo) GetBearerToken(tenantId consts.TenantId, id uint) (po model.DebugInterfaceBearerToken, err error) {
	err = r.GetDB(tenantId).
		Where("interface_id = ?", id).
		First(&po).Error

	return
}
func (r *DebugInterfaceRepo) GetOAuth20(tenantId consts.TenantId, id uint) (po model.DebugInterfaceOAuth20, err error) {
	err = r.GetDB(tenantId).
		Where("interface_id = ?", id).
		First(&po).Error

	return
}
func (r *DebugInterfaceRepo) GetApiKey(tenantId consts.TenantId, id uint) (po model.DebugInterfaceApiKey, err error) {
	err = r.GetDB(tenantId).
		Where("interface_id = ?", id).
		First(&po).Error

	return
}

func (r *DebugInterfaceRepo) SaveInterfaces(tenantId consts.TenantId, interf *model.DebugInterface) (err error) {
	r.GetDB(tenantId).Transaction(func(tx *gorm.DB) error {
		err = r.UpdateInterface(tenantId, interf)
		if err != nil {
			return err
		}

		err = r.UpdateParams(tenantId, interf.ID, interf.QueryParams, interf.PathParams)
		if err != nil {
			return err
		}

		err = r.UpdateHeaders(tenantId, interf.ID, interf.Headers)
		if err != nil {
			return err
		}

		err = r.UpdateCookies(tenantId, interf.ID, interf.Cookies)
		if err != nil {
			return err
		}

		return err
	})

	return
}

func (r *DebugInterfaceRepo) UpdateInterface(tenantId consts.TenantId, interf *model.DebugInterface) (err error) {
	err = r.BaseRepo.Save(tenantId, interf.ID, interf)
	return
}

func (r *DebugInterfaceRepo) PopulateProps(tenantId consts.TenantId, po *model.DebugInterface) (err error) {
	po.QueryParams, po.PathParams, _ = r.ListParams(tenantId, po.ID)
	po.Headers, _ = r.ListHeaders(tenantId, po.ID)
	po.BodyFormData, _ = r.ListBodyFormData(tenantId, po.ID)
	po.BodyFormUrlencoded, _ = r.ListBodyFormUrlencoded(tenantId, po.ID)
	po.BasicAuth, _ = r.GetBasicAuth(tenantId, po.ID)
	po.BearerToken, _ = r.GetBearerToken(tenantId, po.ID)
	po.OAuth20, _ = r.GetOAuth20(tenantId, po.ID)
	po.ApiKey, _ = r.GetApiKey(tenantId, po.ID)

	return
}

func (r *DebugInterfaceRepo) UpdateDebugInfo(tenantId consts.TenantId, id uint, values map[string]interface{}) (err error) {
	err = r.GetDB(tenantId).Model(&model.DebugInterface{}).
		Where("id=?", id).
		Updates(values).
		Error

	return
}

func (r *DebugInterfaceRepo) DeleteByProcessorIds(tenantId consts.TenantId, ids []uint) (err error) {
	if len(ids) > 0 {
		err = r.GetDB(tenantId).Model(&model.DebugInterface{}).
			Where("scenario_processor_id IN (?)", ids).
			Update("deleted", true).
			Error
	}

	return
}

func (r *DebugInterfaceRepo) CreateDefault(tenantId consts.TenantId, src consts.ProcessorInterfaceSrc, projectId uint) (id uint, err error) {
	po := model.DebugInterface{
		ProcessorInterfaceSrc: src,

		InterfaceBase: model.InterfaceBase{
			InterfaceConfigBase: model.InterfaceConfigBase{
				Method: consts.GET,
			},
			ProjectId: projectId,
		},
	}

	serves, _ := r.ServeRepo.ListByProject(tenantId, projectId)
	if len(serves) > 0 {
		po.ServeId = serves[0].ID

		server, _ := r.ServeServerRepo.GetDefaultByServe(tenantId, po.ServeId)
		po.ServerId = server.ID
	}

	err = r.Save(tenantId, &po)

	id = po.ID

	return
}

func (r *DebugInterfaceRepo) GetSourceNameById(tenantId consts.TenantId, id uint) (name string, err error) {
	debugInterface, err := r.Get(tenantId, id)
	if err != nil {
		return
	}

	switch debugInterface.ProcessorInterfaceSrc {
	case consts.InterfaceSrcDefine:
		endpointInterface, err := r.EndpointInterfaceRepo.Get(tenantId, debugInterface.EndpointInterfaceId)
		if err != nil {
			return "", err
		}
		name = endpointInterface.Name
	case consts.InterfaceSrcCase:
		endpointCase, err := r.EndpointCaseRepo.Get(tenantId, debugInterface.CaseInterfaceId)
		if err != nil {
			return "", err
		}
		name = endpointCase.Name
	case consts.InterfaceSrcDiagnose:
		diagnoseInterface, err := r.DiagnoseInterfaceRepo.Get(tenantId, debugInterface.DiagnoseInterfaceId)
		if err != nil {
			return "", err
		}
		name = diagnoseInterface.Title
	}

	return
}

func (r *DebugInterfaceRepo) UpdateGlobalParams(tenantId consts.TenantId, id uint, params []model.DebugInterfaceGlobalParam) (err error) {
	err = r.RemoveGlobalParams(tenantId, id)

	if len(params) == 0 {
		return
	}

	for key, _ := range params {
		params[key].InterfaceId = id
	}

	err = r.GetDB(tenantId).Create(&params).Error

	return
}

func (r *DebugInterfaceRepo) RemoveGlobalParams(tenantId consts.TenantId, id uint) (err error) {
	err = r.GetDB(tenantId).
		Where("interface_id = ?", id).
		Delete(&model.DebugInterfaceGlobalParam{}, "").Error

	return
}

func (r *DebugInterfaceRepo) GetGlobalParams(tenantId consts.TenantId, id uint) (po []model.DebugInterfaceGlobalParam, err error) {
	err = r.GetDB(tenantId).
		Where("interface_id = ?", id).
		Find(&po).Error
	return
}

func (r *DebugInterfaceRepo) SyncPath(tenantId consts.TenantId, endpointId, serveId uint, newPath, oldPath string) {
	if endpointId == 0 {
		return
	}

	interfaceIds, err := r.EndpointInterfaceRepo.ListIdByEndpoint(tenantId, endpointId)
	if err != nil {
		return
	}
	if len(interfaceIds) > 0 {
		r.UpdateDefinePath(tenantId, interfaceIds, newPath, oldPath)
		r.UpdateServeId(tenantId, interfaceIds, serveId)
	}
}

// UpdateDefinePath 如果路径没变更，则更新接口定义-调试-接口定义-用例路径
func (r *DebugInterfaceRepo) UpdateDefinePath(tenantId consts.TenantId, ids []uint, newPath, oldPath string) (err error) {
	err = r.GetDB(tenantId).Model(&model.DebugInterface{}).
		Where("endpoint_interface_id in ? and url = ? and scenario_processor_id = 0 and diagnose_interface_id = 0", ids, oldPath).
		Update("url", newPath).
		Error
	return
}

func (r *DebugInterfaceRepo) UpdateServeId(tenantId consts.TenantId, ids []uint, serveId uint) (err error) {
	err = r.GetDB(tenantId).Model(&model.DebugInterface{}).
		Where("endpoint_interface_id in ?", ids).
		Update("serve_id", serveId).
		Error
	return
}

func (r *DebugInterfaceRepo) SyncServeId(tenantId consts.TenantId, endpointIds []uint, serveId uint) (err error) {
	if len(endpointIds) == 0 {
		return
	}

	interfaceIds, err := r.EndpointInterfaceRepo.ListIdByEndpoints(tenantId, endpointIds)
	if err != nil {
		return
	}
	if len(interfaceIds) > 0 {
		err = r.UpdateServeId(tenantId, interfaceIds, serveId)
	}

	return
}
