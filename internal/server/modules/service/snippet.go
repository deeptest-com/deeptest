package service

import (
	scriptHelper "github.com/aaronchen2k/deeptest/internal/pkg/helper/script"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
)

type SnippetService struct {
	SnippetRepo *repo.SnippetRepo `inject:""`
}

func (s *SnippetService) Get(name scriptHelper.ScriptType) (po model.Snippet, err error) {
	script := scriptHelper.GetScript(name)

	po = model.Snippet{
		Script: script,
	}
	return
}
