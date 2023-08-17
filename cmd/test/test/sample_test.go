package test

import (
	_intUtils "github.com/aaronchen2k/deeptest/pkg/lib/int"
	"log"
	"testing"
)

func TestSample(t *testing.T) {
	indexArr := _intUtils.GenUniqueRandNum(0, 10, 10)

	log.Println(indexArr)
}
