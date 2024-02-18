package repo

import (
	"github.com/aaronchen2k/deeptest/internal/agent/exec/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type LogRepo struct {
	*BaseRepo `inject:""`
	DB        *gorm.DB  `inject:""`
	RoleRepo  *RoleRepo `inject:""`
}

func (r *LogRepo) ListByReport(tenantId consts.TenantId, reportId uint) (logs []*model.ExecLogProcessor, err error) {
	err = r.GetDB(tenantId).
		Where("report_id=?", reportId).
		Where("NOT deleted").
		Order("parent_id ASC, id ASC").
		Find(&logs).Error
	return
}

func (r *LogRepo) Get(tenantId consts.TenantId, id uint) (scenario model.ExecLogProcessor, err error) {
	err = r.GetDB(tenantId).Model(&model.ExecLogProcessor{}).Where("id = ?", id).First(&scenario).Error
	if err != nil {
		logUtils.Errorf("find scenario by id error", zap.String("error:", err.Error()))
		return scenario, err
	}

	return scenario, nil
}

func (r *LogRepo) Save(tenantId consts.TenantId, log *model.ExecLogProcessor) (err error) {
	err = r.GetDB(tenantId).Save(log).Error

	return
}

func (r *LogRepo) DeleteById(tenantId consts.TenantId, id uint) (err error) {
	err = r.GetDB(tenantId).Model(&model.ExecLogProcessor{}).Where("id = ?", id).
		Updates(map[string]interface{}{"deleted": true}).Error
	if err != nil {
		logUtils.Errorf("delete scenario by id error", zap.String("error:", err.Error()))
		return
	}

	return
}

func (r *LogRepo) CreateLogs(tenantId consts.TenantId, rootResult agentDomain.ScenarioExecResult, report *model.ScenarioReport, processorToInvokeIdMap map[uint]uint) (
	err error) {
	r.CreateLog(tenantId, rootResult, 0, report.ID, processorToInvokeIdMap)

	return
}

func (r *LogRepo) CreateLog(tenantId consts.TenantId, result agentDomain.ScenarioExecResult, parentId, reportId uint, processorToInvokeIdMap map[uint]uint) (
	id uint, err error) {

	if result.ProcessorCategory == consts.ProcessorInterface {
		id, _ = r.CreateInterfaceLog(tenantId, result, parentId, reportId, processorToInvokeIdMap[result.ProcessorId])
	} else {
		id, _ = r.CreateCommonLog(tenantId, result, parentId, reportId)
	}

	for _, child2 := range result.Children {
		child := *child2
		r.CreateLog(tenantId, child, id, reportId, processorToInvokeIdMap)
	}

	return
}

func (r *LogRepo) CreateInterfaceLog(tenantId consts.TenantId, result agentDomain.ScenarioExecResult, parentId, reportId, invokeId uint) (id uint, err error) {
	po := model.ExecLogProcessor{
		Name:              result.Name,
		ProcessorCategory: result.ProcessorCategory,
		ProcessorType:     result.ProcessorType,
		ResultStatus:      result.ResultStatus,

		ReqContent:  result.ReqContent,
		RespContent: result.RespContent,

		EndpointInterfaceId: result.EndpointInterfaceId,
		DebugInterfaceId:    result.DebugInterfaceId,

		ScenarioProcessorId: result.ProcessorId,
		ScenarioId:          result.ScenarioId,
		ParentId:            parentId,
		InvokeId:            invokeId,
		ReportId:            reportId,
		Round:               result.Round,
		Detail:              result.Detail,
		Summary:             result.Summary,
	}

	err = r.Save(tenantId, &po)
	id = po.ID

	return
}

func (r *LogRepo) CreateCommonLog(tenantId consts.TenantId, result agentDomain.ScenarioExecResult, parentId, reportId uint) (id uint, err error) {
	po := model.ExecLogProcessor{
		Name:              result.Name,
		ProcessorCategory: result.ProcessorCategory,
		ProcessorType:     result.ProcessorType,
		ResultStatus:      result.ResultStatus,

		Summary: result.Summary,

		ScenarioProcessorId: result.ProcessorId,
		ParentId:            parentId,
		ReportId:            reportId,
		Detail:              result.Detail,
		Round:               result.Round,
	}

	err = r.Save(tenantId, &po)
	id = po.ID

	return
}
