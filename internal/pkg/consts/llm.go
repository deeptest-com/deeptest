package consts

type MetricsType string

const (
	Summarization       MetricsType = "summarization"
	AnswerRelevancy     MetricsType = "answer_relevancy"
	Faithfulness        MetricsType = "faithfulness"
	ContextualPrecision MetricsType = "contextual_precision"
	ContextualRecall    MetricsType = "contextual_recall"
	ContextualRelevancy MetricsType = "contextual_relevancy"
	Hallucination       MetricsType = "hallucination"
	Bias                MetricsType = "bias"
	Toxicity            MetricsType = "toxicity"

	Empty MetricsType = ""
)

func (e MetricsType) String() string {
	return string(e)
}

const (
	StatementsSample = `
{"statements": [
	"Playwright是一个开源的自动化测试库，用于测试浏览器应用。",
	"它支持多种浏览器，如Chrome、Firefox和Safari等。",
	"我是陈琦",
	"Playwright允许开发人员编写用于测试、生成页面截图、自动填写表单、验证和监控等目的的脚本。",
	"这意味着你可以用它来测试你的网站或者Web应用的功能是否按照预期运行。"
	"DeepTest是一个测试工具",
]}`
)
