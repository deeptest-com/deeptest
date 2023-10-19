package thirdPart

type Schema struct {
	FiledName  string
	Type       string `json:"type"`
	Properties Schemas
}

type Schemas map[string]*Schema
