package casesHelper

import (
	"fmt"
)

var (
	ExampleInteger = 100
	ExampleFloat   = 10.01

	ExampleEmpty  = ""
	ExampleString = "string"
	ExampleArray  = "abc,123"

	ExampleStringFloat = fmt.Sprintf("%v", ExampleFloat)

	Category = map[string]string{
		"query":  "查询参数",
		"path":   "路径参数",
		"header": "请求头",
		"body":   "请求体",
	}
)
