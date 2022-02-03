package service

import (
	"github.com/aaronchen2k/deeptest/internal/comm/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/lib/file"
	"github.com/aaronchen2k/deeptest/internal/server/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/model"
	"gopkg.in/yaml.v3"
)

type NluPatternService struct {
}

func NewNluPatternService() *NluPatternService {
	return &NluPatternService{}
}

func (s *NluPatternService) Reload() (project model.Project) {
	files, _ := _fileUtils.ListDir(consts.Pattern)

	for _, f := range files {
		content := _fileUtils.ReadFileBuf(f)

		task := serverDomain.NluTask{}
		yaml.Unmarshal(content, &task)

		serverConsts.NluPatterns = append(serverConsts.NluPatterns, task)
	}

	return
}
