package serverDomain

type SnippetRes struct {
	Label    string       `json:"label""`
	Value    string       `json:"value""`
	Desc     string       `json:"desc""`
	Children []SnippetRes `json:"children""`
	Key      string       `json:"key""`
}
