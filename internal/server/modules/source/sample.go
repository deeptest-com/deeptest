package source

import (
	"github.com/aaronchen2k/deeptest"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	_commUtils "github.com/aaronchen2k/deeptest/pkg/lib/comm"
	"path/filepath"
)

type SampleSource struct {
}

func (s *SampleSource) GetSources() (serve *model.Serve, endpoint *model.Endpoint, err error) {
	serveJson, err := deeptest.ReadResData(filepath.Join("res", "sample", "serve.json"))
	_commUtils.JsonDecode(string(serveJson), serve)

	endpointJson, err := deeptest.ReadResData(filepath.Join("res", "sample", "endpoint.json"))
	_commUtils.JsonDecode(string(endpointJson), endpoint)
	return
}
