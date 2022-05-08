package serverConsts

type NodeCreateMode string

const (
	Brother NodeCreateMode = "brother"
	Child   NodeCreateMode = "child"
)

func (e NodeCreateMode) String() string {
	return string(e)
}

type NodeCreateType string

const (
	Dir  NodeCreateType = "dir"
	Node NodeCreateType = "node"
)

func (e NodeCreateType) String() string {
	return string(e)
}

type DropPos int

const (
	Before DropPos = -1
	Inner  DropPos = 0
	After  DropPos = 1
)

func (e DropPos) Int() int {
	return int(e)
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
