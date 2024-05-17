package service

import (
	"errors"
	"fmt"
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	"strconv"
)

type SummaryBugsService struct {
	SummaryBugsRepo *repo.SummaryBugsRepo `inject:""`
}

func (s *SummaryBugsService) Bugs(tenantId consts.TenantId, projectId int64) (res v1.ResSummaryBugs, err error) {

	var summaryBugsSeverity []model.SummaryBugsSeverity
	if projectId == 0 {
		res.Total, err = s.Count(tenantId)
		summaryBugsSeverity, err = s.FindGroupByBugSeverity(tenantId)
	} else {
		res.Total, err = s.CountByProjectId(tenantId, projectId)
		summaryBugsSeverity, err = s.FindByProjectIdGroupByBugSeverity(tenantId, projectId)
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
			case "suggest":
				res.Suggest = DecimalPer(result.Count, res.Total)
				res.Suggest, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", res.Suggest), 64)
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
	return s.SummaryBugsRepo
}

// FindByProjectId
func (s *SummaryBugsService) FindByProjectIdGroupByBugSeverity(tenantId consts.TenantId, projectId int64) (summaryBugsSeverity []model.SummaryBugsSeverity, err error) {

	summaryBugsSeverity, err = s.HandlerSummaryBugsRepo().FindByProjectIdGroupByBugSeverity(tenantId, projectId)
	return
}

func (s *SummaryBugsService) FindProjectIds(tenantId consts.TenantId) (projectIds []int64, err error) {

	return s.HandlerSummaryBugsRepo().FindProjectIds(tenantId)
}

// FindGroupByBugSeverity
func (s *SummaryBugsService) FindGroupByBugSeverity(tenantId consts.TenantId) (summaryBugsSeverity []model.SummaryBugsSeverity, err error) {

	summaryBugsSeverity, err = s.HandlerSummaryBugsRepo().FindGroupByBugSeverity(tenantId)
	return
}

func (s *SummaryBugsService) Create(tenantId consts.TenantId, req model.SummaryBugs) (err error) {

	return s.HandlerSummaryBugsRepo().Create(tenantId, req)
}

func (s *SummaryBugsService) CreateBug(tenantId consts.TenantId, req model.SummaryBugs) (err error) {
	id, err := s.Existed(tenantId, req.BugId, req.ProjectId)
	if id == 0 {
		err = s.Create(tenantId, req)
	} else {
		err = s.UpdateColumnsByDate(tenantId, req, id)
	}
	return
}

func (s *SummaryBugsService) UpdateColumnsByDate(tenantId consts.TenantId, req model.SummaryBugs, id int64) (err error) {

	return s.HandlerSummaryBugsRepo().UpdateColumnsByDate(tenantId, req, id)
}

func (s *SummaryBugsService) Existed(tenantId consts.TenantId, bugId int64, projectId int64) (id int64, err error) {

	return s.HandlerSummaryBugsRepo().Existed(tenantId, bugId, projectId)
}

// Count
func (s *SummaryBugsService) Count(tenantId consts.TenantId) (count int64, err error) {

	return s.HandlerSummaryBugsRepo().Count(tenantId)
}

// CountByProjectId
func (s *SummaryBugsService) CountByProjectId(tenantId consts.TenantId, projectId int64) (count int64, err error) {

	return s.HandlerSummaryBugsRepo().CountByProjectId(tenantId, projectId)
}

func (s *SummaryBugsService) CreateBugs(tenantId consts.TenantId) {
	bugs, err := s.HandlerSummaryBugsRepo().GetNewBugs(tenantId)
	if err != nil {
		return
	}

	var data []model.SummaryBugs
	for _, bug := range bugs {
		data = append(data, model.SummaryBugs{
			ProjectId:   int64(bug.ProjectId),
			BugId:       int64(bug.ID),
			BugSeverity: bug.Severity.String(),
		})
	}
	s.HandlerSummaryBugsRepo().Creates(tenantId, data)
	return
}
