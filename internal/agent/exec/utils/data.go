package agentUtils

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	"github.com/aaronchen2k/deeptest/pkg/lib/file"
	"github.com/xuri/excelize/v2"
	"path"
	"path/filepath"
	"strings"
)

func ReadDataFromText(url, separator string) (ret []domain.Variables, err error) {
	content := _fileUtils.ReadFile(url)
	arr := strings.Split(strings.ReplaceAll(content, "\r\n", "\n"), "\n")
	if len(arr) < 2 {
		return
	}

	colNameMap := map[int]string{}
	cols := strings.Split(arr[0], separator)
	for index, col := range cols {
		colNameMap[index] = col
	}

	for index, line := range arr {
		if index == 0 {
			continue
		}

		cols := strings.Split(line, separator)
		mp := map[string]interface{}{}
		for index, col := range cols {
			mp[colNameMap[index]] = col
		}
		ret = append(ret, mp)
	}

	return
}
func ReadDataFromExcel(url string) (ret []domain.Variables, err error) {
	excel, err := excelize.OpenFile(url)
	if err != nil {
		return
	}

	if len(excel.GetSheetList()) == 0 {
		return
	}

	firstSheet := excel.GetSheetList()[0]

	rows, err := excel.GetRows(firstSheet)
	if len(rows) < 2 {
		return
	}

	colNameMap := map[int]string{}
	for index, col := range rows[0] {
		col = strings.Replace(col, "'", "''", -1)
		colNameMap[index] = col
	}

	for rowIndex, row := range rows {
		if rowIndex == 0 {
			continue
		}

		mp := map[string]interface{}{}
		for index, col := range row {
			col = strings.Replace(col, "'", "''", -1)
			mp[colNameMap[index]] = col
		}
		ret = append(ret, mp)
	}

	return
}

func DownloadUploadedFile(uri string) (ret string, err error) {
	url := path.Join(consts.ServerUrl, uri)

	_, name := path.Split(uri)
	dist := filepath.Join(consts.TmpDir, "download", name)

	_fileUtils.Download(url, dist)

	return
}
