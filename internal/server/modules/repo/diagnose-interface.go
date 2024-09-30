package repo

import (
	serverDomain "github.com/deeptest-com/deeptest/cmd/server/v1/domain"
	"github.com/deeptest-com/deeptest/internal/pkg/consts"
	serverConsts "github.com/deeptest-com/deeptest/internal/server/consts"
	"github.com/deeptest-com/deeptest/internal/server/modules/model"
	"github.com/kataras/iris/v12"
	"gorm.io/gorm"
)

type DiagnoseInterfaceRepo struct {
	*BaseRepo              `inject:""`
	*DiagnoseInterfaceRepo `inject:""`
	*DebugInterfaceRepo    `inject:""`
	DB                     *gorm.DB `inject:""`
}

func (r *DiagnoseInterfaceRepo) GetTree(tenantId consts.TenantId, projectId uint) (root *serverDomain.DiagnoseInterface, err error) {
	pos, err := r.ListByProject(tenantId, projectId)
	if err != nil {
		return
	}

	tos := r.toTos(tenantId, pos)
	if len(tos) == 0 {
		return
	}

	root = &serverDomain.DiagnoseInterface{}
	r.makeTree(tenantId, tos, root)
	r.mountCount(tenantId, root)

	return
}

func (r *DiagnoseInterfaceRepo) mountCount(tenantId consts.TenantId, root *serverDomain.DiagnoseInterface) (count int) {
	for _, child := range root.Children {
		root.Count += r.mountCount(tenantId, child)
	}
	return root.Count

}

func (r *DiagnoseInterfaceRepo) ListByProject(tenantId consts.TenantId, projectId uint) (pos []*model.DiagnoseInterface, err error) {
	db := r.GetDB(tenantId).
		Where("project_id=?", projectId).
		Where("NOT deleted")

	err = db.
		Order("parent_id ASC, ordr ASC").
		Find(&pos).Error

	return
}

func (r *DiagnoseInterfaceRepo) Get(tenantId consts.TenantId, id uint) (po model.DiagnoseInterface, err error) {
	err = r.GetDB(tenantId).Where("id = ?", id).First(&po).Error
	return
}

func (r *DiagnoseInterfaceRepo) GetDetail(tenantId consts.TenantId, interfId uint) (diagnoseInterface model.DiagnoseInterface, err error) {
	if interfId <= 0 {
		return
	}

	diagnoseInterface, err = r.Get(tenantId, interfId)

	debugInterface, _ := r.DebugInterfaceRepo.Get(tenantId, diagnoseInterface.DebugInterfaceId)

	debugData, _ := r.DebugInterfaceRepo.GetDetail(tenantId, debugInterface.ID)
	diagnoseInterface.DebugData = &debugData

	return
}

func (r *DiagnoseInterfaceRepo) toTos(tenantId consts.TenantId, pos []*model.DiagnoseInterface) (tos []*serverDomain.DiagnoseInterface) {
	for _, po := range pos {
		to := r.ToTo(tenantId, po)

		tos = append(tos, to)
	}

	return
}
func (r *DiagnoseInterfaceRepo) ToTo(tenantId consts.TenantId, po *model.DiagnoseInterface) (to *serverDomain.DiagnoseInterface) {
	to = &serverDomain.DiagnoseInterface{
		Id:               int64(po.ID),
		Title:            po.Title,
		Desc:             po.Desc,
		Type:             po.Type,
		ParentId:         int64(po.ParentId),
		ServeId:          po.ServeId,
		DebugInterfaceId: po.DebugInterfaceId,
		Method:           po.Method,
		IsDir:            true,
	}

	if po.Type == serverConsts.DiagnoseInterfaceTypeInterface {
		to.IsDir = false
		to.Count = 1
	}

	return
}

func (r *DiagnoseInterfaceRepo) makeTree(tenantId consts.TenantId, findIn []*serverDomain.DiagnoseInterface, parent *serverDomain.DiagnoseInterface) { //参数为父节点，添加父节点的子节点指针切片
	children, _ := r.hasChild(tenantId, findIn, parent) // 判断节点是否有子节点并返回

	if children != nil {
		parent.Children = append(parent.Children, children[0:]...) // 添加子节点

		for _, child := range children { // 查询子节点的子节点，并添加到子节点
			_, has := r.hasChild(tenantId, findIn, child)
			if has {
				r.makeTree(tenantId, findIn, child) // 递归添加节点
			}
		}
	}
}

func (r *DiagnoseInterfaceRepo) hasChild(tenantId consts.TenantId, categories []*serverDomain.DiagnoseInterface, parent *serverDomain.DiagnoseInterface) (
	ret []*serverDomain.DiagnoseInterface, yes bool) {

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

func (r *DiagnoseInterfaceRepo) Save(tenantId consts.TenantId, po *model.DiagnoseInterface) (err error) {
	if po.ID == 0 {
		po.Ordr = r.GetMaxOrder(tenantId, po.ParentId)
	}

	err = r.GetDB(tenantId).Save(po).Error

	return
}

func (r *DiagnoseInterfaceRepo) Update(tenantId consts.TenantId, req serverDomain.DiagnoseInterfaceReq) (err error) {
	err = r.GetDB(tenantId).Model(&model.DiagnoseInterface{}).
		Where("id ?", req.Id).
		Updates(map[string]interface{}{"title": req.Title, "desc": req.Desc}).Error

	return
}

func (r *DiagnoseInterfaceRepo) UpdateOrder(tenantId consts.TenantId, pos serverConsts.DropPos, targetId uint, projectId uint) (
	parentId uint, ordr int) {
	if pos == serverConsts.Inner {
		parentId = targetId

		var preChild model.DiagnoseInterface
		r.GetDB(tenantId).Where("parent_id=? AND project_id = ?", parentId, projectId).
			Order("ordr DESC").Limit(1).
			First(&preChild)

		ordr = preChild.Ordr + 1

	} else if pos == serverConsts.Before {
		brother, _ := r.Get(tenantId, targetId)
		parentId = brother.ParentId

		r.GetDB(tenantId).Model(&model.DiagnoseInterface{}).
			Where("NOT deleted AND parent_id=? AND project_id = ? AND ordr >= ?",
				parentId, projectId, brother.Ordr).
			Update("ordr", gorm.Expr("ordr + 1"))

		ordr = brother.Ordr

	} else if pos == serverConsts.After {
		brother, _ := r.Get(tenantId, targetId)
		parentId = brother.ParentId

		r.GetDB(tenantId).Model(&model.DiagnoseInterface{}).
			Where("NOT deleted AND parent_id=? AND project_id = ? AND ordr > ?",
				parentId, parentId, brother.Ordr).
			Update("ordr", gorm.Expr("ordr + 1"))

		ordr = brother.Ordr + 1

	}

	return
}

func (r *DiagnoseInterfaceRepo) UpdateName(tenantId consts.TenantId, id int, name string) (err error) {
	err = r.GetDB(tenantId).Model(&model.DiagnoseInterface{}).
		Where("id = ?", id).
		Update("name", name).Error

	return
}

func (r *DiagnoseInterfaceRepo) Delete(tenantId consts.TenantId, id uint) (err error) {
	err = r.GetDB(tenantId).Model(&model.DiagnoseInterface{}).
		Where("id=?", id).
		Update("deleted", true).
		Error

	return
}

func (r *DiagnoseInterfaceRepo) GetChildren(tenantId consts.TenantId, nodeId uint) (children []*model.DiagnoseInterface, err error) {
	err = r.GetDB(tenantId).Where("parent_id=?", nodeId).Find(&children).Error
	return
}

func (r *DiagnoseInterfaceRepo) UpdateOrdAndParent(tenantId consts.TenantId, node model.DiagnoseInterface) (err error) {
	err = r.GetDB(tenantId).Model(&node).
		Updates(model.DiagnoseInterface{Ordr: node.Ordr, ParentId: node.ParentId}).
		Error

	return
}

func (r *DiagnoseInterfaceRepo) GetMaxOrder(tenantId consts.TenantId, parentId uint) (order int) {
	node := model.DiagnoseInterface{}

	err := r.GetDB(tenantId).Model(&model.DiagnoseInterface{}).
		Where("parent_id=?", parentId).
		Order("ordr DESC").
		First(&node).Error

	if err == nil {
		order = node.Ordr + 1
	}

	return
}

func (r *DiagnoseInterfaceRepo) Remove(tenantId consts.TenantId, id uint, typ serverConsts.DiagnoseInterfaceType) (err error) {
	ids := []uint{}

	if typ == serverConsts.DiagnoseInterfaceTypeInterface {
		ids = append(ids, id)
	} else {
		ids, _ = r.GetAllChildIdsSimple(tenantId, id, model.DiagnoseInterface{}.TableName())
	}

	err = r.GetDB(tenantId).Model(&model.DiagnoseInterface{}).
		Where("id IN (?)", ids).
		Update("deleted", true).Error

	return
}

func (r *DiagnoseInterfaceRepo) SaveDebugData(tenantId consts.TenantId, interf *model.DiagnoseInterface) (err error) {
	r.GetDB(tenantId).Transaction(func(tx *gorm.DB) error {
		err = r.UpdateDebugInfo(tenantId, interf)
		if err != nil {
			return err
		}

		// TODO: save debug data

		return err
	})

	return
}

func (r *DiagnoseInterfaceRepo) UpdateDebugInfo(tenantId consts.TenantId, interf *model.DiagnoseInterface) (err error) {
	values := map[string]interface{}{
		"server_id": interf.DebugData.ServerId,
		"base_url":  interf.DebugData.BaseUrl,
		"url":       interf.DebugData.Url,
		"method":    interf.DebugData.Method,
	}

	err = r.GetDB(tenantId).Model(&model.DiagnoseInterface{}).
		Where("id=?", interf.ID).
		Updates(values).
		Error

	return
}

func (r *DiagnoseInterfaceRepo) CreateInterfaceFromCurl(tenantId consts.TenantId, interf *model.DiagnoseInterface, parent model.DiagnoseInterface) (
	err error) {

	return
}

func (r *DiagnoseInterfaceRepo) UpdateMethod(tenantId consts.TenantId, id uint, method consts.HttpMethod) (err error) {
	err = r.GetDB(tenantId).Model(&model.DiagnoseInterface{}).
		Where("id = ?", id).
		Update("method", method).Error

	return
}
