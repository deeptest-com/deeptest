package jmeterHelper

type JmeterElement string

const (
	Arguments              JmeterElement = "Arguments"
	BeanShellPostProcessor JmeterElement = "BeanShellPostProcessor"
	BeanShellPreProcessor  JmeterElement = "BeanShellPreProcessor"
	CacheManager           JmeterElement = "CacheManager"
	CookieManager          JmeterElement = "CookieManager"
	CounterConfig          JmeterElement = "CounterConfig"
	CSVDataSet             JmeterElement = "CSVDataSet"
	DNSCacheManager        JmeterElement = "DNSCacheManager"
	Fake                   JmeterElement = "Fake"
	ForeachController      JmeterElement = "ForeachController"
	GenericController      JmeterElement = "GenericController"
	HashTree               JmeterElement = "hashTree"
	HTTPSamplerProxy       JmeterElement = "HTTPSamplerProxy"
	IfController           JmeterElement = "IfController"
	InterleaveControl      JmeterElement = "InterleaveControl"
	JmeterTestPlan         JmeterElement = "jmeterTestPlan"
	JSR223PostProcessor    JmeterElement = "JSR223PostProcessor"
	JSR223PreProcessor     JmeterElement = "JSR223PreProcessor"
	SteppingThreadGroup    JmeterElement = "SteppingThreadGroup"
	LoopController         JmeterElement = "LoopController"
	OnceOnlyController     JmeterElement = "OnceOnlyController"
	PostThreadGroup        JmeterElement = "PostThreadGroup"
	RandomController       JmeterElement = "RandomController"
	RandomVariableConfig   JmeterElement = "RandomVariableConfig"
	ResultCollector        JmeterElement = "ResultCollector"
	RunTime                JmeterElement = "RunTime"
	SetupThreadGroup       JmeterElement = "SetupThreadGroup"
	TestPlan               JmeterElement = "TestPlan"
	ThreadGroup            JmeterElement = "ThreadGroup"
	ThroughputController   JmeterElement = "ThroughputController"
	TransactionController  JmeterElement = "TransactionController"
	UniformRandomTimer     JmeterElement = "UniformRandomTimer"
	WhileController        JmeterElement = "WhileController"
	XPath2Extractor        JmeterElement = "XPath2Extractor"
	XPathAssertion         JmeterElement = "XPathAssertion"
	XPathExtractor         JmeterElement = "XPathExtractor"
)

func (e JmeterElement) String() string {
	return string(e)
}
