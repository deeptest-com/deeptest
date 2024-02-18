package service

import (
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	serverConsts "github.com/aaronchen2k/deeptest/internal/server/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	repo "github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	_domain "github.com/aaronchen2k/deeptest/pkg/domain"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"github.com/kataras/iris/v12"
)

type CategoryService struct {
	EndpointService *EndpointService   `inject:""`
	CategoryRepo    *repo.CategoryRepo `inject:""`
	EndpointRepo    *repo.EndpointRepo `inject:""`
	PlanRepo        *repo.PlanRepo     `inject:""`
	ScenarioRepo    *repo.ScenarioRepo `inject:""`
}

func (s *CategoryService) GetTree(tenantId consts.TenantId, typ serverConsts.CategoryDiscriminator, projectId int) (root *v1.Category, err error) {
	root, err = s.CategoryRepo.GetTree(tenantId, typ, uint(projectId))
	root.Children = append(root.Children, &v1.Category{Id: -1, Name: "未分类", ParentId: root.Id, Slots: iris.Map{"icon": "icon"}})
	s.mountCount(tenantId, root, typ, uint(projectId))
	return
}

func (s *CategoryService) Get(tenantId consts.TenantId, id int) (root model.Category, err error) {
	//root, err = s.CategoryRepo.Get(id)

	return
}

func (s *CategoryService) Create(tenantId consts.TenantId, req v1.CategoryCreateReq) (ret model.Category, err *_domain.BizErr) {
	target, _ := s.CategoryRepo.Get(tenantId, req.TargetId)
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

	ret.Ordr = s.CategoryRepo.GetMaxOrder(tenantId, uint(ret.ParentId), req.Type, req.ProjectId)

	s.CategoryRepo.Save(tenantId, &ret)

	if req.Mode == "parent" { // move interface to new folder
		target.ParentId = int(ret.ID)
		s.CategoryRepo.Save(tenantId, &target)
	}

	return
}

func (s *CategoryService) Update(tenantId consts.TenantId, req v1.CategoryReq) (err error) {
	err = s.CategoryRepo.Update(tenantId, req)
	return
}

func (s *CategoryService) UpdateName(tenantId consts.TenantId, req v1.CategoryReq) (err error) {
	err = s.CategoryRepo.UpdateName(tenantId, req.Id, req.Name)
	return
}

func (s *CategoryService) Delete(tenantId consts.TenantId, typ serverConsts.CategoryDiscriminator, projectId, id uint) (err error) {
	err = s.deleteNodeAndChildren(tenantId, typ, projectId, id)
	return
}

func (s *CategoryService) Move(tenantId consts.TenantId, srcId, targetId uint, pos serverConsts.DropPos, typ serverConsts.CategoryDiscriminator, projectId uint) (
	srcScenarioNode model.Category, err error) {
	srcScenarioNode, err = s.CategoryRepo.Get(tenantId, int(srcId))

	srcScenarioNode.ParentId, srcScenarioNode.Ordr = s.CategoryRepo.UpdateOrder(tenantId, pos, int(targetId), typ, projectId)
	err = s.CategoryRepo.UpdateOrdAndParent(tenantId, srcScenarioNode)

	return
}

func (s *CategoryService) deleteNodeAndChildren(tenantId consts.TenantId, typ serverConsts.CategoryDiscriminator, projectId, nodeId uint) (err error) {
	//err = s.CategoryRepo.DeleteWithChildren(nodeId)
	//if err == nil {
	//	children, _ := s.CategoryRepo.GetChildren(nodeId)
	//	for _, child := range children {
	//		s.deleteNodeAndChildren(child.ID)
	//	}
	//}

	categoryIds, err := s.CategoryRepo.GetDescendantIds(tenantId, nodeId, model.Category{}.TableName(), typ, int(projectId))
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

	if err = s.CategoryRepo.BatchDelete(tenantId, categoryIds); err != nil {
		return
	}

	switch typ {
	case serverConsts.EndpointCategory:
		err = s.EndpointService.DeleteByCategories(tenantId, categoryIds)
	case serverConsts.ScenarioCategory:
		err = s.ScenarioRepo.DeleteByCategoryIds(tenantId, categoryIds)
	case serverConsts.PlanCategory:
		err = s.PlanRepo.DeleteByCategoryIds(tenantId, categoryIds)
	}

	return
}

func (s *CategoryService) mountCount(tenantId consts.TenantId, root *v1.Category, typ serverConsts.CategoryDiscriminator, projectId uint) {

	repo := s.getRepo(tenantId, typ)

	var data []v1.CategoryCount
	err := repo.GetCategoryCount(tenantId, &data, projectId)
	if err != nil {
		return
	}

	result := s.convertMap(tenantId, data)

	s.mountCountOnNode(tenantId, root, result)

	//TODO 遍历数据挂载数量。
}

func (s *CategoryService) getRepo(tenantId consts.TenantId, typ serverConsts.CategoryDiscriminator) repo.IRepo {

	repos := map[serverConsts.CategoryDiscriminator]repo.IRepo{
		serverConsts.EndpointCategory: s.EndpointRepo,
		serverConsts.PlanCategory:     s.PlanRepo,
		serverConsts.ScenarioCategory: s.ScenarioRepo,
	}

	return repos[typ]
}
func (s *CategoryService) convertMap(tenantId consts.TenantId, data []v1.CategoryCount) (result map[int64]int64) {
	result = make(map[int64]int64)
	for _, item := range data {
		result[item.CategoryId] = item.Count
	}
	return
}

func (s *CategoryService) mountCountOnNode(tenantId consts.TenantId, root *v1.Category, data map[int64]int64) int64 {
	root.Count = data[root.Id]
	for _, children := range root.Children {
		root.Count += s.mountCountOnNode(tenantId, children, data)
	}
	return root.Count
}

func (s *CategoryService) Copy(tenantId consts.TenantId, targetId, newParentId, userId uint, username string) (err error) {

	category, err := s.CategoryRepo.CopySelf(tenantId, int(targetId), int(newParentId))
	if err != nil {
		return err
	}

	go func() {
		err = s.copyDataByCategoryId(tenantId, category.Type, targetId, category.ID, userId, username)
		if err != nil {
			logUtils.Error(err.Error())
		}
		err = s.copyChildren(tenantId, targetId, category.ID, userId, username)
		if err != nil {
			logUtils.Error(err.Error())
		}
	}()

	return
}

func (s *CategoryService) copyChildren(tenantId consts.TenantId, parentId, newParentId, userId uint, username string) (err error) {
	children, err := s.CategoryRepo.GetChildren(tenantId, parentId)
	if err != nil {
		return err
	}

	for _, child := range children {
		err = s.Copy(tenantId, child.ID, newParentId, userId, username)
		if err != nil {
			return err
		}
	}

	return err
}

func (s *CategoryService) copyDataByCategoryId(tenantId consts.TenantId, typ serverConsts.CategoryDiscriminator, targetId, categoryId, userId uint, username string) (err error) {
	switch typ {
	case serverConsts.EndpointCategory:
		err = s.EndpointService.CopyDataByCategoryId(tenantId, targetId, categoryId, userId, username)
	}

	return err
}
