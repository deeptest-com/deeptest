package service

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/repo"
	"time"
)

type ExecReportService struct {
	ScenarioProcessorRepo *repo.ScenarioProcessorRepo `inject:""`
	ScenarioRepo          *repo.ScenarioRepo          `inject:""`
	TestResultRepo        *repo.ReportRepo            `inject:""`
	TestLogRepo           *repo.LogRepo               `inject:""`
	ReportRepo            *repo.ReportRepo            `inject:""`
}

func (s ExecReportService) UpdateTestReport(rootLog domain.ExecLog) (report model.Report) {
	report, _ = s.ReportRepo.Get(rootLog.ReportId)
	if report.InterfaceStatusMap == nil {
		report.InterfaceStatusMap = map[uint]map[consts.ResultStatus]int{}
	}

	s.countRequest(rootLog, &report)
	s.summarizeInterface(&report)

	now := time.Now()
	report.EndTime = &now
	report.Duration = report.EndTime.Unix() - report.StartTime.Unix()

	s.ReportRepo.UpdateResult(report)

	return
}

func (s ExecReportService) countRequest(log domain.ExecLog, report *model.Report) {
	if log.ProcessorType == consts.ProcessorInterfaceDefault {
		s.countInterface(log.InterfaceId, log.ResultStatus, report)

		report.TotalRequestNum++

		switch log.ResultStatus {
		case consts.Pass:
			report.PassRequestNum++

		case consts.Fail:
			report.FailRequestNum++
			report.ResultStatus = consts.Fail

		default:
		}

	} else if log.ProcessorType == consts.ProcessorAssertionDefault {
		switch log.ResultStatus {
		case consts.Pass:
			report.PassAssertionNum++

		case consts.Fail:
			report.FailAssertionNum++
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

func (s ExecReportService) countInterface(interfaceId uint, status consts.ResultStatus, report *model.Report) {
	if report.InterfaceStatusMap[interfaceId] == nil {
		report.InterfaceStatusMap[interfaceId] = map[consts.ResultStatus]int{}
		report.InterfaceStatusMap[interfaceId][consts.Pass] = 0
		report.InterfaceStatusMap[interfaceId][consts.Fail] = 0
	}

	switch status {
	case consts.Pass:
		report.InterfaceStatusMap[interfaceId][consts.Pass]++

	case consts.Fail:
		report.InterfaceStatusMap[interfaceId][consts.Fail]++

	default:
	}
}

func (s ExecReportService) summarizeInterface(report *model.Report) {
	for _, val := range report.InterfaceStatusMap {
		if val[consts.Fail] > 0 {
			report.FailInterfaceNum++
		} else {
			report.PassInterfaceNum++
		}

		report.TotalInterfaceNum++
	}
}
