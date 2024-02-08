package serverUtils

import (
	ptProto "github.com/aaronchen2k/deeptest/internal/performance/proto"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
)

func GetInterfaceMap(scenarios []*ptProto.Scenario) (ret map[int32]string) {
	ret = map[int32]string{}

	for _, scenario := range scenarios {
		for _, processor := range scenario.Processors {
			if processor.Type != consts.ProcessorInterfaceDefault.ToString() {
				continue
			}

			ret[processor.Id] = processor.Name
		}
	}

	return
}
