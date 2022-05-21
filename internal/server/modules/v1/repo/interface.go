package repo

import (
	"github.com/aaronchen2k/deeptest/internal/server/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/model"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"github.com/kataras/iris/v12"
	"gorm.io/gorm"
)

type InterfaceRepo struct {
	DB *gorm.DB `inject:""`
}

func (r *InterfaceRepo) GetInterfaceTree(projectId int) (root *model.Interface, err error) {
	pos, err := r.ListByProject(projectId)

	if err != nil {
		return
	}

	root = pos[0]
	root.Slots = iris.Map{"icon": "icon"}
	r.makeTree(pos[1:], root)
	return
}

func (r *InterfaceRepo) UpdateOrder(pos serverConsts.DropPos, targetId uint) (parentId uint, ordr int) {
	if pos == serverConsts.Inner {
		parentId = targetId

		var preChild model.Interface
		r.DB.Where("parent_id=?", parentId).
			Order("ordr DESC").Limit(1).
			First(&preChild)

		ordr = preChild.Ordr + 1

	} else if pos == serverConsts.Before {
		brother, _ := r.Get(targetId)
		parentId = brother.ParentId

		r.DB.Model(&model.Interface{}).
			Where("NOT deleted AND parent_id=? AND ordr >= ?", parentId, brother.Ordr).
			Update("ordr", gorm.Expr("ordr + 1"))

		ordr = brother.Ordr

	} else if pos == serverConsts.After {
		brother, _ := r.Get(targetId)
		parentId = brother.ParentId

		r.DB.Model(&model.Interface{}).
			Where("NOT deleted AND parent_id=? AND ordr > ?", parentId, brother.Ordr).
			Update("ordr", gorm.Expr("ordr + 1"))

		ordr = brother.Ordr + 1

	}

	return
}

func (r *InterfaceRepo) ListByProject(projectId int) (pos []*model.Interface, err error) {
	err = r.DB.
		Where("project_id=?", projectId).
		Where("NOT deleted").
		Order("parent_id ASC, ordr ASC").
		Find(&pos).Error
	return
}

func (r *InterfaceRepo) Get(fieldId uint) (field model.Interface, err error) {
	err = r.DB.
		Where("id=?", fieldId).
		Where("NOT deleted").
		First(&field).Error
	return
}

func (r *InterfaceRepo) Save(field *model.Interface) (err error) {
	err = r.DB.Save(field).Error
	return
}

func (r *InterfaceRepo) Update(interf model.Interface) (err error) {
	r.DB.Transaction(func(tx *gorm.DB) error {
		err = r.DB.Updates(interf).Error
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
func (r *InterfaceRepo) UpdateName(id int, name string) (err error) {
	err = r.DB.Model(&model.Interface{}).
		Where("id = ?", id).
		Update("name", name).Error

	return
}
func (r *InterfaceRepo) UpdateDefaultEnvironment(id, envId uint) (err error) {
	err = r.DB.Model(&model.Interface{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{"environment_id": envId}).Error

	if err != nil {
		logUtils.Errorf("update project environment error", err.Error())
		return err
	}

	return
}
func (r *InterfaceRepo) UpdateParams(id uint, params []model.InterfaceParam) (err error) {
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
func (r *InterfaceRepo) UpdateHeaders(id uint, headers []model.InterfaceHeader) (err error) {
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

func (r *InterfaceRepo) UpdateBasicAuth(id uint, payload model.InterfaceBasicAuth) (err error) {
	payload.InterfaceId = id
	err = r.DB.Save(&payload).Error

	return
}

func (r *InterfaceRepo) UpdateBearerToken(id uint, payload model.InterfaceBearerToken) (err error) {
	payload.InterfaceId = id
	err = r.DB.Save(&payload).Error

	return
}

func (r *InterfaceRepo) UpdateOAuth20(interfaceId uint, payload model.InterfaceOAuth20) (err error) {
	r.RemoveOAuth20(interfaceId)

	payload.InterfaceId = interfaceId
	err = r.DB.Save(&payload).Error

	return
}
func (r *InterfaceRepo) RemoveOAuth20(interfaceId uint) (err error) {
	err = r.DB.
		Where("interface_id = ?", interfaceId).
		Delete(&model.InterfaceOAuth20{}).Error

	return
}

func (r *InterfaceRepo) UpdateApiKey(id uint, payload model.InterfaceApiKey) (err error) {
	payload.InterfaceId = id
	err = r.DB.Save(&payload).Error

	return
}

func (r *InterfaceRepo) RemoveParams(id uint) (err error) {
	err = r.DB.
		Where("interface_id = ?", id).
		Delete(&model.InterfaceParam{}, "").Error

	return
}
func (r *InterfaceRepo) RemoveHeaders(id uint) (err error) {
	err = r.DB.
		Where("interface_id = ?", id).
		Delete(&model.InterfaceHeader{}, "").Error

	return
}
func (r *InterfaceRepo) GetBasicAuth(id uint) (po model.InterfaceBasicAuth, err error) {
	err = r.DB.
		Where("interface_id = ?", id).
		First(&po).Error

	return
}
func (r *InterfaceRepo) GetBearerToken(id uint) (po model.InterfaceBearerToken, err error) {
	err = r.DB.
		Where("interface_id = ?", id).
		First(&po).Error

	return
}
func (r *InterfaceRepo) GetOAuth20(id uint) (po model.InterfaceOAuth20, err error) {
	err = r.DB.
		Where("interface_id = ?", id).
		First(&po).Error

	return
}
func (r *InterfaceRepo) GetApiKey(id uint) (po model.InterfaceApiKey, err error) {
	err = r.DB.
		Where("interface_id = ?", id).
		First(&po).Error

	return
}

func (r *InterfaceRepo) makeTree(Data []*model.Interface, node *model.Interface) { //参数为父节点，添加父节点的子节点指针切片
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

func (r *InterfaceRepo) haveChild(Data []*model.Interface, node *model.Interface) (child []*model.Interface, yes bool) {
	for _, v := range Data {
		if v.ParentId == node.ID {
			v.Slots = iris.Map{"icon": "icon"}
			child = append(child, v)
		}
	}
	if child != nil {
		yes = true
	}
	return
}

func (r *InterfaceRepo) Delete(id uint) (err error) {
	err = r.DB.Model(&model.Interface{}).
		Where("id=?", id).
		Update("deleted", true).
		Error

	//field := model.Interface{}
	//field.ID = id
	//err = r.DB.Delete(field).Error

	return
}

func (r *InterfaceRepo) GetChildren(defId, fieldId uint) (children []*model.Interface, err error) {
	err = r.DB.Where("defID=? AND parentID=?", defId, fieldId).Find(&children).Error
	return
}

func (r *InterfaceRepo) SetIsRange(fieldId uint, b bool) (err error) {
	err = r.DB.Model(&model.Interface{}).
		Where("id = ?", fieldId).Update("isRange", b).Error

	return
}

func (r *InterfaceRepo) UpdateOrdAndParent(interf model.Interface) (err error) {
	err = r.DB.Model(&interf).
		Updates(model.Interface{Ordr: interf.Ordr, ParentId: interf.ParentId}).
		Error

	return
}

func (r *InterfaceRepo) SetOAuth2AccessToken(token string, interfaceId int) (err error) {
	err = r.DB.Model(&model.InterfaceOAuth20{}).
		Where("interface_id = ?", interfaceId).
		Update("access_token", token).Error

	return
}

func (r *InterfaceRepo) ListParams(interfaceId uint) (pos []model.InterfaceParam, err error) {
	err = r.DB.
		Where("interface_id=?", interfaceId).
		Where("NOT deleted").
		Order("id ASC").
		Find(&pos).Error
	return
}

func (r *InterfaceRepo) ListHeaders(interfaceId uint) (pos []model.InterfaceHeader, err error) {
	err = r.DB.
		Where("interface_id=?", interfaceId).
		Where("NOT deleted").
		Order("id ASC").
		Find(&pos).Error

	return
}
