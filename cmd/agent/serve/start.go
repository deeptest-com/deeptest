package agentServe

import (
	stdContext "context"
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/pkg/config"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/core/middleware"
	"github.com/aaronchen2k/deeptest/internal/pkg/core/module"
	"github.com/aaronchen2k/deeptest/internal/pkg/log"
	commUtils "github.com/aaronchen2k/deeptest/internal/pkg/utils"
	_i118Utils "github.com/aaronchen2k/deeptest/pkg/lib/i118"
	"sync"
	"testing"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12/context"
	"github.com/snowlyg/helper/str"
	"github.com/snowlyg/helper/tests"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

var client *tests.Client

// Init 初始化web服务
func Init() *AgentServer {
	consts.RunFrom = consts.FromAgent
	consts.WorkDir = commUtils.GetWorkDir()

	config.Init()
	zapLog.Init()
	_i118Utils.Init(consts.Language, "")

	app := iris.New()
	app.Validator = validator.New() //参数验证
	app.Logger().SetLevel(config.CONFIG.System.Level)
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

	return &AgentServer{
		App:               app,
		addr:              config.CONFIG.System.AgentAddress,
		timeFormat:        config.CONFIG.System.TimeFormat,
		idleConnClosed:    idleConnClosed,
		globalMiddlewares: []context.Handler{middleware.Error()},
	}
}

// AddModule 添加模块
func (s *AgentServer) AddModule(module ...module.WebModule) {
	s.modules = append(s.modules, module...)
}

// GetModules 获取模块
func (s *AgentServer) GetModules() []module.WebModule {
	return s.modules
}

// GetTestAuth 获取测试验证客户端
func (s *AgentServer) GetTestAuth(t *testing.T) *tests.Client {
	var once sync.Once
	once.Do(
		func() {
			client = tests.New(str.Join("http://", s.addr), t, s.App)
			if client == nil {
				t.Fatalf("client is nil")
			}
		},
	)

	return client
}

// Init 启动web服务
func (s *AgentServer) Start() {
	s.App.UseGlobal(s.globalMiddlewares...)
	err := s.InitRouter()
	if err != nil {
		fmt.Printf("初始化路由错误： %v\n", err)
		panic(err)
	}
	// 添加上传文件路径
	s.App.Listen(
		s.addr,
		iris.WithoutInterruptHandler,
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations,
		iris.WithTimeFormat(s.timeFormat),
	)
	<-s.idleConnClosed
}
