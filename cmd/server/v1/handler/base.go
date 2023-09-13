package handler

import (
	"fmt"
	mockGenerator "github.com/aaronchen2k/deeptest/internal/pkg/helper/openapi-mock/openapi/generator"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"google.golang.org/protobuf/proto"
	"time"
)

type BaseCtrl struct {
}

func (c *BaseCtrl) WriteRespByContentType(resp mockGenerator.Response, ctx iris.Context) {
	if resp.DelayTime > 0 {
		time.Sleep(time.Duration(resp.DelayTime) * time.Second)
	}

	data := resp.Data
	ctx.StatusCode(resp.StatusCode.Int())
	ctx.ContentType(resp.ContentType.String())

	/** string **/
	switch resp.ContentType {
	case context.ContentTextHeaderValue, context.ContentHTMLHeaderValue:
		str := fmt.Sprintf("%v", data)
		ctx.WriteString(str)

	case context.ContentMarkdownHeaderValue:
		str := fmt.Sprintf("%v", data)
		ctx.Markdown([]byte(str))

	// proto.Message
	case context.ContentProtobufHeaderValue:
		msg, ok := data.(proto.Message)
		if ok {
			ctx.Protobuf(msg)
		}

	// map
	case context.ContentJSONProblemHeaderValue, context.ContentXMLProblemHeaderValue:
		ctx.Problem(data)

	/** object to marshal **/
	case context.ContentJSONHeaderValue:
		ctx.JSON(data)

	case context.ContentJavascriptHeaderValue:
		ctx.JSONP(data)

	case context.ContentXMLHeaderValue, context.ContentXMLUnreadableHeaderValue:
		ctx.XML(data)

	case context.ContentYAMLHeaderValue:
		ctx.YAML(data)

	case context.ContentYAMLTextHeaderValue:
		ctx.TextYAML(data)

	case context.ContentMsgPackHeaderValue, context.ContentMsgPack2HeaderValue:
		ctx.MsgPack(data)

	default:
	}
}
