package mockGenerator

import "github.com/aaronchen2k/deeptest/internal/pkg/consts"

type Response struct {
	StatusCode  consts.HttpRespCode
	ContentType consts.HttpContentType
	Data        interface{}
}
