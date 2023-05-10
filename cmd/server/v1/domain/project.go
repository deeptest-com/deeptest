package serverDomain

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/pkg/domain"
)

type ProjectReq struct {
	_domain.Model
	ProjectBase
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
	Name string `json:"name"`
	Desc string `json:"desc" gorm:"column:descr"`

	SchemaId       uint   `json:"schemaId"`
	OrgId          uint   `json:"orgId"`
	Logo           string `json:"logo"`
	ShortName      string `json:"shortName"`
	IncludeExample bool   `json:"includeExample"`
	AdminId        uint   `json:"adminId"`
	AdminName      string `gorm:"-" json:"adminName"`
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
	ProjectId       uint            `json:"projectId"`
	ProjectRoleName consts.RoleType `json:"projectRoleName"`
	ApplyUserId     uint            `json:"applyUserId"`
}

type AuditProjectReq struct {
	ProjectId uint `json:"projectId"`
	Status    uint `json:"status"`
}
