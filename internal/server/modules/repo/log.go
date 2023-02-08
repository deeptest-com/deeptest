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
	DB       *gorm.DB  `inject:""`
	RoleRepo *RoleRepo `inject:""`
}

func NewLogRepo() *LogRepo {
	return &LogRepo{}
}

func (r *LogRepo) ListByReport(reportId uint) (logs []*model.ExecLogProcessor, err error) {
	err = r.DB.
		Where("report_id=?", reportId).
		Where("NOT deleted").
		Order("parent_id ASC, id ASC").
		Find(&logs).Error
	return
}

func (r *LogRepo) Get(id uint) (scenario model.ExecLogProcessor, err error) {
	err = r.DB.Model(&model.ExecLogProcessor{}).Where("id = ?", id).First(&scenario).Error
	if err != nil {
		logUtils.Errorf("find scenario by id error", zap.String("error:", err.Error()))
		return scenario, err
	}

	return scenario, nil
}

func (r *LogRepo) Save(log *model.ExecLogProcessor) (err error) {
	err = r.DB.Save(log).Error

	return
}

func (r *LogRepo) DeleteById(id uint) (err error) {
	err = r.DB.Model(&model.ExecLogProcessor{}).Where("id = ?", id).
		Updates(map[string]interface{}{"deleted": true}).Error
	if err != nil {
		logUtils.Errorf("delete scenario by id error", zap.String("error:", err.Error()))
		return
	}

	return
}

func (r *LogRepo) CreateLogs(rootResult agentDomain.Result, report *model.Report) (err error) {
	r.CreateLog(rootResult, 0, report.ID)

	return
}

func (r *LogRepo) CreateLog(result agentDomain.Result, parentId, reportId uint) (id uint, err error) {
	if result.ProcessorCategory == consts.ProcessorInterface {
		id, _ = r.CreateInterfaceLog(result, parentId, reportId)
	} else {
		id, _ = r.CreateCommonLog(result, parentId, reportId)
	}

	for _, child2 := range result.Children {
		child := *child2
		r.CreateLog(child, id, reportId)
	}

	return
}

func (r *LogRepo) CreateInterfaceLog(result agentDomain.Result, parentId, reportId uint) (id uint, err error) {
	po := model.ExecLogProcessor{
		Name:              result.Name,
		ProcessorCategory: result.ProcessorCategory,
		ProcessorType:     result.ProcessorType,
		ResultStatus:      result.ResultStatus,

		ReqContent:  result.ReqContent,
		RespContent: result.RespContent,

		InterfaceId: result.InterfaceId,
		ProcessorId: result.ProcessorId,
		ParentId:    parentId,
		ReportId:    reportId,
	}

	err = r.Save(&po)
	id = po.ID

	return
}

func (r *LogRepo) CreateCommonLog(result agentDomain.Result, parentId, reportId uint) (id uint, err error) {
	po := model.ExecLogProcessor{
		Name:              result.Name,
		ProcessorCategory: result.ProcessorCategory,
		ProcessorType:     result.ProcessorType,
		ResultStatus:      result.ResultStatus,

		Summary: result.Summary,

		ProcessorId: result.ProcessorId,
		ParentId:    parentId,
		ReportId:    reportId,
	}

	err = r.Save(&po)
	id = po.ID

	return
}
