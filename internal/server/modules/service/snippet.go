package service

import (
	scriptHelper "github.com/aaronchen2k/deeptest/internal/pkg/helper/script"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	fileUtils "github.com/aaronchen2k/deeptest/pkg/lib/file"
	"github.com/snowlyg/helper/dir"
	"path/filepath"
)

var (
	JslibsDeclares []string
)

type SnippetService struct {
	SnippetRepo *repo.SnippetRepo `inject:""`
	JslibRepo   *repo.JslibRepo   `inject:""`
}

func (s *SnippetService) Get(name scriptHelper.ScriptType) (po model.Snippet, err error) {
	script := scriptHelper.GetScript(name)

	po = model.Snippet{
		Script: script,
	}
	return
}

func (s *SnippetService) GetJslibs() (pos []model.Snippet, err error) {
	//if JslibsDeclares == nil {

	JslibsDeclares = nil
	libs, _ := s.JslibRepo.List("")

	for _, lib := range libs {
		pth := filepath.Join(dir.GetCurrentAbPath(), lib.TypesFile)
		content := fileUtils.ReadFile(pth)

		JslibsDeclares = append(JslibsDeclares, content)
	}
	//}

	for _, item := range JslibsDeclares {
		po := model.Snippet{
			Script: item,
		}
		pos = append(pos, po)
	}

	return
}
