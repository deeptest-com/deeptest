package _fileUtils

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
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
