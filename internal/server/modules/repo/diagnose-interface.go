package repo

import (
	serverDomain "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	serverConsts "github.com/aaronchen2k/deeptest/internal/server/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/kataras/iris/v12"
	"gorm.io/gorm"
)

type DiagnoseInterfaceRepo struct {
	*BaseRepo              `inject:""`
	*DiagnoseInterfaceRepo `inject:""`
	*DebugInterfaceRepo    `inject:""`
	DB                     *gorm.DB `inject:""`
}

func (r *DiagnoseInterfaceRepo) GetTree(projectId uint) (root *serverDomain.DiagnoseInterface, err error) {
	pos, err := r.ListByProject(projectId)
	if err != nil {
		return
	}

	tos := r.toTos(pos)
	if len(tos) == 0 {
		return
	}

	root = &serverDomain.DiagnoseInterface{}
	r.makeTree(tos, root)
	r.mountCount(root)

	return
}

func (r *DiagnoseInterfaceRepo) mountCount(root *serverDomain.DiagnoseInterface) (count int) {
	for _, child := range root.Children {
		root.Count += r.mountCount(child)
	}
	return root.Count

}

func (r *DiagnoseInterfaceRepo) ListByProject(projectId uint) (pos []*model.DiagnoseInterface, err error) {
	db := r.DB.
		Where("project_id=?", projectId).
		Where("NOT deleted")

	err = db.
		Order("parent_id ASC, ordr ASC").
		Find(&pos).Error

	return
}

func (r *DiagnoseInterfaceRepo) Get(id uint) (po model.DiagnoseInterface, err error) {
	err = r.DB.Where("id = ?", id).First(&po).Error
	return
}

func (r *DiagnoseInterfaceRepo) GetDetail(interfId uint) (diagnoseInterface model.DiagnoseInterface, err error) {
	if interfId <= 0 {
		return
	}

	diagnoseInterface, err = r.Get(interfId)

	debugInterface, _ := r.DebugInterfaceRepo.Get(diagnoseInterface.DebugInterfaceId)

	debugData, _ := r.DebugInterfaceRepo.GetDetail(debugInterface.ID)
	diagnoseInterface.DebugData = &debugData

	return
}

func (r *DiagnoseInterfaceRepo) toTos(pos []*model.DiagnoseInterface) (tos []*serverDomain.DiagnoseInterface) {
	for _, po := range pos {
		to := r.ToTo(po)

		tos = append(tos, to)
	}

	return
}
func (r *DiagnoseInterfaceRepo) ToTo(po *model.DiagnoseInterface) (to *serverDomain.DiagnoseInterface) {
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

func (r *DiagnoseInterfaceRepo) makeTree(findIn []*serverDomain.DiagnoseInterface, parent *serverDomain.DiagnoseInterface) { //参数为父节点，添加父节点的子节点指针切片
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

func (r *DiagnoseInterfaceRepo) hasChild(categories []*serverDomain.DiagnoseInterface, parent *serverDomain.DiagnoseInterface) (
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

func (r *DiagnoseInterfaceRepo) Save(po *model.DiagnoseInterface) (err error) {
	po.Ordr = r.GetMaxOrder(po.ParentId)

	err = r.DB.Save(po).Error

	return
}

func (r *DiagnoseInterfaceRepo) Update(req serverDomain.DiagnoseInterfaceReq) (err error) {
	err = r.DB.Model(&model.DiagnoseInterface{}).
		Where("id ?", req.Id).
		Updates(map[string]interface{}{"title": req.Title, "desc": req.Desc}).Error

	return
}

func (r *DiagnoseInterfaceRepo) UpdateOrder(pos serverConsts.DropPos, targetId uint, projectId uint) (
	parentId uint, ordr int) {
	if pos == serverConsts.Inner {
		parentId = targetId

		var preChild model.DiagnoseInterface
		r.DB.Where("parent_id=? AND project_id = ?", parentId, projectId).
			Order("ordr DESC").Limit(1).
			First(&preChild)

		ordr = preChild.Ordr + 1

	} else if pos == serverConsts.Before {
		brother, _ := r.Get(targetId)
		parentId = brother.ParentId

		r.DB.Model(&model.DiagnoseInterface{}).
			Where("NOT deleted AND parent_id=? AND project_id = ? AND ordr >= ?",
				parentId, projectId, brother.Ordr).
			Update("ordr", gorm.Expr("ordr + 1"))

		ordr = brother.Ordr

	} else if pos == serverConsts.After {
		brother, _ := r.Get(targetId)
		parentId = brother.ParentId

		r.DB.Model(&model.DiagnoseInterface{}).
			Where("NOT deleted AND parent_id=? AND project_id = ? AND ordr > ?",
				parentId, parentId, brother.Ordr).
			Update("ordr", gorm.Expr("ordr + 1"))

		ordr = brother.Ordr + 1

	}

	return
}

func (r *DiagnoseInterfaceRepo) UpdateName(id int, name string) (err error) {
	err = r.DB.Model(&model.DiagnoseInterface{}).
		Where("id = ?", id).
		Update("name", name).Error

	return
}

func (r *DiagnoseInterfaceRepo) Delete(id uint) (err error) {
	err = r.DB.Model(&model.DiagnoseInterface{}).
		Where("id=?", id).
		Update("deleted", true).
		Error

	return
}

func (r *DiagnoseInterfaceRepo) GetChildren(nodeId uint) (children []*model.DiagnoseInterface, err error) {
	err = r.DB.Where("parent_id=?", nodeId).Find(&children).Error
	return
}

func (r *DiagnoseInterfaceRepo) UpdateOrdAndParent(node model.DiagnoseInterface) (err error) {
	err = r.DB.Model(&node).
		Updates(model.DiagnoseInterface{Ordr: node.Ordr, ParentId: node.ParentId}).
		Error

	return
}

func (r *DiagnoseInterfaceRepo) GetMaxOrder(parentId uint) (order int) {
	node := model.DiagnoseInterface{}

	err := r.DB.Model(&model.DiagnoseInterface{}).
		Where("parent_id=?", parentId).
		Order("ordr DESC").
		First(&node).Error

	if err == nil {
		order = node.Ordr + 1
	}

	return
}

func (r *DiagnoseInterfaceRepo) Remove(id uint, typ serverConsts.DiagnoseInterfaceType) (err error) {
	ids := []uint{}

	if typ == serverConsts.DiagnoseInterfaceTypeInterface {
		ids = append(ids, id)
	} else {
		ids, _ = r.GetAllChildIdsSimple(id, model.DiagnoseInterface{}.TableName())
	}

	err = r.DB.Model(&model.DiagnoseInterface{}).
		Where("id IN (?)", ids).
		Update("deleted", true).Error

	return
}

func (r *DiagnoseInterfaceRepo) SaveDebugData(interf *model.DiagnoseInterface) (err error) {
	r.DB.Transaction(func(tx *gorm.DB) error {
		err = r.UpdateDebugInfo(interf)
		if err != nil {
			return err
		}

		// TODO: save debug data

		return err
	})

	return
}

func (r *DiagnoseInterfaceRepo) UpdateDebugInfo(interf *model.DiagnoseInterface) (err error) {
	values := map[string]interface{}{
		"server_id": interf.DebugData.ServerId,
		"base_url":  interf.DebugData.BaseUrl,
		"url":       interf.DebugData.Url,
		"method":    interf.DebugData.Method,
	}

	err = r.DB.Model(&model.DiagnoseInterface{}).
		Where("id=?", interf.ID).
		Updates(values).
		Error

	return
}

func (r *DiagnoseInterfaceRepo) CreateInterfaceFromCurl(interf *model.DiagnoseInterface, parent model.DiagnoseInterface) (
	err error) {

	return
}

func (r *DiagnoseInterfaceRepo) UpdateMethod(id uint, method consts.HttpMethod) (err error) {
	err = r.DB.Model(&model.DiagnoseInterface{}).
		Where("id = ?", id).
		Update("method", method).Error

	return
}
