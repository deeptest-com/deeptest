package repo

import (
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"gorm.io/gorm"
)

type SnippetRepo struct {
	DB *gorm.DB `inject:""`
}

func (r *SnippetRepo) Get(name string) (po model.Snippet, err error) {
	err = r.DB.
		Where("name=?", name).
		Where("NOT deleted").
		First(&po).Error
	return
}
