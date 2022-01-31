package repo

import (
	"errors"
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	commonUtils "github.com/aaronchen2k/deeptest/internal/pkg/lib/comm"
	logUtils "github.com/aaronchen2k/deeptest/internal/pkg/lib/log"
	"github.com/aaronchen2k/deeptest/internal/server/core/dao"
	serverDomain "github.com/aaronchen2k/deeptest/internal/server/modules/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type TestScriptRepo struct {
	DB       *gorm.DB  `inject:""`
	RoleRepo *RoleRepo `inject:""`
}

func NewTestScriptRepo() *TestScriptRepo {
	return &TestScriptRepo{}
}

func (r *TestScriptRepo) Paginate(req serverDomain.TestScriptReqPaginate) (data _domain.PageData, err error) {
	var count int64

	db := r.DB.Model(&model.TestScript{}).Where("NOT deleted")

	if req.Keywords != "" {
		db = db.Where("name LIKE ?", fmt.Sprintf("%%%s%%", req.Keywords))
	}
	if req.Enabled != "" {
		db = db.Where("disabled = ?", commonUtils.IsDisable(req.Enabled))
	}

	err = db.Count(&count).Error
	if err != nil {
		logUtils.Errorf("count test script error", zap.String("error:", err.Error()))
		return
	}

	scripts := make([]*model.TestScript, 0)

	err = db.
		Scopes(dao.PaginateScope(req.Page, req.PageSize, req.Order, req.Field)).
		Find(&scripts).Error
	if err != nil {
		logUtils.Errorf("query test script error", zap.String("error:", err.Error()))
		return
	}

	data.Populate(scripts, count, req.Page, req.PageSize)

	return
}

func (r *TestScriptRepo) GetDetail(id uint) (resp serverDomain.TestScriptResp, err error) {
	script := model.TestScript{}

	err = r.DB.Model(&model.TestScript{}).
		Where("id = ?", id).
		Preload("Steps", "NOT deleted").
		First(&script).Error

	if err != nil {
		logUtils.Errorf("find test script by id error", zap.String("error:", err.Error()))
		return
	}

	resp.TestScript = script

	return
}

func (r *TestScriptRepo) FindById(id uint) (script serverDomain.TestScriptResp, err error) {
	err = r.DB.Model(&model.TestScript{}).Where("id = ?", id).First(&script).Error
	if err != nil {
		logUtils.Errorf("find test script by id error", zap.String("error:", err.Error()))
		return script, err
	}

	return script, nil
}

func (r *TestScriptRepo) FindByName(scriptName string, ids ...uint) (script serverDomain.TestScriptResp, err error) {
	db := r.DB.Model(&model.TestScript{}).Where("name = ?", scriptName)
	if len(ids) == 1 {
		db.Where("id != ?", ids[0])
	}
	err = db.First(&script).Error
	if err != nil {
		logUtils.Errorf("find test script by name error", zap.String("name:", scriptName), zap.Uints("ids:", ids), zap.String("error:", err.Error()))
		return script, err
	}

	return script, nil
}

func (r *TestScriptRepo) Create(req serverDomain.TestScriptReq) (id uint, err error) {
	if _, err := r.FindByName(req.Name); !errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, fmt.Errorf("%d", _domain.BizErrNameExist.Code)
	}
	script := req.TestScript

	err = r.DB.Model(&model.TestScript{}).Create(&script).Error
	if err != nil {
		logUtils.Errorf("add test script error", zap.String("error:", err.Error()))
		return 0, err
	}

	id = script.ID
	return
}

func (r *TestScriptRepo) Update(id uint, req serverDomain.TestScriptReq) (err error) {
	script := req.TestScript

	err = r.DB.Model(&model.TestScript{}).Where("id = ?", id).Updates(&script).Error
	if err != nil {
		logUtils.Errorf("update test script error", zap.String("error:", err.Error()))
		return err
	}

	return
}

func (r *TestScriptRepo) DeleteById(id uint) (err error) {
	err = r.DB.Model(&model.TestScript{}).Where("id = ?", id).
		Updates(map[string]interface{}{"deleted": true}).Error
	if err != nil {
		logUtils.Errorf("delete test script by id error", zap.String("error:", err.Error()))
		return
	}

	return
}
