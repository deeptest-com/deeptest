package repo

import (
	"github.com/aaronchen2k/deeptest/internal/agent/exec/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	model "github.com/aaronchen2k/deeptest/internal/server/modules/model"
	_domain "github.com/aaronchen2k/deeptest/pkg/domain"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"github.com/jinzhu/copier"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"strings"
)

type ExtractorRepo struct {
	DB *gorm.DB `inject:""`
}

func (r *ExtractorRepo) List(debugInterfaceId, endpointInterfaceId uint) (pos []model.DebugInterfaceExtractor, err error) {
	db := r.DB.
		Where("NOT deleted").
		Order("created_at ASC")

	if debugInterfaceId > 0 {
		db.Where("debug_interface_id=?", debugInterfaceId)
	} else {
		db.Where("endpoint_interface_id=? AND debug_interface_id=?", endpointInterfaceId, 0)
	}

	err = db.Find(&pos).Error

	return
}

func (r *ExtractorRepo) ListTo(debugInterfaceId, endpointInterfaceId uint) (ret []agentDomain.Extractor, err error) {
	pos, _ := r.List(debugInterfaceId, endpointInterfaceId)

	for _, po := range pos {
		extractor := agentDomain.Extractor{}
		copier.CopyWithOption(&extractor, po, copier.Option{DeepCopy: true})

		ret = append(ret, extractor)
	}

	return
}

func (r *ExtractorRepo) Get(id uint) (extractor model.DebugInterfaceExtractor, err error) {
	err = r.DB.
		Where("id=?", id).
		Where("NOT deleted").
		First(&extractor).Error
	return
}

func (r *ExtractorRepo) GetByInterfaceVariable(variable string, id, interfaceId uint) (extractor model.DebugInterfaceExtractor, err error) {
	db := r.DB.Model(&extractor).
		Where("variable = ? AND endpoint_interface_id =? AND used_by = ? AND not deleted",
			variable, interfaceId, consts.InterfaceDebug)

	if id > 0 {
		db.Where("id != ?", id)
	}

	db.First(&extractor)

	return
}

func (r *ExtractorRepo) Save(extractor *model.DebugInterfaceExtractor) (id uint, bizErr _domain.BizErr) {
	po, _ := r.GetByInterfaceVariable(extractor.Variable, extractor.ID, extractor.EndpointInterfaceId)
	if po.ID > 0 {
		bizErr.Code = _domain.ErrNameExist.Code
		return
	}

	err := r.DB.Save(extractor).Error
	if err != nil {
		bizErr.Code = _domain.SystemErr.Code
		return
	}

	id = extractor.ID

	return
}

func (r *ExtractorRepo) Update(extractor *model.DebugInterfaceExtractor) (err error) {
	err = r.DB.Updates(extractor).Error
	if err != nil {
		return
	}

	return
}

func (r *ExtractorRepo) CreateOrUpdateResult(extractor *model.DebugInterfaceExtractor, usedBy consts.UsedBy) (err error) {
	po, _ := r.GetByInterfaceVariable(extractor.Variable, extractor.ID, extractor.EndpointInterfaceId)
	if po.ID > 0 {
		extractor.ID = po.ID
		r.UpdateResult(*extractor, usedBy)
		return
	}

	err = r.DB.Save(extractor).Error
	if err != nil {
		return
	}

	return
}

func (r *ExtractorRepo) Delete(id uint) (err error) {
	err = r.DB.Model(&model.DebugInterfaceExtractor{}).
		Where("id=?", id).
		Update("deleted", true).
		Error

	return
}

func (r *ExtractorRepo) UpdateResult(extractor model.DebugInterfaceExtractor, usedBy consts.UsedBy) (err error) {
	extractor.Result = strings.TrimSpace(extractor.Result)

	values := map[string]interface{}{}
	if extractor.Result != "" {
		values["result"] = extractor.Result
	}
	if extractor.Scope != "" {
		values["scope"] = extractor.Scope
	}

	err = r.DB.Model(&extractor).
		Where("id = ?", extractor.ID).
		Updates(values).Error

	if err != nil {
		logUtils.Errorf("update scenario error", zap.String("error:", err.Error()))
		return err
	}

	return
}
func (r *ExtractorRepo) UpdateResultToExecLog(extractor model.DebugInterfaceExtractor, log *model.ExecLogProcessor) (
	logExtractor model.ExecLogExtractor, err error) {

	copier.CopyWithOption(&logExtractor, extractor, copier.Option{DeepCopy: true})

	logExtractor.ID = 0
	logExtractor.LogId = log.ID
	logExtractor.CreatedAt = nil
	logExtractor.UpdatedAt = nil

	err = r.DB.Save(&logExtractor).Error

	return
}

func (r *ExtractorRepo) ListExtractorVariableByInterface(interfaceId uint) (variables []domain.Variable, err error) {
	err = r.DB.Model(&model.DebugInterfaceExtractor{}).
		Select("id, variable AS name, result AS value").
		Where("endpoint_interface_id=?", interfaceId).
		Where("NOT deleted AND NOT disabled").
		Order("created_at ASC").
		Find(&variables).Error

	return
}

func (r *ExtractorRepo) ListValidExtractorVariableForInterface(interfaceId, projectId uint, usedBy consts.UsedBy) (
	variables []domain.Variable, err error) {

	q := r.DB.Model(&model.DebugInterfaceExtractor{}).
		Select("id, variable AS name, result AS value, " +
			"endpoint_interface_id AS endpointInterfaceId, scope AS scope").
		Where("NOT deleted AND NOT disabled")

	//if usedBy == consts.InterfaceDebug {
	//	q.Where("project_id=?", projectId)
	//
	//} else {
	//	processorInterface, _ := r.ProcessorInterfaceRepo.Get(interfaceId)
	//
	//	var parentIds []uint
	//	r.GetParentIds(processorInterface.ProcessorId, &parentIds)
	//
	//	q.Where("scenario_id=?", processorInterface.ScenarioId).
	//		Where("scope = ? OR scenario_processor_id IN(?)", consts.Public, parentIds)
	//}

	err = q.Order("created_at ASC").
		Find(&variables).Error

	return
}

func (r *ExtractorRepo) GetParentIds(processorId uint, ids *[]uint) {
	var po model.Processor

	r.DB.Where("id = ?", processorId).
		Where("NOT deleted AND NOT disabled").
		First(&po)

	if po.ID > 0 {
		*ids = append(*ids, processorId)
	}

	if po.ParentId > 0 {
		r.GetParentIds(po.ParentId, ids)
	}

	return
}
