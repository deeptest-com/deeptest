package repo

import (
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
	db := r.DB.Model(&model.Endpoint{}).Where("project_id = ? AND NOT deleted", req.ProjectId)

	/*
		if req.Keywords != "" {
			db = db.Where("name LIKE ?", fmt.Sprintf("%%%s%%", req.Keywords))
		}
		if req.ScenarioId != 0 {
			db = db.Where("scenario_id = ?", req.ScenarioId)
		}
	*/

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

func (r *EndpointRepo) SaveAll(endpoint model.Endpoint) (err error) {
	r.DB.Transaction(func(tx *gorm.DB) error {
		//更新终端
		err = r.saveEndpoint(&endpoint)
		if err != nil {
			return err
		}
		//保存路径参数
		err = r.saveEndpointParams(endpoint.Id, endpoint.PathParams)
		if err != nil {
			return err
		}
		//保存接口
		err = r.saveInterfaces(endpoint.Id, endpoint.Interfaces)
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
func (r *EndpointRepo) saveEndpointParams(endpointId int64, params []model.EndpointPathParam) (err error) {
	for _, item := range params {
		item.EndpointId = endpointId
		err = r.Save(item.ID, &item)
		if err != nil {
			return
		}
	}
	return
}

//保存接口信息
func (r *EndpointRepo) saveInterfaces(endpointId int64, interfaces []model.Interface) (err error) {
	for _, item := range interfaces {
		item.EndpointId = endpointId
		err = r.InterfaceRepo.SaveInterface(item)
		//fmt.Println(item)
		if err != nil {
			return err
		}
	}
	return
}
