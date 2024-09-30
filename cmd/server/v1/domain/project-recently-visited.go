package serverDomain

import _domain "github.com/deeptest-com/deeptest/pkg/domain"

type ProjectRecentlyVisitedBase struct {
	UserId    uint `gorm:"index:user_id_index;not null" json:"userId"`
	ProjectId uint `gorm:"index:project_id_index;not null" json:"projectId"`
}

type ProjectRecentlyVisitedReq struct {
	_domain.Model
	ProjectRecentlyVisitedBase
}
