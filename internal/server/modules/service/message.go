package service

import (
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	_domain "github.com/aaronchen2k/deeptest/pkg/domain"
)

type MessageService struct {
	MessageRepo     *repo.MessageRepo     `inject:""`
	MessageReadRepo *repo.MessageReadRepo `inject:""`
}

func (s *MessageService) GetScope(userId uint) (scope map[int][]string) {
	return s.MessageRepo.GetScope(userId)
}

func (s *MessageService) Create(req v1.MessageReq) (uint, *_domain.BizErr) {
	return s.MessageRepo.Create(req)
}

func (s *MessageService) Paginate(req v1.MessageReqPaginate, userId uint) (ret _domain.PageData, err error) {
	req.Scope = s.MessageRepo.GetScope(userId)

	ret, err = s.MessageRepo.Paginate(req)

	if err != nil {
		return
	}

	return
}

func (s *MessageService) UnreadCount(userId uint) (count int64, err error) {
	scope := s.MessageRepo.GetScope(userId)
	req := v1.MessageScope{Scope: scope}

	count, err = s.MessageRepo.GetUnreadCount(req)

	if err != nil {
		return
	}

	return
}

func (s *MessageService) OperateRead(req v1.MessageReadReq) (uint, error) {
	return s.MessageReadRepo.Create(req)
}
