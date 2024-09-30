package repo

import (
	"github.com/deeptest-com/deeptest/internal/pkg/consts"
	"github.com/deeptest-com/deeptest/internal/server/modules/model"
	"gorm.io/gorm"
)

type SnippetRepo struct {
	*BaseRepo `inject:""`
	DB        *gorm.DB `inject:""`
}

func (r *SnippetRepo) Get(tenantId consts.TenantId, name string) (po model.Snippet, err error) {
	err = r.GetDB(tenantId).
		Where("name=?", name).
		Where("NOT deleted").
		First(&po).Error
	return
}
