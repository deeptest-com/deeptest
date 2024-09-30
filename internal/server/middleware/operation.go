package middleware

import (
	"bytes"
	"github.com/deeptest-com/deeptest/internal/pkg/consts"
	"github.com/deeptest-com/deeptest/internal/server/core/dao"
	"github.com/deeptest-com/deeptest/internal/server/modules/model"
	logUtils "github.com/deeptest-com/deeptest/pkg/lib/log"
	"github.com/deeptest-com/deeptest/saas/common"
	"io/ioutil"
	"strings"
	"time"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/snowlyg/multi"
	"go.uber.org/zap"
)

// OperationRecord 操作日志中间件
func OperationRecord() iris.Handler {
	return func(ctx iris.Context) {
		var body []byte
		var err error

		// 上传文件记录日志文件数据太大
		if !strings.Contains(ctx.Path(), "/api/v1/upload") {
			body, err = ctx.GetBody()
			if err != nil {
				logUtils.Errorf("获取请求内容错误 %s", zap.String("错误:", err.Error()))
			} else {
				ctx.Request().Body = ioutil.NopCloser(bytes.NewBuffer(body))
			}
		}

		userId := multi.GetUserId(ctx)
		data := string(body)
		if len(data) > 1000 {
			data = "BIG DATA"
		}
		record := model.Oplog{
			Ip:     ctx.RemoteAddr(),
			Method: ctx.Method(),
			Path:   ctx.Path(),
			Agent:  ctx.Request().UserAgent(),
			Body:   data,
			UserID: userId,
		}

		writer := responseBodyWriter{
			ResponseWriter: ctx.ResponseWriter().Clone(),
			body:           &bytes.Buffer{},
		}
		ctx.ResetResponseWriter(writer)
		now := time.Now()

		ctx.Next()

		latency := time.Since(now)
		errorMessage := ""
		if ctx.GetErr() != nil {
			errorMessage = ctx.GetErr().Error()
		}
		record.ErrorMessage = errorMessage
		record.Status = ctx.GetStatusCode()
		record.Latency = latency
		record.Resp = writer.body.String()
		tenantId := common.GetTenantId(ctx)
		if err := CreateOplog(tenantId, record); err != nil {
			logUtils.Errorf("生成日志错误 %s", zap.String("错误:", err.Error()))
		}
	}
}

type responseBodyWriter struct {
	context.ResponseWriter
	body *bytes.Buffer
}

func (r responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}

// CreateOplog
func CreateOplog(tenantId consts.TenantId, ol model.Oplog) error {
	err := dao.GetDB(tenantId).Model(&model.Oplog{}).Create(&ol).Error
	if err != nil {
		logUtils.Errorf("生成系统日志错误 %s", zap.String("错误:", err.Error()))
		return err
	}
	return nil
}
