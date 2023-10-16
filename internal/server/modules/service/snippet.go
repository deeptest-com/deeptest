package service

import (
	jslibHelper "github.com/aaronchen2k/deeptest/internal/pkg/helper/jslib"
	scriptHelper "github.com/aaronchen2k/deeptest/internal/pkg/helper/script"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	fileUtils "github.com/aaronchen2k/deeptest/pkg/lib/file"
	"github.com/snowlyg/helper/dir"
	"path/filepath"
	"time"
)

var (
	JslibsDeclares []string
)

type SnippetService struct {
	SnippetRepo *repo.SnippetRepo `inject:""`
	JslibRepo   *repo.JslibRepo   `inject:""`
}

func (s *SnippetService) ListJslibNames(projectId int) (names []string, err error) {
	libs, _ := s.JslibRepo.List("", projectId, true)

	for _, po := range libs {
		names = append(names, po.Name)
	}

	return
}

func (s *SnippetService) Get(name scriptHelper.ScriptType) (po jslibHelper.Jslib, err error) {
	script := scriptHelper.GetScript(name)

	po = jslibHelper.Jslib{
		Script: script,
	}
	return
}

func (s *SnippetService) GetJslibs(projectId int) (pos []jslibHelper.Jslib, err error) {
	//if JslibsDeclares == nil {

	JslibsDeclares = nil
	libs, _ := s.JslibRepo.List("", projectId, true)

	for _, lib := range libs {
		pth := filepath.Join(dir.GetCurrentAbPath(), lib.TypesFile)
		content := fileUtils.ReadFile(pth)

		JslibsDeclares = append(JslibsDeclares, content)
	}
	//}

	for _, item := range JslibsDeclares {
		po := jslibHelper.Jslib{
			Script: item,
		}
		pos = append(pos, po)
	}

	return
}

func (s *SnippetService) GetJslibsForAgent(loadedLibs map[uint]time.Time, projectId int) (tos []jslibHelper.Jslib, err error) {
	pos, _ := s.JslibRepo.List("", projectId, true)

	for _, po := range pos {
		pth := filepath.Join(dir.GetCurrentAbPath(), po.ScriptFile)
		content := fileUtils.ReadFile(pth)

		updateTime := po.UpdatedAt
		if updateTime == nil {
			updateTime = po.CreatedAt
		}

		loadTime, found := loadedLibs[po.ID]
		if !found || loadTime.Before(*updateTime) {
			to := jslibHelper.Jslib{
				Id:        po.ID,
				Name:      po.Name,
				Script:    content,
				UpdatedAt: *updateTime,
			}
			tos = append(tos, to)
		}
	}

	return
}
