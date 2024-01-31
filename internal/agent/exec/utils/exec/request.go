package execUtils

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
)

func IsNotUseBaseUrl(usedBy consts.UsedBy, src consts.ProcessorInterfaceSrc) (ret bool) {
	ret = usedBy == consts.DiagnoseDebug ||
		(usedBy == consts.ScenarioDebug && (src == consts.InterfaceSrcDiagnose ||
			src == consts.InterfaceSrcCustom ||
			src == consts.InterfaceSrcCurl))

	return
}
