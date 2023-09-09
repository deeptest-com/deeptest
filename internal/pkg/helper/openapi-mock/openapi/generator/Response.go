package mockGenerator

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
)

type Response struct {
	StatusCode  consts.HttpRespCode
	ContentType consts.HttpContentType
	Data        interface{}

	// used by adv mock
	UseAdvMockMock bool
	Content        string
	Headers        []model.EndpointMockExpectResponseHeader
}
