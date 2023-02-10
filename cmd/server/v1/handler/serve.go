package handler

import (
	"github.com/kataras/iris/v12"
)

type ServeCtrl struct {
	//ServeService *service.ServeService `inject:""`
}

// Index 服务列表
func (c *ServeCtrl) Index(ctx iris.Context) {
	return
}

// Save 保存服务
func (c *ServeCtrl) Save(ctx iris.Context) {
	return
}

// Detail 服务详情
func (c *ServeCtrl) Detail(ctx iris.Context) {

}

// Delete 删除服务
func (c *ServeCtrl) Delete(ctx iris.Context) {

}

// SaveVersion 保存版本
func (c *ServeCtrl) SaveVersion(ctx iris.Context) {

}

func (c *ServeCtrl) ListVersion(ctx iris.Context) {

}

func (c *ServeCtrl) SaveSchema(ctx iris.Context) {

}

/*
//构造参数构造auth，BasicAuth,BearerToken,OAuth20,ApiKey
func (c *ServeCtrl) requestParser(req *v1.ServeReq) {
	for _, item := range req.Interfaces {
		fmt.Println(_commUtils.JsonEncode(item.ResponseBodies))
		//req.Interfaces[key].RequestBody.SchemaItem.Content = _commUtils.JsonEncode(item.RequestBody.SchemaItem.Content)
	}
}

*/
