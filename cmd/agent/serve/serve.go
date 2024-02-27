package agentServe

import (
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/pkg/config"
	"github.com/aaronchen2k/deeptest/internal/pkg/core/middleware"
	"github.com/aaronchen2k/deeptest/internal/pkg/core/module"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"strings"
	"sync"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/core/router"
	"github.com/kataras/iris/v12/middleware/pprof"
	"github.com/snowlyg/helper/arr"
)

// AgentServer 服务器
type AgentServer struct {
	App               *iris.Application
	modules           []module.WebModule
	idleConnClosed    chan struct{}
	addr              string
	timeFormat        string
	globalMiddlewares []context.Handler
	wg                sync.WaitGroup
	staticPath        string
	webPath           string
}

// InitRouter 初始化模块路由
func (s *AgentServer) InitRouter() error {
	s.App.UseRouter(middleware.CrsAuth("agent"))

	app := s.App.Party("/").AllowMethods(iris.MethodOptions)
	{
		if config.CONFIG.System.Level == "debug" {
			debug := DebugParty()
			app.PartyFunc(debug.RelativePath, debug.Handler)
		}
		s.initModule()
		err := s.App.Build()
		if err != nil {
			return fmt.Errorf("build router %w", err)
		}

		config.PermRoutes = s.GetSources()

		return nil
	}
}

// GetSources 获取web服务需要认证的权限
func (s *AgentServer) GetSources() []map[string]string {
	routeLen := len(s.App.GetRoutes())
	ch := make(chan map[string]string, routeLen)
	for _, r := range s.App.GetRoutes() {
		if strings.Index(r.Path, "test123") > -1 {
			logUtils.Info("")
		}

		r := r
		// 去除非接口路径
		handerNames := context.HandlersNames(r.Handlers)
		if !arr.InArrayS([]string{"GET", "POST", "PUT", "DELETE"}, r.Method) ||
			!arr.InArrayS(strings.Split(handerNames, ","), "github.com/snowlyg/multi.(*Verifier).Verify") {
			routeLen--
			continue
		}
		go func(r *router.Route) {
			route := map[string]string{
				"path": r.Path,
				"name": r.Name,
				"act":  r.Method,
			}
			ch <- route
		}(r)
	}

	routes := make([]map[string]string, routeLen)
	for i := 0; i < routeLen; i++ {
		routes[i] = <-ch
	}
	return routes
}

// initModule 初始化web服务模块，包括子模块
func (s *AgentServer) initModule() {
	if len(s.modules) > 0 {
		for _, mod := range s.modules {
			mod := mod
			s.wg.Add(1)
			go func(mod module.WebModule) {
				sub := s.App.PartyFunc(mod.RelativePath, mod.Handler)
				if len(mod.Modules) > 0 {
					for _, subModule := range mod.Modules {
						sub.PartyFunc(subModule.RelativePath, subModule.Handler)
					}
				}
				s.wg.Done()
			}(mod)
		}
		s.wg.Wait()
	}
}

// Party 调试模块
func DebugParty() module.WebModule {
	handler := func(index iris.Party) {
		index.Get("/", func(ctx iris.Context) {
			ctx.HTML("<h1>请点击<a href='/debug/pprof'>这里</a>打开调试页面")
		})
		index.Any("/pprof", pprof.New())
		index.Any("/pprof/{action:path}", pprof.New())
	}
	return module.NewModule("/debug", handler)
}
