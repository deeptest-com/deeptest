package test

import (
	"github.com/Knetic/govaluate"
	"testing"
)

func TestGovaluate(t *testing.T) {
	_, err := govaluate.NewEvaluableExpression("https://pms.deeptest.com/api/v1")
	if err != nil {
		return
	}
}
