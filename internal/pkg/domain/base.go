package domain

import (
	"time"
)

type BaseObj struct {
	ID uint `json:"id"`

	CreatedAt *time.Time `json:"createdAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`

	CreatedBy uint `json:"createdBy"`
	UpdatedBy uint `json:"updatedBy"`

	Disabled bool `json:"disabled,omitempty" gorm:"default:false"`
}
