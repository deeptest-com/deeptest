package repo

import (
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/server/core/dao"
	serverDomain "github.com/aaronchen2k/deeptest/internal/server/modules/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/model"
	"github.com/aaronchen2k/deeptest/pkg/domain"
	commonUtils "github.com/aaronchen2k/deeptest/pkg/lib/comm"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ScenarioRepo struct {
	DB       *gorm.DB  `inject:""`
	RoleRepo *RoleRepo `inject:""`
}

func NewScenarioRepo() *ScenarioRepo {
	return &ScenarioRepo{}
}

func (r *ScenarioRepo) Paginate(req serverDomain.ScenarioReqPaginate) (data _domain.PageData, err error) {
	var count int64

	db := r.DB.Model(&model.TestScenario{}).Where("NOT deleted")

	if req.Keywords != "" {
		db = db.Where("name LIKE ?", fmt.Sprintf("%%%s%%", req.Keywords))
	}
	if req.Enabled != "" {
		db = db.Where("disabled = ?", commonUtils.IsDisable(req.Enabled))
	}

	err = db.Count(&count).Error
	if err != nil {
		logUtils.Errorf("count scenario error", zap.String("error:", err.Error()))
		return
	}

	scenarios := make([]*model.TestScenario, 0)

	err = db.
		Scopes(dao.PaginateScope(req.Page, req.PageSize, req.Order, req.Field)).
		Find(&scenarios).Error
	if err != nil {
		logUtils.Errorf("query scenario error", zap.String("error:", err.Error()))
		return
	}

	data.Populate(scenarios, count, req.Page, req.PageSize)

	return
}

func (r *ScenarioRepo) Get(id uint) (scenario model.TestScenario, err error) {
	err = r.DB.Model(&model.TestScenario{}).Where("id = ?", id).First(&scenario).Error
	if err != nil {
		logUtils.Errorf("find scenario by id error", zap.String("error:", err.Error()))
		return scenario, err
	}

	return scenario, nil
}

func (r *ScenarioRepo) FindByName(scenarioName string, id uint) (scenario model.TestScenario, err error) {
	db := r.DB.Model(&model.TestScenario{}).
		Where("name = ?", scenarioName)

	if id > 0 {
		db.Where("id != ?", id)
	}

	db.First(&scenario)

	return
}

func (r *ScenarioRepo) Create(scenario model.TestScenario) (ret model.TestScenario, err error) {
	po, err := r.FindByName(scenario.Name, 0)
	if po.Name != "" {
		err = fmt.Errorf("%d", _domain.BizErrNameExist.Code)
		return
	}

	err = r.DB.Model(&model.TestScenario{}).Create(&scenario).Error
	if err != nil {
		logUtils.Errorf("add scenario error", zap.String("error:", err.Error()))
		err = fmt.Errorf("%d", _domain.BizErrNameExist.Code)

		return
	}

	ret = scenario

	return
}

func (r *ScenarioRepo) Update(req model.TestScenario) error {
	values := map[string]interface{}{
		"name":     req.Name,
		"desc":     req.Desc,
		"disabled": req.Disabled,
	}
	err := r.DB.Model(&req).Where("id = ?", req.ID).Updates(values).Error
	if err != nil {
		logUtils.Errorf("update scenario error", zap.String("error:", err.Error()))
		return err
	}

	return nil
}

func (r *ScenarioRepo) DeleteById(id uint) (err error) {
	err = r.DB.Model(&model.TestScenario{}).Where("id = ?", id).
		Updates(map[string]interface{}{"deleted": true}).Error
	if err != nil {
		logUtils.Errorf("delete scenario by id error", zap.String("error:", err.Error()))
		return
	}

	return
}

func (r *ScenarioRepo) DeleteChildren(ids []int, tx *gorm.DB) (err error) {
	err = tx.Model(&model.TestScenario{}).Where("id IN (?)", ids).
		Updates(map[string]interface{}{"deleted": true}).Error
	if err != nil {
		logUtils.Errorf("batch delete scenario error", zap.String("error:", err.Error()))
		return err
	}

	return nil
}

func (r *ScenarioRepo) GetChildrenIds(id uint) (ids []int, err error) {
	tmpl := `
		WITH RECURSIVE scenario AS (
			SELECT * FROM biz_scenario WHERE id = %d
			UNION ALL
			SELECT child.* FROM biz_scenario child, scenario WHERE child.parent_id = scenario.id
		)
		SELECT id FROM scenario WHERE id != %d
    `
	sql := fmt.Sprintf(tmpl, id, id)
	err = r.DB.Raw(sql).Scan(&ids).Error
	if err != nil {
		logUtils.Errorf("get children scenario error", zap.String("error:", err.Error()))
		return
	}

	return
}
