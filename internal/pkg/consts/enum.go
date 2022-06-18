package consts

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

	ContentTypeFormData       HttpContentType = "application/form-data"
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
	LangJSON      HttpRespLangType = "json"
	LangXML       HttpRespLangType = "xml"
	LangHTML      HttpRespLangType = "html"
	LangPlainTEXT HttpRespLangType = "plaintext"
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

type ValueType string

const (
	Int       FieldType = "int"
	String    FieldType = "string"
	Bool      FieldType = "bool"
	IntArr    FieldType = "intArr"
	StringArr FieldType = "stringArr"
	BoolArr   FieldType = "boolArr"
)

func (e ValueType) String() string {
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

type Progress string

const (
	New        Progress = "new"
	InProgress Progress = "inProgress"
	Completed  Progress = "completed"
	Cancel     Progress = "cancel"
)

func (e Progress) String() string {
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

type ExtractorSrc string

const (
	Header ExtractorSrc = "header"
	Body   ExtractorSrc = "body"
)

type ExtractorType string

const (
	Boundary  ExtractorType = "boundary"
	JsonQuery ExtractorType = "jsonquery"
	HtmlQuery ExtractorType = "htmlquery"
	XmlQuery  ExtractorType = "xmlquery"
	//Regular   ExtractorType = "regular"
	//FullText  ExtractorType = "fulltext"
)

type CheckpointType string

const (
	ResponseStatus CheckpointType = "responseStatus"
	ResponseHeader CheckpointType = "responseHeader"
	ResponseBody   CheckpointType = "responseBody"
	Extractor      CheckpointType = "extractor"
)

type CheckpointOperator string

const (
	Contain            CheckpointOperator = "contain"
	Equal              CheckpointOperator = "equal"
	NotEqual           CheckpointOperator = "notEqual"
	GreaterThan        CheckpointOperator = "greaterThan"
	LessThan           CheckpointOperator = "lessThan"
	GreaterThanOrEqual CheckpointOperator = "greaterThanOrEqual"
	LessThanOrEqual    CheckpointOperator = "lessThanOrEqual"
)

func (e CheckpointOperator) String() string {
	return string(e)
}

type CheckpointResult string

const (
	Pass CheckpointResult = "PASS"
	Fail CheckpointResult = "FAIL"
)

type ProcessorCategory string

const (
	ProcessorInterface ProcessorCategory = "processor_interface"

	ProcessorRoot        ProcessorCategory = "processor_root"
	ProcessorThreadGroup ProcessorCategory = "processor_thread_group"
	ProcessorSimple      ProcessorCategory = "processor_simple"
	ProcessorLogic       ProcessorCategory = "processor_logic"
	ProcessorLoop        ProcessorCategory = "processor_loop"
	ProcessorTimer       ProcessorCategory = "processor_timer"
	ProcessorVariable    ProcessorCategory = "processor_variable"
	ProcessorAssertion   ProcessorCategory = "processor_assertion"
	ProcessorExtractor   ProcessorCategory = "processor_extractor"

	ProcessorCookie ProcessorCategory = "processor_cookie"
	ProcessorData   ProcessorCategory = "processor_data"
)

func (e ProcessorCategory) ToString() string {
	return string(e)
}

type ProcessorType string

const (
	ProcessorThreadDefault ProcessorType = "processor_thread_default"
	ProcessorSimpleDefault ProcessorType = "processor_simple_default"
	ProcessorTimerDefault  ProcessorType = "processor_time_default"

	ProcessorLogicIf   ProcessorType = "processor_logic_if"
	ProcessorLogicElse ProcessorType = "processor_logic_else"

	ProcessorLoopRepeatTime  ProcessorType = "processor_loop_repeat_time"
	ProcessorLoopRepeatUntil ProcessorType = "processor_loop_repeat_until"
	ProcessorLoopRepeatIn    ProcessorType = "processor_loop_repeat_in"
	ProcessorLoopRepeatRange ProcessorType = "processor_loop_range"
	ProcessorLoopRepeatBreak ProcessorType = "processor_loop_break"

	ProcessorVariableSet   ProcessorType = "processor_variable_set"
	ProcessorVariableGet   ProcessorType = "processor_variable_get"
	ProcessorVariableClear ProcessorType = "processor_variable_clear"

	ProcessorAssertionEqual      ProcessorType = "processor_assertion_equal"
	ProcessorAssertionNotEqual   ProcessorType = "processor_assertion_not_equal"
	ProcessorAssertionContain    ProcessorType = "processor_assertion_contain"
	ProcessorAssertionNotContain ProcessorType = "processor_assertion_not_contain"

	ProcessorExtractorBoundary  ProcessorType = "processor_extractor_boundary"
	ProcessorExtractorJsonQuery ProcessorType = "processor_extractor_jsonquery"
	ProcessorExtractorHtmlQuery ProcessorType = "processor_extractor_htmlquery"
	ProcessorExtractorXmlQuery  ProcessorType = "processor_extractor_xmlquery"

	ProcessorCookieGet   ProcessorType = "processor_cookie_get"
	ProcessorCookieSet   ProcessorType = "processor_cookie_set"
	ProcessorCookieClear ProcessorType = "processor_cookie_clear"

	ProcessorDataText    ProcessorType = "processor_data_text"
	ProcessorDataExcel   ProcessorType = "processor_data_excel"
	ProcessorDataZenData ProcessorType = "processor_data_zendata"
)

func (e ProcessorType) ToString() string {
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

type DataSource string

const (
	CSV     DataSource = "csv"
	Excel   DataSource = "excel"
	ZenData DataSource = "zendata"
)

func (e DataSource) ToString() string {
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

type ValueAction string

const (
	Get   ValueAction = "get"
	Set   ValueAction = "set"
	Clear ValueAction = "clear"
)

func (e ValueAction) ToString() string {
	return string(e)
}
