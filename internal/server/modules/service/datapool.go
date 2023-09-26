package service

import (
	"encoding/json"
	"fmt"
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	_domain "github.com/aaronchen2k/deeptest/pkg/domain"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"github.com/xuri/excelize/v2"
	"strings"
)

type DatapoolService struct {
	DatapoolRepo *repo.DatapoolRepo `inject:""`
}

func (s *DatapoolService) Paginate(req v1.DatapoolReqPaginate) (ret _domain.PageData, err error) {
	ret, err = s.DatapoolRepo.Paginate(req)
	return
}

func (s *DatapoolService) Get(id uint) (model.Datapool, error) {
	return s.DatapoolRepo.Get(id)
}
func (s *DatapoolService) GetByName(name string, projectId uint) (model.Datapool, error) {
	return s.DatapoolRepo.GetByName(name, projectId)
}

func (s *DatapoolService) Save(req *model.Datapool, userId uint) (err error) {
	return s.DatapoolRepo.Save(req, userId)
}

func (s *DatapoolService) Delete(id uint) (err error) {
	return s.DatapoolRepo.Delete(id)
}
func (s *DatapoolService) Disable(id uint) (err error) {
	return s.DatapoolRepo.Disable(id)
}

func (s *DatapoolService) ReadExcel(pth string) (ret [][]interface{}, err error) {
	excl, err := excelize.OpenFile(pth)
	if err != nil {
		logUtils.Info("read upload file as excel failed")
		return
	}

	sht := excl.GetSheetList()[0]
	lines, err := excl.GetRows(sht)

	for _, line := range lines {
		var row []interface{}
		for _, col := range line {
			row = append(row, strings.TrimSpace(col))
		}

		ret = append(ret, row)
	}

	return
}

func (s *DatapoolService) ListForExec(projectId uint) (ret domain.Datapools, error interface{}) {
	ret = domain.Datapools{}

	datapools, err := s.DatapoolRepo.ListForExec(projectId)
	if err != nil {
		return
	}

	for _, datapool := range datapools {
		var arr [][]interface{}
		json.Unmarshal([]byte(datapool.Data), &arr)
		if len(arr) == 0 {
			continue
		}
		var headers []string
		for _, col := range arr[0] {
			headers = append(headers, fmt.Sprintf("%v", col))
		}

		var items []domain.VarKeyValuePair

		for rowIndex, row := range arr {
			if rowIndex == 0 {
				continue
			}

			item := map[string]interface{}{}
			for colIndex, col := range row {
				item[headers[colIndex]] = col
			}

			items = append(items, item)
		}

		ret[datapool.Name] = items
	}

	return
}
