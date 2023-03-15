package service

import (
	"errors"
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

func (s *SummaryBugsService) Bugs(projectId int64) (res v1.ResSummaryBugs, err error) {

	var summaryBugsSeverity []model.SummaryBugsSeverity
	if projectId == 0 {
		res.Total, err = s.Count()
		summaryBugsSeverity, err = s.FindGroupByBugSeverity()
	} else {
		res.Total, err = s.CountByProjectId(projectId)
		summaryBugsSeverity, err = s.FindByProjectIdGroupByBugSeverity(projectId)
	}

	if err == nil {
		for _, result := range summaryBugsSeverity {
			switch result.BugSeverity {
			case "critical":
				res.Critical = result.Count
			case "blocker":
				res.Blocker = result.Count
			case "deadly":
				res.Deadly = result.Count
			case "major":
				res.Major = result.Count
			case "minor":
				res.Minor = result.Count
			default:
				errors.New("Bug严重程度错误,请检查数据")
			}
		}
	}

	return
}

// FindByProjectId
func (s *SummaryBugsService) FindByProjectIdGroupByBugSeverity(projectId int64) (summaryBugsSeverity []model.SummaryBugsSeverity, err error) {
	summaryBugsSeverity, err = s.SummaryBugsRepo.FindByProjectIdGroupByBugSeverity(projectId)
	return
}

// FindGroupByBugSeverity
func (s *SummaryBugsService) FindGroupByBugSeverity() (summaryBugsSeverity []model.SummaryBugsSeverity, err error) {
	summaryBugsSeverity, err = s.SummaryBugsRepo.FindGroupByBugSeverity()
	return
}

// Create
func (s *SummaryBugsService) Create(req v1.ReqSummaryBugs) (err error) {
	var summaryBugs model.SummaryBugs
	copier.CopyWithOption(&summaryBugs, req, copier.Option{DeepCopy: true})
	return s.SummaryBugsRepo.Create(&summaryBugs)
}

// Count
func (s *SummaryBugsService) Count() (count int64, err error) {
	return s.SummaryBugsRepo.Count()
}

// CountByProjectId
func (s *SummaryBugsService) CountByProjectId(projectId int64) (count int64, err error) {
	return s.SummaryBugsRepo.CountByProjectId(projectId)
}
