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

type EndpointRepo struct {
	*BaseRepo     `inject:""`
	InterfaceRepo *InterfaceRepo `inject:""`
}

func NewEndpointRepo() *EndpointRepo {
	return &EndpointRepo{}
}

func (r *EndpointRepo) Paginate(req v1.EndpointReqPaginate) (ret _domain.PageData, err error) {
	//fmt.Println(r.DB.Model(&model.SysUser{}))
	//err = r.DB.Where("id=?", id).Where("name=?", name).Find(&res).Error
	var count int64
	db := r.DB.Model(&model.Endpoint{}).Where("project_id = ? AND NOT deleted AND NOT disabled", req.ProjectId)

	if req.Title != "" {
		db = db.Where("title LIKE ?", fmt.Sprintf("%%%s%%", req.Title))
	}
	if req.UserId != 0 {
		db = db.Where("user_id = ?", req.UserId)
	}
	if req.Status != 0 {
		db = db.Where("status = ?", req.UserId)
	}

	err = db.Count(&count).Error
	if err != nil {
		logUtils.Errorf("count report error %s", err.Error())
		return
	}

	results := make([]*model.Endpoint, 0)

	err = db.Scopes(dao.PaginateScope(req.Page, req.PageSize, req.Order, req.Field)).Find(&results).Error
	if err != nil {
		logUtils.Errorf("query report error %s", err.Error())
		return
	}
	ret.Populate(results, count, req.Page, req.PageSize)

	return
}

func (r *EndpointRepo) SaveAll(endpoint *model.Endpoint) (err error) {
	r.DB.Transaction(func(tx *gorm.DB) error {
		//更新终端
		err = r.saveEndpoint(endpoint)
		if err != nil {
			return err
		}
		//保存路径参数
		err = r.saveEndpointParams(endpoint.ID, endpoint.PathParams)
		if err != nil {
			return err
		}
		//保存接口
		err = r.saveInterfaces(endpoint.ID, endpoint.Interfaces)
		if err != nil {
			return err
		}

		return nil
	})
	return
}

//保存终端信息
func (r *EndpointRepo) saveEndpoint(endpoint *model.Endpoint) (err error) {
	err = r.Save(endpoint.ID, endpoint)
	return
}

//保存路径参数
func (r *EndpointRepo) saveEndpointParams(endpointId uint, params []model.EndpointPathParam) (err error) {
	err = r.removeEndpointParams(endpointId)
	if err != nil {
		return
	}
	for _, item := range params {
		item.EndpointId = endpointId
		err = r.Save(item.ID, &item)
		if err != nil {
			return
		}
	}
	return
}

func (r *EndpointRepo) removeEndpointParams(endpointId uint) (err error) {
	err = r.DB.
		Where("endpoint_id = ?", endpointId).
		Delete(&model.EndpointPathParam{}, "").Error

	return
}

//保存接口信息
func (r *EndpointRepo) saveInterfaces(endpointId uint, interfaces []model.Interface) (err error) {
	err = r.removeInterfaces(endpointId)
	if err != nil {
		return
	}
	for _, item := range interfaces {
		item.EndpointId = endpointId
		err = r.InterfaceRepo.SaveInterfaces(item)
		if err != nil {
			return err
		}
	}
	return
}

func (r *EndpointRepo) removeInterfaces(endpointId uint) (err error) {
	err = r.DB.
		Where("endpoint_id = ?", endpointId).
		Delete(&model.Interface{}, "").Error

	return
}

func (r *EndpointRepo) GetAll(id uint) (endpoint model.Endpoint, err error) {
	endpoint, err = r.Get(id)
	if err != nil {
		return
	}
	endpoint.PathParams, _ = r.GetEndpointParams(id)
	endpoint.Interfaces, _ = r.InterfaceRepo.GetByEndpointId(id)

	return
}

func (r *EndpointRepo) Get(id uint) (res model.Endpoint, err error) {
	err = r.DB.First(&res, id).Error
	return
}

func (r *EndpointRepo) GetEndpointParams(endpointId uint) (pathParam []model.EndpointPathParam, err error) {
	err = r.DB.Find(&pathParam, "endpoint_id=?", endpointId).Error
	return
}

func (r *EndpointRepo) DeleteById(id uint) error {
	return r.DB.Model(&model.Endpoint{}).Where("id = ?", id).Update("deleted", 1).Error
}

func (r *EndpointRepo) DisableById(id uint) error {
	return r.DB.Model(&model.Endpoint{}).Where("id = ?", id).Update("status", 4).Error
}
