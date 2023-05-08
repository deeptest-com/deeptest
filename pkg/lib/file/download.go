package _fileUtils

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
)

func Download(url string, dst string) (err error) {
	fmt.Printf("DownloadToFile From: %s to %s.\n", url, dst)

	MkDirIfNeeded(filepath.Dir(dst))

	var data []byte
	data, err = HTTPDownload(url)
	if err == nil {
		fmt.Printf("file downloaded %s", url)

		RmDir(dst)

		err = WriteDownloadFile(dst, data)
		if err == nil {
			fmt.Printf("file %s saved to %s", url, dst)
		}
	}

	return
}

func HTTPDownload(uri string) ([]byte, error) {
	res, err := http.Get(uri)
	if err != nil {
		fmt.Printf(err.Error())
	}
	defer res.Body.Close()
	d, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf(err.Error())
	}
	return d, err
}

func WriteDownloadFile(dst string, d []byte) error {
	err := ioutil.WriteFile(dst, d, 0444)
	if err != nil {
		fmt.Printf(err.Error())
	}
	return err
}
