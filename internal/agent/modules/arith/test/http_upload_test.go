package rpc

import (
	"encoding/json"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	fileUtils "github.com/aaronchen2k/deeptest/pkg/lib/file"
	httpUtils "github.com/aaronchen2k/deeptest/pkg/lib/http"
	"testing"
)

func TestUpload(t *testing.T) {
	//_zap.Init()
	result := domain.TestResult{Name: "RasaResult Title"}

	zipFile := "/Users/aaron/testResult.zip"

	result.Payload = nil
	uploadResultUrl := httpUtils.GenUrlWithParams("http://localhost:8085/", nil, "client/build/uploadResult")

	files := []string{zipFile}
	extraParams := map[string]string{}
	json, _ := json.Marshal(result)
	extraParams["result"] = string(json)

	fileUtils.Upload(uploadResultUrl, files, extraParams)
}
