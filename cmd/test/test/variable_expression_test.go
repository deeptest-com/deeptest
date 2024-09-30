package test

import (
	"fmt"
	commUtils "github.com/deeptest-com/deeptest/internal/pkg/utils"
	"log"
	"regexp"
	"strings"
	"testing"
)

func Test123(t *testing.T) {
	//str1 := "${var_name}"
	//
	//str2 := `${_mock("@date('yyyy-MM-dd')")}` // 不支持变量参数
	//
	//str3 := "${_url_encode(${url})}"
	//str4 := "${_url_encode('http://baidu.com')}"
	//str5 := `${_url_encode("http://baidu.com")}`

	str6 := "${math.add(${p1}, 1)} + ${_url_encode(${url})}"

	//isFunc, name, params := parseSingleVariableToken(str1)
	//log.Println(isFunc, name, params)
	//
	//isFunc, name, params = parseSingleVariableToken(str2)
	//log.Println(isFunc, name, params)
	//
	//isFunc, name, params = parseSingleVariableToken(str3)
	//log.Println(isFunc, name, params)
	//
	//isFunc, name, params = parseSingleVariableToken(str4)
	//log.Println(isFunc, name, params)
	//
	//isFunc, name, params = parseSingleVariableToken(str5)
	//log.Println(isFunc, name, params)

	value1 := replaceVariableToken(str6)
	variablePlaceholders := commUtils.GetVariablesInExpressionPlaceholder(value1)

	for _, placeholder := range variablePlaceholders {
		isFunc, name, params := parseSingleVariableToken(placeholder)
		log.Println(isFunc, name, params)
	}
}

func parseSingleVariableToken(str string) (isFunc bool, name string, params []TokenParam) {
	regx0 := regexp.MustCompile(`(?U)\${(\+?[A-Za-z0-9]+)}`)
	arr0 := regx0.FindAllStringSubmatch(str, -1)
	if len(arr0) > 0 {
		for _, item := range arr0 {
			str = strings.Replace(str,
				fmt.Sprintf("${%s}", item[1]),
				fmt.Sprintf("#[%s]", item[1]), 1)
		}
	}

	regx1 := regexp.MustCompile(`\${(\+?.+)}`)
	arr1 := regx1.FindAllStringSubmatch(str, -1)
	if len(arr1) == 0 {
		return
	}

	token := strings.TrimSpace(arr1[0][1])

	regx2 := regexp.MustCompile(`^((?U).+)\((\+?.+)\)$`)
	arr2 := regx2.FindAllStringSubmatch(token, -1)

	if len(arr2) == 0 { // is not func
		name = token
		return
	}

	// is func
	isFunc = true
	name = strings.TrimSpace(arr2[0][1])
	arr3 := strings.Split(arr2[0][2], ",")
	for _, item := range arr3 {
		paramName := strings.TrimSpace(item)
		if paramName == "" {
			continue
		}

		paramType := TokenParamTypeString

		if name == "_mock" {
			paramType = "mock"
			paramName = strings.Trim(paramName, `"'`)
		} else {
			regx4 := regexp.MustCompile(`#\[(\+?[A-Za-z0-9]+)\]`)
			arr4 := regx4.FindAllStringSubmatch(item, -1)

			if len(arr4) > 0 {
				paramType = TokenParamTypeVariable
				paramName = arr4[0][1]
			} else {
				paramNameTrim := strings.Trim(paramName, `"'`)
				if paramNameTrim == paramName { // number
					paramType = TokenParamTypeNumber
				}
			}
		}

		p := TokenParam{
			Name: paramName,
			Type: paramType,
		}

		params = append(params, p)
	}

	return
}

func replaceVariableToken(str string) (ret string) {
	ret = str

	reg := regexp.MustCompile(`(?U)\${(\+?[A-Za-z0-9]+)}`)
	ret = reg.ReplaceAllString(ret, "#[$1]")

	return
}

type TokenParam struct {
	Name string
	Type TokenParamType
}

type TokenParamType string

const (
	TokenParamTypeVariable TokenParamType = "variable"
	TokenParamTypeString   TokenParamType = "string"
	TokenParamTypeNumber   TokenParamType = "number"
)
