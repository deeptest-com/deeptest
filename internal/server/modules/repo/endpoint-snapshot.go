package repo

import (
	"encoding/json"
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"go.uber.org/zap"
)

type EndpointSnapshotRepo struct {
	*BaseRepo              `inject:""`
	EndpointInterfaceRepo  *EndpointInterfaceRepo  `inject:""`
	ServeRepo              *ServeRepo              `inject:""`
	ProcessorInterfaceRepo *ProcessorInterfaceRepo `inject:""`
	ProjectRepo            *ProjectRepo            `inject:""`
	EndpointRepo           *EndpointRepo           `inject:""`
	EndpointDocumentRepo   *EndpointDocumentRepo   `inject:""`
}

func NewEndpointSnapshotRepo() *EndpointSnapshotRepo {
	return &EndpointSnapshotRepo{}
}

func (r *EndpointSnapshotRepo) BatchCreateSnapshot(req v1.DocumentVersionReq, projectId uint) (err error) {
	documentId, err := r.EndpointDocumentRepo.GetIdByVersionAndProject(req, projectId)
	if err != nil {
		return
	}

	if err = r.BatchDeleteByEndpointId(req.EndpointIds); err != nil {
		return
	}

	snapshots := make([]*model.EndpointSnapshot, 0)
	for _, v := range req.EndpointIds {
		endpoint, err := r.EndpointRepo.GetAll(v, "v0.1.0")
		if err != nil {
			logUtils.Errorf("create endpoint snapshot error", zap.String("error:", err.Error()), zap.Uint("endpointId:", v))
			continue
		}
		content, _ := json.Marshal(endpoint)

		snapshotTmp := model.EndpointSnapshot{
			EndpointId: endpoint.ID,
			DocumentId: documentId,
			Content:    string(content),
		}
		snapshots = append(snapshots, &snapshotTmp)
	}

	err = r.DB.Create(snapshots).Error

	return
}

func (r *EndpointSnapshotRepo) GetByDocumentId(documentId uint) (endpoints []*model.Endpoint, err error) {
	var snapshots []model.EndpointSnapshot
	err = r.DB.Where("document_id = ? and not deleted and not disabled", documentId).Find(&snapshots).Error
	if err != nil {
		return
	}

	for _, v := range snapshots {
		var endpoint model.Endpoint
		_ = json.Unmarshal([]byte(v.Content), &endpoint)
		endpoints = append(endpoints, &endpoint)
	}

	return
}

func (r *EndpointSnapshotRepo) DeleteById(id uint) (err error) {
	err = r.DB.
		Where("id = ?", id).
		Delete(&model.EndpointSnapshot{}).Error

	if err != nil {
		logUtils.Errorf("delete endpoint snapshot by id error", zap.String("error:", err.Error()))
		return
	}

	return
}

func (r *EndpointSnapshotRepo) BatchDeleteByEndpointId(endpointIds []uint) (err error) {
	err = r.DB.
		Where("endpoint_id IN (?)", endpointIds).
		Delete(&model.EndpointSnapshot{}).Error

	if err != nil {
		logUtils.Errorf("delete endpoint snapshot by id error", zap.String("error:", err.Error()))
		return
	}

	return
}

func (r *EndpointSnapshotRepo) UpdateContent(id uint, endpoint model.Endpoint) (err error) {
	content, _ := json.Marshal(endpoint)
	err = r.DB.Model(model.EndpointSnapshot{}).
		Where("id = ?", id).
		Update("content = ?", string(content)).Error

	return
}
