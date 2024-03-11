package service

import (
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	serverConsts "github.com/aaronchen2k/deeptest/internal/server/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	repo "github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	_domain "github.com/aaronchen2k/deeptest/pkg/domain"
	_commUtils "github.com/aaronchen2k/deeptest/pkg/lib/comm"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"github.com/jinzhu/copier"
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

func (s *CategoryService) GetTree(tenantId consts.TenantId, typ serverConsts.CategoryDiscriminator, projectId int) (root *v1.Category, err error) {
	root, err = s.CategoryRepo.GetTree(tenantId, typ, uint(projectId))
	if typ != serverConsts.SchemaCategory {
		root.Children = append(root.Children, &v1.Category{Id: -1, Name: "未分类", ParentId: root.Id, Slots: iris.Map{"icon": "icon"}})
	}
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
		EntityId:  req.EntityId,
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

	if req.IsEntity {
		repo := s.getRepo(tenantId, req.Type)
		repo.SaveEntity(tenantId, &ret)
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
	if err != nil {
		return
	}

	if typ == serverConsts.SchemaCategory && srcScenarioNode.EntityId != 0 {
		s.ComponentSchemaRepo.ChangeRef(tenantId, srcScenarioNode.EntityId, srcScenarioNode.ID)
	}
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

	case serverConsts.SchemaCategory:
		entityIds, err := s.CategoryRepo.GetEntityIdsByIds(tenantId, categoryIds)
		if err != nil {
			return err
		}

		err = s.ComponentSchemaRepo.DeleteByIds(tenantId, entityIds)

	}

	return
}

func (s *CategoryService) mountCount(tenantId consts.TenantId, root *v1.Category, typ serverConsts.CategoryDiscriminator, projectId uint) {

	repo := s.getRepo(tenantId, typ)

	var data []v1.CategoryCount

	err := repo.GetCategoryCount(tenantId, &data, projectId)
	if err != nil || len(data) == 0 {
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
		serverConsts.SchemaCategory:   s.ComponentSchemaRepo,
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

func (s *CategoryService) GetJoinedPath(tenantId consts.TenantId, typ serverConsts.CategoryDiscriminator, projectId, categoryId uint) (path string, err error) {
	categories, err := s.CategoryRepo.ListByProject(tenantId, typ, projectId)
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

func (s *CategoryService) BatchAddSchemaRoot1(tenantId consts.TenantId, projectIds []uint) (err error) {
	if len(projectIds) == 0 {
		projects, err := s.ProjectRepo.ListAll(tenantId)
		if err != nil {
			return err
		}

		for _, v := range projects {
			projectIds = append(projectIds, v.ID)
		}
	}

	projectServeMap := make(map[uint][]uint)
	serves, err := s.ServeRepo.ListByProjects(tenantId, projectIds)
	if err != nil {
		return
	}

	for _, v := range serves {
		projectServeMap[v.ProjectId] = append(projectServeMap[v.ProjectId], v.ID)
	}

	var serveIds []uint

	for _, projectId := range projectIds {
		category, err := s.CategoryRepo.GetByItem(tenantId, 0, serverConsts.SchemaCategory, projectId, "分类")
		if err != gorm.ErrRecordNotFound {
			continue
		}

		err = s.ProjectRepo.AddProjectRootSchemaCategory(tenantId, projectId)
		if err != nil {
		}

		category, _ = s.CategoryRepo.GetByItem(tenantId, 0, serverConsts.SchemaCategory, projectId, "分类")

		rootId := category.ID

		//开始处理历史schema数据
		if v, ok := projectServeMap[projectId]; ok {
			serveIds = v
		}

		if len(serveIds) > 0 {
			//schema表给project_id赋值
			err = s.ServeRepo.BatchUpdateSchemaProjectByServeId(tenantId, serveIds, projectId)
			if err != nil {
				continue
			}

			schemas, err := s.ServeRepo.GetSchemas(tenantId, serveIds)
			if err != nil {
				continue
			}

			for _, schema := range schemas {
				//先查后创建，避免重复增加分类数据
				_, err = s.CategoryRepo.GetByEntityId(tenantId, schema.ID, serverConsts.SchemaCategory)
				if err != gorm.ErrRecordNotFound {
					continue
				}

				createCategoryReq := v1.CategoryCreateReq{Name: schema.Name, TargetId: int(rootId), ProjectId: projectId, Type: serverConsts.SchemaCategory, Mode: "child", EntityId: schema.ID}
				_, _ = s.Create(tenantId, createCategoryReq)
			}
		}

	}

	return

}

func (s *CategoryService) BatchAddSchemaRoot(tenantId consts.TenantId, projectIds []uint) (err error) {
	if len(projectIds) == 0 {
		projects, err := s.ProjectRepo.ListAll(tenantId)
		if err != nil {
			return err
		}

		for _, v := range projects {
			projectIds = append(projectIds, v.ID)
		}
	}

	//已经有根结点的项目
	existedRootProjects, err := s.CategoryRepo.BatchGetRootNodeProjectIds(tenantId, projectIds, serverConsts.SchemaCategory)
	if err != nil {
		return
	}

	//需要创建根结点的项目
	needCreateRootProjects := _commUtils.DifferenceUint(projectIds, existedRootProjects)
	if len(needCreateRootProjects) > 0 {
		if err = s.CategoryRepo.BatchAddProjectRootSchemaCategory(tenantId, needCreateRootProjects); err != nil {
			return err
		}
	}

	rootNodes, err := s.CategoryRepo.BatchGetRootNodes(tenantId, projectIds, serverConsts.SchemaCategory)
	if err != nil {
		return
	}

	projectRootNodesMap := make(map[uint]model.Category)
	for _, v := range rootNodes {
		projectRootNodesMap[v.ProjectId] = v
	}

	projectServeMap := make(map[uint][]uint)
	serves, err := s.ServeRepo.ListByProjects(tenantId, projectIds)
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
			err = s.ServeRepo.BatchUpdateSchemaProjectByServeId(tenantId, serveIds, projectId)
			if err != nil {
				continue
			}
		}

	}

	schemas, err := s.ComponentSchemaRepo.GetSchemasNotExistedInCategory(tenantId, projectIds)
	if err != nil {
		return
	}

	var rootId uint
	for _, schema := range schemas {
		if v, ok := projectRootNodesMap[schema.ProjectId]; ok {
			rootId = v.ID
		}
		createCategoryReq := v1.CategoryCreateReq{Name: schema.Name, TargetId: int(rootId), ProjectId: schema.ProjectId, Type: serverConsts.SchemaCategory, Mode: "child", EntityId: schema.ID}
		_, _ = s.Create(tenantId, createCategoryReq)
	}

	logUtils.Infof("batchAddSchemaRootEnd, projectIds:%+v", projectIds)
	return

}

func (s *CategoryService) Copy(tenantId consts.TenantId, targetId, newParentId, userId uint, username string) (category model.Category, err error) {

	category, err = s.CategoryRepo.CopySelf(tenantId, int(targetId), int(newParentId))
	if err != nil {
		return
	}

	entityId, err := s.copyDataByCategoryId(tenantId, category.Type, targetId, category, userId, username)
	if err != nil {
		return
	}
	category.EntityId = entityId

	go func() {
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

		_, err = s.Copy(tenantId, child.ID, newParentId, userId, username)

		if err != nil {
			return err
		}
	}

	return err
}

func (s *CategoryService) copyDataByCategoryId(tenantId consts.TenantId, typ serverConsts.CategoryDiscriminator, targetId uint, category model.Category, userId uint, username string) (entityId uint, err error) {
	switch typ {
	case serverConsts.EndpointCategory:
		err = s.EndpointService.CopyDataByCategoryId(tenantId, targetId, category.ID, userId, username)
	case serverConsts.SchemaCategory:
		entityId, err = s.ServeService.CopySchemaOther(tenantId, category.EntityId)
	}

	//更新实体信息
	if entityId > 0 {
		s.CategoryRepo.UpdateEntityId(tenantId, category.ID, entityId)
	}

	return
}

func (s *CategoryService) GetChildrenNodes(tenantId consts.TenantId, categoryId int) (ret []v1.Category, err error) {
	nodes, err := s.CategoryRepo.GetChildrenNodes(tenantId, uint(categoryId))

	for _, node := range nodes {
		var category v1.Category
		copier.CopyWithOption(&category, node, copier.Option{IgnoreEmpty: true, DeepCopy: true})
		category.Count = s.CategoryRepo.GetEntityCountByCategoryId(tenantId, uint(categoryId))
		ret = append(ret, category)
	}

	return
}
