package repo

import (
	"github.com/deeptest-com/deeptest/internal/server/modules/model"
	"gorm.io/gorm"
)

type ComponentSchemaSecurityRepo struct {
	DB *gorm.DB `inject:""`
}

func (r *ComponentSchemaSecurityRepo) Save(ComponentSchema model.ComponentSchema) (err error) {

	return
}
