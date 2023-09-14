package main

import (
	"fmt"
	curlHelper "github.com/aaronchen2k/deeptest/internal/pkg/helper/gcurl"
	"log"
)

func main() {
	scurl := `curl -X POST 'http://127.0.0.1:8085/api/v1/mock?reqType=%s' \
  -H 'Cookie: BIDUPSID=88B7FC40D50C2F811E57590167144216;' `

	// 1. json
	scurl = fmt.Sprintf(scurl, "json")
	scurl += ` -H 'Content-Type: application/json'`
	scurl += ` -d '{"name": "aaron", "key": "1"}'` // ` --data-raw '{"name": "aaron", "key": "1"}'`

	// 2. form
	//scurl = fmt.Sprintf(scurl, "form")
	//scurl += ` -H 'Content-Type: application/x-www-form-urlencoded'`
	//scurl += ` -d 'name=aaron&password=123'`

	// 3. file
	//scurl = fmt.Sprintf(scurl, "file")
	////scurl += `-H 'Content-Type: multipart/form-data'`
	//scurl += ` -F name=aaron -F myFile=@/Users/aaron/rd/project/gudi/deeptest/cmd/test/curl/files/file.txt;type=text/plain`

	curl := curlHelper.Parse(scurl)

	wf := curl.CreateTemporary(curl.CreateSession())
	resp, err := wf.Execute()

	log.Println(curl)
	log.Println(resp)
	log.Println(err)
}
