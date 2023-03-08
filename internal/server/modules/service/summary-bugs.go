package service

import (
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	"github.com/jinzhu/copier"
)

type SummaryBugsService struct {
	SummaryBugsRepo *repo.SummaryBugsRepo `inject:""`
}

func NewSummaryBugsService() *SummaryBugsService {
	return &SummaryBugsService{}
}

// FindByProjectId
func (s *SummaryBugsService) FindByProjectIdGroupByBugSeverity(req v1.SummaryBugsReq) (summaryBugs []model.SummaryBugs, err error) {
	return s.SummaryBugsRepo.FindByProjectIdGroupByBugSeverity(req.ProjectId)
}

// FindGroupByBugSeverity
func (s *SummaryBugsService) FindGroupByBugSeverity() (summaryBugs []model.SummaryBugs, err error) {
	return s.SummaryBugsRepo.FindGroupByBugSeverity()
}

// Create
func (s *SummaryBugsService) Create(req v1.SummaryBugsReq) (err error) {
	var summaryBugs model.SummaryBugs
	copier.CopyWithOption(&summaryBugs, req, copier.Option{DeepCopy: true})
	return s.SummaryBugsRepo.Create(&summaryBugs)
}

// Count
func (s *SummaryBugsService) Count() (count int64, err error) {
	return s.SummaryBugsRepo.Count()
}

// CountByProjectId
func (s *SummaryBugsService) CountByProjectId(req v1.SummaryBugsReq) (count int64, err error) {
	return s.SummaryBugsRepo.CountByProjectId(req.ProjectId)
}
