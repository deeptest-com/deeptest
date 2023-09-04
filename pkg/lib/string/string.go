package _stringUtils

import (
	"encoding/base64"
	"fmt"
	"github.com/jameskeane/bcrypt"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode"
)

func FindInArr(val string, arr []string) bool {
	for _, i := range arr {
		if val == i {
			return true
		}
	}

	return false
}

func TrimAll(str string) string {
	str = strings.Trim(str, "\n")
	str = strings.TrimSpace(str)

	return str
}

func UcFirst(str string) string {
	for i, v := range str {
		return string(unicode.ToUpper(v)) + str[i+1:]
	}
	return ""
}

func StrInArr(str string, arr []string) bool {
	for _, s := range arr {
		if str == s {
			return true
		}
	}

	return false
}

func RandStr(n int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	var src = rand.NewSource(time.Now().UnixNano())

	const (
		letterIdxBits = 6
		letterIdxMask = 1<<letterIdxBits - 1
		letterIdxMax  = 63 / letterIdxBits
	)
	b := make([]byte, n)
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return string(b)
}

func Base64Decode(str string) string {
	s, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return ""
	}
	return string(s)
}

func HashPassword(pwd string) string {
	salt, err := bcrypt.Salt(10)
	if err != nil {
		return ""
	}
	hash, err := bcrypt.Hash(pwd, salt)
	if err != nil {
		return ""
	}

	return hash
}

func ParseInt(str string) (ret int) {
	ret, _ = strconv.Atoi(str)
	return
}
func ParseUint(str string) (ret uint) {
	i := ParseInt(str)
	ret = uint(i)
	return
}

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func SnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

func UnescapeUnicode(raw []byte) ([]byte, error) {
	str, err := strconv.Unquote(strings.Replace(strconv.Quote(string(raw)), `\\u`, `\u`, -1))
	if err != nil {
		return nil, err
	}
	return []byte(str), nil
}

func unicode2utf8(source string) string {
	var res = []string{""}
	sUnicode := strings.Split(source, "\\u")
	var context = ""
	for _, v := range sUnicode {
		var additional = ""
		if len(v) < 1 {
			continue
		}
		if len(v) > 4 {
			rs := []rune(v)
			v = string(rs[:4])
			additional = string(rs[4:])
		}
		temp, err := strconv.ParseInt(v, 16, 32)
		if err != nil {
			context += v
		}
		context += fmt.Sprintf("%c", temp)
		context += additional
	}
	res = append(res, context)
	return strings.Join(res, "")
}

func StrToInt(s string) (ret int) {
	ret, _ = strconv.Atoi(s)
	return
}
func IntToStr(i int) (ret string) {
	ret = strconv.Itoa(i)
	return
}
func InterfToStr(i interface{}) (ret string) {
	ret = fmt.Sprintf("%v", i)
	return
}
