package repo

import (
	v1 "github.com/deeptest-com/deeptest/cmd/server/v1/domain"
	"github.com/deeptest-com/deeptest/internal/pkg/consts"
	"github.com/deeptest-com/deeptest/internal/server/modules/model"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type GrpcInterfaceRepo struct {
	DiagnoseInterfaceRepo *DiagnoseInterfaceRepo `inject:""`

	*BaseRepo `inject:""`
	DB        *gorm.DB `inject:""`
}

func (r *GrpcInterfaceRepo) Get(tenantId consts.TenantId, id uint) (po model.GrpcInterface, err error) {
	err = r.GetDB(tenantId).Where("id = ?", id).First(&po).Error
	return
}

func (r *GrpcInterfaceRepo) SaveDebugData(po model.GrpcInterface, tenantId consts.TenantId) (err error) {
	err = r.Save(tenantId, po.ID, &po)

	return
}

func (r *GrpcInterfaceRepo) Create(tenantId consts.TenantId, req v1.DiagnoseInterfaceSaveReq) (
	diagnoseInterface model.DiagnoseInterface, err error) {

	if req.ID == 0 { // create both Diagnose and gRPC interface
		grpcInterface := model.GrpcInterface{
			Name:      req.Title,
			ProjectId: req.ProjectId,
			CreatedBy: req.CreatedBy,
		}
		err = r.Save(tenantId, 0, &grpcInterface)

		copier.CopyWithOption(&diagnoseInterface, req, copier.Option{
			DeepCopy: true,
		})
		diagnoseInterface.DebugInterfaceId = grpcInterface.ID
		err = r.DiagnoseInterfaceRepo.Save(tenantId, &diagnoseInterface)

		err = r.UpdateDiagnoseInterfaceId(tenantId, grpcInterface.ID, diagnoseInterface.ID)

	} else { // update title
		diagnoseInterface, _ = r.DiagnoseInterfaceRepo.Get(tenantId, req.ID)
		diagnoseInterface.Title = req.Title

		err = r.DiagnoseInterfaceRepo.Save(tenantId, &diagnoseInterface)
	}

	return
}

func (r *GrpcInterfaceRepo) UpdateDiagnoseInterfaceId(tenantId consts.TenantId, grpcInterfaceId,
	diagnoseInterfaceId uint) (err error) {

	values := map[string]interface{}{
		"diagnose_interface_id": diagnoseInterfaceId,
	}

	err = r.GetDB(tenantId).Model(&model.GrpcInterface{}).
		Where("id=?", grpcInterfaceId).
		Updates(values).
		Error

	return
}
