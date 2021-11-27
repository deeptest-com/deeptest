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

type ProjectRepo struct {
	DB       *gorm.DB  `inject:""`
	RoleRepo *RoleRepo `inject:""`
}

func NewProjectRepo() *ProjectRepo {
	return &ProjectRepo{}
}

func (r *ProjectRepo) Paginate(req serverDomain.ProjectReqPaginate) (data domain.PageData, err error) {
	var count int64

	db := r.DB.Model(&model.Project{}).
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
		logUtils.Errorf("count product error", zap.String("error:", err.Error()))
		return
	}

	projects := make([]*model.Project, 0)

	err = db.
		Scopes(dao.PaginateScope(req.Page, req.PageSize, req.Order, req.Field)).
		Find(&projects).Error
	if err != nil {
		logUtils.Errorf("query product error", zap.String("error:", err.Error()))
		return
	}

	data.Populate(projects, count, req.Page, req.PageSize)

	return
}

func (r *ProjectRepo) FindById(id uint) (serverDomain.ProjectResponse, error) {
	product := serverDomain.ProjectResponse{}
	err := r.DB.Model(&model.Project{}).Where("id = ?", id).First(&product).Error
	if err != nil {
		logUtils.Errorf("find product by id error", zap.String("error:", err.Error()))
		return product, err
	}

	return product, nil
}

func (r *ProjectRepo) FindByName(productname string, ids ...uint) (serverDomain.ProjectResponse, error) {
	product := serverDomain.ProjectResponse{}
	db := r.DB.Model(&model.Project{}).Where("name = ?", productname)
	if len(ids) == 1 {
		db.Where("id != ?", ids[0])
	}
	err := db.First(&product).Error
	if err != nil {
		logUtils.Errorf("find product by name error", zap.String("name:", productname), zap.Uints("ids:", ids), zap.String("error:", err.Error()))
		return product, err
	}

	return product, nil
}

func (r *ProjectRepo) Create(req serverDomain.ProjectRequest) (uint, error) {
	if _, err := r.FindByName(req.Name); !errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, fmt.Errorf("%d", domain.BizErrNameExist.Code)
	}
	product := req.Project

	err := r.DB.Model(&model.Project{}).Create(&product).Error
	if err != nil {
		logUtils.Errorf("add product error", zap.String("error:", err.Error()))
		return 0, err
	}

	return product.ID, nil
}

func (r *ProjectRepo) Update(id uint, req serverDomain.ProjectRequest) error {
	product := req.Project
	err := r.DB.Model(&model.Project{}).Where("id = ?", id).Updates(&product).Error
	if err != nil {
		logUtils.Errorf("update product error", zap.String("error:", err.Error()))
		return err
	}

	return nil
}

func (r *ProjectRepo) BatchDelete(id uint) (err error) {
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

func (r *ProjectRepo) DeleteById(id uint, tx *gorm.DB) (err error) {
	err = tx.Model(&model.Project{}).Where("id = ?", id).
		Updates(map[string]interface{}{"deleted": true}).Error
	if err != nil {
		logUtils.Errorf("delete product by id error", zap.String("error:", err.Error()))
		return
	}

	return
}

func (r *ProjectRepo) DeleteChildren(ids []int, tx *gorm.DB) (err error) {
	err = tx.Model(&model.Project{}).Where("id IN (?)", ids).
		Updates(map[string]interface{}{"deleted": true}).Error
	if err != nil {
		logUtils.Errorf("batch delete product error", zap.String("error:", err.Error()))
		return err
	}

	return nil
}

func (r *ProjectRepo) GetChildrenIds(id uint) (ids []int, err error) {
	tmpl := `
		WITH RECURSIVE product AS (
			SELECT * FROM biz_product WHERE id = %d
			UNION ALL
			SELECT child.* FROM biz_product child, product WHERE child.parent_id = product.id
		)
		SELECT id FROM product WHERE id != %d
    `
	sql := fmt.Sprintf(tmpl, id, id)
	err = r.DB.Raw(sql).Scan(&ids).Error
	if err != nil {
		logUtils.Errorf("get children product error", zap.String("error:", err.Error()))
		return
	}

	return
}
