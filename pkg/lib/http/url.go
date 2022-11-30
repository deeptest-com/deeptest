package _httpUtils

import "net/url"

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
