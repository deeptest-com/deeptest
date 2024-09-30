package pkg

import (
	"github.com/deeptest-com/deeptest/internal/pkg/config"
	"github.com/deeptest-com/deeptest/internal/pkg/domain"
	commUtils "github.com/deeptest-com/deeptest/internal/pkg/utils"
	_commUtils "github.com/deeptest-com/deeptest/pkg/lib/comm"
	"strconv"
	"time"
)

func GetHeaders(body string) (headers []domain.Header) {
	xNancalTimestamp := strconv.FormatInt(time.Now().UnixMilli(), 10)
	xNancalNonceStr := _commUtils.RandStr(8)

	if body != "" {
		body = commUtils.CompressedJson(body)
	}

	headers = []domain.Header{
		{
			Name:  "x-nancal-appkey",
			Value: config.CONFIG.Saas.ApiSign.AppKey,
		},
		{
			Name:  "x-nancal-timestamp",
			Value: xNancalTimestamp,
		},
		{
			Name:  "x-nancal-nonce-str",
			Value: xNancalNonceStr,
		},
		{
			Name:  "x-nancal-sign",
			Value: _commUtils.GetSign(config.CONFIG.Saas.ApiSign.AppKey, config.CONFIG.Saas.ApiSign.AppSecret, xNancalNonceStr, xNancalTimestamp, body),
		},
	}

	return
}
