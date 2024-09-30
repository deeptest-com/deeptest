package repo

import (
	"fmt"
	serverDomain "github.com/deeptest-com/deeptest/cmd/server/v1/domain"
	"github.com/deeptest-com/deeptest/internal/pkg/consts"
	"github.com/deeptest-com/deeptest/internal/server/core/dao"
	"github.com/deeptest-com/deeptest/internal/server/modules/model"
	_domain "github.com/deeptest-com/deeptest/pkg/domain"
	_commUtils "github.com/deeptest-com/deeptest/pkg/lib/comm"
	logUtils "github.com/deeptest-com/deeptest/pkg/lib/log"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"strconv"
)

type EndpointCaseRepo struct {
	*BaseRepo `inject:""`
	DB        *gorm.DB `inject:""`

	EndpointRepo       *EndpointRepo       `inject:""`
	DebugInterfaceRepo *DebugInterfaceRepo `inject:""`
	ProjectRepo        *ProjectRepo        `inject:""`
	CategoryRepo       *CategoryRepo       `inject:""`
}

func (r *EndpointCaseRepo) Paginate(tenantId consts.TenantId, req serverDomain.EndpointCaseReqPaginate) (data _domain.PageData, err error) {
	var count int64

	db := r.GetDB(tenantId).Model(&model.EndpointCase{}).
		Where("endpoint_id = ? AND NOT deleted", req.EndpointId).
		Where("case_type != ?", consts.CaseAlternative)

	if req.Keywords != "" {
		db = db.Where("name LIKE ?", fmt.Sprintf("%%%s%%", req.Keywords))
	}
	if req.Enabled != "" {
		db = db.Where("disabled = ?", _commUtils.IsDisable(req.Enabled))
	}

	err = db.Count(&count).Error
	if err != nil {
		logUtils.Errorf("count project error", zap.String("error:", err.Error()))
		return
	}

	cases := make([]*model.EndpointCase, 0)
	err = db.
		Scopes(dao.PaginateScope(req.Page, req.PageSize, req.Order, req.Field)).
		Find(&cases).Error
	if err != nil {
		logUtils.Errorf("query case error", zap.String("error:", err.Error()))
		return
	}

	r.MountChildrenForBenchmark(tenantId, cases)
	data.Populate(cases, count, req.Page, req.PageSize)

	return
}

func (r *EndpointCaseRepo) MountChildrenForBenchmark(tenantId consts.TenantId, cases []*model.EndpointCase) {
	for _, v := range cases {
		children, err := r.ListByCaseTypeAndBaseCase(tenantId, consts.CaseAlternative, v.ID)
		if err != nil {
			continue
		}

		v.Children = children
	}
}

func (r *EndpointCaseRepo) ListByCaseTypeAndBaseCase(tenantId consts.TenantId, caseType consts.CaseType, baseCase uint) (cases []model.EndpointCase, err error) {
	err = r.GetDB(tenantId).
		Where("case_type=?", caseType).
		Where("base_case=?", baseCase).
		Where("NOT deleted").Order("created_at desc").
		Find(&cases).Error

	return

}
func (r *EndpointCaseRepo) List(tenantId consts.TenantId, endpointId uint) (pos []model.EndpointCase, err error) {
	err = r.GetDB(tenantId).
		Where("endpoint_id=?", endpointId).
		Where("NOT deleted").Order("created_at desc").
		Find(&pos).Error

	return
}

func (r *EndpointCaseRepo) Get(tenantId consts.TenantId, id uint) (po model.EndpointCase, err error) {
	err = r.GetDB(tenantId).Where("id = ?", id).First(&po).Error
	return
}

func (r *EndpointCaseRepo) GetDetail(tenantId consts.TenantId, caseId uint) (endpointCase model.EndpointCase, err error) {
	if caseId <= 0 {
		return
	}

	endpointCase, err = r.Get(tenantId, caseId)

	debugInterface, _ := r.DebugInterfaceRepo.Get(tenantId, endpointCase.DebugInterfaceId)

	debugData, _ := r.DebugInterfaceRepo.GetDetail(tenantId, debugInterface.ID)
	endpointCase.DebugData = &debugData

	return
}

func (r *EndpointCaseRepo) Save(tenantId consts.TenantId, po *model.EndpointCase) (err error) {
	err = r.GetDB(tenantId).Save(po).Error

	err = r.UpdateSerialNumber(tenantId, po.ID, po.ProjectId)

	return
}

func (r *EndpointCaseRepo) UpdateName(tenantId consts.TenantId, req serverDomain.EndpointCaseSaveReq) (err error) {
	err = r.GetDB(tenantId).Model(&model.EndpointCase{}).
		Where("id=?", req.ID).
		Update("name", req.Name).Error

	return
}

func (r *EndpointCaseRepo) Remove(tenantId consts.TenantId, id uint) (err error) {
	err = r.GetDB(tenantId).Model(&model.EndpointCase{}).
		Where("id = ?", id).
		Update("deleted", true).Error

	return
}

func (r *EndpointCaseRepo) SaveDebugData(tenantId consts.TenantId, interf *model.EndpointCase) (err error) {
	r.GetDB(tenantId).Transaction(func(tx *gorm.DB) error {
		err = r.UpdateDebugInfo(tenantId, interf)
		if err != nil {
			return err
		}

		// TODO: save debug data

		return err
	})

	return
}

func (r *EndpointCaseRepo) UpdateDebugInfo(tenantId consts.TenantId, interf *model.EndpointCase) (err error) {
	values := map[string]interface{}{
		"server_id": interf.DebugData.ServerId,
		"base_url":  interf.DebugData.BaseUrl,
		"url":       interf.DebugData.Url,
		"method":    interf.DebugData.Method,
	}

	err = r.GetDB(tenantId).Model(&model.EndpointCase{}).
		Where("id=?", interf.ID).
		Updates(values).
		Error

	return
}

func (r *EndpointCaseRepo) UpdateInfo(tenantId consts.TenantId, id uint, values map[string]interface{}) (err error) {
	err = r.GetDB(tenantId).Model(&model.EndpointCase{}).
		Where("id=?", id).
		Updates(values).
		Error

	return
}

func (r *EndpointCaseRepo) UpdateSerialNumber(tenantId consts.TenantId, id, projectId uint) (err error) {
	var project model.Project
	project, err = r.ProjectRepo.Get(tenantId, projectId)
	if err != nil {
		return
	}

	err = r.GetDB(tenantId).Model(&model.EndpointCase{}).
		Where("id=?", id).
		Update("serial_number", project.ShortName+"-TC-"+strconv.Itoa(int(id))).Error
	return
}

func (r *EndpointCaseRepo) ListByProjectIdAndServeId(tenantId consts.TenantId, projectId, serveId uint) (endpointCases []*serverDomain.InterfaceCase, err error) {
	err = r.GetDB(tenantId).Model(&model.EndpointCase{}).
		Joins("left join biz_debug_interface d on biz_endpoint_case.debug_interface_id=d.id").
		Select("biz_endpoint_case.*, d.method as method").
		Where("biz_endpoint_case.project_id = ? and biz_endpoint_case.serve_id = ? and processor_interface_src = '' and not biz_endpoint_case.deleted and not biz_endpoint_case.disabled", projectId, serveId).
		Find(&endpointCases).Error
	//err = r.GetDB(tenantId).Where("project_id = ? and serve_id = ? and not deleted and not disabled", projectId, serveId).Find(&endpointCases).Error
	return
}

func (r *EndpointCaseRepo) GetEndpointCount(tenantId consts.TenantId, projectId uint, serveIds consts.Integers) (result []serverDomain.EndpointCount, err error) {
	err = r.GetDB(tenantId).Raw("select count(id) count,endpoint_id from "+model.EndpointCase{}.TableName()+" where not deleted and not disabled and project_id=? and serve_id in ? group by endpoint_id", projectId, serveIds).Scan(&result).Error
	return
}

func (r *EndpointCaseRepo) GetCategoryEndpointCase(tenantId consts.TenantId, projectId uint, serveIds consts.Integers) (result []serverDomain.CategoryEndpointCase, err error) {

	sql := fmt.Sprintf("select concat('case_',ec.id) as case_unique_id,concat('endpoint_',e.id) as endpoint_unique_id,ec.id as case_id,ec.name as case_name,i.method,ec.`desc` as case_desc,ec.endpoint_id as case_endpoint_id,ec.debug_interface_id as case_debug_interface_id,ec.project_id,ec.serve_id,e.id as endpoint_id,e.title as endpoint_title,e.description as endpoint_description,e.category_id as category_id from biz_endpoint_case ec left join biz_endpoint e on ec.endpoint_id=e.id left join biz_debug_interface i on ec.debug_interface_id=i.id Where ec.project_id= %d and not e.deleted and not ec.deleted", projectId)
	if len(serveIds) != 0 {
		sql = fmt.Sprintf("%s and i.serve_id in (%s)", sql, serveIds.ToString())
	}
	err = r.GetDB(tenantId).Raw(sql).Scan(&result).Error
	return
}

func (r *EndpointCaseRepo) UpdateDebugInterfaceId(tenantId consts.TenantId, debugInterfaceId, id uint) (err error) {
	err = r.GetDB(tenantId).Model(&model.EndpointCase{}).
		Where("id=?", id).
		Update("debug_interface_id", debugInterfaceId).Error

	return
}

func (r *EndpointCaseRepo) ListByCaseType(tenantId consts.TenantId, endpointId uint, caseTypes []consts.CaseType) (pos []model.EndpointCase, err error) {
	err = r.GetDB(tenantId).
		Where("endpoint_id=?", endpointId).
		Where("case_type in ?", caseTypes).
		Where("NOT deleted").Order("created_at desc").
		Find(&pos).Error

	return
}
