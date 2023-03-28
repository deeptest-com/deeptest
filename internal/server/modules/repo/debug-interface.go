package repo

import (
	"github.com/aaronchen2k/deeptest/internal/server/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/kataras/iris/v12"
	"gorm.io/gorm"
)

type DebugInterfaceRepo struct {
	*BaseRepo `inject:""`
	DB        *gorm.DB `inject:""`
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

func (r *DebugInterfaceRepo) GetInterfaceTree(projectId int) (root *model.DebugInterface, err error) {
	pos, err := r.ListByProject(projectId)

	if err != nil || len(pos) == 0 {
		return
	}

	root = pos[0]
	root.Slots = iris.Map{"icon": "icon"}
	r.makeTree(pos[1:], root)
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

func (r *DebugInterfaceRepo) Get(interfaceId uint) (field model.DebugInterface, err error) {
	err = r.DB.
		Where("id=?", interfaceId).
		Where("NOT deleted").
		First(&field).Error
	return
}

func (r *DebugInterfaceRepo) GetDetail(interfId uint) (interf model.DebugInterface, err error) {
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

func (r *DebugInterfaceRepo) Save(interf *model.DebugInterface) (err error) {
	err = r.DB.Save(interf).Error
	return
}

func (r *DebugInterfaceRepo) Update(interf model.DebugInterface) (err error) {
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
func (r *DebugInterfaceRepo) UpdateName(id int, name string) (err error) {
	err = r.DB.Model(&model.DebugInterface{}).
		Where("id = ?", id).
		Update("name", name).Error

	return
}

func (r *DebugInterfaceRepo) UpdateHeaders(id uint, headers []model.DebugInterfaceHeader) (err error) {
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

func (r *DebugInterfaceRepo) RemoveHeaders(id uint) (err error) {
	err = r.DB.
		Where("interface_id = ?", id).
		Delete(&model.DebugInterfaceHeader{}, "").Error

	return
}

func (r *DebugInterfaceRepo) UpdateParams(id uint, params []model.DebugInterfaceParam) (err error) {
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

	for idx, _ := range cookies {
		cookies[idx].ID = 0
		cookies[idx].InterfaceId = id
	}

	err = r.DB.Create(&cookies).Error

	return
}

func (r *DebugInterfaceRepo) RemoveCookie(id uint) (err error) {
	err = r.DB.
		Where("interface_id = ?", id).
		Delete(&model.DebugInterfaceCookie{}, "").Error

	return
}

func (r *DebugInterfaceRepo) UpdateBodyFormData(id uint, items []model.DebugInterfaceBodyFormDataItem) (err error) {
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

	for idx, _ := range items {
		items[idx].ID = 0
		items[idx].InterfaceId = id
	}

	err = r.DB.Create(&items).Error

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

func (r *DebugInterfaceRepo) makeTree(Data []*model.DebugInterface, node *model.DebugInterface) { //参数为父节点，添加父节点的子节点指针切片
	children, _ := r.haveChild(Data, node) //判断节点是否有子节点并返回
	if children != nil {
		node.Children = append(node.Children, children[0:]...) //添加子节点
		for _, v := range children {                           //查询子节点的子节点，并添加到子节点
			_, has := r.haveChild(Data, v)
			if has {
				r.makeTree(Data, v) //递归添加节点
			}
		}
	}
}

func (r *DebugInterfaceRepo) haveChild(Data []*model.DebugInterface, node *model.DebugInterface) (children []*model.DebugInterface, yes bool) {
	for _, v := range Data {
		if v.ParentId == node.ID {
			v.Slots = iris.Map{"icon": "icon"}
			children = append(children, v)
		}
	}
	if children != nil {
		yes = true
	}
	return
}

func (r *DebugInterfaceRepo) Delete(id uint) (err error) {
	err = r.DB.Model(&model.DebugInterface{}).
		Where("id=?", id).
		Update("deleted", true).
		Error

	//field := model.UsedByInterface{}
	//field.ID = id
	//err = r.DB.Delete(field).SendErrorMsg

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

func (r *DebugInterfaceRepo) ListParams(interfaceId uint) (pos []model.DebugInterfaceParam, err error) {
	err = r.DB.
		Where("interface_id=?", interfaceId).
		Where("NOT deleted").
		Order("id ASC").
		Find(&pos).Error
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

		err = r.UpdateParams(interf.ID, interf.Params)
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

func (r *DebugInterfaceRepo) GetByEndpointId(endpointId uint, version string) (interfaces []model.DebugInterface, err error) {

	interfaces, err = r.GetEndpointId(endpointId, version)
	for key, interf := range interfaces {
		interfaces[key].Params, _ = r.ListParams(interf.ID)
		interfaces[key].Headers, _ = r.ListHeaders(interf.ID)
		interfaces[key].Cookies, _ = r.ListCookies(interf.ID)
	}

	return
}

func (r *DebugInterfaceRepo) GetEndpointId(endpointId uint, version string) (field []model.DebugInterface, err error) {
	err = r.DB.
		Where("endpoint_id=? and version=?", endpointId, version).
		Where("NOT deleted").
		Find(&field).Error
	return
}

func (r *DebugInterfaceRepo) GetById(interfId uint) (interf model.DebugInterface, err error) {
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
