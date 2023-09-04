package consts

import "time"

const (
	App = "deeptest"

	ContentType   = "Content-Type"
	ContentLength = "Content-Length"
	Server        = "Server"
	Allow         = "Allow"
	Connection    = "Connection"
	Host          = "Host"

	Authorization = "Authorization"

	UserAgentChrome = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36"

	ApiPathServer = "/api/v1"
	ApiPathAgent  = "/api/v1"
	ApiPathMock   = "/mocks"
	WsPath        = ApiPathAgent + "/ws"

	WsDefaultNameSpace = "default"
	WsDefaultRoom      = "default"
	WsChatEvent        = "OnChat"

	WebCheckInterval         = 60 * 60
	SummaryDataCheckInterval = 60 * 60
	MaxNum                   = 10000

	ConfigFileName = "server.yaml"
	CasbinFileName = "rbac_model.conf" // casbin 规则文件名称

	SupportEmail = "chenqi@deeptest.com"
	Sys          = "DeepTest"
	Url          = "https://deeptest.com/"
	UrlDev       = "http://localhost:8000/"

	ExtractorErr = "extractor_err"
	ContentErr   = "content_err"

	EmailSmtpAddress = "smtp.exmail.qq.com"
	EmailSmtpPort    = 465
	EmailAccount     = "chenqi@deeptest.com"
	EmailPassword    = ""

	HttpRequestTimeout = 60 * time.Second
	DeepestKey         = "com_deeptest_prop_for_selection"
	KEY_BASE_URL       = "_base_url_"

	MaxLoopTimeForInterfaceTest   = 1000
	MaxLoopTimeForPerformanceTest = 1000000
)

var (
	DirUpload = "upload"
)
