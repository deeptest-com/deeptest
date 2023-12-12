package k6Converter

import (
	k6Comm "github.com/aaronchen2k/deeptest/internal/pkg/helper/performance/k6/comm"
	k6Domain "github.com/aaronchen2k/deeptest/internal/pkg/helper/performance/k6/domain"
	"os"
	"strconv"
	"strings"
)

func ScenarioOptions(scenarios []k6Domain.PerfScenario) (script string, err error) {
	var arr []string

	for _, sc := range scenarios {
		code, _ := ScenarioOption(sc)
		arr = append(arr, code)
	}

	script = strings.Join(arr, "\n")

	return
}

func ScenarioOption(scenario k6Domain.PerfScenario) (script string, err error) {
	data := map[string]string{
		"name":        scenario.Name,
		"exec":        scenario.Func,
		"constantVus": scenario.Executor.String(),
		"vues":        strconv.Itoa(scenario.Vues),
		"duration":    strconv.Itoa(scenario.Duration),
	}

	content := k6Comm.GetTmpl("scenario")
	script = os.Expand(content, func(k string) string { return data[k] })

	return
}

func ScenarioFuncs(scenarios []k6Domain.PerfScenario) (script string, err error) {
	var arr []string

	for _, sc := range scenarios {
		code, _ := ScenarioFunc(sc)
		arr = append(arr, code)
	}

	script = strings.Join(arr, "\n")

	return
}

func ScenarioFunc(scenario k6Domain.PerfScenario) (script string, err error) {

	return
}
