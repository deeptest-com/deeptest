package interfaceHelper

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"strings"
)

func GetUrl(bodyType consts.HttpContentType) bool {
	return strings.HasPrefix(bodyType.String(), consts.ContentTypeFormUrlencoded.String())
}
