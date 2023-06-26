package repo

import (
	"fmt"
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/core/dao"
	model "github.com/aaronchen2k/deeptest/internal/server/modules/model"
	_domain "github.com/aaronchen2k/deeptest/pkg/domain"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"gorm.io/gorm"
)

type DatapoolRepo struct {
	DB       *gorm.DB  `inject:""`
	UserRepo *UserRepo `inject:""`
}

func (r *DatapoolRepo) Paginate(req v1.DatapoolReqPaginate) (ret _domain.PageData, err error) {
	var count int64
	db := r.DB.Model(&model.Datapool{}).Where("project_id = ? AND NOT deleted", req.ProjectId)

	if req.Name != "" {
		db = db.Where("name LIKE ?", fmt.Sprintf("%%%s%%", req.Name))
	}

	err = db.Count(&count).Error
	if err != nil {
		logUtils.Errorf("count report error %s", err.Error())
		return
	}

	results := make([]*model.Datapool, 0)
	err = db.Scopes(dao.PaginateScope(req.Page, req.PageSize, req.Order, req.Field)).Find(&results).Error
	if err != nil {
		logUtils.Errorf("query report error %s", err.Error())
		return
	}

	ret.Populate(results, count, req.Page, req.PageSize)

	return
}

func (r *DatapoolRepo) ListForExec(projectId uint) (ret []v1.DatapoolReq, err error) {
	var pos []model.Datapool
	err = r.DB.Model(&model.Datapool{}).
		Where("project_id = ? AND NOT deleted", projectId).
		Find(&pos).Error

	for _, po := range pos {
		to := v1.DatapoolReq{
			Model: _domain.Model{
				Id: po.ID,
			},
			Name: po.Name,
			Desc: po.Desc,
			Data: po.Data,
		}

		ret = append(ret, to)
	}

	return
}

func (r *DatapoolRepo) Get(id uint) (project model.Datapool, err error) {
	err = r.DB.Model(&model.Datapool{}).
		Where("id = ?", id).First(&project).Error

	return
}

func (r *DatapoolRepo) GetByName(name string) (po model.Datapool, err error) {
	err = r.DB.Model(&model.Datapool{}).
		Where("name = ?", name).
		First(&po).Error

	return
}

func (r *DatapoolRepo) Save(po *model.Datapool, userId uint) (err error) {
	user, _ := r.UserRepo.FindById(userId)
	if po.CreateUser == "" {
		po.CreateUser = user.Username
	}

	err = r.DB.Save(po).Error

	return
}

func (r *DatapoolRepo) SaveData(req v1.DatapoolReq) (err error) {
	err = r.DB.Model(&model.Datapool{}).
		Where("id = ?", req.Id).
		Updates(map[string]interface{}{"data": req.Data}).Error

	return nil
}

func (r *DatapoolRepo) Delete(id uint) (err error) {
	err = r.DB.Model(&model.Datapool{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{"deleted": true}).Error

	return
}

func (r *DatapoolRepo) Disable(id uint) (err error) {
	err = r.DB.Model(&model.Datapool{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{"disabled": gorm.Expr("NOT disabled")}).Error

	return
}
