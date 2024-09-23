package service

import (
	"encoding/json"
	"fmt"
	serverDomain "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/kataras/iris/v12"
	"net/http"
)

type StreamTestService struct {
}

func (s *StreamTestService) Chat(req serverDomain.StreamTestObj, flusher http.Flusher, ctx iris.Context) (err error) {
	bts, err := json.Marshal(req)
	ctx.Writef("%s\n\n", string(bts))

	for i := 0; i < req.Count; i++ {
		fmt.Printf("\n>>> %d \n", i)

		// must with prefix "data:" for openai response
		// must add a postfix "\n\n"
		ctx.Writef("index = %d\n\n", i)
		flusher.Flush()
	}

	return
}
