package _floatUtils

import (
	"fmt"
	"strconv"
)

func PointNumb(value float64, points int) (ret float64) {
	ret, _ = strconv.ParseFloat(fmt.Sprintf("%."+strconv.Itoa(points)+"f", value), 64)

	return
}
