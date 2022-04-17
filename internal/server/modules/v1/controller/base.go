package controller

import (
	"strconv"
)

type BaseCtrl struct {
}

func (c *BaseCtrl) ErrCode(err error) (code int64) {
	codeInt, _ := strconv.Atoi(err.Error())
	code = int64(codeInt)

	return
}
