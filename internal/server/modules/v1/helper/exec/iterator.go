package execHelper

import (
	"container/list"
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
)

var (
	IteratorContextStack *list.List
	IteratorContextIndex int
)

func GenerateLoopTimes(log domain.Log) (
	ret domain.ExecIterator, err error) {
	ret.ProcessorCategory = log.ProcessorCategory
	ret.ProcessorType = log.ProcessorType

	if log.Output.Times > 0 {
		for i := 0; i < log.Output.Times; i++ {
			ret.Times = append(ret.Times, i+1)
		}
	}

	return
}

func RetrieveIteratorsVal() (item interface{}, desc string, err error) {
	elem := IteratorContextStack.Front()
	if elem == nil {
		return
	}

	it := elem.Value.(domain.ExecIterator)

	if it.ProcessorType == consts.ProcessorLoopTime {
		items := it.Times
		if IteratorContextIndex > len(items)-1 {
			IteratorContextIndex = 0
		}

		if len(items) == 0 {
			return
		}

		item = items[IteratorContextIndex]
		desc = fmt.Sprintf("(%v / %d)", item, len(items))

		IteratorContextIndex++
	}

	return
}
