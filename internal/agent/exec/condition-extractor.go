package agentExec

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	extractorHelper "github.com/aaronchen2k/deeptest/internal/pkg/helper/extractor"
)

func ExecExtract(extractor domain.ExtractorBase, resp domain.DebugResponse) (result string, err error) {
	result, err = extractorHelper.Extract(extractor, resp)

	return
}
