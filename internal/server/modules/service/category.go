package service

import (
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	serverConsts "github.com/aaronchen2k/deeptest/internal/server/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	repo "github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	_domain "github.com/aaronchen2k/deeptest/pkg/domain"

	_commUtils "github.com/aaronchen2k/deeptest/pkg/lib/comm"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"github.com/kataras/iris/v12"
	"gorm.io/gorm"
	"strings"
)

type CategoryService struct {
	EndpointService     *EndpointService          `inject:""`
	ServeService        *ServeService             `inject:""`
	CategoryRepo        *repo.CategoryRepo        `inject:""`
	EndpointRepo        *repo.EndpointRepo        `inject:""`
	PlanRepo            *repo.PlanRepo            `inject:""`
	ScenarioRepo        *repo.ScenarioRepo        `inject:""`
	ComponentSchemaRepo *repo.ComponentSchemaRepo `inject:""`
	ProjectRepo         *repo.ProjectRepo         `inject:""`
	ServeRepo           *repo.ServeRepo           `inject:""`
}

func (s *CategoryService) GetTree(typ serverConsts.CategoryDiscriminator, projectId int) (root *v1.Category, err error) {
	root, err = s.CategoryRepo.GetTree(typ, uint(projectId))
	if typ != serverConsts.SchemaCategory {
		root.Children = append(root.Children, &v1.Category{Id: -1, Name: "未分类", ParentId: root.Id, Slots: iris.Map{"icon": "icon"}})
	}
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
		EntityId:  req.EntityId,
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

	if req.IsEntity {
		repo := s.getRepo(req.Type)
		repo.SaveEntity(&ret)
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
	if err != nil {
		return
	}

	if typ == serverConsts.SchemaCategory && srcScenarioNode.EntityId != 0 {
		s.ComponentSchemaRepo.ChangeRef(srcScenarioNode.EntityId, srcScenarioNode.ID)
	}
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
	case serverConsts.SchemaCategory:
		entityIds, err := s.CategoryRepo.GetEntityIdsByIds(categoryIds)
		if err != nil {
			return err
		}

		err = s.ComponentSchemaRepo.DeleteByIds(entityIds)
	}

	return
}

func (s *CategoryService) mountCount(root *v1.Category, typ serverConsts.CategoryDiscriminator, projectId uint) {

	repo := s.getRepo(typ)

	var data []v1.CategoryCount
	err := repo.GetCategoryCount(&data, projectId)
	if err != nil || len(data) == 0 {
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
		serverConsts.SchemaCategory:   s.ComponentSchemaRepo,
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

func (s *CategoryService) GetJoinedPath(typ serverConsts.CategoryDiscriminator, projectId, categoryId uint) (path string, err error) {
	categories, err := s.CategoryRepo.ListByProject(typ, projectId)
	categoryIdParentMap := make(map[uint]uint)
	categoryIdNameMap := make(map[uint]string)
	for _, v := range categories {
		categoryIdParentMap[v.ID] = uint(v.ParentId)
		categoryIdNameMap[v.ID] = v.Name
	}

	if name, ok := categoryIdNameMap[categoryId]; ok {
		path = "/" + name
	}

	s.doGetJoinedPath(categoryIdParentMap, categoryIdNameMap, categoryId, &path)

	index := strings.Index(path, "/")

	if index != -1 {
		path = path[index:]
	}
	return
}

func (s *CategoryService) doGetJoinedPath(categoryIdParentMap map[uint]uint, categoryIdNameMap map[uint]string, categoryId uint, path *string) {
	if parentId, ok := categoryIdParentMap[categoryId]; ok {
		if parentName, ok1 := categoryIdNameMap[parentId]; ok1 {
			*path = "/" + parentName

			s.doGetJoinedPath(categoryIdParentMap, categoryIdNameMap, parentId, path)
		}
	}

	return
}

func (s *CategoryService) BatchAddSchemaRoot1(projectIds []uint) (err error) {
	if len(projectIds) == 0 {
		projects, err := s.ProjectRepo.ListAll()
		if err != nil {
			return err
		}

		for _, v := range projects {
			projectIds = append(projectIds, v.ID)
		}
	}

	projectServeMap := make(map[uint][]uint)
	serves, err := s.ServeRepo.ListByProjects(projectIds)
	if err != nil {
		return
	}

	for _, v := range serves {
		projectServeMap[v.ProjectId] = append(projectServeMap[v.ProjectId], v.ID)
	}

	var serveIds []uint

	for _, projectId := range projectIds {
		category, err := s.CategoryRepo.GetByItem(0, serverConsts.SchemaCategory, projectId, "分类")
		if err != gorm.ErrRecordNotFound {
			continue
		}

		err = s.ProjectRepo.AddProjectRootSchemaCategory(projectId)
		if err != nil {
		}

		category, _ = s.CategoryRepo.GetByItem(0, serverConsts.SchemaCategory, projectId, "分类")

		rootId := category.ID

		//开始处理历史schema数据
		if v, ok := projectServeMap[projectId]; ok {
			serveIds = v
		}

		if len(serveIds) > 0 {
			//schema表给project_id赋值
			err = s.ServeRepo.BatchUpdateSchemaProjectByServeId(serveIds, projectId)
			if err != nil {
				continue
			}

			schemas, err := s.ServeRepo.GetSchemas(serveIds)
			if err != nil {
				continue
			}

			for _, schema := range schemas {
				//先查后创建，避免重复增加分类数据
				_, err = s.CategoryRepo.GetByEntityId(schema.ID, serverConsts.SchemaCategory)
				if err != gorm.ErrRecordNotFound {
					continue
				}

				createCategoryReq := v1.CategoryCreateReq{Name: schema.Name, TargetId: int(rootId), ProjectId: projectId, Type: serverConsts.SchemaCategory, Mode: "child", EntityId: schema.ID}
				_, _ = s.Create(createCategoryReq)
			}
		}

	}

	return

}

func (s *CategoryService) BatchAddSchemaRoot(projectIds []uint) (err error) {
	if len(projectIds) == 0 {
		projects, err := s.ProjectRepo.ListAll()
		if err != nil {
			return err
		}

		for _, v := range projects {
			projectIds = append(projectIds, v.ID)
		}
	}

	//已经有根结点的项目
	existedRootProjects, err := s.CategoryRepo.BatchGetRootNodeProjectIds(projectIds, serverConsts.SchemaCategory)
	if err != nil {
		return
	}

	//需要创建根结点的项目
	needCreateRootProjects := _commUtils.DifferenceUint(projectIds, existedRootProjects)
	if len(needCreateRootProjects) > 0 {
		if err = s.CategoryRepo.BatchAddProjectRootSchemaCategory(needCreateRootProjects); err != nil {
			return err
		}
	}

	rootNodes, err := s.CategoryRepo.BatchGetRootNodes(projectIds, serverConsts.SchemaCategory)
	if err != nil {
		return
	}

	projectRootNodesMap := make(map[uint]model.Category)
	for _, v := range rootNodes {
		projectRootNodesMap[v.ProjectId] = v
	}

	projectServeMap := make(map[uint][]uint)
	serves, err := s.ServeRepo.ListByProjects(projectIds)
	if err != nil {
		return
	}

	for _, v := range serves {
		projectServeMap[v.ProjectId] = append(projectServeMap[v.ProjectId], v.ID)
	}

	var serveIds []uint
	for _, projectId := range projectIds {
		//开始处理历史schema数据
		if v, ok := projectServeMap[projectId]; ok {
			serveIds = v
		}

		if len(serveIds) > 0 {
			//schema表给project_id赋值
			err = s.ServeRepo.BatchUpdateSchemaProjectByServeId(serveIds, projectId)
			if err != nil {
				continue
			}
		}

	}

	schemas, err := s.ComponentSchemaRepo.GetSchemasNotExistedInCategory(projectIds)
	if err != nil {
		return
	}

	var rootId uint
	for _, schema := range schemas {
		if v, ok := projectRootNodesMap[schema.ProjectId]; ok {
			rootId = v.ID
		}
		createCategoryReq := v1.CategoryCreateReq{Name: schema.Name, TargetId: int(rootId), ProjectId: schema.ProjectId, Type: serverConsts.SchemaCategory, Mode: "child", EntityId: schema.ID}
		_, _ = s.Create(createCategoryReq)
	}
	return

}

func (s *CategoryService) Copy(targetId, newParentId, userId uint, username string) (err error) {

	category, err := s.CategoryRepo.CopySelf(int(targetId), int(newParentId))
	if err != nil {
		return err
	}

	go func() {
		err = s.copyDataByCategoryId(category.Type, targetId, category, userId, username)
		if err != nil {
			logUtils.Error(err.Error())
		}
		err = s.copyChildren(targetId, category.ID, userId, username)
		if err != nil {
			logUtils.Error(err.Error())
		}
	}()

	return
}

func (s *CategoryService) copyChildren(parentId, newParentId, userId uint, username string) (err error) {
	children, err := s.CategoryRepo.GetChildren(parentId)
	if err != nil {
		return err
	}

	for _, child := range children {
		err = s.Copy(child.ID, newParentId, userId, username)
		if err != nil {
			return err
		}
	}

	return err
}

func (s *CategoryService) copyDataByCategoryId(typ serverConsts.CategoryDiscriminator, targetId uint, category model.Category, userId uint, username string) (err error) {
	var entityId uint
	switch typ {
	case serverConsts.EndpointCategory:
		err = s.EndpointService.CopyDataByCategoryId(targetId, category.ID, userId, username)
	case serverConsts.SchemaCategory:
		entityId, err = s.ServeService.CopySchemaOther(category.EntityId)
	}

	//更新实体信息
	if entityId > 0 {
		s.CategoryRepo.UpdateEntityId(category.ID, entityId)
	}

	return err

}
