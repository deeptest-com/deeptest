package service

import (
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	serverConsts "github.com/aaronchen2k/deeptest/internal/server/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	repo "github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	_domain "github.com/aaronchen2k/deeptest/pkg/domain"
	"github.com/kataras/iris/v12"
)

type CategoryService struct {
	EndpointService *EndpointService   `inject:""`
	CategoryRepo    *repo.CategoryRepo `inject:""`
	EndpointRepo    *repo.EndpointRepo `inject:""`
	PlanRepo        *repo.PlanRepo     `inject:""`
	ScenarioRepo    *repo.ScenarioRepo `inject:""`
}

func (s *CategoryService) GetTree(typ serverConsts.CategoryDiscriminator, projectId int) (root *v1.Category, err error) {
	root, err = s.CategoryRepo.GetTree(typ, uint(projectId))
	root.Children = append(root.Children, &v1.Category{Id: -1, Name: "未分类", ParentId: root.Id, Slots: iris.Map{"icon": "icon"}})
	s.mountCount(root, typ, uint(projectId))
	return
}

func (s *CategoryService) Get(id int) (root model.Category, err error) {
	//root, err = s.CategoryRepo.Get(id)

	return
}

func (s *CategoryService) Create(req v1.CategoryCreateReq) (ret model.Category, err *_domain.BizErr) {
	target, _ := s.CategoryRepo.Get(req.TargetId)
	if target.ID == 0 {
		return
	}

	ret = model.Category{
		Name:      req.Name,
		ParentId:  req.TargetId,
		ProjectId: req.ProjectId,
		Type:      req.Type,
	}

	if req.Mode == "child" {
		ret.ParentId = int(target.ID)
	} else if req.Mode == "brother" {
		ret.ParentId = target.ParentId
	}

	ret.Ordr = s.CategoryRepo.GetMaxOrder(uint(ret.ParentId), req.Type, req.ProjectId)

	s.CategoryRepo.Save(&ret)

	if req.Mode == "parent" { // move interface to new folder
		target.ParentId = int(ret.ID)
		s.CategoryRepo.Save(&target)
	}

	return
}

func (s *CategoryService) Update(req v1.CategoryReq) (err error) {
	err = s.CategoryRepo.Update(req)
	return
}

func (s *CategoryService) UpdateName(req v1.CategoryReq) (err error) {
	err = s.CategoryRepo.UpdateName(req.Id, req.Name)
	return
}

func (s *CategoryService) Delete(typ serverConsts.CategoryDiscriminator, projectId, id uint) (err error) {
	err = s.deleteNodeAndChildren(typ, projectId, id)
	return
}

func (s *CategoryService) Move(srcId, targetId uint, pos serverConsts.DropPos, typ serverConsts.CategoryDiscriminator, projectId uint) (
	srcScenarioNode model.Category, err error) {
	srcScenarioNode, err = s.CategoryRepo.Get(int(srcId))

	srcScenarioNode.ParentId, srcScenarioNode.Ordr = s.CategoryRepo.UpdateOrder(pos, int(targetId), typ, projectId)
	err = s.CategoryRepo.UpdateOrdAndParent(srcScenarioNode)

	return
}

func (s *CategoryService) deleteNodeAndChildren(typ serverConsts.CategoryDiscriminator, projectId, nodeId uint) (err error) {
	//err = s.CategoryRepo.DeleteWithChildren(nodeId)
	//if err == nil {
	//	children, _ := s.CategoryRepo.GetChildren(nodeId)
	//	for _, child := range children {
	//		s.deleteNodeAndChildren(child.ID)
	//	}
	//}

	categoryIds, err := s.CategoryRepo.GetDescendantIds(nodeId, model.Category{}.TableName(), typ, int(projectId))
	if err != nil {
		return
	}
	/*
			child, err := s.CategoryRepo.GetAllChild(typ, projectId, int(nodeId))
			if err != nil {
				return
			}


		categoryIds := make([]uint, 0)
		for _, v := range child {
			categoryIds = append(categoryIds, v.ID)
		}
		categoryIds = append(categoryIds, nodeId)
	*/

	if err = s.CategoryRepo.BatchDelete(categoryIds); err != nil {
		return
	}

	switch typ {
	case serverConsts.EndpointCategory:
		err = s.EndpointService.DeleteByCategories(categoryIds)
	case serverConsts.ScenarioCategory:
		err = s.ScenarioRepo.DeleteByCategoryIds(categoryIds)
	case serverConsts.PlanCategory:
		err = s.PlanRepo.DeleteByCategoryIds(categoryIds)
	}

	return
}

func (s *CategoryService) mountCount(root *v1.Category, typ serverConsts.CategoryDiscriminator, projectId uint) {

	repo := s.getRepo(typ)

	var data []v1.CategoryCount
	err := repo.GetCategoryCount(&data, projectId)
	if err != nil {
		return
	}

	result := s.convertMap(data)

	s.mountCountOnNode(root, result)

	//TODO 遍历数据挂载数量。
}

func (s *CategoryService) getRepo(typ serverConsts.CategoryDiscriminator) repo.IRepo {

	repos := map[serverConsts.CategoryDiscriminator]repo.IRepo{
		serverConsts.EndpointCategory: s.EndpointRepo,
		serverConsts.PlanCategory:     s.PlanRepo,
		serverConsts.ScenarioCategory: s.ScenarioRepo,
	}

	return repos[typ]
}
func (s *CategoryService) convertMap(data []v1.CategoryCount) (result map[int64]int64) {
	result = make(map[int64]int64)
	for _, item := range data {
		result[item.CategoryId] = item.Count
	}
	return
}

func (s *CategoryService) mountCountOnNode(root *v1.Category, data map[int64]int64) int64 {
	root.Count = data[root.Id]
	for _, children := range root.Children {
		root.Count += s.mountCountOnNode(children, data)
	}
	return root.Count
}
