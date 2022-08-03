package domain

import "github.com/aaronchen2k/deeptest/internal/pkg/consts"

type ExecIterator struct {
	ProcessorCategory consts.ProcessorCategory
	ProcessorType     consts.ProcessorType

	// loop
	Times []int `json:"times"`
}
