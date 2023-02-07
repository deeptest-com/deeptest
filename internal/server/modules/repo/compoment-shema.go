package repo

import (
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"gorm.io/gorm"
)

type ComponentSchemaRepo struct {
	DB *gorm.DB `inject:""`
}

func NewComponentSchemaRepo() *ComponentSchemaRepo {
	return &ComponentSchemaRepo{}
}

func (r *ComponentSchemaRepo) Save(ComponentSchema model.ComponentSchema) (err error) {

	return
}

