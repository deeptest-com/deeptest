package repo

import (
	serverDomain "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"gorm.io/gorm"
	"strconv"
)

type EndpointCaseRepo struct {
	*BaseRepo `inject:""`
	DB        *gorm.DB `inject:""`

	EndpointRepo       *EndpointRepo       `inject:""`
	DebugInterfaceRepo *DebugInterfaceRepo `inject:""`
	ProjectRepo        *ProjectRepo        `inject:""`
}

func (r *EndpointCaseRepo) List(endpointId uint) (pos []model.EndpointCase, err error) {
	err = r.DB.
		Where("endpoint_id=?", endpointId).
		Where("NOT deleted").
		Find(&pos).Error

	return
}

func (r *EndpointCaseRepo) Get(id uint) (po model.EndpointCase, err error) {
	err = r.DB.Where("id = ?", id).First(&po).Error
	return
}

func (r *EndpointCaseRepo) GetDetail(caseId uint) (endpointCase model.EndpointCase, err error) {
	if caseId <= 0 {
		return
	}

	endpointCase, err = r.Get(caseId)

	debugInterface, _ := r.DebugInterfaceRepo.Get(endpointCase.DebugInterfaceId)

	debugData, _ := r.DebugInterfaceRepo.GetDetail(debugInterface.ID)
	endpointCase.DebugData = &debugData

	return
}

func (r *EndpointCaseRepo) Save(po *model.EndpointCase) (err error) {
	err = r.DB.Save(po).Error

	err = r.UpdateSerialNumber(po.ID, po.ProjectId)

	return
}

func (r *EndpointCaseRepo) UpdateName(req serverDomain.EndpointCaseSaveReq) (err error) {
	err = r.DB.Model(&model.EndpointCase{}).
		Where("id=?", req.ID).
		Update("name", req.Name).Error

	return
}

func (r *EndpointCaseRepo) Remove(id uint) (err error) {
	err = r.DB.Model(&model.EndpointCase{}).
		Where("id = ?", id).
		Update("deleted", true).Error

	return
}

func (r *EndpointCaseRepo) SaveDebugData(interf *model.EndpointCase) (err error) {
	r.DB.Transaction(func(tx *gorm.DB) error {
		err = r.UpdateDebugInfo(interf)
		if err != nil {
			return err
		}

		// TODO: save debug data

		return err
	})

	return
}

func (r *EndpointCaseRepo) UpdateDebugInfo(interf *model.EndpointCase) (err error) {
	values := map[string]interface{}{
		"server_id": interf.DebugData.ServerId,
		"base_url":  interf.DebugData.BaseUrl,
		"url":       interf.DebugData.Url,
		"method":    interf.DebugData.Method,
	}

	err = r.DB.Model(&model.EndpointCase{}).
		Where("id=?", interf.ID).
		Updates(values).
		Error

	return
}

func (r *EndpointCaseRepo) UpdateSerialNumber(id, projectId uint) (err error) {
	var project model.Project
	project, err = r.ProjectRepo.Get(projectId)
	if err != nil {
		return
	}

	err = r.DB.Model(&model.EndpointCase{}).
		Where("id=?", id).
		Update("serial_number", project.ShortName+"-TS-"+strconv.Itoa(int(id))).Error
	return
}
