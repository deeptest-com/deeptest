package repo

import (
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type WebsocketInterfaceRepo struct {
	DiagnoseInterfaceRepo *DiagnoseInterfaceRepo `inject:""`

	*BaseRepo `inject:""`
	DB        *gorm.DB `inject:""`
}

func (r *WebsocketInterfaceRepo) Get(tenantId consts.TenantId, id uint) (po model.WebsocketInterface, err error) {
	err = r.GetDB(tenantId).Where("id = ?", id).First(&po).Error

	return
}
func (r *WebsocketInterfaceRepo) SaveDebugData(po model.WebsocketInterface, tenantId consts.TenantId) (err error) {
	err = r.Save(tenantId, po.ID, &po)

	return
}

func (r *WebsocketInterfaceRepo) Create(tenantId consts.TenantId, req v1.DiagnoseInterfaceSaveReq) (
	diagnoseInterface model.DiagnoseInterface, err error) {

	if req.ID == 0 { // create both Diagnose and Websocket interface
		websocketInterface := model.WebsocketInterface{
			Name:      req.Title,
			ProjectId: req.ProjectId,
			CreatedBy: req.CreatedBy,
		}
		err = r.Save(tenantId, 0, &websocketInterface)

		copier.CopyWithOption(&diagnoseInterface, req, copier.Option{
			DeepCopy: true,
		})
		diagnoseInterface.DebugInterfaceId = websocketInterface.ID
		err = r.DiagnoseInterfaceRepo.Save(tenantId, &diagnoseInterface)

		err = r.UpdateDiagnoseInterfaceId(tenantId, websocketInterface.ID, diagnoseInterface.ID)

	} else { // update title
		diagnoseInterface, _ = r.DiagnoseInterfaceRepo.Get(tenantId, req.ID)
		diagnoseInterface.Title = req.Title

		err = r.DiagnoseInterfaceRepo.Save(tenantId, &diagnoseInterface)
	}

	return
}

func (r *WebsocketInterfaceRepo) UpdateDiagnoseInterfaceId(tenantId consts.TenantId, websocketInterfaceId,
	diagnoseInterfaceId uint) (err error) {

	values := map[string]interface{}{
		"diagnose_interface_id": diagnoseInterfaceId,
	}

	err = r.GetDB(tenantId).Model(&model.WebsocketInterface{}).
		Where("id=?", websocketInterfaceId).
		Updates(values).
		Error

	return
}

func (r *WebsocketInterfaceRepo) ListParams(tenantId consts.TenantId, interfaceId uint) (
	params []model.WebsocketInterfaceParam, err error) {

	pos := []model.WebsocketInterfaceParam{}

	err = r.GetDB(tenantId).
		Where("interface_id=?", interfaceId).
		Where("NOT deleted").
		Order("id ASC").
		Find(&pos).Error

	for _, po := range pos {
		params = append(params, po)
	}

	return
}
func (r *WebsocketInterfaceRepo) ListHeaders(tenantId consts.TenantId, interfaceId uint) (
	pos []model.WebsocketInterfaceHeader, err error) {

	err = r.GetDB(tenantId).
		Where("interface_id=?", interfaceId).
		Where("NOT deleted").
		Order("id ASC").
		Find(&pos).Error

	return
}

func (r *WebsocketInterfaceRepo) SaveParams(params *[]domain.Param, id uint, tenantId consts.TenantId) (err error) {
	err = r.RemoveParams(tenantId, id)

	if params == nil || len(*params) == 0 {
		return
	}

	var pos []model.WebsocketInterfaceParam

	for _, p := range *params {
		if p.Name == "" {
			continue
		}

		po := model.WebsocketInterfaceParam{
			InterfaceParamBase: model.InterfaceParamBase{
				Name:        p.Name,
				Value:       p.Value,
				Type:        p.Type,
				InterfaceId: id,
			},
		}

		pos = append(pos, po)
	}

	err = r.GetDB(tenantId).Create(&pos).Error

	return
}

func (r *WebsocketInterfaceRepo) RemoveParams(tenantId consts.TenantId, id uint) (err error) {
	err = r.GetDB(tenantId).
		Where("interface_id = ?", id).
		Delete(&model.WebsocketInterfaceParam{}, "").Error

	return
}

func (r *WebsocketInterfaceRepo) SaveHeaders(headers *[]domain.Header, id uint, tenantId consts.TenantId) (err error) {
	err = r.RemoveHeaders(tenantId, id)

	if len(*headers) == 0 {
		return
	}

	var pos []model.WebsocketInterfaceHeader

	for _, p := range *headers {
		if p.Name == "" {
			continue
		}

		po := model.WebsocketInterfaceHeader{
			InterfaceHeaderBase: model.InterfaceHeaderBase{
				Name:        p.Name,
				Value:       p.Value,
				Type:        p.Type,
				InterfaceId: id,
			},
		}

		pos = append(pos, po)
	}

	err = r.GetDB(tenantId).Create(&pos).Error

	return
}

func (r *WebsocketInterfaceRepo) RemoveHeaders(tenantId consts.TenantId, id uint) (err error) {
	err = r.GetDB(tenantId).
		Where("interface_id = ?", id).
		Delete(&model.WebsocketInterfaceHeader{}, "").Error

	return
}
