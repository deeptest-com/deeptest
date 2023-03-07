package repo

import (
	"fmt"
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/core/dao"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	_domain "github.com/aaronchen2k/deeptest/pkg/domain"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"gorm.io/gorm"
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

func (r *ServeRepo) PaginateVersion(req v1.ServeVersionPaginate) (ret _domain.PageData, err error) {
	var count int64
	db := r.DB.Model(&model.ServeVersion{}).Where("serve_id = ? AND NOT deleted AND NOT disabled", req.ServeId)

	if req.Version != "" {
		db = db.Where("value LIKE ?", fmt.Sprintf("%%%s%%", req.Version))
	}

	if req.CreateUser != "" {
		db = db.Where("creat_user = ?", fmt.Sprintf("%s", req.CreateUser))
	}

	err = db.Count(&count).Error
	if err != nil {
		logUtils.Errorf("count report error %s", err.Error())
		return
	}

	results := make([]*model.ServeVersion, 0)

	err = db.Scopes(dao.PaginateScope(req.Page, req.PageSize, req.Order, req.Field)).Find(&results).Error
	if err != nil {
		logUtils.Errorf("query report error %s", err.Error())
		return
	}
	ret.Populate(results, count, req.Page, req.PageSize)

	return
}

func (r *ServeRepo) PaginateSchema(req v1.ServeSchemaPaginate) (ret _domain.PageData, err error) {
	var count int64
	db := r.DB.Model(&model.ComponentSchema{}).Where("serve_id = ? AND NOT deleted AND NOT disabled", req.ServeId)

	err = db.Count(&count).Error
	if err != nil {
		logUtils.Errorf("count report error %s", err.Error())
		return
	}

	results := make([]*model.ComponentSchema, 0)

	err = db.Scopes(dao.PaginateScope(req.Page, req.PageSize, req.Order, req.Field)).Find(&results).Error
	if err != nil {
		logUtils.Errorf("query report error %s", err.Error())
		return
	}
	ret.Populate(results, count, req.Page, req.PageSize)

	return
}

func (r *ServeRepo) Get(id uint) (res model.Serve, err error) {
	//err = r.DB.Where("NOT deleted AND not disabled").First(&res, id).Error
	err = r.DB.Where("NOT deleted").First(&res, id).Error
	return
}

func (r *ServeRepo) GetSchema(id uint) (res model.ComponentSchema, err error) {
	err = r.DB.Where("NOT deleted AND not disabled").First(&res, id).Error
	return
}

func (r *ServeRepo) GetSchemasByServeId(id uint) (res []model.ComponentSchema, err error) {
	err = r.DB.Where("NOT deleted AND not disabled").Find(&res, id).Error
	return
}

func (r *ServeRepo) DeleteById(id uint) error {
	return r.DB.Model(&model.Serve{}).Where("id = ?", id).Update("deleted", 1).Error
}

func (r *ServeRepo) DisableById(id uint) error {
	return r.DB.Model(&model.Serve{}).Where("id = ?", id).Update("status", 4).Error
}

func (r *ServeRepo) DeleteVersionById(id uint) error {
	return r.DB.Model(&model.ServeVersion{}).Where("id = ?", id).Update("deleted", 1).Error
}

func (r *ServeRepo) DisableVersionById(id uint) error {
	return r.DB.Model(&model.ServeVersion{}).Where("id = ?", id).Update("disabled", 1).Error
}

func (r *ServeRepo) ListServer(serveId uint) (res []model.ServeServer, err error) {
	err = r.DB.Where("serve_id = ? AND NOT deleted AND not disabled", serveId).Find(&res).Error
	return
}

func (r *ServeRepo) DeleteSchemaById(id uint) error {
	return r.DB.Model(&model.ComponentSchema{}).Where("id = ?", id).Update("deleted", 1).Error
}

func (r *ServeRepo) SaveServeEndpointVersions(versions []model.ServeEndpointVersion) error {
	return r.DB.Create(&versions).Error
}

func (r *ServeRepo) DeleteSaveServeEndpointVersions(serveId int64, version string) error {
	return r.DB.Delete(&model.ServeEndpointVersion{}, "serve_id=? and serve_version=?", serveId, version).Error
}

func (r *ServeRepo) BindEndpoint(serveId int64, serveVersion string, serveEndpointVersion []model.ServeEndpointVersion) (err error) {
	r.DB.Transaction(func(tx *gorm.DB) error {
		err = r.DeleteSaveServeEndpointVersions(serveId, serveVersion)
		if err != nil {
			return err
		}
		err = r.SaveServeEndpointVersions(serveEndpointVersion)
		if err != nil {
			return err
		}
		return nil
	})
	return
}
