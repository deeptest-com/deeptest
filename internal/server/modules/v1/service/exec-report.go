package service

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/business"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/repo"
	"time"
)

type ExecReportService struct {
	ScenarioProcessorRepo *repo.ScenarioProcessorRepo  `inject:""`
	ScenarioRepo          *repo.ScenarioRepo           `inject:""`
	TestResultRepo        *repo.ReportRepo             `inject:""`
	TestLogRepo           *repo.LogRepo                `inject:""`
	ReportRepo            *repo.ReportRepo             `inject:""`
	ExecRequestService    *business.ExecRequestService `inject:""`
}

func (s ExecReportService) UpdateTestReport(rootLog domain.Log) {
	report, _ := s.ReportRepo.Get(rootLog.ReportId)
	s.countRequest(rootLog, &report)

	now := time.Now()
	report.EndTime = &now
	report.Duration = report.EndTime.Unix() - report.StartTime.Unix()

	s.ReportRepo.UpdateResult(report)
}

func (s ExecReportService) countRequest(log domain.Log, report *model.Report) {
	if log.ProcessorType == consts.ProcessorInterfaceDefault {
		switch log.ResultStatus {
		case consts.Pass:
			report.PassNum++

		case consts.Fail:
			report.FailNum++
			report.ResultStatus = consts.Fail

		default:
		}
	}

	if log.Logs == nil {
		return
	}

	for _, log := range *log.Logs {
		s.countRequest(*log, report)
	}
}
