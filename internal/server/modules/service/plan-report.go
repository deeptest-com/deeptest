package service

import (
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	repo2 "github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	"github.com/aaronchen2k/deeptest/pkg/domain"
)

type PlanReportService struct {
	ReportRepo *repo2.PlanReportRepo `inject:""`
	LogRepo    *repo2.LogRepo        `inject:""`
}

func (s *PlanReportService) Paginate(tenantId consts.TenantId, req v1.PlanReportReqPaginate, projectId int) (ret _domain.PageData, err error) {
	ret, err = s.ReportRepo.Paginate(tenantId, req, projectId)
	return
}

func (s *PlanReportService) GetById(tenantId consts.TenantId, id uint) (report model.PlanReportDetail, err error) {
	report, err = s.ReportRepo.Get(tenantId, id)
	return
}

func (s *PlanReportService) DeleteById(tenantId consts.TenantId, id uint) error {
	return s.ReportRepo.DeleteById(tenantId, id)
}
