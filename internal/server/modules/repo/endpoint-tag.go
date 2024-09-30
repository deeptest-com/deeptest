package repo

import (
	"github.com/deeptest-com/deeptest/internal/pkg/consts"
	"github.com/deeptest-com/deeptest/internal/server/modules/model"
	logUtils "github.com/deeptest-com/deeptest/pkg/lib/log"
	"go.uber.org/zap"
)

type EndpointTagRepo struct {
	*BaseRepo             `inject:""`
	EndpointInterfaceRepo *EndpointInterfaceRepo `inject:""`
	ServeRepo             *ServeRepo             `inject:""`
	ProjectRepo           *ProjectRepo           `inject:""`
}

func NewEndpointTagRepo() *EndpointTagRepo {
	return &EndpointTagRepo{}
}

func (r *EndpointTagRepo) ListByProject(tenantId consts.TenantId, projectId uint) (tags []model.EndpointTag, err error) {
	err = r.GetDB(tenantId).Model(&model.EndpointTag{}).
		Where("project_id = ? AND NOT deleted AND NOT disabled", projectId).
		Find(&tags).Error

	return
}

func (r *EndpointTagRepo) Create(tenantId consts.TenantId, name string, projectId uint) (id uint, err error) {
	tag := model.EndpointTag{
		Name:      name,
		ProjectId: projectId,
	}
	err = r.GetDB(tenantId).Model(&model.EndpointTag{}).Create(&tag).Error
	if err != nil {
		logUtils.Errorf("add endpoint tag error", zap.String("error:", err.Error()))
		return
	}

	id = tag.ID
	return
}

func (r *EndpointTagRepo) BatchCreate(tenantId consts.TenantId, names []string, projectId uint) (err error) {
	tags := make([]model.EndpointTag, 0)
	for _, v := range names {
		tag := model.EndpointTag{
			Name:      v,
			ProjectId: projectId,
		}
		tags = append(tags, tag)
	}

	err = r.GetDB(tenantId).Model(&model.EndpointTag{}).
		Create(tags).Error

	if err != nil {
		logUtils.Errorf("batch add endpoint tag error", zap.String("error:", err.Error()))
		return
	}

	return
}

func (r *EndpointTagRepo) DeleteById(tenantId consts.TenantId, id uint) (err error) {
	err = r.GetDB(tenantId).
		Where("id = ?", id).
		Delete(&model.EndpointTag{}).Error

	if err != nil {
		logUtils.Errorf("delete endpoint tag by id error", zap.String("error:", err.Error()))
		return
	}

	return
}

func (r *EndpointTagRepo) BatchDeleteByIds(tenantId consts.TenantId, ids []uint) (err error) {
	err = r.GetDB(tenantId).
		Where("id IN (?)", ids).
		Delete(&model.EndpointTag{}).Error

	if err != nil {
		logUtils.Errorf("batch delete endpoint tag by ids error", zap.String("error:", err.Error()))
		return
	}

	return
}

func (r *EndpointTagRepo) BatchGetByName(tenantId consts.TenantId, names []string, projectId uint) (tags []model.EndpointTag, err error) {
	err = r.GetDB(tenantId).Model(&model.EndpointTag{}).
		Where("name IN (?)", names).
		Where("project_id = ? AND NOT deleted AND NOT disabled", projectId).
		Find(&tags).Error

	return
}

func (r *EndpointTagRepo) BatchGetIdsByName(tenantId consts.TenantId, names []string, projectId uint) (ids []uint, err error) {
	tags, err := r.BatchGetByName(tenantId, names, projectId)
	if err != nil {
		return
	}

	for _, tag := range tags {
		ids = append(ids, tag.ID)
	}

	return
}

func (r *EndpointTagRepo) BatchGetById(tenantId consts.TenantId, ids []string, projectId uint) (tags []model.EndpointTag, err error) {
	err = r.GetDB(tenantId).Model(&model.EndpointTag{}).
		Where("id IN (?)", ids).
		Where("project_id = ? AND NOT deleted AND NOT disabled", projectId).
		Find(&tags).Error

	return
}

func (r *EndpointTagRepo) ListRelByTagId(tenantId consts.TenantId, tagId uint) (rel []model.EndpointTagRel, err error) {
	err = r.GetDB(tenantId).Model(&model.EndpointTagRel{}).
		Where("tag_id = ?", tagId).
		Find(&rel).Error

	if err != nil {
		logUtils.Errorf("get endpoint tag relation by tag_id error", zap.String("error:", err.Error()))
		return
	}

	return
}

func (r *EndpointTagRepo) ListRelByEndpointId(tenantId consts.TenantId, endpointId uint) (rel []model.EndpointTagRel, err error) {
	err = r.GetDB(tenantId).Model(&model.EndpointTagRel{}).
		Where("endpoint_id = ?", endpointId).
		Find(&rel).Error

	if err != nil {
		logUtils.Errorf("get endpoint tag relation by endpoint_id error", zap.String("error:", err.Error()))
		return
	}

	return
}

func (r *EndpointTagRepo) GetTagIdsByEndpointId(tenantId consts.TenantId, endpointId uint) (tagIds []uint, err error) {
	err = r.GetDB(tenantId).Model(&model.EndpointTagRel{}).
		Where("endpoint_id = ?", endpointId).
		Select("tag_id").
		Find(&tagIds).Error

	if err != nil {
		logUtils.Errorf("get tag ids by endpoint_id error", zap.String("error:", err.Error()))
		return
	}

	return
}

func (r *EndpointTagRepo) DeleteRelByEndpointId(tenantId consts.TenantId, endpointId uint) (err error) {
	err = r.GetDB(tenantId).
		Where("endpoint_id = ?", endpointId).
		Delete(&model.EndpointTagRel{}).Error

	if err != nil {
		logUtils.Errorf("delete endpoint tag relation by endpoint_id error", zap.String("error:", err.Error()))
		return
	}

	return
}

func (r *EndpointTagRepo) AddRel(tenantId consts.TenantId, endpointId uint, tagIds []uint) (err error) {
	relations := make([]model.EndpointTagRel, 0)
	for _, v := range tagIds {
		relation := model.EndpointTagRel{
			EndpointId: endpointId,
			TagId:      v,
		}
		relations = append(relations, relation)
	}

	err = r.GetDB(tenantId).Model(&model.EndpointTagRel{}).
		Create(relations).Error

	if err != nil {
		logUtils.Errorf("batch add endpoint and tag relation error", zap.String("error:", err.Error()))
		return
	}

	return
}

func (r *EndpointTagRepo) GetEndpointIdsByTagNames(tenantId consts.TenantId, tagNames []string, projectId int64) (endpointIds []uint, err error) {
	//err = r.GetDB(tenantId).Model(&model.EndpointTagRel{}).
	//	Joins("LEFT JOIN biz_endpoint_tag t ON biz_endpoint_tag_rel.tag_id=t.id").
	//	Where("t.project_id = ?", projectId).
	//	Where("t.name IN (?) AND NOT t.deleted AND NOT t.disabled", tagNames).
	//	Select("biz_endpoint_tag_rel.endpoint_id").
	//	Find(&endpointIds).Error
	err = r.GetDB(tenantId).Model(&model.EndpointTagRel{}).
		Where("project_id = ?", projectId).
		Where("tag_name IN (?) AND NOT deleted AND NOT disabled", tagNames).
		Select("endpoint_id").
		Find(&endpointIds).Error

	return
}

func (r *EndpointTagRepo) GetTagNamesByEndpointId(tenantId consts.TenantId, endpointId, projectId uint) (tagNames []string, err error) {
	//err = r.GetDB(tenantId).Model(&model.EndpointTag{}).
	//	Joins("LEFT JOIN biz_endpoint_tag_rel l ON biz_endpoint_tag.id=l.tag_id").
	//	Where("l.endpoint_id = ?", endpointId).
	//	Where("biz_endpoint_tag.project_id = ? AND NOT biz_endpoint_tag.deleted AND NOT biz_endpoint_tag.disabled", projectId).
	//	Select("biz_endpoint_tag.name").
	//	Find(&tagNames).Error
	err = r.GetDB(tenantId).Model(&model.EndpointTagRel{}).
		Where("endpoint_id = ?", endpointId).
		Where("project_id = ? AND NOT deleted AND NOT disabled", projectId).
		Select("tag_name").
		Find(&tagNames).Error

	return
}

func (r *EndpointTagRepo) DeleteRelByEndpointAndProject(tenantId consts.TenantId, endpointId, projectId uint) (err error) {
	err = r.GetDB(tenantId).
		Where("endpoint_id = ?", endpointId).
		Where("project_id = ?", projectId).
		Delete(&model.EndpointTagRel{}).Error

	if err != nil {
		logUtils.Errorf("delete endpoint tag relation by endpoint and project error", zap.String("error:", err.Error()))
		return
	}

	return
}

func (r *EndpointTagRepo) BatchAddRel(tenantId consts.TenantId, endpointId, projectId uint, tagNames []string) (err error) {
	relations := make([]model.EndpointTagRel, 0)
	for _, v := range tagNames {
		relation := model.EndpointTagRel{
			EndpointId: endpointId,
			TagName:    v,
			ProjectId:  projectId,
		}
		relations = append(relations, relation)
	}

	err = r.GetDB(tenantId).Model(&model.EndpointTagRel{}).
		Create(relations).Error

	if err != nil {
		logUtils.Errorf("batch add endpoint and tag relation error", zap.String("error:", err.Error()))
		return
	}

	return
}

func (r *EndpointTagRepo) ListRelByProject(tenantId consts.TenantId, projectId uint) (tags []model.EndpointTagRel, err error) {
	err = r.GetDB(tenantId).Model(&model.EndpointTagRel{}).
		Distinct("tag_name").
		Where("project_id = ? AND NOT deleted AND NOT disabled", projectId).
		Find(&tags).Error

	return
}

func (r *EndpointTagRepo) BatchAddRelForTag(tagName string, endpointIds []uint, projectId uint) (err error) {
	relations := make([]model.EndpointTagRel, 0)
	for _, v := range endpointIds {
		relation := model.EndpointTagRel{
			EndpointId: v,
			TagName:    tagName,
			ProjectId:  projectId,
		}
		relations = append(relations, relation)
	}

	err = r.DB.Model(&model.EndpointTagRel{}).
		Create(relations).Error

	if err != nil {
		logUtils.Errorf("batch add tag relation for endpoint error", zap.String("error:", err.Error()))
		return
	}

	return
}

func (r *EndpointTagRepo) BatchGetEndpointIdsByTag(tagName string, endpointIds []uint, projectId uint) (res []uint, err error) {
	err = r.DB.Model(&model.EndpointTagRel{}).
		Where("tag_name = ?", tagName).
		Where("endpoint_id IN (?)", endpointIds).
		Where("project_id = ? AND NOT deleted AND NOT disabled", projectId).
		Select("endpoint_id").
		Find(&res).Error

	return
}
