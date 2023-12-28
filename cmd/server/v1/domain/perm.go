package serverDomain

import (
	"github.com/aaronchen2k/deeptest/pkg/domain"
)

type PermReq struct {
	_domain.Model
	PermBase
}

type PermReqPaginate struct {
	_domain.PaginateReq
	Name string `json:"name"`
}

type PermResp struct {
	_domain.Model
	PermReq
}

type PermBase struct {
	Name        string `gorm:"index:perm_index,unique;not null ;type:varchar(200)" json:"name" validate:"required,gte=4,lte=50"`
	DisplayName string `gorm:"type:varchar(256)" json:"displayName"`
	Description string `gorm:"type:text" json:"description"`
	Act         string `gorm:"index:perm_index;type:varchar(50)" json:"act" validate:"required"`
}

type PermStruct struct {
	Name string `json:"name"`
	Act  string `json:"act"`
}
