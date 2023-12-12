package k6Converter

import (
	k6Comm "github.com/aaronchen2k/deeptest/internal/pkg/helper/performance/k6/comm"
	keDomain "github.com/aaronchen2k/deeptest/internal/pkg/helper/performance/k6/domain"
	"os"
)

func PlanCode(plan keDomain.PerfPlan) (script string, err error) {
	scenarios, err := ScenarioOptions(plan.Scenarios)

	codes, err := ScenarioCodes(plan.Scenarios)

	data := map[string]string{
		"scenarios": scenarios,
		"codes":     codes,
	}

	content := k6Comm.GetTmpl("plan")
	script = os.Expand(content, func(k string) string { return data[k] })

	return
}
