package repo

import (
	"errors"
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type EndpointDocumentRepo struct {
	*BaseRepo              `inject:""`
	EndpointInterfaceRepo  *EndpointInterfaceRepo  `inject:""`
	ServeRepo              *ServeRepo              `inject:""`
	ProcessorInterfaceRepo *ProcessorInterfaceRepo `inject:""`
	ProjectRepo            *ProjectRepo            `inject:""`
}

func NewEndpointDocumentRepo() *EndpointDocumentRepo {
	return &EndpointDocumentRepo{}
}

func (r *EndpointDocumentRepo) ListByProject(projectId uint) (documents []model.EndpointDocument, err error) {
	err = r.DB.Model(&model.EndpointDocument{}).
		Where("project_id = ?", projectId).
		Find(&documents).Error

	return
}

func (r *EndpointDocumentRepo) GetByVersionAndProject(version string, projectId uint) (document model.EndpointDocument, err error) {
	err = r.DB.Model(&model.EndpointDocument{}).
		Where("version = ?", version).
		Where("project_id = ?", projectId).
		First(&document).Error

	return
}

func (r *EndpointDocumentRepo) Create(req v1.DocumentVersionReq, projectId uint) (id uint, err error) {
	document := model.EndpointDocument{
		Name:      req.Name,
		Version:   req.Version,
		ProjectId: projectId,
	}
	err = r.DB.Model(&model.EndpointDocument{}).Create(&document).Error
	if err != nil {
		logUtils.Errorf("add endpoint document error", zap.String("error:", err.Error()))
		return
	}

	id = document.ID
	return
}

func (r *EndpointDocumentRepo) GetIdByVersionAndProject(req v1.DocumentVersionReq, projectId uint) (id uint, err error) {
	document, err := r.GetByVersionAndProject(req.Version, projectId)
	if err != nil && err != gorm.ErrRecordNotFound {
		return
	}

	if document.ID != 0 {
		id = document.ID
		return id, nil
	}
	id, err = r.Create(req, projectId)

	return
}
func (r *EndpointDocumentRepo) Update(req v1.UpdateDocumentVersionReq) (err error) {
	updateFields := map[string]interface{}{}
	if req.Name != "" {
		updateFields["name"] = req.Name
	}
	if req.Version != "" {
		updateFields["version"] = req.Version
	}
	if len(updateFields) == 0 {
		return errors.New("update field can't be empty")
	}

	err = r.DB.Model(&model.EndpointDocument{}).Where("id = ?", req.Id).Updates(updateFields).Error
	if err != nil {
		logUtils.Errorf("update endpoint document error", zap.String("error:", err.Error()))
		return err
	}

	return
}

func (r *EndpointDocumentRepo) DeleteById(id uint) (err error) {
	err = r.DB.
		Where("id = ?", id).
		Delete(&model.EndpointDocument{}).Error

	if err != nil {
		logUtils.Errorf("delete endpoint document by id error", zap.String("error:", err.Error()))
		return
	}

	return
}
