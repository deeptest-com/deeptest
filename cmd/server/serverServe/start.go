package serverServe

import (
	stdContext "context"
	"fmt"
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1"
	"github.com/aaronchen2k/deeptest/cmd/server/v1/handler"
	"github.com/aaronchen2k/deeptest/internal/pkg/config"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/core/cron"
	"github.com/aaronchen2k/deeptest/internal/pkg/core/middleware"
	"github.com/aaronchen2k/deeptest/internal/pkg/core/module"
	"github.com/aaronchen2k/deeptest/internal/pkg/log"
	commUtils "github.com/aaronchen2k/deeptest/internal/pkg/utils"
	"github.com/aaronchen2k/deeptest/internal/server/core/cache"
	"github.com/aaronchen2k/deeptest/internal/server/core/dao"
	"github.com/aaronchen2k/deeptest/internal/server/modules/service"
	fileUtils "github.com/aaronchen2k/deeptest/pkg/lib/file"
	_i118Utils "github.com/aaronchen2k/deeptest/pkg/lib/i118"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"github.com/facebookgo/inject"
	"github.com/kataras/iris/v12/websocket"
	"github.com/sirupsen/logrus"
	"path/filepath"
	"sync"
	"testing"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12/context"
	"github.com/snowlyg/helper/dir"
	"github.com/snowlyg/helper/str"
	"github.com/snowlyg/helper/tests"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"

	_ "github.com/aaronchen2k/deeptest/cmd/server/docs"
	"github.com/iris-contrib/swagger"
	"github.com/iris-contrib/swagger/swaggerFiles"
)

var client *tests.Client

// Start 初始化web服务
func Start() {
	inits()

	idleConnClosed := make(chan struct{})
	irisApp := createIrisApp(&idleConnClosed)

	//irisApp.Use(func(ctx iris.Context) {
	//	ctx.Posts().Header.Del("Origin")
	//	ctx.Next()
	//})

	initWebSocket(irisApp)

	server := &WebServer{
		app:               irisApp,
		addr:              config.CONFIG.System.ServerAddress,
		timeFormat:        config.CONFIG.System.TimeFormat,
		staticPath:        config.CONFIG.System.StaticPath,
		webPath:           config.CONFIG.System.WebPath,
		idleConnClosed:    idleConnClosed,
		globalMiddlewares: []context.Handler{middleware.Error()},
	}

	server.InjectModule()
	server.Start()
}

func inits() {
	consts.RunFrom = consts.FromServer
	consts.WorkDir = commUtils.GetWorkDir()

	config.Init()
	zapLog.Init()
	_i118Utils.Init(consts.Language, "")

	err := cache.Init()
	if err != nil {
		logUtils.Errorf("init redis cache failed, error %s", err.Error())
		return
	}
}

func createIrisApp(idleConnClosed *chan struct{}) (irisApp *iris.Application) {
	irisApp = iris.New()
	irisApp.Validator = validator.New() //参数验证
	irisApp.Logger().SetLevel(config.CONFIG.System.Level)

	iris.RegisterOnInterrupt(func() { //优雅退出
		timeout := 10 * time.Second
		ctx, cancel := stdContext.WithTimeout(stdContext.Background(), timeout)
		defer cancel()
		irisApp.Shutdown(ctx) // close all hosts

		close(*idleConnClosed)
	})

	return
}

// injectWebsocketModule 注册组件
func injectWebsocketModule(websocketCtrl *handler.WebSocketCtrl) {
	var g inject.Graph
	g.Logger = logrus.StandardLogger()

	if err := g.Provide(
		&inject.Object{Value: dao.GetDB()},
		&inject.Object{Value: websocketCtrl},
	); err != nil {
		logrus.Fatalf("provide usecase objects to the Graph: %v", err)
	}
	err := g.Populate()
	if err != nil {
		logrus.Fatalf("populate the incomplete Objects: %v", err)
	}
}

// Init 启动web服务
func (webServer *WebServer) Start() {
	webServer.app.UseGlobal(webServer.globalMiddlewares...)
	err := webServer.InitRouter()
	if err != nil {
		fmt.Printf("初始化路由错误： %v\n", err)
		panic(err)
	}
	// 添加上传文件路径
	webServer.app.Listen(
		webServer.addr,
		iris.WithoutInterruptHandler,
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations,
		iris.WithTimeFormat(webServer.timeFormat),
	)
	<-webServer.idleConnClosed
}

// GetAddr 获取web服务地址
func (webServer *WebServer) GetAddr() string {
	return webServer.addr
}

// AddModule 添加模块
func (webServer *WebServer) AddModule(module ...module.WebModule) {
	webServer.modules = append(webServer.modules, module...)
}

// AddWebUi 添加前端页面访问
func (webServer *WebServer) AddWebUi() {
	pth := filepath.Join(dir.GetCurrentAbPath(), "ui", "dist")
	fileUtils.MkDirIfNeeded(pth)
	logUtils.Infof("*** ui dir: %s", pth)

	webServer.app.HandleDir("/", iris.Dir(pth), iris.DirOptions{
		IndexName: "index.html",
		ShowList:  false,
		SPA:       true,
	})
}

// AddUpload 添加上传文件访问
func (webServer *WebServer) AddUpload() {
	pth := filepath.Join(dir.GetCurrentAbPath(), consts.DirUpload)
	fileUtils.MkDirIfNeeded(pth)
	logUtils.Infof("*** upload dir: %s", pth)

	webServer.app.HandleDir("/upload", iris.Dir(pth))
}

// AddTest 添加测试文件访问
func (webServer *WebServer) AddTest() {
	pth := filepath.Join(dir.GetCurrentAbPath(), filepath.Join(webServer.staticPath, "test"))
	fileUtils.MkDirIfNeeded(pth)
	logUtils.Infof("*** test dir: %s", pth)

	webServer.app.HandleDir("/test", iris.Dir(pth))
}

func (webServer *WebServer) AddSwagger() {
	swaggerConfig := swagger.Config{
		URL:          fmt.Sprintf("swagger/doc.json"),
		DeepLinking:  true,
		DocExpansion: "list",
		DomID:        "#swagger-ui",
		Prefix:       "/swagger",
	}

	swaggerUI := swagger.Handler(swaggerFiles.Handler, swaggerConfig)
	webServer.app.Get("/swagger", swaggerUI)
	webServer.app.Get("/swagger/{any:path}", swaggerUI)

}

// GetModules 获取模块
func (webServer *WebServer) GetModules() []module.WebModule {
	return webServer.modules
}

// GetTestAuth 获取测试验证客户端
func (webServer *WebServer) GetTestAuth(t *testing.T) *tests.Client {
	var once sync.Once
	once.Do(
		func() {
			client = tests.New(str.Join("http://", webServer.addr), t, webServer.app)
			if client == nil {
				t.Fatalf("client is nil")
			}
		},
	)

	return client
}

// Init 加载模块
func initWebSocket(irisApp *iris.Application) {
	websocketCtrl := handler.NewWebsocketCtrl()
	injectWebsocketModule(websocketCtrl)

	mvc.New(irisApp)

	websocketAPI := irisApp.Party(consts.WsPath)
	m := mvc.New(websocketAPI)
	m.Register(
		&service.PrefixedLogger{Prefix: ""},
	)
	m.HandleWebsocket(websocketCtrl)
	websocketServer := websocket.New(
		middleware.DefaultUpgrader,
		//gorilla.Upgrader(gorillaWs.Upgrader{CheckOrigin: func(*http.Posts) bool { return true }}),
		m)

	websocketAPI.Get("/", websocket.Handler(websocketServer))
}

// Init 加载模块
func (webServer *WebServer) InjectModule() {
	var g inject.Graph
	g.Logger = logrus.StandardLogger()

	cron := cron.NewServerCron()
	cron.Init()
	indexModule := v1.NewIndexModule()

	// inject objects
	if err := g.Provide(
		&inject.Object{Value: dao.GetDB()},
		&inject.Object{Value: cron},
		&inject.Object{Value: indexModule},
	); err != nil {
		logrus.Fatalf("provide usecase objects to the Graph: %v", err)
	}
	err := g.Populate()
	if err != nil {
		logrus.Fatalf("populate the incomplete Objects: %v", err)
	}

	webServer.AddModule(indexModule.ApiParty())
	//webServer.AddModule(indexModule.MockParty())
}
