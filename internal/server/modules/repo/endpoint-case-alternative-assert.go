package repo

import (
	"fmt"
	serverDomain "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	checkpointHelpper "github.com/aaronchen2k/deeptest/internal/pkg/helper/checkpoint"
	model "github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"gorm.io/gorm"
)

type EndpointCaseAlternativeAssertRepo struct {
	DB *gorm.DB `inject:""`
}

func (r *EndpointCaseAlternativeAssertRepo) List(alternativeCaseId uint) (pos []model.EndpointCaseAlternativeAssert, err error) {
	db := r.DB.Where("NOT deleted").
		Where("alternative_case_id=?", alternativeCaseId).
		Order("ordr ASC")

	err = db.Find(&pos).Error

	return
}

func (r *EndpointCaseAlternativeAssertRepo) Get(id uint) (po model.EndpointCaseAlternativeAssert, err error) {
	err = r.DB.
		Where("id=?", id).
		Where("NOT deleted").
		First(&po).Error
	return
}

func (r *EndpointCaseAlternativeAssertRepo) Save(po *model.EndpointCaseAlternativeAssert) (err error) {
	if po.Type == "" {
		po.Type = consts.ResponseStatus
		po.Operator = consts.Equal
	}

	po.Desc = checkpointHelpper.GenDesc(po.Type, po.Operator, po.Value, po.Expression, po.ExtractorVariable)

	err = r.DB.Save(po).Error
	return
}

func (r *EndpointCaseAlternativeAssertRepo) Delete(id uint) (err error) {
	err = r.DB.Model(&model.EndpointCaseAlternativeAssert{}).
		Where("id=?", id).
		Update("deleted", true).
		Error

	return
}

func (r *EndpointCaseAlternativeAssertRepo) Disable(id uint) (err error) {
	err = r.DB.Model(&model.EndpointCaseAlternativeAssert{}).
		Where("id=?", id).
		Update("disabled", gorm.Expr("NOT disabled")).
		Error

	return
}

func (r *EndpointCaseAlternativeAssertRepo) UpdateOrders(req serverDomain.ConditionMoveReq) (err error) {
	return r.DB.Transaction(func(tx *gorm.DB) error {
		for index, id := range req.Data {
			sql := fmt.Sprintf("UPDATE %s SET ordr = %d WHERE id = %d",
				model.EndpointCaseAlternativeAssert{}.TableName(), index+1, id)

			err = r.DB.Exec(sql).Error
			if err != nil {
				return err
			}
		}

		return nil
	})
}
