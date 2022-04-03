package consts

type HttpMethod string

const (
	GET    HttpMethod = "GET"
	POST   HttpMethod = "POST"
	PUT    HttpMethod = "PUT"
	DELETE HttpMethod = "DELETE"
)

func (e HttpMethod) string() string {
	return string(e)
}

type HttpRespCode int

const (
	OK HttpRespCode = 200
)

func (e HttpRespCode) Int() int {
	return int(e)
}

type HttpReqBodyType string

const (
	ReqJSON HttpReqBodyType = "application/json"
	ReqXML  HttpReqBodyType = "application/xml"
	ReqHTML HttpReqBodyType = "text/html"
	ReqTEXT HttpReqBodyType = "text/text"

	ReqFormData       HttpReqBodyType = "application/form-data"
	ReqFormUrlencoded HttpReqBodyType = "application/x-www-form-urlencoded"
)

func (e HttpReqBodyType) String() string {
	return string(e)
}

type HttpRespBodyType string

const (
	RespJSON HttpRespBodyType = "json"
	RespXML  HttpRespBodyType = "xml"
	RespHTML HttpRespBodyType = "html"
	RespTEXT HttpRespBodyType = "text"
)

func (e HttpRespBodyType) String() string {
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
	Created    Progress = "created"
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
