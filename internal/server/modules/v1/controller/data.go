package controller

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	logUtils "github.com/aaronchen2k/deeptest/internal/pkg/lib/log"
	serverConfig "github.com/aaronchen2k/deeptest/internal/server/config"
	"github.com/aaronchen2k/deeptest/internal/server/core/web/validate"
	serverDomain "github.com/aaronchen2k/deeptest/internal/server/modules/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/service"
	"strings"

	"github.com/kataras/iris/v12"
	"github.com/snowlyg/helper/str"
	"go.uber.org/zap"
)

type DataCtrl struct {
	DataService *service.DataService `inject:""`
}

func NewDataCtrl() *DataCtrl {
	return &DataCtrl{}
}

// InitDB 初始化项目接口
func (c *DataCtrl) Init(ctx iris.Context) {
	req := serverDomain.DataReq{}
	if err := ctx.ReadJSON(&req); err != nil {
		errs := validate.ValidRequest(err)
		if len(errs) > 0 {
			logUtils.Errorf("参数验证失败", zap.String("错误", strings.Join(errs, ";")))
			ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Data: nil, Msg: strings.Join(errs, ";")})
			return
		}
	}
	err := c.DataService.InitDB(req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Data: nil, Msg: _domain.SystemErr.Msg})
		return
	}
	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: nil, Msg: _domain.NoErr.Msg})
}

// Check 检测是否需要初始化项目
func (c *DataCtrl) Check(ctx iris.Context) {
	if c.DataService.DataRepo.DB == nil {
		ctx.JSON(_domain.Response{Code: _domain.NeedInitErr.Code, Data: iris.Map{
			"needInit": true,
		}, Msg: str.Join(_domain.NeedInitErr.Msg, ":数据库初始化失败")})
		return
	} else if serverConfig.CONFIG.System.CacheType == "redis" && serverConfig.CACHE == nil {
		ctx.JSON(_domain.Response{Code: _domain.NeedInitErr.Code, Data: iris.Map{
			"needInit": true,
		}, Msg: str.Join(_domain.NeedInitErr.Msg, ":缓存驱动初始化失败")})
		return
	}
	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: iris.Map{
		"needInit": false,
	}, Msg: _domain.NoErr.Msg})
}
