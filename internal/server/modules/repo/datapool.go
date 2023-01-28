package repo

import (
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	model "github.com/aaronchen2k/deeptest/internal/server/modules/model"
	_domain "github.com/aaronchen2k/deeptest/pkg/domain"
	"gorm.io/gorm"
)

type DatapoolRepo struct {
	DB *gorm.DB `inject:""`
}

func NewDatapoolRepo() *DatapoolRepo {
	return &DatapoolRepo{}
}

func (r *DatapoolRepo) List(projectId uint) (ret []v1.DatapoolReq, err error) {
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

func (r *DatapoolRepo) Save(po *model.Datapool) (err error) {
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
