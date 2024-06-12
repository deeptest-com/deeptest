package serverDomain

type KnowledgeBaseChatReq struct {
	Query             string `json:"query"`
	KnowledgeBaseName string `json:"knowledge_base_name"`
	TopK              int    `json:"top_k"`
	ScoreThreshold    int    `json:"score_threshold"`
	History           []struct {
		Role    string `json:"role"`
		Content string `json:"content"`
	} `json:"history"`
	Stream      bool    `json:"stream"`
	ModelName   string  `json:"model_name"`
	Temperature float64 `json:"temperature"`
	MaxTokens   int     `json:"max_tokens"`
	PromptName  string  `json:"prompt_name"`
}

type ChatChatModelReq struct {
	ControllerAddress string `json:"controller_address"`
	Placeholder       string `json:"placeholder"`
}

type ChatChatModel struct {
	Type string `json:"type"`
	Code string `json:"code"`
	Name string `json:"name"`

	Host       string `json:"host"`
	Port       int    `json:"port"`
	Device     string `json:"device"`
	InferTurbo bool   `json:"infer_turbo"`
	ModelPath  string `json:"model_path"`
}
