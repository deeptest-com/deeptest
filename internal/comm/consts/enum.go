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
