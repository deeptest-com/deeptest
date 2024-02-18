package handler

import (
	"github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/config"
	"github.com/aaronchen2k/deeptest/internal/server/core/web/validate"
	"github.com/aaronchen2k/deeptest/internal/server/modules/service"
	"github.com/aaronchen2k/deeptest/pkg/domain"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"strings"

	"github.com/kataras/iris/v12"
	"github.com/snowlyg/helper/str"
	"go.uber.org/zap"
)

type DataCtrl struct {
	BaseCtrl
	DataService *service.DataService `inject:""`
}

// Init  初始化项目接口
// @Tags	初始化模块
// @summary	初始化项目
// @accept	application/json
// @Produce	application/json
// @Param 	DataReq	body	serverDomain.DataReq true 	"初始化项目的请求体"
// @success	200	{object}	_domain.Response
// @Router	/api/v1/init/initdb	[post]
func (c *DataCtrl) Init(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	req := serverDomain.DataReq{}
	if err := ctx.ReadJSON(&req); err != nil {
		errs := validate.ValidRequest(err)
		if len(errs) > 0 {
			logUtils.Errorf("参数验证失败", zap.String("错误", strings.Join(errs, ";")))
			ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: strings.Join(errs, ";")})
			return
		}
	}

	err := c.DataService.InitDB(tenantId, req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: _domain.SystemErr.Msg})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
}

// Check 检测是否需要初始化项目
// @Tags	初始化模块
// @summary	检测是否需要初始化项目
// @accept	application/json
// @Produce	application/json
// @success	200	{object}	_domain.Response{data=object{needInit=bool}}
// @Router	/api/v1/init/checkdb	[get]
func (c *DataCtrl) Check(ctx iris.Context) {

	if c.DataService.DataRepo.DB == nil {
		ctx.JSON(_domain.Response{Code: _domain.NeedInitErr.Code, Data: iris.Map{
			"needInit": true,
		}, Msg: str.Join(_domain.NeedInitErr.Msg, ":数据库初始化失败")})
		return
	} else if config.CONFIG.System.CacheType == "redis" && config.CACHE == nil {
		ctx.JSON(_domain.Response{Code: _domain.NeedInitErr.Code, Data: iris.Map{
			"needInit": true,
		}, Msg: str.Join(_domain.NeedInitErr.Msg, ":缓存驱动初始化失败")})
		return
	}
	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: iris.Map{
		"needInit": false,
	}, Msg: _domain.NoErr.Msg})
}
