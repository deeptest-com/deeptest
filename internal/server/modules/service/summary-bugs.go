package service

import (
	"errors"
	"fmt"
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	"strconv"
)

type SummaryBugsService struct {
	SummaryBugsRepo *repo.SummaryBugsRepo `inject:""`
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
				res.Critical = DecimalPer(result.Count, res.Total)
				res.Critical, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", res.Critical), 64)
			case "blocker":
				res.Blocker = DecimalPer(result.Count, res.Total)
				res.Blocker, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", res.Blocker), 64)
			case "deadly":
				res.Deadly = DecimalPer(result.Count, res.Total)
				res.Deadly, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", res.Deadly), 64)
			case "major":
				res.Major = DecimalPer(result.Count, res.Total)
				res.Major, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", res.Major), 64)
			case "minor":
				res.Minor = DecimalPer(result.Count, res.Total)
				res.Minor, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", res.Minor), 64)
			default:
				errors.New("Bug严重程度错误,请检查数据")
			}
		}
	}

	return
}

func DecimalPer(number int64, total int64) float64 {
	value := float64(number) / float64(total)
	return value * 100.0
}

func (s *SummaryBugsService) HandlerSummaryBugsRepo() *repo.SummaryBugsRepo {
	return repo.NewSummaryBugsRepo()
}

// FindByProjectId
func (s *SummaryBugsService) FindByProjectIdGroupByBugSeverity(projectId int64) (summaryBugsSeverity []model.SummaryBugsSeverity, err error) {

	summaryBugsSeverity, err = s.HandlerSummaryBugsRepo().FindByProjectIdGroupByBugSeverity(projectId)
	return
}

func (s *SummaryBugsService) FindProjectIds() (projectIds []int64, err error) {

	return s.HandlerSummaryBugsRepo().FindProjectIds()
}

// FindGroupByBugSeverity
func (s *SummaryBugsService) FindGroupByBugSeverity() (summaryBugsSeverity []model.SummaryBugsSeverity, err error) {

	summaryBugsSeverity, err = s.HandlerSummaryBugsRepo().FindGroupByBugSeverity()
	return
}

func (s *SummaryBugsService) Create(req model.SummaryBugs) (err error) {

	return s.HandlerSummaryBugsRepo().Create(req)
}

func (s *SummaryBugsService) CreateBug(req model.SummaryBugs) (err error) {
	id, err := s.Existed(req.BugId, req.ProjectId)
	if id == 0 {
		err = s.Create(req)
	} else {
		err = s.UpdateColumnsByDate(req, id)
	}
	return
}

func (s *SummaryBugsService) UpdateColumnsByDate(req model.SummaryBugs, id int64) (err error) {

	return s.HandlerSummaryBugsRepo().UpdateColumnsByDate(req, id)
}

func (s *SummaryBugsService) Existed(bugId int64, projectId int64) (id int64, err error) {

	return s.HandlerSummaryBugsRepo().Existed(bugId, projectId)
}

// Count
func (s *SummaryBugsService) Count() (count int64, err error) {

	return s.HandlerSummaryBugsRepo().Count()
}

// CountByProjectId
func (s *SummaryBugsService) CountByProjectId(projectId int64) (count int64, err error) {

	return s.HandlerSummaryBugsRepo().CountByProjectId(projectId)
}
