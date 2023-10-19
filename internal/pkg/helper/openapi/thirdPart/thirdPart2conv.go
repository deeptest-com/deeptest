package thirdPart

import "github.com/aaronchen2k/deeptest/internal/pkg/helper/openapi"

type thirdPart2conv struct {
}

func NewThirdPart2conv() *thirdPart2conv {
	return new(thirdPart2conv)
}

func (t *thirdPart2conv) Convert(schemas Schemas) (ret openapi.Schema) {

	return

}
