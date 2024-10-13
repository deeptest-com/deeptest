package main

import (
	agentExec "github.com/deeptest-com/deeptest/internal/agent/exec"
	"github.com/deeptest-com/deeptest/internal/pkg/domain"
	"log"
)

func main() {
	body := `
{
    "stream": true,
    "model": "qwen2.5:0.5b-instruct",
    "messages": [
        {
            "role": "system",
            "content": "You are a helpful assistant."
        },
        {
            "role": "user",
            "content": "你好"
        }
    ]
}
`

	req := domain.BaseRequest{
		Method: "POST",
		Url:    "http://localhost:11434/v1/chat/completions",
		Body:   body,
	}

	ret, err := agentExec.Post(req)
	if err != nil {
		return
	}

	log.Println(ret)
}
