package repo

import (
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"gorm.io/gorm"
)

type ComponentSchemaSecurityRepo struct {
	DB *gorm.DB `inject:""`
}

func NewComponentSchemaSecurityRepo() *ComponentSchemaSecurityRepo {
	return &ComponentSchemaSecurityRepo{}
}

func (r *ComponentSchemaSecurityRepo) Save(ComponentSchema model.ComponentSchema) (err error) {

	return
}
