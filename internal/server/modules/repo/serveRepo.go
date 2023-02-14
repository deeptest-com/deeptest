package repo

import (
	"fmt"
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/core/dao"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	_domain "github.com/aaronchen2k/deeptest/pkg/domain"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
)

type ServeRepo struct {
	*BaseRepo `inject:""`
}

func NewServeRepo() *ServeRepo {
	return &ServeRepo{}
}

func (r *ServeRepo) ListVersion(serveId uint) (res []model.ServeVersion, err error) {
	err = r.DB.Where("serve_id = ? AND NOT deleted AND not disabled", serveId).Find(&res).Error
	return
}

func (r *ServeRepo) Paginate(req v1.ServeReqPaginate) (ret _domain.PageData, err error) {
	var count int64
	db := r.DB.Model(&model.Serve{}).Where("project_id = ? AND NOT deleted AND NOT disabled", req.ProjectId)

	if req.Name != "" {
		db = db.Where("name LIKE ?", fmt.Sprintf("%%%s%%", req.Name))
	}

	err = db.Count(&count).Error
	if err != nil {
		logUtils.Errorf("count report error %s", err.Error())
		return
	}

	results := make([]*model.Serve, 0)

	err = db.Scopes(dao.PaginateScope(req.Page, req.PageSize, req.Order, req.Field)).Find(&results).Error
	if err != nil {
		logUtils.Errorf("query report error %s", err.Error())
		return
	}
	ret.Populate(results, count, req.Page, req.PageSize)

	return
}

func (r *ServeRepo) Get(id uint) (res model.Serve, err error) {
	err = r.DB.Where("NOT deleted AND not disabled").First(&res, id).Error
	return
}

func (r *ServeRepo) DeleteById(id uint) error {
	return r.DB.Model(&model.Serve{}).Where("id = ?", id).Update("deleted", 1).Error
}

func (r *ServeRepo) DisableById(id uint) error {
	return r.DB.Model(&model.Serve{}).Where("id = ?", id).Update("disabled", 1).Error
}

func (r *ServeRepo) DeleteVersionById(id uint) error {
	return r.DB.Model(&model.ServeVersion{}).Where("id = ?", id).Update("deleted", 1).Error
}

func (r *ServeRepo) DisableVersionById(id uint) error {
	return r.DB.Model(&model.ServeVersion{}).Where("id = ?", id).Update("disabled", 1).Error
}
