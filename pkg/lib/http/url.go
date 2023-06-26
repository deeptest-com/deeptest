package _httpUtils

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
		//utl += "/"
	}
	return utl
}
