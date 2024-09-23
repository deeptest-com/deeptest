package _http

import (
	"net/url"
	"strings"
)

func AddDefaultUrlSchema(urlStr string) (ret string, err error) {
	ret = urlStr

	u, err := url.Parse(urlStr)
	if err != nil {
		return
	}
	if u.Scheme == "" {
		u.Scheme = "https"
		ret = u.String()
	}

	return
}

func AddSepIfNeeded(utl string) string {
	if strings.LastIndex(utl, "/") < len(utl)-1 {
		utl += "/"
	}
	return utl
}

func CombineUrls(base, uri string) (ret string) {
	if base == "" {
		return uri
	}

	if uri == "" {
		return base
	}

	last := base[len(base)-1:]
	first := uri[:1]

	if last == "/" && first == "/" {
		ret = base[:len(base)-1] + uri
	} else if last == "/" || first == "/" {
		ret = base + uri
	} else {
		ret = AddSepIfNeeded(base) + uri
	}

	return
}
