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

	WsDefaultNamespace = "default"
	WsDefaultRoom      = "default_room"
	WsChatEvent        = "OnChat"

	WebCheckInterval         = 60 * 60
	SummaryDataCheckInterval = 60 * 60
	MaxNum                   = 10000

	AgentConfigFileName  = "agent.yaml"
	ServerConfigFileName = "server.yaml"
	CasbinFileName       = "rbac_model.conf" // casbin 规则文件名称

	WebsiteDev = "http://localhost:8000/"

	ExtractorErr = "extractor_err"
	ContentErr   = "content_err"

	HttpRequestTimeout = 60 * time.Second
	DeepestKey         = "com_deeptest_prop_for_selection"
	KEY_BASE_URL       = "_base_url_"

	MaxLoopTimeForInterfaceTest   = 1000
	MaxLoopTimeForPerformanceTest = 1000000

	KEY_MOCKJS = "x-mock-type"

	INVALID_VALUE = "N/A"

	TmpDirRelative       = "tmp"
	TmpDirRelativeServer = TmpDirRelative + "/server"
	TmpDirRelativeAgent  = TmpDirRelative + "/agent"
)

var (
	DirUpload     = "upload"
	HeaderOptions = []string{"Accept", "Accept-Encoding", "Accept-Language", "Connection", "Host", "Referer", "User-Agent", "Cache-Control", "Cookie", "Range"}
)
