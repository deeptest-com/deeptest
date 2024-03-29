package thirdPart

type Schema struct {
	FiledName   string  `json:"filedName"`
	Type        string  `json:"type"`
	Required    bool    `json:"required"`
	Properties  Schemas `json:"properties"`
	Description string  `json:"description"`
	Items       *Schema `json:"items"`
}

type Schemas map[string]*Schema
