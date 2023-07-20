package service

import (
	repo "github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	"log"
)

type ExecConditionService struct {
	PreConditionRepo  *repo.PreConditionRepo  `inject:""`
	PostConditionRepo *repo.PostConditionRepo `inject:""`

	ExtractorRepo  *repo.ExtractorRepo  `inject:""`
	CheckpointRepo *repo.CheckpointRepo `inject:""`
	ScriptRepo     *repo.ScriptRepo     `inject:""`
}

func (s *ExecConditionService) ExecPreCondition(debugInterfaceId, endpointInterfaceId uint) (err error) {
	conditions, err := s.PreConditionRepo.List(debugInterfaceId, endpointInterfaceId)

	for _, condition := range conditions {
		log.Println(condition)
	}

	return
}

func (s *ExecConditionService) ExecPostCondition(debugInterfaceId, endpointInterfaceId uint) (err error) {
	conditions, err := s.PostConditionRepo.List(debugInterfaceId, endpointInterfaceId)

	for _, condition := range conditions {
		log.Println(condition)
	}

	return
}
