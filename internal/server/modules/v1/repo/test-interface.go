package repo

import (
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/model"
	"gorm.io/gorm"
)

type InterfaceRepo struct {
	DB *gorm.DB `inject:""`
}

func (r *InterfaceRepo) GetInterfaceTree(projectId int) (root *model.TestInterface, err error) {
	pos, err := r.ListByProject(projectId)

	if err != nil {
		return
	}

	root = &model.TestInterface{Name: "root", IsDir: true}
	root.ID = 0

	r.makeTree(pos, root)
	return
}

func (r *InterfaceRepo) CreateTreeNode(defId, targetId uint, name string, mode string) (field *model.TestInterface, err error) {
	field = &model.TestInterface{}
	field.Name = name
	if mode == "root" {
	} else {
		var target model.TestInterface

		err = r.DB.Where("id=?", targetId).First(&target).Error

		if mode == "child" {
			field.ParentId = target.ID
		} else {
			field.ParentId = target.ParentId
		}
		field.Ordr = r.GetMaxOrder(field.ParentId)
	}

	err = r.DB.Save(field).Error
	return
}

func (r *InterfaceRepo) GetMaxOrder(parentId uint) (ord int) {
	var preChild model.TestInterface
	err := r.DB.
		Where("parent_id=?", parentId).
		Order("ord DESC").Limit(1).
		First(&preChild).Error

	if err != nil {
		ord = 1
	}
	ord = preChild.Ordr + 1

	return
}

func (r *InterfaceRepo) ListByProject(projectId int) (pos []*model.TestInterface, err error) {
	err = r.DB.
		Where("project_id=?", projectId).
		Where("NOT deleted").
		Order("parent_id ASC, ordr ASC").
		Find(&pos).Error
	return
}

func (r *InterfaceRepo) Get(fieldId uint) (field model.TestInterface, err error) {
	err = r.DB.
		Where("id=?", fieldId).
		Where("NOT deleted").
		First(&field).Error
	return
}

func (r *InterfaceRepo) Save(field *model.TestInterface) (err error) {
	err = r.DB.Save(field).Error
	return
}
func (r *InterfaceRepo) UpdateRange(rang string, id uint) (err error) {
	err = r.DB.Model(&model.TestInterface{}).Where("id=?", id).Update("range", rang).Error

	return
}

func (r *InterfaceRepo) makeTree(Data []*model.TestInterface, node *model.TestInterface) { //参数为父节点，添加父节点的子节点指针切片
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

func (r *InterfaceRepo) haveChild(Data []*model.TestInterface, node *model.TestInterface) (child []*model.TestInterface, yes bool) {
	for _, v := range Data {
		if v.ParentId == node.ID {
			child = append(child, v)
		}
	}
	if child != nil {
		yes = true
	}
	return
}

func (r *InterfaceRepo) Delete(id uint) (err error) {
	err = r.DB.Model(&model.TestInterface{}).
		Where("id=?", id).
		Update("deleted", true).
		Error

	//field := model.TestInterface{}
	//field.ID = id
	//err = r.DB.Delete(field).Error

	return
}

func (r *InterfaceRepo) GetChildren(defId, fieldId uint) (children []*model.TestInterface, err error) {
	err = r.DB.Where("defID=? AND parentID=?", defId, fieldId).Find(&children).Error
	return
}

func (r *InterfaceRepo) SetIsRange(fieldId uint, b bool) (err error) {
	err = r.DB.Model(&model.TestInterface{}).
		Where("id = ?", fieldId).Update("isRange", b).Error

	return
}

func (r *InterfaceRepo) AddOrderForTargetAndNextCases(srcID uint, targetOrder int, targetParentID uint) (err error) {
	sql := fmt.Sprintf(`update %s set ord = ord + 1 where ord >= %d and parentID = %d and id!=%d`,
		(&model.TestInterface{}).TableName(), targetOrder, targetParentID, srcID)
	err = r.DB.Exec(sql).Error

	return
}

func (r *InterfaceRepo) AddOrderForNextCases(srcID uint, targetOrder int, targetParentID uint) (err error) {
	sql := fmt.Sprintf(`update %s set ord = ord + 1 where ord > %d and parentID = %d and id!=%d`,
		(&model.TestInterface{}).TableName(), targetOrder, targetParentID, srcID)
	err = r.DB.Exec(sql).Error

	return
}

func (r *InterfaceRepo) UpdateOrdAndParent(field model.TestInterface) (err error) {
	//err = r.DB.Model(&field).UpdateColumn(model.TestInterface{Ordr: field.Ordr, ParentId: field.ParentId}).Error

	return
}

func NewInterfaceRepo(db *gorm.DB) *InterfaceRepo {
	return &InterfaceRepo{DB: db}
}
