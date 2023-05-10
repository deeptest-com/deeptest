package model

import "github.com/aaronchen2k/deeptest/internal/pkg/consts"

type ProjectMemberAudit struct {
	BaseModel
	ProjectId       uint               `json:"projectId"`
	ProjectRoleName consts.RoleType    `json:"projectRoleName"`
	AuditUserId     uint               `json:"auditUserId"`
	ApplyUserId     uint               `json:"applyUserId"`
	Status          consts.AuditStatus `gorm:"default:0" json:"status"`
	Description     string             `json:"description" gorm:"type:text"`
}

func (ProjectMemberAudit) TableName() string {
	return "biz_project_member_audit"
}
