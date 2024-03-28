package serve

import (
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/pkg/config"
	"github.com/aaronchen2k/deeptest/internal/pkg/core/middleware"
	"github.com/aaronchen2k/deeptest/internal/pkg/core/module"
	serverMiddleware "github.com/aaronchen2k/deeptest/internal/server/middleware"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"log"
	"strings"
	"sync"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/core/router"
	"github.com/kataras/iris/v12/middleware/pprof"
	"github.com/snowlyg/helper/arr"
)

// WebServer 服务器
type WebServer struct {
	app               *iris.Application
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
func (webServer *WebServer) InitRouter() error {
	webServer.app.UseRouter(middleware.CrsAuth("server"))

	app := webServer.app.Party("/").AllowMethods(iris.MethodOptions)
	{
		app.Use(serverMiddleware.InitCheck())
		if config.CONFIG.System.Level == "debug" {
			debug := DebugParty()
			app.PartyFunc(debug.RelativePath, debug.Handler)
		}
		webServer.initModule()

		webServer.AddWebUi()
		webServer.AddUpload()
		webServer.AddTest()
		webServer.AddSwagger()

		err := webServer.app.Build()
		if err != nil {
			return fmt.Errorf("build router %w", err)
		}

		config.PermRoutes = webServer.GetSources()
		for _, r := range config.PermRoutes {
			if strings.Contains(r["path"], "summary/details") {
				log.Print("")
			}
		}

		return nil
	}
}

// GetSources 获取web服务需要认证的权限
func (webServer *WebServer) GetSources() []map[string]string {
	serverRoutes := webServer.app.GetRoutes()
	routeLen := len(serverRoutes)
	logUtils.Infof("routes len = %d", routeLen)

	ch := make(chan map[string]string, routeLen)
	for _, r := range serverRoutes {
		r := r

		// ignore no api OR no perm methods
		handlerNames := context.HandlersNames(r.Handlers)

		if !isApiMethod(r.Method) || !hasPerm(handlerNames) {
			logUtils.Infof("NoPerm NAME: %s, PATH: %s, METHOD: %s, NAMES: %s ", r.Name, r.Path, r.Method, handlerNames)

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

	for _, r := range routes {
		if strings.Contains(r["path"], "summary/details") {
			log.Print("")
		}
	}

	return routes
}

// initModule 初始化web服务模块，包括子模块
func (webServer *WebServer) initModule() {
	if len(webServer.modules) > 0 {
		for _, mod := range webServer.modules {
			mod := mod
			webServer.wg.Add(1)
			func(mod module.WebModule) {
				sub := webServer.app.PartyFunc(mod.RelativePath, mod.Handler)
				if len(mod.Modules) > 0 {
					for _, subModule := range mod.Modules {
						sub.PartyFunc(subModule.RelativePath, subModule.Handler)
					}
				}
				webServer.wg.Done()
			}(mod)
		}
		webServer.wg.Wait()
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

func isApiMethod(method string) bool {
	return arr.InArrayS([]string{"GET", "POST", "PUT", "DELETE"}, method)
}
func hasPerm(handlerNames string) bool {
	names := strings.Split(handlerNames, ",")

	hasPerm := false
	for _, name := range names {
		if strings.Index(name, ".Casbin") > -1 {
			hasPerm = true
			break
		}
	}

	return hasPerm
}
