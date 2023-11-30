package curlHelper

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"regexp"
	"strings"

	"github.com/474420502/requests"
)

// CURL 信息结构
type CURL struct {
	ParsedURL *url.URL
	Method    string
	Header    http.Header
	CookieJar http.CookieJar
	Cookies   []*http.Cookie

	ContentType string
	Body        *bytes.Buffer

	Auth     *requests.BasicAuth
	Timeout  int // second
	Insecure bool

	// ITask   string
	// Crontab string
	// Name    string
}

// New new 一个 curl 出来
func New() *CURL {

	u := &CURL{}
	u.Insecure = false

	u.Header = make(http.Header)
	u.CookieJar, _ = cookiejar.New(nil)
	u.Body = bytes.NewBuffer(nil)
	u.Timeout = 30

	return u
}

func (curl *CURL) String() string {
	if curl != nil {
		return fmt.Sprintf("Method: %s\nParsedURL: %s\nHeader: %s\nCookie: %s",
			curl.Method, curl.ParsedURL.String(), curl.Header, curl.Cookies)
	}
	return ""
}

// Execute 直接执行curlbash
func Execute(curlbash string) (*requests.Response, error) {
	return Parse(curlbash).CreateTemporary(nil).Execute()
}

// CreateSession 创建Session
func (curl *CURL) CreateSession() *requests.Session {
	ses := requests.NewSession()
	ses.SetHeader(curl.Header)
	ses.SetCookies(curl.ParsedURL, curl.Cookies)

	ses.Config().SetTimeout(curl.Timeout)

	if curl.Auth != nil {
		ses.Config().SetBasicAuth(curl.Auth)
	}

	if curl.Insecure {
		ses.Config().SetInsecure(curl.Insecure)
	}

	return ses
}

// CreateTemporary 根据Session 创建Temporary
func (curl *CURL) CreateTemporary(ses *requests.Session) *requests.Temporary {
	var wf *requests.Temporary

	if ses == nil {
		ses = curl.CreateSession()
	}

	curl.Method = strings.ToUpper(curl.Method)

	switch curl.Method {
	case "HEAD":
		wf = ses.Head(curl.ParsedURL.String())
	case "GET":
		wf = ses.Get(curl.ParsedURL.String())
	case "POST":
		wf = ses.Post(curl.ParsedURL.String())
	case "PUT":
		wf = ses.Put(curl.ParsedURL.String())
	case "PATCH":
		wf = ses.Patch(curl.ParsedURL.String())
	case "OPTIONS":
		wf = ses.Options(curl.ParsedURL.String())
	case "DELETE":
		wf = ses.Delete(curl.ParsedURL.String())
	}

	wf.SetHeader(curl.Header)
	wf.AddCookies(curl.Cookies)
	wf.SetContentType(curl.ContentType)
	wf.SetBody(curl.Body)
	wf.MergeQuery(curl.ParsedURL.Query())
	return wf
}

// Temporary 根据自己CreateSession 创建Temporary
func (curl *CURL) Temporary() *requests.Temporary {
	return curl.CreateTemporary(curl.CreateSession())
}

// Parse curl_bash
func Parse(scurl string) (cURL *CURL) {
	executor := newPQueueExecute()
	curl := New()

	if len(scurl) <= 4 {
		panic("scurl error:" + scurl)
	}

	if scurl[0] == '"' && scurl[len(scurl)-1] == '"' {
		scurl = strings.Trim(scurl, `"`)
	} else if scurl[0] == '\'' && scurl[len(scurl)-1] == '\'' {
		scurl = strings.Trim(scurl, `'`)
	}

	scurl = strings.TrimSpace(scurl)
	scurl = strings.TrimLeft(scurl, "curl")

	scurl = strings.ReplaceAll(scurl, "\n", "")

	if strings.HasPrefix(scurl, "http") {
		var parseurl []rune
		for _, v := range scurl {
			if v == ' ' {
				break
			}
			parseurl = append(parseurl, v)
		}

		purl, err := url.Parse(string(parseurl))
		if err != nil {
			panic(err)
		}
		curl.ParsedURL = purl
	}

	mathches := regexp.MustCompile(
		`--data-binary +\$.+--\\r\\n'([\n \t]|$)|`+
			`--[^ ]+ +'[^']+'([\n \t]|$)|`+
			`--[^ ]+ +"[^"]+"([\n \t]|$)|`+
			`--[^ ]+ +[^ ]+|`+
			`-[A-Za-z] +'[^']+'([\n \t]|$)|`+
			`-[A-Za-z] +"[^"]+"([\n \t]|$)|`+
			`-[A-Za-z] +[^ ]+|`+
			`[\n \t]'[^']+'([\n \t]|$)|`+
			`[\n \t]"[^"]+"([\n \t]|$)|`+
			`--[a-z]+ {0,}`,
	).FindAllString(scurl, -1)

	for _, m := range mathches {
		m = strings.Trim(m, " \n\t")
		switch v := m[0]; v {
		case '\'':
			purl, err := url.Parse(m[1 : len(m)-1])
			if err != nil {
				panic(err)
			}
			curl.ParsedURL = purl
		case '"':
			purl, err := url.Parse(m[1 : len(m)-1])
			if err != nil {
				panic(err)
			}
			curl.ParsedURL = purl
		case '-':
			exec := judgeOptions(curl, m)
			if exec != nil {
				executor.Push(exec)
			}
		}
	}

	for executor.Len() > 0 {
		exec := executor.Pop()
		exec.Execute()
	}

	if curl.Method == "" {
		curl.Method = "GET"
	}

	return curl
}
