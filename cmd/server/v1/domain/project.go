package serverDomain

import (
	integrationDomain "github.com/deeptest-com/deeptest/integration/domain"
	"github.com/deeptest-com/deeptest/internal/pkg/consts"
	serverConsts "github.com/deeptest-com/deeptest/internal/server/consts"
	"github.com/deeptest-com/deeptest/pkg/domain"
)

type ProjectReq struct {
	_domain.Model
	ProjectBase
	integrationDomain.ProjectReq
}

type ProjectReqPaginate struct {
	_domain.PaginateReq
	Keywords string `json:"keywords"`
	Enabled  string `json:"enabled"`
}

type ProjectResp struct {
	_domain.PaginateReq
	ProjectBase
}

type ProjectMemberRemoveReq struct {
	UserId    int `json:"userId"`
	ProjectId int `json:"projectId"`
}

type ProjectBase struct {
	Name string                   `json:"name"`
	Desc string                   `json:"desc" gorm:"column:descr;type:text"`
	Type serverConsts.ProjectType `json:"type"`

	SchemaId       uint                       `json:"schemaId"`
	OrgId          uint                       `json:"orgId"`
	Logo           string                     `json:"logo"`
	ShortName      string                     `json:"shortName"`
	IncludeExample bool                       `json:"includeExample"`
	AdminId        uint                       `json:"adminId"`
	AdminName      string                     `gorm:"-" json:"adminName"`
	Source         serverConsts.ProjectSource `json:"source"`
}

type ProjectUserPermsPaginate struct {
	_domain.PaginateReq
}

type UpdateProjectMemberReq struct {
	ProjectId     uint `json:"projectId"`
	ProjectRoleId uint `json:"projectRoleId"`
	UserId        uint `json:"userId"`
}

type ApplyProjectReq struct {
	ProjectId       uint            `json:"projectId" validate:"required"`
	ProjectRoleName consts.RoleType `json:"projectRoleName" validate:"required"`
	ApplyUserId     uint            `json:"applyUserId"`
	Description     string          `json:"description"`
}

type AuditProjectReq struct {
	Id     uint               `json:"id" validate:"required"`
	Status consts.AuditStatus `json:"status" validate:"required"`
}

type AuditProjectPaginate struct {
	_domain.PaginateReq
	AuditUserId uint `json:"auditUserId"`
	ApplyUserId uint `json:"applyUserId"`
	Type        uint `json:"type"`
}
