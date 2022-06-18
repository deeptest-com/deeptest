package repo

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/model"
	"github.com/kataras/iris/v12"
	"gorm.io/gorm"
)

type ScenarioProcessorRepo struct {
	DB *gorm.DB `inject:""`
}

func (r *ScenarioProcessorRepo) GetTree(scenarioId int) (root *model.TestProcessor, err error) {
	pos, err := r.ListByScenario(scenarioId)
	if err != nil {
		return
	}

	root = pos[0]
	root.Slots = iris.Map{"icon": "icon"}
	r.makeTree(pos[1:], root)

	return
}

func (r *ScenarioProcessorRepo) ListByScenario(scenarioId int) (pos []*model.TestProcessor, err error) {
	err = r.DB.
		Where("scenario_id=?", scenarioId).
		Where("NOT deleted").
		Order("parent_id ASC, ordr ASC").
		Find(&pos).Error
	return
}

func (r *ScenarioProcessorRepo) makeTree(Data []*model.TestProcessor, node *model.TestProcessor) { //参数为父节点，添加父节点的子节点指针切片
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

func (r *ScenarioProcessorRepo) haveChild(Data []*model.TestProcessor, node *model.TestProcessor) (child []*model.TestProcessor, yes bool) {
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

func (r *ScenarioProcessorRepo) CreateDefault(scenarioId uint) (po model.TestProcessor, err error) {
	po = model.TestProcessor{
		ScenarioId:     scenarioId,
		Name:           "root",
		EntityCategory: consts.ProcessorRoot,
		EntityId:       0,
	}
	err = r.DB.Create(&po).Error

	return
}

func (r *ScenarioProcessorRepo) Get(id uint) (processor model.TestProcessor, err error) {
	err = r.DB.Where("id = ?", id).First(&processor).Error
	return
}

func (r *ScenarioProcessorRepo) Save(processor *model.TestProcessor) (err error) {
	err = r.DB.Save(processor).Error

	return
}
