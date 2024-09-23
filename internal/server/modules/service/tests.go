package service

import (
	"encoding/json"
	"fmt"
	serverDomain "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/kataras/iris/v12"
	"net/http"
	"time"
)

type StreamTestService struct {
}

func (s *StreamTestService) Chat(req serverDomain.StreamTestObj, flusher http.Flusher, ctx iris.Context) (err error) {
	bts, err := json.Marshal(req)
	ctx.Writef("data: request is %s\n\n", string(bts))

	for i := 0; i < req.Count; i++ {
		fmt.Printf("\n>>> %d \n", i)

		bts, _ := json.Marshal(iris.Map{"index": i + 1})

		// must add prefix "data:" for openai response
		// must add postfix "\n\n"
		ctx.Writef("data:%s\n\n", string(bts))
		flusher.Flush()

		time.Sleep(1 * time.Second)
	}

	return
}
