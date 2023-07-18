package service

import (
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	_commonUtils "github.com/aaronchen2k/deeptest/pkg/lib/comm"
)

type EndpointTagService struct {
	EndpointTagRepo *repo.EndpointTagRepo `inject:""`
}

func (s *EndpointTagService) ListTagsByProject(projectId uint) (tags []model.EndpointTag, err error) {
	tags, err = s.EndpointTagRepo.ListByProject(projectId)
	return
}

func (s *EndpointTagService) Create(name string, projectId uint) (id uint, err error) {
	id, err = s.EndpointTagRepo.Create(name, projectId)
	return
}

// GetTagIdsNyName 存在的tagName查出来tagId,不存在的写入数据库
func (s *EndpointTagService) GetTagIdsNyName(tagNames []string, projectId uint) (tagIds []uint, err error) {
	tags, err := s.EndpointTagRepo.BatchGetByName(tagNames, projectId)
	if err != nil {
		return
	}

	tagNamesExisted := make([]string, 0)
	for _, tag := range tags {
		tagNamesExisted = append(tagNamesExisted, tag.Name)
	}

	tagNamesNeedInsert := _commonUtils.Difference(tagNames, tagNamesExisted)

	err s.EndpointTagRepo.BatchCreate(tagNamesNeedInsert, projectId)
}
