package execUtils

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
)

func DatapoolToVarKeyValuePairs(dpData []map[string]interface{}) (ret []domain.VarKeyValuePair) {
	for _, item := range dpData {
		ret = append(ret, item)
	}

	return
}
