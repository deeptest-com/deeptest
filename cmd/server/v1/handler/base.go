package handler

import (
	"fmt"
	mockGenerator "github.com/aaronchen2k/deeptest/internal/pkg/helper/openapi-mock/openapi/generator"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"google.golang.org/protobuf/proto"
)

type BaseCtrl struct {
}

func (c *MockCtrl) WriteResp(resp mockGenerator.Response, ctx iris.Context) {
	data := resp.Data
	ctx.StatusCode(resp.StatusCode)
	ctx.ContentType(resp.ContentType)

	str := fmt.Sprintf("%v", resp.Data)

	switch resp.ContentType {
	case context.ContentTextHeaderValue, context.ContentHTMLHeaderValue:
		ctx.WriteString(str)

	case context.ContentMarkdownHeaderValue:
		ctx.Markdown([]byte(str))

	case context.ContentJSONHeaderValue:
		ctx.JSON(data)

	case context.ContentJSONProblemHeaderValue, context.ContentXMLProblemHeaderValue:
		ctx.Problem(data)

	case context.ContentJavascriptHeaderValue:
		ctx.JSONP(data)

	case context.ContentXMLHeaderValue, context.ContentXMLUnreadableHeaderValue:
		ctx.XML(data)

	case context.ContentYAMLHeaderValue:
		ctx.YAML(data)

	case context.ContentYAMLTextHeaderValue:
		ctx.TextYAML(data)

	case context.ContentProtobufHeaderValue:
		msg, ok := data.(proto.Message)
		if ok {
			ctx.Protobuf(msg)
		}

	case context.ContentMsgPackHeaderValue, context.ContentMsgPack2HeaderValue:
		ctx.MsgPack(data)

	default:
	}
}
