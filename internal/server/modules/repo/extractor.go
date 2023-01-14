package repo

import (
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/agent/exec/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	model "github.com/aaronchen2k/deeptest/internal/server/modules/model"
	_domain "github.com/aaronchen2k/deeptest/pkg/domain"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"github.com/jinzhu/copier"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"strings"
)

type ExtractorRepo struct {
	DB                     *gorm.DB                `inject:""`
	InterfaceRepo          *InterfaceRepo          `inject:""`
	ProcessorInterfaceRepo *ProcessorInterfaceRepo `inject:""`
}

func (r *ExtractorRepo) List(interfaceId uint, usedBy consts.UsedBy) (pos []model.InterfaceExtractor, err error) {
	err = r.DB.
		Where("interface_id=?", interfaceId).
		Where("used_by = ? AND NOT deleted", usedBy).
		Order("created_at ASC").
		Find(&pos).Error
	return
}

func (r *ExtractorRepo) ListTo(interfaceId uint, usedBy consts.UsedBy) (ret []domain.Extractor, err error) {
	pos, _ := r.List(interfaceId, usedBy)

	for _, po := range pos {
		extractor := domain.Extractor{}
		copier.CopyWithOption(&extractor, po, copier.Option{DeepCopy: true})

		ret = append(ret, extractor)
	}

	return
}

func (r *ExtractorRepo) Get(id uint) (extractor model.InterfaceExtractor, err error) {
	err = r.DB.
		Where("id=?", id).
		Where("NOT deleted").
		First(&extractor).Error
	return
}

func (r *ExtractorRepo) GetByInterfaceVariable(variable string, id, interfaceId uint) (extractor model.InterfaceExtractor, err error) {
	db := r.DB.Model(&extractor).
		Where("variable = ? AND interface_id =? AND used_by = ? AND not deleted",
			variable, interfaceId, consts.Interface)

	if id > 0 {
		db.Where("id != ?", id)
	}

	db.First(&extractor)

	return
}

func (r *ExtractorRepo) Save(extractor *model.InterfaceExtractor) (id uint, bizErr _domain.BizErr) {
	po, _ := r.GetByInterfaceVariable(extractor.Variable, extractor.ID, extractor.InterfaceId)
	if po.ID > 0 && extractor.UsedBy == consts.Interface {
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

func (r *ExtractorRepo) Update(extractor *model.InterfaceExtractor) (err error) {
	err = r.DB.Updates(extractor).Error
	if err != nil {
		return
	}

	return
}

func (r *ExtractorRepo) CreateOrUpdateResult(extractor *model.InterfaceExtractor, usedBy consts.UsedBy) (err error) {
	po, _ := r.GetByInterfaceVariable(extractor.Variable, extractor.ID, extractor.InterfaceId)
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
	err = r.DB.Model(&model.InterfaceExtractor{}).
		Where("id=?", id).
		Update("deleted", true).
		Error

	return
}

func (r *ExtractorRepo) UpdateResult(extractor model.InterfaceExtractor, usedBy consts.UsedBy) (err error) {
	extractor.Result = strings.TrimSpace(extractor.Result)
	if extractor.Result == "" {
		return
	}

	values := map[string]interface{}{
		"result": extractor.Result,
	}
	err = r.DB.Model(&extractor).
		Where("id = ? AND used_by=?", extractor.ID, usedBy).
		Updates(values).Error
	if err != nil {
		logUtils.Errorf("update scenario error", zap.String("error:", err.Error()))
		return err
	}

	return
}
func (r *ExtractorRepo) UpdateResultToExecLog(extractor model.InterfaceExtractor, log *model.ExecLogProcessor) (
	logExtractor model.ExecLogExtractor, err error) {

	copier.CopyWithOption(&logExtractor, extractor, copier.Option{DeepCopy: true})

	logExtractor.ID = 0
	logExtractor.LogId = log.ID
	logExtractor.CreatedAt = nil
	logExtractor.UpdatedAt = nil

	err = r.DB.Save(&logExtractor).Error

	return
}

func (r *ExtractorRepo) ListExtractorVariableByInterface(interfaceId uint) (variables []v1.Variable, err error) {
	err = r.DB.Model(&model.InterfaceExtractor{}).
		Select("id, variable AS name, result AS value").
		Where("interface_id=?", interfaceId).
		Where("NOT deleted AND NOT disabled").
		Order("created_at ASC").
		Find(&variables).Error

	return
}

func (r *ExtractorRepo) ListValidExtractorVariableForInterface(interfaceId, projectId uint, usedBy consts.UsedBy) (
	variables []v1.Variable, err error) {

	q := r.DB.Model(&model.InterfaceExtractor{}).
		Select("id, variable AS name, result AS value, "+
			"interface_id AS interfaceId, scope AS scope").
		Where("used_by = ?", usedBy).
		Where("NOT deleted AND NOT disabled")

	if usedBy == consts.Interface {
		q.Where("project_id=?", projectId)

	} else {
		processorInterface, _ := r.ProcessorInterfaceRepo.Get(interfaceId)

		var parentIds []uint
		r.GetParentIds(processorInterface.ProcessorId, &parentIds)

		q.Where("scenario_id=?", processorInterface.ScenarioId).
			Where("scope = ? OR processor_id IN(?)", consts.Global, parentIds)
	}

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
