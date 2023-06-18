package repo

import (
	serverDomain "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	serverConsts "github.com/aaronchen2k/deeptest/internal/server/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/kataras/iris/v12"
	"gorm.io/gorm"
)

type TestInterfaceRepo struct {
	*BaseRepo           `inject:""`
	*DebugInterfaceRepo `inject:""`
	DB                  *gorm.DB `inject:""`
}

func (r *TestInterfaceRepo) GetTree(projectId, serveId uint) (root *serverDomain.TestInterface, err error) {
	pos, err := r.ListByProject(projectId, serveId)
	if err != nil {
		return
	}

	tos := r.toTos(pos)
	if len(tos) == 0 {
		return
	}

	root = &serverDomain.TestInterface{}
	r.makeTree(tos, root)

	return
}

func (r *TestInterfaceRepo) ListByProject(projectId, serveId uint) (pos []*model.TestInterface, err error) {
	db := r.DB.
		Where("project_id=?", projectId).
		Where("NOT deleted")

	if serveId > 0 {
		db.Where("serve_id=?", serveId)
	}

	err = db.
		Order("parent_id ASC, ordr ASC").
		Find(&pos).Error

	return
}

func (r *TestInterfaceRepo) Get(id uint) (po model.TestInterface, err error) {
	err = r.DB.Where("id = ?", id).First(&po).Error
	return
}

func (r *TestInterfaceRepo) toTos(pos []*model.TestInterface) (tos []*serverDomain.TestInterface) {
	for _, po := range pos {
		to := r.toTo(po)

		tos = append(tos, to)
	}

	return
}
func (r *TestInterfaceRepo) toTo(po *model.TestInterface) (to *serverDomain.TestInterface) {

	to = &serverDomain.TestInterface{
		Id:       int64(po.ID),
		Title:    po.Title,
		Desc:     po.Desc,
		Type:     po.Type,
		ParentId: int64(po.ParentId),
	}

	if po.Type == serverConsts.TestInterfaceTypeInterface {
		to.IsLeaf = true
	}

	return
}

func (r *TestInterfaceRepo) makeTree(findIn []*serverDomain.TestInterface, parent *serverDomain.TestInterface) { //参数为父节点，添加父节点的子节点指针切片
	children, _ := r.hasChild(findIn, parent) // 判断节点是否有子节点并返回

	if children != nil {
		parent.Children = append(parent.Children, children[0:]...) // 添加子节点

		for _, child := range children { // 查询子节点的子节点，并添加到子节点
			_, has := r.hasChild(findIn, child)
			if has {
				r.makeTree(findIn, child) // 递归添加节点
			}
		}
	}
}

func (r *TestInterfaceRepo) hasChild(categories []*serverDomain.TestInterface, parent *serverDomain.TestInterface) (
	ret []*serverDomain.TestInterface, yes bool) {

	for _, item := range categories {
		if item.ParentId == parent.Id {
			item.Slots = iris.Map{"icon": "icon"}
			//item.Parent = parent // loop json

			ret = append(ret, item)
		}
	}

	if ret != nil {
		yes = true
	}

	return
}

func (r *TestInterfaceRepo) Save(po *model.TestInterface) (err error) {
	po.Ordr = r.GetMaxOrder(po.ParentId, po.ProjectId)

	err = r.DB.Save(po).Error

	return
}

func (r *TestInterfaceRepo) Update(req serverDomain.TestInterfaceReq) (err error) {
	err = r.DB.Model(&model.TestInterface{}).
		Where("id ?", req.Id).
		Updates(map[string]interface{}{"title": req.Title, "desc": req.Desc}).Error

	return
}

func (r *TestInterfaceRepo) UpdateOrder(pos serverConsts.DropPos, targetId uint, projectId uint) (
	parentId uint, ordr int) {
	if pos == serverConsts.Inner {
		parentId = targetId

		var preChild model.TestInterface
		r.DB.Where("parent_id=? AND project_id = ?", parentId, projectId).
			Order("ordr DESC").Limit(1).
			First(&preChild)

		ordr = preChild.Ordr + 1

	} else if pos == serverConsts.Before {
		brother, _ := r.Get(targetId)
		parentId = brother.ParentId

		r.DB.Model(&model.TestInterface{}).
			Where("NOT deleted AND parent_id=? AND project_id = ? AND ordr >= ?",
				parentId, projectId, brother.Ordr).
			Update("ordr", gorm.Expr("ordr + 1"))

		ordr = brother.Ordr

	} else if pos == serverConsts.After {
		brother, _ := r.Get(targetId)
		parentId = brother.ParentId

		r.DB.Model(&model.TestInterface{}).
			Where("NOT deleted AND parent_id=? AND project_id = ? AND ordr > ?",
				parentId, parentId, brother.Ordr).
			Update("ordr", gorm.Expr("ordr + 1"))

		ordr = brother.Ordr + 1

	}

	return
}

func (r *TestInterfaceRepo) UpdateName(id int, name string) (err error) {
	err = r.DB.Model(&model.TestInterface{}).
		Where("id = ?", id).
		Update("name", name).Error

	return
}

func (r *TestInterfaceRepo) Delete(id uint) (err error) {
	err = r.DB.Model(&model.TestInterface{}).
		Where("id=?", id).
		Update("deleted", true).
		Error

	return
}

func (r *TestInterfaceRepo) GetChildren(nodeId uint) (children []*model.TestInterface, err error) {
	err = r.DB.Where("parent_id=?", nodeId).Find(&children).Error
	return
}

func (r *TestInterfaceRepo) UpdateOrdAndParent(node model.TestInterface) (err error) {
	err = r.DB.Model(&node).
		Updates(model.TestInterface{Ordr: node.Ordr, ParentId: node.ParentId}).
		Error

	return
}

func (r *TestInterfaceRepo) GetMaxOrder(parentId uint, projectId uint) (order int) {
	node := model.TestInterface{}

	err := r.DB.Model(&model.TestInterface{}).
		Where("parent_id=? AND project_id = ?", parentId, projectId).
		Order("ordr DESC").
		First(&node).Error

	if err == nil {
		order = node.Ordr + 1
	}

	return
}

func (r *TestInterfaceRepo) Remove(id uint, typ serverConsts.TestInterfaceType) (err error) {
	ids := []uint{}

	if typ == serverConsts.TestInterfaceTypeDir {
		ids, _ = r.GetAllChildIdsSimple(id, model.TestInterface{}.TableName())
	} else {
		ids = append(ids, id)
	}

	err = r.DB.Model(&model.TestInterface{}).
		Where("id IN (?)", ids).
		Update("deleted", true).Error

	return
}

func (r *TestInterfaceRepo) SaveDebugData(interf *model.TestInterface) (err error) {
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

func (r *TestInterfaceRepo) UpdateParams(id uint, queryParams, pathParams []model.ScenarioInterfaceParam) (err error) {
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

func (r *TestInterfaceRepo) RemoveParams(id uint) (err error) {
	err = r.DB.
		Where("interface_id = ?", id).
		Delete(&model.ScenarioInterfaceParam{}, "").Error

	return
}

func (r *TestInterfaceRepo) UpdateBodyFormData(id uint, items []model.ScenarioInterfaceBodyFormDataItem) (err error) {
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

func (r *TestInterfaceRepo) RemoveBodyFormData(id uint) (err error) {
	err = r.DB.
		Where("interface_id = ?", id).
		Delete(&model.ScenarioInterfaceBodyFormDataItem{}, "").Error

	return
}

func (r *TestInterfaceRepo) UpdateBodyFormUrlencoded(id uint, items []model.ScenarioInterfaceBodyFormUrlEncodedItem) (err error) {
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

func (r *TestInterfaceRepo) RemoveBodyFormUrlencoded(id uint) (err error) {
	err = r.DB.
		Where("interface_id = ?", id).
		Delete(&model.ScenarioInterfaceBodyFormUrlEncodedItem{}, "").Error

	return
}

func (r *TestInterfaceRepo) UpdateHeaders(id uint, headers []model.ScenarioInterfaceHeader) (err error) {
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

func (r *TestInterfaceRepo) RemoveHeaders(id uint) (err error) {
	err = r.DB.
		Where("interface_id = ?", id).
		Delete(&model.ScenarioInterfaceHeader{}, "").Error

	return
}

func (r *TestInterfaceRepo) UpdateBasicAuth(id uint, payload model.ScenarioInterfaceBasicAuth) (err error) {
	if err = r.RemoveBasicAuth(id); err != nil {
		return
	}

	payload.InterfaceId = id
	err = r.DB.Save(&payload).Error

	return
}

func (r *TestInterfaceRepo) RemoveBasicAuth(id uint) (err error) {
	err = r.DB.
		Where("interface_id = ?", id).
		Delete(&model.ScenarioInterfaceBasicAuth{}, "").Error

	return
}

func (r *TestInterfaceRepo) UpdateBearerToken(id uint, payload model.ScenarioInterfaceBearerToken) (err error) {
	if err = r.RemoveBearerToken(id); err != nil {
		return
	}

	payload.InterfaceId = id
	err = r.DB.Save(&payload).Error

	return
}

func (r *TestInterfaceRepo) RemoveBearerToken(id uint) (err error) {
	err = r.DB.
		Where("interface_id = ?", id).
		Delete(&model.ScenarioInterfaceBearerToken{}, "").Error

	return
}

func (r *TestInterfaceRepo) UpdateOAuth20(interfaceId uint, payload model.ScenarioInterfaceOAuth20) (err error) {
	if err = r.RemoveOAuth20(interfaceId); err != nil {
		return
	}

	payload.InterfaceId = interfaceId
	err = r.DB.Save(&payload).Error

	return
}

func (r *TestInterfaceRepo) RemoveOAuth20(interfaceId uint) (err error) {
	err = r.DB.
		Where("interface_id = ?", interfaceId).
		Delete(&model.ScenarioInterfaceOAuth20{}).Error

	return
}

func (r *TestInterfaceRepo) UpdateApiKey(id uint, payload model.ScenarioInterfaceApiKey) (err error) {
	if err = r.RemoveApiKey(id); err != nil {
		return
	}

	payload.InterfaceId = id
	err = r.DB.Save(&payload).Error

	return
}

func (r *TestInterfaceRepo) RemoveApiKey(id uint) (err error) {
	err = r.DB.
		Where("interface_id = ?", id).
		Delete(&model.ScenarioInterfaceApiKey{}, "").Error

	return
}
