package source

import (
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	_commUtils "github.com/aaronchen2k/deeptest/pkg/lib/comm"
	_fileUtils "github.com/aaronchen2k/deeptest/pkg/lib/file"
)

type SampleSource struct {
}

func (s *SampleSource) GetSources() (serve *model.Serve, endpoint *model.Endpoint, err error) {
	serveJson := _fileUtils.ReadFile("./config/sample/serve.json")
	_commUtils.JsonDecode(serveJson, serve)
	endpointJson := _fileUtils.ReadFile("./config/sample/endpoint.json")
	_commUtils.JsonDecode(endpointJson, endpoint)
	return
}
