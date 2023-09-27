package _fileUtils

import (
	"bytes"
	"errors"
	"fmt"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"github.com/oklog/ulid/v2"
	"github.com/snowlyg/helper/str"
	"io"
	"io/ioutil"
	"math/rand"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
	"time"
)

func Upload(url string, files []string, extraParams map[string]string) {
	bodyBuffer := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuffer)

	for _, file := range files {
		fw, _ := bodyWriter.CreateFormFile("file", file)
		f, _ := os.Open(file)
		defer f.Close()
		io.Copy(fw, f)
	}

	for key, value := range extraParams {
		_ = bodyWriter.WriteField(key, value)
	}

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	resp, err := http.Post(url, contentType, bodyBuffer)
	defer resp.Body.Close()

	if err != nil {
		fmt.Printf("failed to upload file %s", err.Error())
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("failed to parse upload response %s", err.Error())
	}

	fmt.Printf("upload status %s, body %s", resp.Status, string(respBody))
}

func GetUploadFileName(name string) (ret string, err error) {
	fns := strings.Split(strings.TrimPrefix(name, "./"), ".")
	if len(fns) != 2 {
		msg := fmt.Sprintf("文件名错误 %s", name)

		logUtils.Info(msg)
		err = errors.New(msg)

		return
	}

	base := fns[0]
	ext := fns[1]

	entropy := rand.New(rand.NewSource(time.Now().UnixNano()))
	ms := ulid.Timestamp(time.Now())
	rand, _ := ulid.New(ms, entropy)

	ret = str.Join(base, "-", strings.ToLower(rand.String()), ".", ext)

	return
}
