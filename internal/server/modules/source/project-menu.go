package source

import (
	repo2 "github.com/aaronchen2k/deeptest/internal/server/modules/repo"
)

type ProjectMenuSource struct {
	ProjectMenuRepo *repo2.ProjectMenuRepo `inject:""`
}

func NewProjectMenuSource() *ProjectMenuSource {
	return &ProjectMenuSource{}
}

func (s *ProjectMenuSource) Init() (err error) {
	defer s.ProjectMenuRepo.BatchInitData("buttonLevel")
	defer s.ProjectMenuRepo.BatchInitData("secondLevel")
	defer s.ProjectMenuRepo.BatchInitData("firstLevel")
	defer s.ProjectMenuRepo.DeleteAllData()

	return
}
