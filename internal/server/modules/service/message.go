package service

import (
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
)

type MessageService struct {
	MessageRepo *repo.MessageRepo `inject:""`
}

func NewMessageService() *MessageService {
	return &MessageService{}
}

func (s *MessageService) GetScope(userId uint) (scope map[int][]string) {
	return s.MessageRepo.GetScope(userId)
}
