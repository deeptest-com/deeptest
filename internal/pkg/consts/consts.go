package consts

const (
	App     = "deeptest"
	Pattern = "pattern"

	ContentType   = "Content-Type"
	ContentLength = "Content-Length"
	Server        = "Server"
	Allow         = "Allow"
	Connection    = "Connection"
	Host          = "Host"

	Authorization = "Authorization"

	UserAgentChrome = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36"

	ApiPath = "/api/v1"
	WsPath  = ApiPath + "/ws"

	WsDefaultNameSpace = "default"
	WsDefaultRoom      = "default"
	WsChatEvent        = "OnChat"

	WebCheckInterval = 60 * 60
	MaxNum           = 10000

	ConfigFileName = "server.yaml"
	CasbinFileName = "rbac_model.conf" // casbin 规则文件名称

	SupportEmail = "support@deeptest.com"
	Sys          = "DeepTest"
	Url          = "https://deeptest.com"
)
