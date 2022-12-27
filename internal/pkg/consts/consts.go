package consts

const (
	App = "deeptest"

	ContentType   = "NodeContent-Type"
	ContentLength = "NodeContent-Length"
	Server        = "Server"
	Allow         = "Allow"
	Connection    = "Connection"
	Host          = "Host"

	Authorization = "Authorization"

	UserAgentChrome = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36"

	ApiPath      = "/api/v1"
	ApiPathAgent = "/api/v1"
	WsPath       = ApiPathAgent + "/ws"

	WsDefaultNameSpace = "default"
	WsDefaultRoom      = "default"
	WsChatEvent        = "OnChat"

	WebCheckInterval = 60 * 60
	MaxNum           = 10000

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

	DeepestKey = "com_deeptest_prop_for_selection"
)
