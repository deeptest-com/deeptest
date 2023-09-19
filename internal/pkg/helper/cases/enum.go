package casesHelper

type OasFieldType string

const (
	OasFieldTypeAny     OasFieldType = "any"
	OasFieldTypeString  OasFieldType = "string"
	OasFieldTypeNumber  OasFieldType = "number"
	OasFieldTypeInteger OasFieldType = "integer"
	OasFieldTypeBoolean OasFieldType = "boolean"
	OasFieldTypeArray   OasFieldType = "array"

	OasFieldTypeObject OasFieldType = "object"
)

func (e OasFieldType) String() string {
	return string(e)
}

type OasFieldFormat string

const (
	// for integer
	OasFieldFormatInt32 OasFieldFormat = "int32"
	OasFieldFormatInt64 OasFieldFormat = "int64"

	// for number
	OasFieldFormatFloat  OasFieldFormat = "float"
	OasFieldFormatDouble OasFieldFormat = "double"

	// for string
	OasFieldFormatDataTime     OasFieldFormat = "data-time"
	OasFieldFormatTime         OasFieldFormat = "time"
	OasFieldFormatEmail        OasFieldFormat = "email"
	OasFieldFormatIdnEmail     OasFieldFormat = "idn-email"
	OasFieldFormatHostname     OasFieldFormat = "hostname"
	OasFieldFormatIdnHostname  OasFieldFormat = "idn-hostname"
	OasFieldFormatIpv4         OasFieldFormat = "ipv4"
	OasFieldFormatIpv6         OasFieldFormat = "ipv6"
	OasFieldFormatUri          OasFieldFormat = "uri"
	OasFieldFormatUriReference OasFieldFormat = "uri-reference"
	OasFieldFormatIri          OasFieldFormat = "iri"
	OasFieldFormatIriReference OasFieldFormat = "iri-reference"
	OasFieldFormatUriTemplate  OasFieldFormat = "uri-template"
	OasFieldFormatJsonPointer  OasFieldFormat = "json-pointer"
	OasFieldFormatRegex        OasFieldFormat = "regex"
	OasFieldFormatUuid         OasFieldFormat = "uuid"
	OasFieldFormatPassword     OasFieldFormat = "password"
	OasFieldFormatByte         OasFieldFormat = "byte"
)

func (e OasFieldFormat) String() string {
	return string(e)
}
