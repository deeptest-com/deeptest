package agent

import (
	stdContext "context"
	"fmt"
	"github.com/aaronchen2k/deeptest/cmd/agent/v1/handler"
	"github.com/aaronchen2k/deeptest/internal/pkg/config"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/core/module"
	"github.com/aaronchen2k/deeptest/internal/pkg/log"
	"github.com/aaronchen2k/deeptest/internal/pkg/service"
	_i118Utils "github.com/aaronchen2k/deeptest/pkg/lib/i118"
	"github.com/facebookgo/inject"
	gorillaWs "github.com/gorilla/websocket"
	"github.com/kataras/iris/v12/websocket"
	"github.com/kataras/neffos/gorilla"
	"github.com/sirupsen/logrus"
	"net/http"
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
)

var client *tests.Client

// WebServer 服务器
type WebServer struct {
	app               *iris.Application
	modules           []module.WebModule
	idleConnClosed    chan struct{}
	addr              string
	timeFormat        string
	globalMiddlewares []context.Handler
	wg                sync.WaitGroup
	staticPrefix      string
	staticPath        string
	webPath           string
}

// Init 初始化web服务
func Init() *WebServer {
	serverConfig.Init("agent")
	zapLog.Init("agent")
	_i118Utils.Init(consts.Language, "")

	app := iris.New()
	app.Validator = validator.New() //参数验证
	app.Logger().SetLevel(serverConfig.CONFIG.System.Level)
	idleConnClosed := make(chan struct{})
	iris.RegisterOnInterrupt(func() { //优雅退出
		timeout := 10 * time.Second
		ctx, cancel := stdContext.WithTimeout(stdContext.Background(), timeout)
		defer cancel()
		app.Shutdown(ctx) // close all hosts
		close(idleConnClosed)
	})

	// init grpc
	mvc.New(app)

	// init websocket
	websocketCtrl := handler.NewWsCtrl()
	injectWsModule(websocketCtrl)

	websocketAPI := app.Party(consts.WsPath)
	m := mvc.New(websocketAPI)
	m.Register(
		&service.PrefixedLogger{Prefix: ""},
	)
	m.HandleWebsocket(websocketCtrl)
	websocketServer := websocket.New(
		gorilla.Upgrader(gorillaWs.Upgrader{CheckOrigin: func(*http.Request) bool { return true }}), m)
	websocketAPI.Get("/", websocket.Handler(websocketServer))

	return &WebServer{
		app:               app,
		addr:              serverConfig.CONFIG.System.AgentAddress,
		timeFormat:        serverConfig.CONFIG.System.TimeFormat,
		staticPrefix:      serverConfig.CONFIG.System.StaticPrefix,
		staticPath:        serverConfig.CONFIG.System.StaticPath,
		webPath:           serverConfig.CONFIG.System.WebPath,
		idleConnClosed:    idleConnClosed,
		globalMiddlewares: []context.Handler{},
	}
}

// GetStaticPath 获取静态路径
func (webServer *WebServer) GetStaticPath() string {
	return webServer.staticPath
}

// GetWebPath 获取前端路径
func (webServer *WebServer) GetWebPath() string {
	return webServer.webPath
}

// GetAddr 获取web服务地址
func (webServer *WebServer) GetAddr() string {
	return webServer.addr
}

// AddModule 添加模块
func (webServer *WebServer) AddModule(module ...module.WebModule) {
	webServer.modules = append(webServer.modules, module...)
}

// AddStatic 添加静态文件
func (webServer *WebServer) AddStatic(requestPath string, fsOrDir interface{}, opts ...iris.DirOptions) {
	webServer.app.HandleDir(requestPath, fsOrDir, opts...)
}

// AddWebStatic 添加前端访问地址
func (webServer *WebServer) AddWebStatic(requestPath string) {
	fsOrDir := iris.Dir(filepath.Join(dir.GetCurrentAbPath(), webServer.webPath))
	webServer.AddStatic(requestPath, fsOrDir, iris.DirOptions{
		IndexName: "index.html",
		SPA:       true,
	})
}

// AddUploadStatic 添加上传文件访问地址
func (webServer *WebServer) AddUploadStatic() {
	fsOrDir := iris.Dir(filepath.Join(dir.GetCurrentAbPath(), webServer.staticPath))
	webServer.AddStatic(webServer.staticPrefix, fsOrDir)
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

// GetTestLogin 测试登录web服务
func (webServer *WebServer) GetTestLogin(t *testing.T, url string, res tests.Responses, datas ...map[string]interface{}) *tests.Client {
	client := webServer.GetTestAuth(t)
	err := client.Login(url, res, datas...)
	if err != nil {
		t.Fatal(err)
	}
	return client
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

func injectWsModule(websocketCtrl *handler.WebSocketCtrl) {
	var g inject.Graph
	g.Logger = logrus.StandardLogger()

	if err := g.Provide(
		&inject.Object{Value: websocketCtrl},
	); err != nil {
		logrus.Fatalf("provide usecase objects to the Graph: %v", err)
	}
	err := g.Populate()
	if err != nil {
		logrus.Fatalf("populate the incomplete Objects: %v", err)
	}
}
