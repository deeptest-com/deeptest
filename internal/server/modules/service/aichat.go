package service

import (
	"bytes"
	"encoding/json"
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/config"
	"github.com/aaronchen2k/deeptest/pkg/domain"
	_httpUtils "github.com/aaronchen2k/deeptest/pkg/lib/http"
	"github.com/kataras/iris/v12"
	"net/http"
)

type AiChatService struct {
}

func (s *AiChatService) KnowledgeBaseChat(req v1.KnowledgeBaseChatReq, flusher http.Flusher, ctx iris.Context) (ret _domain.PageData, err error) {
	url := _httpUtils.AddSepIfNeeded(config.CONFIG.ChatChatUrl) + "chat/knowledge_base_chat"

	bts, err := json.Marshal(req)

	reader := bytes.NewReader(bts)
	request, err := http.NewRequest("POST", url, reader)
	if err != nil {
		return
	}

	request.Header.Set("Cache-Control", "no-cache")
	request.Header.Set("Accept", "text/event-stream")
	request.Header.Set("Connection", "keep-alive")

	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return
	}

	for {
		data := make([]byte, 1024)
		_, err1 := resp.Body.Read(data)
		if err1 != nil {
			break
		}

		// must with prefix "data:" which is from openai response msg,
		// must add a postfix "\n\n"
		ctx.Writef("%s\n\n", string(data))

		flusher.Flush()
	}

	return
}

func (s *AiChatService) ListValidModel(v1.ChatChatModelReq) (ret []v1.ChatChatModel, err error) {
	// now no local models, just return one online openai-api llm
	ret = append(ret, v1.ChatChatModel{
		Type: "online",
		Code: "openai-api",
		Name: "OpenAI",
	})

	//url := _httpUtils.AddSepIfNeeded(config.CONFIG.ChatChatUrl) + "llm_model/list_running_models"
	//
	//body := v1.ChatChatModelReq{
	//	ControllerAddress: config.CONFIG.ChatChatControllerAddress,
	//}
	//
	//bytes, err := _httpUtils.Post(url, body)
	//if err != nil {
	//	return
	//}
	//
	// resp := iris.Map{}
	//err = json.Unmarshal(bytes, &resp)
	//if err != nil {
	//	return
	//}
	//
	// ret = resp["data"].([]v1.ChatChatModel)

	return
}

func (s *AiChatService) ListKnowledgeBase() (ret []interface{}, err error) {
	url := _httpUtils.AddSepIfNeeded(config.CONFIG.ChatChatUrl) + "knowledge_base/list_knowledge_bases"

	bytes, err := _httpUtils.Get(url)
	if err != nil {
		return
	}

	resp := iris.Map{}
	err = json.Unmarshal(bytes, &resp)
	if err != nil {
		return
	}

	ret = resp["data"].([]interface{})

	return
}
