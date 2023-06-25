package repo

import (
	serverDomain "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	serverConsts "github.com/aaronchen2k/deeptest/internal/server/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/kataras/iris/v12"
	"gorm.io/gorm"
)

type TestInterfaceRepo struct {
	*BaseRepo           `inject:""`
	*TestInterfaceRepo  `inject:""`
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

func (r *TestInterfaceRepo) GetDetail(interfId uint) (testInterface model.TestInterface, err error) {
	if interfId <= 0 {
		return
	}

	testInterface, err = r.Get(interfId)

	debugInterface, _ := r.DebugInterfaceRepo.Get(testInterface.DebugInterfaceId)

	testInterface.DebugData, _ = r.DebugInterfaceRepo.GetDetail(debugInterface.ID)

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
	po.Ordr = r.GetMaxOrder(po.ParentId)

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

func (r *TestInterfaceRepo) GetMaxOrder(parentId uint) (order int) {
	node := model.TestInterface{}

	err := r.DB.Model(&model.TestInterface{}).
		Where("parent_id=?", parentId).
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
		err = r.UpdateDebugInfo(interf)
		if err != nil {
			return err
		}

		// TODO: save debug data

		return err
	})

	return
}

func (r *TestInterfaceRepo) UpdateDebugInfo(interf *model.TestInterface) (err error) {
	values := map[string]interface{}{
		"server_id": interf.DebugData.ServerId,
		"base_url":  interf.DebugData.BaseUrl,
		"url":       interf.DebugData.Url,
		"method":    interf.DebugData.Method,
	}

	err = r.DB.Model(&model.TestInterface{}).
		Where("id=?", interf.ID).
		Updates(values).
		Error

	return
}
