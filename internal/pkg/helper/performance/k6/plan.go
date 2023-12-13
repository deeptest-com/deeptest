package k6Converter

import (
	k6Comm "github.com/aaronchen2k/deeptest/internal/pkg/helper/performance/k6/comm"
	keDomain "github.com/aaronchen2k/deeptest/internal/pkg/helper/performance/k6/domain"
	"github.com/flosch/pongo2/v4"
)

func PlanCode(plan keDomain.PerfPlan) (script string, err error) {
	content := k6Comm.GetTmpl("plan")

	tpl, err := pongo2.FromString(content)
	if err != nil {
		return
	}

	script, err = tpl.Execute(pongo2.Context{"plan": plan})
	if err != nil {
		return
	}

	return
}
