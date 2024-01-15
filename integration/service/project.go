package service

import (
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	integrationDomain "github.com/aaronchen2k/deeptest/integration/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
)

type ProjectService struct {
	RemoteService   *RemoteService        `inject:""`
	IntegrationRepo *repo.IntegrationRepo `inject:""`
	ProjectRepo     *repo.ProjectRepo     `inject:""`
}

func (s *ProjectService) GetUserProductList(page, pageSize int, username string) (ret []integrationDomain.ProductItem, err error) {
	return s.RemoteService.GetUserProductList(page, pageSize, username)
}

func (s *ProjectService) GetSpacesByUsername(username string) (ret []integrationDomain.SpaceItem, err error) {
	return s.RemoteService.GetSpacesByUsername(username)
}

func (s *ProjectService) Create(req integrationDomain.ProjectReq, projectId uint) (err error) {
	if err = s.SaveProjectRelatedProducts(projectId, req.Products); err != nil {
		return
	}

	if err = s.SaveProjectRelatedSpaces(projectId, req.Spaces); err != nil {
		return
	}

	err = s.SyncSpaceMembers(projectId, req.Spaces)

	return
}

func (s *ProjectService) SaveProjectRelatedProducts(projectId uint, products []uint) (err error) {
	relations := make([]model.ProjectProductRel, 0)
	for _, product := range products {
		relations = append(relations, model.ProjectProductRel{
			ProjectId: projectId,
			ProductId: product,
		})
	}

	if len(relations) > 0 {
		err = s.IntegrationRepo.BatchCreateProjectProductRel(relations)
	}

	return
}

func (s *ProjectService) SaveProjectRelatedSpaces(projectId uint, spaces []string) (err error) {
	relations := make([]model.ProjectSpaceRel, 0)
	for _, space := range spaces {
		relations = append(relations, model.ProjectSpaceRel{
			ProjectId: projectId,
			SpaceCode: space,
		})
	}

	if len(relations) > 0 {
		err = s.IntegrationRepo.BatchCreateProjectSpaceRel(relations)
	}

	return
}

// SyncSpaceMembers TODO 存疑
func (s *ProjectService) SyncSpaceMembers(projectId uint, spaces []string) (err error) {

	return
}

func (s *ProjectService) Save(req integrationDomain.ProjectReq, projectId uint) (err error) {
	if err = s.IntegrationRepo.DeleteProductByProject(projectId); err != nil {
		return
	}

	if err = s.SaveProjectRelatedProducts(projectId, req.Products); err != nil {
		return
	}

	if err = s.IntegrationRepo.DeleteSpaceByProject(projectId); err != nil {
		return
	}

	if err = s.SaveProjectRelatedSpaces(projectId, req.Spaces); err != nil {
		return
	}

	err = s.SyncSpaceMembers(projectId, req.Spaces)

	return
}

func (s *ProjectService) GetProductsByProject(projectId uint) (res []integrationDomain.ProductBaseItem, err error) {
	productIds, err := s.IntegrationRepo.GetProductsByProject(projectId)
	if err != nil {
		return
	}

	res, err = s.RemoteService.GetProductListById(productIds)

	return
}

func (s *ProjectService) GetSpacesByProject(projectId uint) (res []integrationDomain.SpaceItem, err error) {
	spaceCodes, err := s.IntegrationRepo.GetSpacesByProject(projectId)
	if err != nil {
		return
	}

	res, err = s.RemoteService.BatchGetSpacesByCode(spaceCodes)

	return
}

func (s *ProjectService) Detail(projectId uint) (res integrationDomain.ProjectDetail, err error) {
	products, err := s.GetProductsByProject(projectId)
	if err != nil {
		return
	}

	spaces, err := s.GetSpacesByProject(projectId)
	if err != nil {
		return
	}

	res.Products = products
	res.Spaces = spaces

	return
}

func (s *ProjectService) GetListWithRoleBySpace(spaceCode string) (res []v1.ProjectListWithRole, err error) {
	return s.IntegrationRepo.GetProjectListWithRoleBySpace(spaceCode)
}

func (s *ProjectService) SaveSpaceRelatedProjects(spaceCode string, shortNames []string) (err error) {
	err = s.IntegrationRepo.DeleteBySpaceCode(spaceCode)
	if err != nil {
		return
	}

	projectShortNameIdMap, err := s.GetProjectShortNameAndIdMap(shortNames)
	if err != nil {
		return
	}

	relations := make([]model.ProjectSpaceRel, 0)
	for _, shortName := range shortNames {
		relTmp := model.ProjectSpaceRel{
			SpaceCode: spaceCode,
		}
		if projectId, ok := projectShortNameIdMap[shortName]; ok {
			relTmp.ProjectId = projectId
		}

		relations = append(relations, relTmp)
	}

	err = s.IntegrationRepo.BatchCreateProjectSpaceRel(relations)
	return
}

func (s *ProjectService) GetProjectShortNameAndIdMap(shortNames []string) (res map[string]uint, err error) {
	projects, err := s.ProjectRepo.BatchGetByShortNames(shortNames)
	if err != nil {
		return
	}

	res = make(map[string]uint)
	for _, project := range projects {
		res[project.ShortName] = project.ID
	}

	return
}
