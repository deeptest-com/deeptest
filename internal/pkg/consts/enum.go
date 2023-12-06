package consts

type RunType string

const (
	FromServer RunType = "server"
	FromAgent  RunType = "agent"
)

func (e RunType) String() string {
	return string(e)
}

type ExecFromType string

const (
	FromCmd    ExecFromType = "cmd"
	FromClient ExecFromType = "client"
)

func (e ExecFromType) String() string {
	return string(e)
}

type WsMsgCategory string

const (
	ProgressInProgress WsMsgCategory = "in_progress"
	ProgressEnd        WsMsgCategory = "end"
	ProgressResult     WsMsgCategory = "result"
	Processor          WsMsgCategory = "processor"
	Initialize         WsMsgCategory = "initialize"
	Statistic          WsMsgCategory = "stat"
	Exception          WsMsgCategory = "exception"
)

func (e WsMsgCategory) String() string {
	return string(e)
}

type HttpMethod string

const (
	GET    HttpMethod = "GET"
	POST   HttpMethod = "POST"
	PUT    HttpMethod = "PUT"
	DELETE HttpMethod = "DELETE"

	PATCH   HttpMethod = "PATCH"
	HEAD    HttpMethod = "HEAD"
	CONNECT HttpMethod = "CONNECT"
	OPTIONS HttpMethod = "OPTIONS"
	TRACE   HttpMethod = "TRACE"
)

func (e HttpMethod) String() string {
	return string(e)
}

type FormDataType string

const (
	FormDataTypeText FormDataType = "text"
	FormDataTypeFile FormDataType = "file"
)

func (e FormDataType) String() string {
	return string(e)
}

type HttpRespCode int

const (
	Continue          HttpRespCode = 100
	SwitchingProtocol HttpRespCode = 101

	OK                          HttpRespCode = 200
	Created                     HttpRespCode = 201
	Accepted                    HttpRespCode = 202
	NonAuthoritativeInformation HttpRespCode = 203
	NoContent                   HttpRespCode = 204
	ResetContent                HttpRespCode = 205
	PartialContent              HttpRespCode = 206

	MultipleChoice    HttpRespCode = 300
	MovedPermanently  HttpRespCode = 301
	Found             HttpRespCode = 302
	SeeOther          HttpRespCode = 303
	NotModified       HttpRespCode = 304
	UseProxy          HttpRespCode = 305
	unused            HttpRespCode = 306
	TemporaryRedirect HttpRespCode = 307
	PermanentRedirect HttpRespCode = 308

	BadRequest                   HttpRespCode = 400
	Unauthorized                 HttpRespCode = 401
	PaymentRequired              HttpRespCode = 402
	Forbidden                    HttpRespCode = 403
	NotFound                     HttpRespCode = 404
	MethodNotAllowed             HttpRespCode = 405
	NotAcceptable                HttpRespCode = 406
	ProxyAuthenticationRequired  HttpRespCode = 407
	RequestTimeout               HttpRespCode = 408
	Conflict                     HttpRespCode = 409
	Gone                         HttpRespCode = 410
	LengthRequired               HttpRespCode = 411
	PreconditionFailed           HttpRespCode = 412
	RequestEntityTooLarge        HttpRespCode = 413
	RequestURITooLong            HttpRespCode = 414
	UnsupportedMediaType         HttpRespCode = 415
	RequestedRangeNotSatisfiable HttpRespCode = 416
	ExpectationFailed            HttpRespCode = 417

	InternalServerError     HttpRespCode = 500
	Implemented             HttpRespCode = 501
	BadGateway              HttpRespCode = 502
	ServiceUnavailable      HttpRespCode = 503
	GatewayTimeout          HttpRespCode = 504
	HTTPVersionNotSupported HttpRespCode = 505
)

func (e HttpRespCode) Int() int {
	return int(e)
}

type HttpContentType string

const (
	ContentTypeJSON HttpContentType = "application/json"
	ContentTypeXML  HttpContentType = "application/xml"
	ContentTypeHTML HttpContentType = "text/html"
	ContentTypeTEXT HttpContentType = "text/text"

	ContentTypeFormData       HttpContentType = "multipart/form-data"
	ContentTypeFormUrlencoded HttpContentType = "application/x-www-form-urlencoded"

	ContentTypeUnixDir HttpContentType = "httpd/unix-directory"
)

func (e HttpContentType) String() string {
	return string(e)
}

type AuthorType string

const (
	BasicAuth   AuthorType = "basicAuth"
	BearerToken AuthorType = "bearerToken"
	OAuth2      AuthorType = "oAuth2"
	ApiKey      AuthorType = "apiKey"
)

func (e AuthorType) String() string {
	return string(e)
}

type GrantType string

const (
	AuthorizationCode         GrantType = "authorizationCode"
	AuthorizationCodeWithPKCE GrantType = "authorizationCodeWithPKCE"
	Implicit                  GrantType = "implicit"
	PasswordCredential        GrantType = "passwordCredential"
	ClientCredential          GrantType = "clientCredential"
)

func (e GrantType) String() string {
	return string(e)
}

type ClientAuthenticationWay string

const (
	SendAsBasicAuthHeader       ClientAuthenticationWay = "sendAsBasicAuthHeader"
	SendClientCredentialsInBody ClientAuthenticationWay = "sendClientCredentialsInBody"
)

func (e ClientAuthenticationWay) String() string {
	return string(e)
}

type HttpRespLangType string

const (
	LangJSON HttpRespLangType = "json"
	LangXML  HttpRespLangType = "xml"
	LangHTML HttpRespLangType = "html"
	LangTEXT HttpRespLangType = "text"
)

func (e HttpRespLangType) String() string {
	return string(e)
}

type HttpRespCharset string

const (
	UTF8 HttpRespCharset = "utf-8"
)

func (e HttpRespCharset) String() string {
	return string(e)
}

type FieldSource string

const (
	System FieldSource = "requirement"
	Custom FieldSource = "task"
)

func (e FieldSource) String() string {
	return string(e)
}

type FieldType string

const (
	Input       FieldType = "input"
	TextArea    FieldType = "textarea"
	Password    FieldType = "password"
	Checkbox    FieldType = "checkbox"
	Radio       FieldType = "radio"
	File        FieldType = "file"
	image       FieldType = "image"
	Hidden      FieldType = "hidden"
	Select      FieldType = "select"
	MultiSelect FieldType = "multiselect"

	Button FieldType = "button"
)

func (e FieldType) String() string {
	return string(e)
}

type FieldFormat string

const (
	PlainText FieldFormat = "plainText"
	RichText  FieldFormat = "richText"
)

func (e FieldFormat) String() string {
	return string(e)
}

type ProductStatus string

const (
	Active ProductStatus = "active"
	Closed ProductStatus = "closed"
)

func (e ProductStatus) String() string {
	return string(e)
}

type UsedBy string

const (
	InterfaceDebug UsedBy = "interface_debug"
	CaseDebug      UsedBy = "case_debug"
	//AlternativeCaseDebug UsedBy = "alternative_case_debug"
	DiagnoseDebug UsedBy = "diagnose_debug"
	ScenarioDebug UsedBy = "scenario_debug"
)

type CaseType string

const (
	CaseDefault     CaseType = "default"
	CaseBenchmark   CaseType = "benchmark"   // for alternative cases design
	CaseAlternative CaseType = "alternative" // saved as independent case
)

type ProcessorInterfaceSrc string

const (
	InterfaceSrcDefine   ProcessorInterfaceSrc = "define"
	InterfaceSrcCase     ProcessorInterfaceSrc = "case"
	InterfaceSrcDiagnose ProcessorInterfaceSrc = "diagnose"
	InterfaceSrcCurl     ProcessorInterfaceSrc = "curl"
	InterfaceSrcCustom   ProcessorInterfaceSrc = "custom"
)

type ConditionSrc string

const (
	ConditionSrcPre  ConditionSrc = "pre"
	ConditionSrcPost ConditionSrc = "post"
)

type ConditionType string

const (
	ConditionTypeExtractor      ConditionType = "extractor"
	ConditionTypeCheckpoint     ConditionType = "checkpoint"
	ConditionTypeScript         ConditionType = "script"
	ConditionTypeDatabase       ConditionType = "databaseOpt"
	ConditionTypeResponseDefine ConditionType = "responseDefine"
)

type ConditionCategory string

const (
	ConditionCategoryResult   ConditionCategory = "result"
	ConditionCategoryConsole  ConditionCategory = "console"
	ConditionCategoryAssert   ConditionCategory = "assert"
	ConditionCategoryAll      ConditionCategory = "all"
	ConditionCategoryResponse ConditionCategory = "response"
)

type ExtractorSrc string

const (
	Header ExtractorSrc = "header"
	Body   ExtractorSrc = "body"
	Cookie ExtractorSrc = "cookie"
)

type ExtractorType string

const (
	Boundary  ExtractorType = "boundary"
	JSONPath  ExtractorType = "jsonpath"
	JsonQuery ExtractorType = "jsonquery"
	HtmlQuery ExtractorType = "htmlquery"
	XmlQuery  ExtractorType = "xmlquery"
	Regx      ExtractorType = "regx"
	//FullText  ExtractorType = "fulltext"
)

type CheckpointType string

const (
	ResponseStatus CheckpointType = "responseStatus"
	ResponseHeader CheckpointType = "responseHeader"
	ResponseBody   CheckpointType = "responseBody"
	Judgement      CheckpointType = "judgement"
	ExtractorVari  CheckpointType = "extractorVari"
	Extractor      CheckpointType = "extractor"

	Script CheckpointType = "script"
)

type ExtractorScope string

const (
	Private ExtractorScope = "private" // in current interface
	Public  ExtractorScope = "public"  // shared by other interfaces in serve OR scenario
)

type ExtractorResultType string

const (
	ExtractorResultTypeString ExtractorResultType = "string"
	ExtractorResultTypeNumber ExtractorResultType = "number"
	ExtractorResultTypeObject ExtractorResultType = "object"
)

type ComparisonOperator string

const (
	Equal              ComparisonOperator = "equal"
	NotEqual           ComparisonOperator = "notEqual"
	GreaterThan        ComparisonOperator = "greaterThan"
	GreaterThanOrEqual ComparisonOperator = "greaterThanOrEqual"
	LessThan           ComparisonOperator = "lessThan"
	LessThanOrEqual    ComparisonOperator = "lessThanOrEqual"

	Contain    ComparisonOperator = "contain"
	NotContain ComparisonOperator = "notContain"

	RegularMatch ComparisonOperator = "regularMatch"

	Exist    ComparisonOperator = "exist"
	NotExist ComparisonOperator = "notExist"
)

func (e ComparisonOperator) String() string {
	return string(e)
}

var StringComparisons = []ComparisonOperator{
	Equal,
	NotEqual,
	GreaterThan,
	GreaterThanOrEqual,
	LessThan,
	LessThanOrEqual,

	Contain,
	NotContain,

	RegularMatch,
}
var TextComparisons = []ComparisonOperator{
	Equal,
	NotEqual,

	Contain,
	NotContain,

	RegularMatch,
}

type ValueOperator string

const (
	Get   ValueOperator = "get"
	Set   ValueOperator = "set"
	Clear ValueOperator = "clear"
)

func (e ValueOperator) String() string {
	return string(e)
}

type ProgressStatus string

const (
	Start      ProgressStatus = "start"
	InProgress ProgressStatus = "in_progress"
	End        ProgressStatus = "end"
	Cancel     ProgressStatus = "cancel"
	Error      ProgressStatus = "error"
)

func (e ProgressStatus) String() string {
	return string(e)
}

type ResultStatus string

const (
	Pass    ResultStatus = "pass"
	Fail    ResultStatus = "fail"
	Skip    ResultStatus = "skip"
	Block   ResultStatus = "block"
	Unknown ResultStatus = "unknown"
)

func (e ResultStatus) String() string {
	return string(e)
}

type ProcessorCategory string

const (
	ProcessorRoot ProcessorCategory = "processor_root"
	//ProcessorThreadGroup ProcessorCategory = "processor_thread_group"

	ProcessorInterface ProcessorCategory = "processor_interface"
	ProcessorGroup     ProcessorCategory = "processor_group"
	ProcessorLogic     ProcessorCategory = "processor_logic"
	ProcessorLoop      ProcessorCategory = "processor_loop"
	ProcessorTimer     ProcessorCategory = "processor_timer"
	ProcessorPrint     ProcessorCategory = "processor_print"
	ProcessorVariable  ProcessorCategory = "processor_variable"
	ProcessorAssertion ProcessorCategory = "processor_assertion"

	ProcessorCookie     ProcessorCategory = "processor_cookie"
	ProcessorData       ProcessorCategory = "processor_data"
	ProcessorCustomCode ProcessorCategory = "processor_custom_code"
)

func (e ProcessorCategory) ToString() string {
	return string(e)
}

type ProcessorType string

const (
	ProcessorRootDefault ProcessorType = "processor_root_default"
	//ProcessorThreadDefault ProcessorType = "processor_thread_default"

	ProcessorInterfaceDefault ProcessorType = "processor_interface_default"
	ProcessorGroupDefault     ProcessorType = "processor_group_default"
	ProcessorTimerDefault     ProcessorType = "processor_timer_default"
	ProcessorPrintDefault     ProcessorType = "processor_print_default"

	ProcessorLogicIf   ProcessorType = "processor_logic_if"
	ProcessorLogicElse ProcessorType = "processor_logic_else"

	ProcessorLoopTime  ProcessorType = "processor_loop_time"
	ProcessorLoopIn    ProcessorType = "processor_loop_in"
	ProcessorLoopRange ProcessorType = "processor_loop_range"
	ProcessorLoopUntil ProcessorType = "processor_loop_until"

	ProcessorVariableSet   ProcessorType = "processor_variable_set"
	ProcessorVariableClear ProcessorType = "processor_variable_clear"

	ProcessorAssertionDefault ProcessorType = "processor_assertion_default"

	ProcessorExtractorBoundary  ProcessorType = "processor_extractor_boundary"
	ProcessorExtractorJsonQuery ProcessorType = "processor_extractor_jsonquery"
	ProcessorExtractorHtmlQuery ProcessorType = "processor_extractor_htmlquery"
	ProcessorExtractorXmlQuery  ProcessorType = "processor_extractor_xmlquery"

	ProcessorCookieSet   ProcessorType = "processor_cookie_set"
	ProcessorCookieClear ProcessorType = "processor_cookie_clear"

	ProcessorDataDefault       ProcessorType = "processor_data_default"
	ProcessorCustomCodeDefault ProcessorType = "processor_custom_code"
)

func (e ProcessorType) ToString() string {
	return string(e)
}

type LogType string

const (
	LogRoot      LogType = "root"
	LogInterface LogType = "interface"
	LogProcessor LogType = "processor"
)

func (e LogType) ToString() string {
	return string(e)
}

type ErrorAction string

const (
	ActionContinue        ErrorAction = "continue"
	ActionStartNextThread ErrorAction = "start_next_thread"
	ActionLoop            ErrorAction = "loop"
	ActionStopThread      ErrorAction = "stop_thread"
	ActionStopTest        ErrorAction = "stop_test"
	ActionStopTestNow     ErrorAction = "stop_test_now"
)

func (e ErrorAction) ToString() string {
	return string(e)
}

type DataItSrc string

const (
	SrcFileUpload DataItSrc = "fileUpload"
	SrcDatapool   DataItSrc = "datapool"
)

func (e DataItSrc) ToString() string {
	return string(e)
}

type DataItType string

const (
	Text  DataItType = "text"
	Excel DataItType = "excel"
	//ZenData DataItType = "zendata"
)

func (e DataItType) ToString() string {
	return string(e)
}

type DataFileFormat string

const (
	FormatText    DataFileFormat = "text"
	FormatExcel   DataFileFormat = "excel"
	FormatCsv     DataFileFormat = "csv"
	FormatUnknown DataFileFormat = "unknown"
)

func (e DataFileFormat) ToString() string {
	return string(e)
}

type TimeUnit string

const (
	Second TimeUnit = "sec"
	Minute TimeUnit = "min"
	Hour   TimeUnit = "hour"
)

func (e TimeUnit) ToString() string {
	return string(e)
}

type ExecType string

const (
	ExecStart ExecType = "start"
	ExecStop  ExecType = "stop"

	ExecScenario ExecType = "execScenario"
	ExecPlan     ExecType = "execPlan"
	ExecCase     ExecType = "execCases"
	ExecMessage  ExecType = "execMessage"
)

func (e ExecType) String() string {
	return string(e)
}

type DataType string

const (
	Int    DataType = "int"
	Float  DataType = "float"
	String DataType = "string"
)

func (e DataType) String() string {
	return string(e)
}

type RoleType string

const (
	Admin          RoleType = "admin"
	User           RoleType = "user"
	Tester         RoleType = "tester"
	Developer      RoleType = "developer"
	ProductManager RoleType = "product_manager"
)

func (e RoleType) String() string {
	return string(e)
}

type NodeType string

const (
	NodeElem    NodeType = "elem"
	NodeProp    NodeType = "prop"
	NodeContent NodeType = "content"
	NodeText    NodeType = "text"
)

func (e NodeType) String() string {
	return string(e)
}

type PlaceholderPrefix string

const (
	PlaceholderPrefixDatapool PlaceholderPrefix = "_dp"
	PlaceholderPrefixFunction PlaceholderPrefix = "_func"
)

func (e PlaceholderPrefix) String() string {
	return string(e)
}

type PlaceholderType string

const (
	PlaceholderTypeEnvironmentVariable PlaceholderType = "environment_variable"
	PlaceholderTypeVariable            PlaceholderType = "variable"
	PlaceholderTypeDatapool            PlaceholderType = "datapool"
	PlaceholderTypeFunction            PlaceholderType = "function"
)

func (e PlaceholderType) String() string {
	return string(e)
}

type ParamType string

const (
	ParamTypeString  ParamType = "string"
	ParamTypeNumber  ParamType = "number"
	ParamTypeInteger ParamType = "integer"
)

func (e ParamType) String() string {
	return string(e)
}

type ParamIn string

const (
	ParamInPath   ParamIn = "path"
	ParamInQuery  ParamIn = "query"
	ParamInHeader ParamIn = "header"
	ParamInCookie ParamIn = "cookie"
	ParamInBody   ParamIn = "body"
)

func (e ParamIn) String() string {
	return string(e)
}

type AuditStatus uint

const (
	Init    AuditStatus = 0
	Agreed  AuditStatus = 1
	Refused AuditStatus = 2
)

type TestStatus string

const (
	Draft     TestStatus = "draft"      //草稿
	Disabled  TestStatus = "disabled"   //已禁用
	ToExecute TestStatus = "to_execute" //待执行
	Executed  TestStatus = "executed"   //已执行
)

func (e TestStatus) String() string {
	return string(e)
}

type TestStage string

const (
	UintTest        TestStage = "unit_test"        //单元测试
	IntegrationTest TestStage = "integration_test" //集成测试
	SystemTest      TestStage = "system_test"      //系统测试
	AcceptanceTest  TestStage = "acceptance_test"  //验收测试
)

func (e TestStage) String() string {
	return string(e)
}

type TestType string

//冒烟测试、逻辑验证、异常测试、性能测试、接口合规测试

const (
	ApiTest         TestType = "api_test"         //接口测试
	PerformanceTest TestType = "performance_test" //性能测试
	SmokeTest       TestType = "smoke_test"       //冒烟测试
)

func (e TestType) String() string {
	return string(e)
}

type SwitchStatus uint

const (
	SwitchON  SwitchStatus = 1
	SwitchOFF SwitchStatus = 2
)

type SourceType uint

const (
	SwaggerSync    SourceType = 1
	SwaggerImport  SourceType = 2
	ThirdPartySync SourceType = 3
)

type MockPriority string

const (
	MockPrioritySmart   MockPriority = "smart"
	MockPriorityExample MockPriority = "example"
)

type DataSyncType uint

const (
	FullCover DataSyncType = 1 //完全覆盖
	AutoAdd   DataSyncType = 2 //智能合并
	Add       DataSyncType = 3 //新增
)

func (e DataSyncType) String() string {
	return string(e)
}

type ExpectRequestSelectType string

const (
	KeyValue ExpectRequestSelectType = "keyValue"
	Xpath    ExpectRequestSelectType = "xPath"
	FullText ExpectRequestSelectType = "fullText"
)

func (e ExpectRequestSelectType) String() string {
	return string(e)
}

type AlternativeCaseCategories string

const (
	AlternativeCaseRoot     AlternativeCaseCategories = "root"
	AlternativeCaseCategory AlternativeCaseCategories = "category"
	AlternativeCaseDir      AlternativeCaseCategories = "dir"
	AlternativeCaseParam    AlternativeCaseCategories = "param"
	AlternativeCaseObject   AlternativeCaseCategories = "object"
	AlternativeCaseArray    AlternativeCaseCategories = "array"
	AlternativeCaseProp     AlternativeCaseCategories = "prop"
	AlternativeCaseCase     AlternativeCaseCategories = "case"
)

type AlternativeCaseTypes string

const (
	AlternativeCaseRequired AlternativeCaseTypes = "required"
	AlternativeCaseTyped    AlternativeCaseTypes = "typed"
	AlternativeCaseEnum     AlternativeCaseTypes = "enum"
	AlternativeCaseFormat   AlternativeCaseTypes = "format"
	AlternativeCaseRule     AlternativeCaseTypes = "rule"
)

type AlternativeCaseRules string

const (
	AlternativeCaseRulesMin          AlternativeCaseRules = "min"
	AlternativeCaseRulesMax          AlternativeCaseRules = "max"
	AlternativeCaseRulesMaxLength    AlternativeCaseRules = "maxLength"
	AlternativeCaseRulesMinLength    AlternativeCaseRules = "minLength"
	AlternativeCaseRulesMultipleOf   AlternativeCaseRules = "multipleOf"
	AlternativeCaseRulesExclusiveMin AlternativeCaseRules = "exclusiveMin"
	AlternativeCaseRulesExclusiveMax AlternativeCaseRules = "exclusiveMax"

	AlternativeCaseRulesPattern AlternativeCaseRules = "pattern"
)

func (e AlternativeCaseRules) String() string {
	return string(e)
}

type AlternativeCaseType string

const (
	QueryParam  AlternativeCaseType = "query_param"
	PathParam   AlternativeCaseType = "path_param"
	HeaderParam AlternativeCaseType = "header_param"

	BodyField AlternativeCaseType = "body_field"
	FormField AlternativeCaseType = "form_field"
)

func (e AlternativeCaseType) String() string {
	return string(e)
}

type DatabaseType string

const (
	DbTypeMySql      DatabaseType = "mysql"
	DbTypeSqlServer  DatabaseType = "sqlserver"
	DbTypePostgreSql DatabaseType = "postgreSql"
	DbTypeOracle     DatabaseType = "oracle"
)

func (e DatabaseType) String() string {
	return string(e)
}

type MessageSendStatus string

const (
	MessageCreated            MessageSendStatus = "created"
	MessageSendSuccess        MessageSendStatus = "send_success"
	MessageSendFailed         MessageSendStatus = "send_failed"
	MessageApprovalInProgress MessageSendStatus = "approval_in_progress"
	MessageApprovalAgreed     MessageSendStatus = "approval_agreed"
	MessageApprovalReject     MessageSendStatus = "approval_reject"
)

func (e MessageSendStatus) String() string {
	return string(e)
}

type MessageServiceType string

const (
	ServiceTypeApproval MessageServiceType = "approval"
	ServiceTypeInfo     MessageServiceType = "info"
)

func (e MessageServiceType) String() string {
	return string(e)
}

type MessageSource string

const (
	MessageSourceEndpoint        MessageSource = "endpoint"
	MessageSourceJoinProject     MessageSource = "join_project"
	MessageSourceAuditProjectRes MessageSource = "audit_project_res"
)

func (e MessageSource) String() string {
	return string(e)
}

type ChangedStatus uint

const (
	NoChanged     ChangedStatus = 1
	Changed       ChangedStatus = 2
	IgnoreChanged ChangedStatus = 3
)
