package model

import (
	"github.com/aaronchen2k/deeptest/internal/server/middleware"
)

var (
	Models = []interface{}{
		&middleware.Oplog{},
		&Permission{},
		&SysRole{},
		&User{},

		&ProjectPermission{},
		&Org{},
		&Product{},
		&Project{},
		&ProjectRole{},
		&ProjectUserRole{},
		&Iteration{},

		&Feature{},
		&Issue{},
		&Label{},
		&ItemLink{},
		&ItemLinkType{},
		&Story{},
		&Task{},

		&CustomWorkitem{},
		&CustomField{},
		&CustomWorkitemField{},

		&CustomSchema{},
		&CustomSchemaItem{},
		&CustomStatus{},
		&CustomAction{},
		&CustomTransaction{},
		&CustomTransactionItem{},
		&CustomPage{},
		&CustomPageField{},
	}
)
