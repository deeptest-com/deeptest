package handler

import (
	"github.com/kataras/iris/v12"
)

type ComponentCtrl struct {
}

func NewComponentCtrl() *ComponentCtrl {
	return &ComponentCtrl{}
}

func (c *ComponentCtrl) Detail(ctx iris.Context) {
	return
}

func (c *ComponentCtrl) Save(ctx iris.Context) {
	return
}
