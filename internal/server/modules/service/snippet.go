package service

import (
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
)

type SnippetService struct {
	SnippetRepo *repo.SnippetRepo `inject:""`
}

func (s *SnippetService) Get(name string) (po model.Snippet, err error) {
	po, err = s.SnippetRepo.Get(name)
	return
}
