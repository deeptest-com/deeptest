package controllerService

import (
	"github.com/aaronchen2k/deeptest/internal/performance/pkg/domain"
	"strconv"
	"strings"
)

func IsGoalMet(req ptdomain.PerformanceTestReq,
	avgResponseTime, avgQps float64, failed, total int32) (ret bool) {

	if req.GoalAvgResponseTime > 0 {
		if avgResponseTime > req.GoalAvgResponseTime*1000 {
			return true
		}
	}

	if req.GoalAvgQps > 0 {
		if avgQps > req.GoalAvgQps {
			return true
		}
	}

	if req.GoalFailed != "" {
		val, isPercent := ParseFailTarget(req.GoalFailed)
		if val <= 0 {
			return
		}

		if isPercent {
			percent := failed / total
			if percent > val {
				return true
			}

		} else {
			if failed > val {
				return true
			}
		}
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
