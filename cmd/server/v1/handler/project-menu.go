package handler

import (
	"github.com/aaronchen2k/deeptest/internal/server/modules/service"
	_domain "github.com/aaronchen2k/deeptest/pkg/domain"
	"github.com/kataras/iris/v12"
	"github.com/snowlyg/multi"
)

type ProjectMenuCtrl struct {
	ProjectMenuService *service.ProjectMenuService `inject:""`
	BaseCtrl
}

// UserMenuList
// @Tags	项目菜单
// @summary	项目中用户的左侧菜单栏列表
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string							true	"Authentication header"
// @Param 	currProjectId	query	int								true	"当前项目ID"
// @success	200	{object}	_domain.Response{data=object{result=[]model.ProjectMenu}}
// @Router	/api/v1/projects/menus/userMenuList	[get]
func (c *ProjectMenuCtrl) UserMenuList(ctx iris.Context) {
	userId := multi.GetUserId(ctx)

	projectId, err := ctx.URLParamInt("projectId")
	if projectId == 0 {
		projectId, err = ctx.URLParamInt("currProjectId")
	}

	data, err := c.ProjectMenuService.GetUserMenuList(uint(projectId), userId)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ret := iris.Map{"result": data}
	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: ret, Msg: _domain.NoErr.Msg})
}
