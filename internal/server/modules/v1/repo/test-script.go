package repo

import (
	"errors"
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
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

func (r *TestScriptRepo) Paginate(req serverDomain.TestScriptReqPaginate) (data domain.PageData, err error) {
	var count int64

	db := r.DB.Model(&model.TestScript{}).
		Where("NOT deleted")
	if req.Name != "" {
		db = db.Where("name LIKE ?", fmt.Sprintf("%s%%", req.Name))
	}
	if req.Category != "" {
		db = db.Where("category = ?", req.Category)
	}
	if req.Status != "" {
		db = db.Where("status = ?", req.Status)
	}

	err = db.Count(&count).Error
	if err != nil {
		logUtils.Errorf("count testScript error", zap.String("error:", err.Error()))
		return
	}

	TestScripts := make([]*model.TestScript, 0)

	err = db.
		Scopes(dao.PaginateScope(req.Page, req.PageSize, req.Order, req.Field)).
		Find(&TestScripts).Error
	if err != nil {
		logUtils.Errorf("query testScript error", zap.String("error:", err.Error()))
		return
	}

	data.Populate(TestScripts, count, req.Page, req.PageSize)

	return
}

func (r *TestScriptRepo) FindById(id uint) (serverDomain.TestScriptResponse, error) {
	testScript := serverDomain.TestScriptResponse{}
	err := r.DB.Model(&model.TestScript{}).Where("id = ?", id).First(&testScript).Error
	if err != nil {
		logUtils.Errorf("find testScript by id error", zap.String("error:", err.Error()))
		return testScript, err
	}

	return testScript, nil
}

func (r *TestScriptRepo) FindByName(testScriptname string, ids ...uint) (serverDomain.TestScriptResponse, error) {
	testScript := serverDomain.TestScriptResponse{}
	db := r.DB.Model(&model.TestScript{}).Where("name = ?", testScriptname)
	if len(ids) == 1 {
		db.Where("id != ?", ids[0])
	}
	err := db.First(&testScript).Error
	if err != nil {
		logUtils.Errorf("find testScript by name error", zap.String("name:", testScriptname), zap.Uints("ids:", ids), zap.String("error:", err.Error()))
		return testScript, err
	}

	return testScript, nil
}

func (r *TestScriptRepo) Create(req serverDomain.TestScriptRequest) (uint, error) {
	if _, err := r.FindByName(req.Name); !errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, fmt.Errorf("%d", domain.BizErrNameExist.Code)
	}
	testScript := req.TestScript

	err := r.DB.Model(&model.TestScript{}).Create(&testScript).Error
	if err != nil {
		logUtils.Errorf("add testScript error", zap.String("error:", err.Error()))
		return 0, err
	}

	return testScript.ID, nil
}

func (r *TestScriptRepo) Update(id uint, req serverDomain.TestScriptRequest) error {
	testScript := req.TestScript
	err := r.DB.Model(&model.TestScript{}).Where("id = ?", id).Updates(&testScript).Error
	if err != nil {
		logUtils.Errorf("update testScript error", zap.String("error:", err.Error()))
		return err
	}

	return nil
}

func (r *TestScriptRepo) BatchDelete(id uint) (err error) {
	ids, err := r.GetChildrenIds(id)
	if err != nil {
		return err
	}

	r.DB.Transaction(func(tx *gorm.DB) (err error) {
		err = r.DeleteChildren(ids, tx)
		if err != nil {
			return
		}

		err = r.DeleteById(id, tx)
		if err != nil {
			return
		}

		return
	})

	return
}

func (r *TestScriptRepo) DeleteById(id uint, tx *gorm.DB) (err error) {
	err = tx.Model(&model.TestScript{}).Where("id = ?", id).
		Updates(map[string]interface{}{"deleted": true}).Error
	if err != nil {
		logUtils.Errorf("delete testScript by id error", zap.String("error:", err.Error()))
		return
	}

	return
}

func (r *TestScriptRepo) DeleteChildren(ids []int, tx *gorm.DB) (err error) {
	err = tx.Model(&model.TestScript{}).Where("id IN (?)", ids).
		Updates(map[string]interface{}{"deleted": true}).Error
	if err != nil {
		logUtils.Errorf("batch delete testScript error", zap.String("error:", err.Error()))
		return err
	}

	return nil
}

func (r *TestScriptRepo) GetChildrenIds(id uint) (ids []int, err error) {
	tmpl := `
		WITH RECURSIVE testScript AS (
			SELECT * FROM biz_testScript WHERE id = %d
			UNION ALL
			SELECT child.* FROM biz_testScript child, testScript WHERE child.parent_id = testScript.id
		)
		SELECT id FROM testScript WHERE id != %d
    `
	sql := fmt.Sprintf(tmpl, id, id)
	err = r.DB.Raw(sql).Scan(&ids).Error
	if err != nil {
		logUtils.Errorf("get children testScript error", zap.String("error:", err.Error()))
		return
	}

	return
}
