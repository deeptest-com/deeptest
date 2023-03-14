package service

import (
	"message/domain"
	"message/repo"
)

type MessageServiceV1 struct {
	MessageRepo     *repo.MessageRepo     `inject:""`
	MessageReadRepo *repo.MessageReadRepo `inject:""`
}

func NewMessageService() *MessageServiceV1 {
	return &MessageServiceV1{}
}

func (s *MessageServiceV1) Create(req domain.MessageReq) (uint, error) {
	return s.MessageRepo.Create(req)
}

func (s *MessageServiceV1) Paginate(req domain.MessageReqPaginate) (ret domain.PageData, err error) {
	ret, err = s.MessageRepo.Paginate(req)

	if err != nil {
		return
	}

	return
}

func (s *MessageServiceV1) UnreadCount(scope domain.MessageScope) (count int64, err error) {
	count, err = s.MessageRepo.GetUnreadCount(scope)

	if err != nil {
		return
	}

	return
}

func (s *MessageServiceV1) OperateRead(req domain.MessageReadReq) (uint, error) {
	return s.MessageReadRepo.Create(req)
}
