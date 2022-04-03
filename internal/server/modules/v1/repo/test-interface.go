package repo

import (
	"github.com/aaronchen2k/deeptest/internal/server/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/model"
	"github.com/kataras/iris/v12"
	"gorm.io/gorm"
)

type TestInterfaceRepo struct {
	DB *gorm.DB `inject:""`
}

func NewInterfaceRepo(db *gorm.DB) *TestInterfaceRepo {
	return &TestInterfaceRepo{DB: db}
}

func (r *TestInterfaceRepo) GetInterfaceTree(projectId int) (root *model.TestInterface, err error) {
	pos, err := r.ListByProject(projectId)

	if err != nil {
		return
	}

	root = pos[0]
	root.Slots = iris.Map{"icon": "icon"}
	r.makeTree(pos[1:], root)
	return
}

func (r *TestInterfaceRepo) UpdateOrder(pos serverConsts.DropPos, targetId uint) (parentId uint, ordr int) {
	if pos == serverConsts.Inner {
		parentId = targetId

		var preChild model.TestInterface
		r.DB.Where("parent_id=?", parentId).
			Order("ordr DESC").Limit(1).
			First(&preChild)

		ordr = preChild.Ordr + 1

	} else if pos == serverConsts.Before {
		brother, _ := r.Get(targetId)
		parentId = brother.ParentId

		r.DB.Model(&model.TestInterface{}).
			Where("NOT deleted AND parent_id=? AND ordr >= ?", parentId, brother.Ordr).
			Update("ordr", gorm.Expr("ordr + 1"))

		ordr = brother.Ordr

	} else if pos == serverConsts.After {
		brother, _ := r.Get(targetId)
		parentId = brother.ParentId

		r.DB.Model(&model.TestInterface{}).
			Where("NOT deleted AND parent_id=? AND ordr > ?", parentId, brother.Ordr).
			Update("ordr", gorm.Expr("ordr + 1"))

		ordr = brother.Ordr + 1

	}

	return
}

func (r *TestInterfaceRepo) ListByProject(projectId int) (pos []*model.TestInterface, err error) {
	err = r.DB.
		Where("project_id=?", projectId).
		Where("NOT deleted").
		Order("parent_id ASC, ordr ASC").
		Find(&pos).Error
	return
}

func (r *TestInterfaceRepo) Get(fieldId uint) (field model.TestInterface, err error) {
	err = r.DB.
		Where("id=?", fieldId).
		Where("NOT deleted").
		First(&field).Error
	return
}

func (r *TestInterfaceRepo) Save(field *model.TestInterface) (err error) {
	err = r.DB.Save(field).Error
	return
}
func (r *TestInterfaceRepo) UpdateRange(rang string, id uint) (err error) {
	err = r.DB.Model(&model.TestInterface{}).Where("id=?", id).Update("range", rang).Error

	return
}

func (r *TestInterfaceRepo) makeTree(Data []*model.TestInterface, node *model.TestInterface) { //参数为父节点，添加父节点的子节点指针切片
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

func (r *TestInterfaceRepo) haveChild(Data []*model.TestInterface, node *model.TestInterface) (child []*model.TestInterface, yes bool) {
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

func (r *TestInterfaceRepo) Delete(id uint) (err error) {
	err = r.DB.Model(&model.TestInterface{}).
		Where("id=?", id).
		Update("deleted", true).
		Error

	//field := model.TestInterface{}
	//field.ID = id
	//err = r.DB.Delete(field).Error

	return
}

func (r *TestInterfaceRepo) GetChildren(defId, fieldId uint) (children []*model.TestInterface, err error) {
	err = r.DB.Where("defID=? AND parentID=?", defId, fieldId).Find(&children).Error
	return
}

func (r *TestInterfaceRepo) SetIsRange(fieldId uint, b bool) (err error) {
	err = r.DB.Model(&model.TestInterface{}).
		Where("id = ?", fieldId).Update("isRange", b).Error

	return
}

func (r *TestInterfaceRepo) UpdateOrdAndParent(interf model.TestInterface) (err error) {
	err = r.DB.Model(&interf).
		Updates(model.TestInterface{Ordr: interf.Ordr, ParentId: interf.ParentId}).
		Error

	return
}
