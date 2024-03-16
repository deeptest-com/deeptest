package conductorExec

import (
	"github.com/aaronchen2k/deeptest/internal/agent/performance/pkg/domain"
	"strconv"
	"strings"
)

func IsGoalMet(req ptdomain.PerformanceTestData,
	avgResponseTime, avgQps float64, failed, total int32) (ret bool) {

	if req.Goal.ResponseTime > 0 {
		if avgResponseTime > float64(req.Goal.ResponseTime*1000) {
			return true
		}
	}

	if req.Goal.Qps > 0 {
		if avgQps > float64(req.Goal.Qps) {
			return true
		}
	}

	if req.Goal.FailRate > 0 {
		// 改成比率
		//if failed > req.GoalFailRate {
		//	return true
		//}
	}

	return
}

func ParseFailTarget(str string) (ret int32, isPercent bool) {
	if strings.Index(str, "%") == len(str)-1 {
		numStr := strings.TrimRight(str, "%")
		val, err := strconv.Atoi(numStr)

		if err == nil {
			isPercent = true
			ret = int32(val)
		}

		return
	}

	retInt, _ := strconv.Atoi(str)
	ret = int32(retInt)

	return
}
