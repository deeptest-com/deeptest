package test

import (
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	_commUtils "github.com/aaronchen2k/deeptest/pkg/lib/comm"
	_fileUtils "github.com/aaronchen2k/deeptest/pkg/lib/file"
	"testing"
)

func TestSample(t *testing.T) {
	var endpoint model.Endpoint
	endpointJson := _fileUtils.ReadFile("../../../config/sample/endpoint.json")
	_commUtils.JsonDecode(endpointJson, &endpoint)
	fmt.Println(_fileUtils.GetWorkDir(), endpointJson, endpoint)
}
