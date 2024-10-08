package service

import (
	"github.com/deeptest-com/deeptest/internal/pkg/consts"
	"github.com/deeptest-com/deeptest/internal/server/modules/model"
	"github.com/deeptest-com/deeptest/internal/server/modules/repo"
	_commonUtils "github.com/deeptest-com/deeptest/pkg/lib/comm"
)

type EndpointTagService struct {
	EndpointTagRepo *repo.EndpointTagRepo `inject:""`
}

func (s *EndpointTagService) ListTagsByProject(tenantId consts.TenantId, projectId uint) (tags []model.EndpointTagRel, err error) {
	tags, err = s.EndpointTagRepo.ListRelByProject(tenantId, projectId)
	return
}

func (s *EndpointTagService) Create(tenantId consts.TenantId, name string, projectId uint) (id uint, err error) {
	id, err = s.EndpointTagRepo.Create(tenantId, name, projectId)
	return
}

// GetTagIdsNyName 存在的tagName查出来tagId,不存在的写入数据库
func (s *EndpointTagService) GetTagIdsNyName(tenantId consts.TenantId, tagNames []string, projectId uint) (tagIds []uint, err error) {
	tags, err := s.EndpointTagRepo.BatchGetByName(tenantId, tagNames, projectId)
	if err != nil {
		return
	}

	tagNamesExisted := make([]string, 0)
	for _, tag := range tags {
		tagNamesExisted = append(tagNamesExisted, tag.Name)
		tagIds = append(tagIds, tag.ID)
	}

	tagNamesNeedInsert := _commonUtils.Difference(tagNames, tagNamesExisted)
	if len(tagNamesNeedInsert) == 0 {
		return tagIds, nil
	}

	if err = s.EndpointTagRepo.BatchCreate(tenantId, tagNamesNeedInsert, projectId); err != nil {
		return
	}

	tagIds, err = s.EndpointTagRepo.BatchGetIdsByName(tenantId, tagNames, projectId)
	return
}

func (s *EndpointTagService) BatchAddEndpointForTag(tagName string, endpointIds []uint, projectId uint) (err error) {
	existedEndpointIds, err := s.EndpointTagRepo.BatchGetEndpointIdsByTag(tagName, endpointIds, projectId)
	if err != nil {
		return
	}

	endpointNeedAdd := _commonUtils.DifferenceUint(endpointIds, existedEndpointIds)

	if len(endpointNeedAdd) == 0 {
		return
	}

	err = s.EndpointTagRepo.BatchAddRelForTag(tagName, endpointNeedAdd, projectId)

	return
}
